package plugin

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/loomnetwork/loom-plugin/testdata"
	"github.com/loomnetwork/loom-plugin/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var callArgs = testdata.CallArgs{
	Key:   "test",
	Value: 12345,
}

var staticCallArgs = testdata.StaticCallArgs{
	Index: 5,
	Name:  "john",
}

type fakeRequestDispatcherContract struct {
	RequestDispatcher
	t *testing.T
}

func (c *fakeRequestDispatcherContract) Meta() (Meta, error) {
	return Meta{
		Name:    "fakecontract",
		Version: "0.0.1",
	}, nil
}

func (c *fakeRequestDispatcherContract) Init(ctx Context, req *Request) error {
	return nil
}

func (c *fakeRequestDispatcherContract) HandleTx(ctx Context, args *testdata.CallArgs) error {
	require.Equal(c.t, &callArgs, args)
	return nil
}

func (c *fakeRequestDispatcherContract) HandleQuery(ctx Context, args *testdata.StaticCallArgs) (*testdata.StaticCallResult, error) {
	require.Equal(c.t, &staticCallArgs, args)
	return &testdata.StaticCallResult{
		Code:   123,
		Result: "QueryResult",
	}, nil
}

func newFakeRequestDispatcherContract(t *testing.T) *fakeRequestDispatcherContract {
	c := &fakeRequestDispatcherContract{t: t}
	c.RequestDispatcher.Init(c)
	return c
}

func TestEmbeddedRequestDispatcherDoesNotRegisterOwnMethods(t *testing.T) {
	c := newFakeRequestDispatcherContract(t)
	var err error
	_, _, err = c.RequestDispatcher.callbacks.Get("fakecontract.Call", false)
	require.NotNil(t, err)
	_, _, err = c.RequestDispatcher.callbacks.Get("fakecontract.StaticCall", true)
	require.NotNil(t, err)
}

func TestRequestDispatcherCallMethod(t *testing.T) {
	c := newFakeRequestDispatcherContract(t)
	meta, err := c.Meta()
	require.Nil(t, err)

	encodings := []EncodingType{EncodingType_JSON, EncodingType_PROTOBUF3}

	for _, encoding := range encodings {
		marshaler, err := marshalerFactory(encoding)
		require.Nil(t, err)

		var argsBuffer bytes.Buffer
		require.Nil(t, marshaler.Marshal(&argsBuffer, &callArgs))

		msg := &types.ContractMethodCall{
			Method: fmt.Sprintf("%s.HandleTx", meta.Name),
			Args:   argsBuffer.Bytes(),
		}

		var msgBuffer bytes.Buffer
		require.Nil(t, marshaler.Marshal(&msgBuffer, msg))

		req := &Request{
			ContentType: encoding,
			Accept:      encoding,
			Body:        msgBuffer.Bytes(),
		}

		resp, err := c.Call(CreateFakeContext(), req)
		require.Nil(t, err)
		require.Equal(t, req.Accept, resp.ContentType)
	}
}

func TestRequestDispatcherStaticCallMethod(t *testing.T) {
	c := newFakeRequestDispatcherContract(t)
	meta, err := c.Meta()
	require.Nil(t, err)

	encodings := []EncodingType{EncodingType_JSON, EncodingType_PROTOBUF3}

	for _, encoding := range encodings {
		marshaler, err := marshalerFactory(encoding)
		require.Nil(t, err)

		var argsBuffer bytes.Buffer
		require.Nil(t, marshaler.Marshal(&argsBuffer, &staticCallArgs))

		msg := &types.ContractMethodCall{
			Method: fmt.Sprintf("%s.HandleQuery", meta.Name),
			Args:   argsBuffer.Bytes(),
		}

		var msgBuffer bytes.Buffer
		require.Nil(t, marshaler.Marshal(&msgBuffer, msg))

		req := &Request{
			ContentType: encoding,
			Accept:      encoding,
			Body:        msgBuffer.Bytes(),
		}

		resp, err := c.StaticCall(CreateFakeContext(), req)
		require.Nil(t, err)
		require.Equal(t, req.Accept, resp.ContentType)

		unmarshaler, err := unmarshalerFactory(resp.ContentType)
		require.Nil(t, err)

		var callResult testdata.StaticCallResult
		err = unmarshaler.Unmarshal(bytes.NewBuffer(resp.Body), &callResult)
		require.Nil(t, err)
		assert.Equal(t, int32(123), callResult.Code)
		assert.Equal(t, "QueryResult", callResult.Result)
	}
}
