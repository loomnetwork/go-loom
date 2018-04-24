package plugin

import (
	"encoding/json"
	"errors"
	"reflect"

	proto "github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
	lt "github.com/loomnetwork/loom-plugin/types"
)

type (
	ContractMethodCall     = lt.ContractMethodCall
	ContractMethodCallJSON = lt.ContractMethodCallJSON
)

// RequestDispatcher dispatches Request(s) to contract methods.
// The dispatcher takes care of unmarshalling requests and marshalling responses from/to protobufs
// or JSON - based on the content type specified in the Request.ContentType/Accept fields.
type RequestDispatcher struct {
	callbacks *serviceMap
}

func (s *RequestDispatcher) Init(contract Contract) error {
	s.callbacks = new(serviceMap)
	meta, err := contract.Meta()
	if err != nil {
		return err
	}
	return s.callbacks.Register(contract, meta.Name)
}

func (s *RequestDispatcher) StaticCall(ctx StaticContext, req *Request) (*Response, error) {
	var result []reflect.Value
	if req.ContentType == EncodingType_JSON {
		var query lt.ContractMethodCallJSON
		if err := json.Unmarshal(req.Body, &query); err != nil {
			return nil, err
		}
		serviceSpec, methodSpec, err := s.callbacks.Get(query.Method, true)
		if err != nil {
			return nil, err
		}
		queryParams := reflect.New(methodSpec.argsType)
		if err := json.Unmarshal(query.Data, queryParams.Interface()); err != nil {
			return nil, err
		}
		result = methodSpec.method.Func.Call([]reflect.Value{
			serviceSpec.rcvr,
			reflect.ValueOf(ctx),
			queryParams,
		})
	} else if req.ContentType == EncodingType_PROTOBUF3 {
		var query lt.ContractMethodCall
		if err := proto.Unmarshal(req.Body, &query); err != nil {
			return nil, err
		}
		serviceSpec, methodSpec, err := s.callbacks.Get(query.Method, true)
		if err != nil {
			return nil, err
		}
		queryParams := reflect.New(methodSpec.argsType)
		if err := types.UnmarshalAny(query.Data, queryParams.Interface().(proto.Message)); err != nil {
			return nil, err
		}
		result = methodSpec.method.Func.Call([]reflect.Value{
			serviceSpec.rcvr,
			reflect.ValueOf(ctx),
			queryParams,
		})
	} else {
		return nil, errors.New("unsupported content type")
	}

	// If the method returned an error, extract & return it
	var err error
	errInter := result[1].Interface()
	if errInter != nil {
		err = errInter.(error)
	}
	if err != nil {
		return nil, err
	}

	resp := &Response{ContentType: req.Accept}
	if req.Accept == EncodingType_JSON {
		resp.Body, err = json.Marshal(result[0].Interface())
	} else if req.Accept == EncodingType_PROTOBUF3 {
		resp.Body, err = proto.Marshal(result[0].Interface().(proto.Message))
	} else {
		return nil, errors.New("unsupported accept type")
	}
	return resp, err
}

func (s *RequestDispatcher) Call(ctx Context, req *Request) (*Response, error) {
	// TODO: handle req.ContentType/Accept == JSON
	var tx lt.ContractMethodCall
	if err := proto.Unmarshal(req.Body, &tx); err != nil {
		return nil, err
	}

	serviceSpec, methodSpec, err := s.callbacks.Get(tx.Method, false)
	if err != nil {
		return nil, err
	}

	txData := reflect.New(methodSpec.argsType)

	if err := types.UnmarshalAny(tx.Data, txData.Interface().(proto.Message)); err != nil {
		return nil, err
	}

	//Lookup the method we need to call
	errValue := methodSpec.method.Func.Call([]reflect.Value{
		serviceSpec.rcvr,
		reflect.ValueOf(ctx),
		txData,
	})

	// Cast the result to error if needed.
	var errResult error
	errInter := errValue[0].Interface()
	if errInter != nil {
		errResult = errInter.(error)
	}

	if errResult != nil {
		return nil, errResult
	}
	return &Response{}, nil
}
