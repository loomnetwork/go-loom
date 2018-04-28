package contractpb

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"

	"github.com/gogo/protobuf/proto"
	"github.com/loomnetwork/go-loom/plugin"
	"github.com/loomnetwork/go-loom/types"
)

// RequestDispatcher dispatches Request(s) to contract methods.
// The dispatcher takes care of unmarshalling requests and marshalling responses from/to protobufs
// or JSON - based on the content type specified in the Request.ContentType/Accept fields.
type RequestDispatcher struct {
	Contract
	callbacks *serviceMap
}

func NewRequestDispatcher(contract Contract) (*RequestDispatcher, error) {
	s := &RequestDispatcher{
		Contract:  contract,
		callbacks: new(serviceMap),
	}
	meta, err := contract.Meta()
	if err != nil {
		return nil, err
	}
	err = s.callbacks.Register(contract, meta.Name)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (s *RequestDispatcher) Init(ctx plugin.Context, req *plugin.Request) error {
	_, err := s.doCall(methodSigInit, &wrappedPluginContext{Context: ctx}, req)
	return err
}

func (s *RequestDispatcher) StaticCall(ctx plugin.StaticContext, req *plugin.Request) (*plugin.Response, error) {
	return s.doCall(methodSigStaticCall, &wrappedPluginStaticContext{StaticContext: ctx}, req)
}

func (s *RequestDispatcher) Call(ctx plugin.Context, req *plugin.Request) (*plugin.Response, error) {
	return s.doCall(methodSigCall, &wrappedPluginContext{Context: ctx}, req)
}

func (s *RequestDispatcher) doCall(sig methodSig, ctx interface{}, req *plugin.Request) (*plugin.Response, error) {
	body := bytes.NewBuffer(req.Body)
	unmarshaler, err := unmarshalerFactory(req.ContentType)
	if err != nil {
		return nil, err
	}
	marshaler, err := marshalerFactory(req.Accept)
	if err != nil {
		return nil, err
	}

	var query types.ContractMethodCall
	err = unmarshaler.Unmarshal(body, &query)
	if err != nil {
		return nil, err
	}

	serviceSpec, methodSpec, err := s.callbacks.Get(query.Method)
	if err != nil {
		return nil, err
	}

	if methodSpec.methodSig != sig {
		fmt.Printf("%v %v\n", sig, methodSpec.methodSig)
		return nil, errors.New("method call does not match method signature type")
	}

	queryParams := reflect.New(methodSpec.argsType)
	err = unmarshaler.Unmarshal(bytes.NewBuffer(query.Args), queryParams.Interface().(proto.Message))
	if err != nil {
		return nil, err
	}

	resultTypes := methodSpec.method.Func.Call([]reflect.Value{
		serviceSpec.rcvr,
		reflect.ValueOf(ctx),
		queryParams,
	})

	var resp bytes.Buffer
	if len(resultTypes) > 0 {
		err, _ = resultTypes[len(resultTypes)-1].Interface().(error)
		if err != nil {
			return nil, err
		}

		pb, _ := resultTypes[0].Interface().(proto.Message)
		if pb != nil {
			err = marshaler.Marshal(&resp, pb)
			if err != nil {
				return nil, err
			}
		}
	}

	return &plugin.Response{
		ContentType: req.Accept,
		Body:        resp.Bytes(),
	}, nil
}
