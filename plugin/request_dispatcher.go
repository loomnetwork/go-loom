package plugin

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"reflect"

	"github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	lt "github.com/loomnetwork/go-loom/types"
)

var (
	errUnknownEncodingType = errors.New("unknown encoding type")
)

// RequestDispatcher dispatches Request(s) to contract methods.
// The dispatcher takes care of unmarshalling requests and marshalling responses from/to protobufs
// or JSON - based on the content type specified in the Request.ContentType/Accept fields.
type RequestDispatcher struct {
	callbacks *serviceMap
}

type PBMarshaler interface {
	Marshal(w io.Writer, pb proto.Message) error
}

type PBUnmarshaler interface {
	Unmarshal(r io.Reader, pb proto.Message) error
}

type BinaryPBMarshaler struct {
}

func (m *BinaryPBMarshaler) Marshal(w io.Writer, pb proto.Message) error {
	b, err := proto.Marshal(pb)
	if err != nil {
		return err
	}
	_, err = w.Write(b)
	return err
}

type BinaryPBUnmarshaler struct {
}

func (m *BinaryPBUnmarshaler) Unmarshal(r io.Reader, pb proto.Message) error {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	return proto.Unmarshal(b, pb)
}

func marshalerFactory(encoding EncodingType) (PBMarshaler, error) {
	switch encoding {
	case EncodingType_JSON:
		return &jsonpb.Marshaler{}, nil
	case EncodingType_PROTOBUF3:
		return &BinaryPBMarshaler{}, nil
	}

	return nil, errUnknownEncodingType
}

func unmarshalerFactory(encoding EncodingType) (PBUnmarshaler, error) {
	switch encoding {
	case EncodingType_JSON:
		return &jsonpb.Unmarshaler{}, nil
	case EncodingType_PROTOBUF3:
		return &BinaryPBUnmarshaler{}, nil
	}

	return nil, errUnknownEncodingType
}

func (s *RequestDispatcher) Init(contract Contract) error {
	s.callbacks = new(serviceMap)
	meta, err := contract.Meta()
	if err != nil {
		return err
	}
	return s.callbacks.Register(contract, meta.Name)
}

func unmarshalBody(req *Request, pb proto.Message) error {
	body := bytes.NewBuffer(req.Body)
	unmarshaler, err := unmarshalerFactory(req.ContentType)
	if err != nil {
		return err
	}

	return unmarshaler.Unmarshal(body, pb)
}

func (s *RequestDispatcher) StaticCall(ctx StaticContext, req *Request) (*Response, error) {
	var query lt.ContractMethodCall
	err := unmarshalBody(req, &query)
	if err != nil {
		return nil, err
	}
	serviceSpec, methodSpec, err := s.callbacks.Get(query.Method, true)
	if err != nil {
		return nil, err
	}
	queryParams := reflect.New(methodSpec.argsType)
	unmarshaler, err := unmarshalerFactory(req.ContentType)
	if err != nil {
		return nil, err
	}

	if err := unmarshaler.Unmarshal(bytes.NewBuffer(query.Args), queryParams.Interface().(proto.Message)); err != nil {
		return nil, err
	}
	result := methodSpec.method.Func.Call([]reflect.Value{
		serviceSpec.rcvr,
		reflect.ValueOf(ctx),
		queryParams,
	})

	// If the method returned an error, extract & return it
	errInter := result[1].Interface()
	if errInter != nil {
		err = errInter.(error)
	}
	if err != nil {
		return nil, err
	}

	marshaler, err := marshalerFactory(req.Accept)
	if err != nil {
		return nil, err
	}

	var respBody bytes.Buffer
	err = marshaler.Marshal(&respBody, result[0].Interface().(proto.Message))
	if err != nil {
		return nil, err
	}

	return &Response{
		ContentType: req.Accept,
		Body:        respBody.Bytes(),
	}, nil
}

func (s *RequestDispatcher) Call(ctx Context, req *Request) (*Response, error) {
	var tx lt.ContractMethodCall
	err := unmarshalBody(req, &tx)
	if err != nil {
		return nil, err
	}

	serviceSpec, methodSpec, err := s.callbacks.Get(tx.Method, false)
	if err != nil {
		return nil, err
	}

	txData := reflect.New(methodSpec.argsType)
	unmarshaler, err := unmarshalerFactory(req.ContentType)
	if err != nil {
		return nil, err
	}

	if err := unmarshaler.Unmarshal(bytes.NewBuffer(tx.Args), txData.Interface().(proto.Message)); err != nil {
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
	return &Response{
		ContentType: req.Accept,
	}, nil
}
