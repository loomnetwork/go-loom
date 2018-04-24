package plugin

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type fakeRequestDispatcherContract struct {
	RequestDispatcher
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

func (c *fakeRequestDispatcherContract) HandleTx(ctx Context, req *Request) error {
	return nil
}

func (c *fakeRequestDispatcherContract) HandleQuery(ctx Context, req *Request) (*Response, error) {
	return nil, nil
}

func newFakeRequestDispatcherContract() *fakeRequestDispatcherContract {
	c := &fakeRequestDispatcherContract{}
	c.RequestDispatcher.Init(c)
	return c
}

func TestEmbeddedRequestDispatcherDoesNotRegisterOwnMethods(t *testing.T) {
	c := newFakeRequestDispatcherContract()
	var err error
	_, _, err = c.RequestDispatcher.callbacks.Get("fakecontract.Call", false)
	require.NotNil(t, err)
	_, _, err = c.RequestDispatcher.callbacks.Get("fakecontract.StaticCall", true)
	require.NotNil(t, err)
}
