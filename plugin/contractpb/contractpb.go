package contractpb

import (
	"errors"
	"time"

	proto "github.com/gogo/protobuf/proto"

	loom "github.com/loomnetwork/go-loom"
	"github.com/loomnetwork/go-loom/plugin"
)

var (
	ErrNotFound = errors.New("not found")
)

type StaticContext interface {
	plugin.StaticAPI
	Get(key []byte, pb proto.Message) error
	Has(key []byte) bool
	Block() loom.BlockHeader
	Now() time.Time
	Message() plugin.Message
	ContractAddress() loom.Address
}

type Context interface {
	plugin.VolatileAPI
	StaticContext
	Set(key []byte, pb proto.Message) error
	Delete(key []byte)
}

type Contract interface {
	Meta() (plugin.Meta, error)
}

type wrappedPluginStaticContext struct {
	plugin.StaticContext
}

var _ StaticContext = &wrappedPluginStaticContext{}

func (c *wrappedPluginStaticContext) Get(key []byte, pb proto.Message) error {
	data := c.StaticContext.Get(key)
	if len(data) == 0 {
		return ErrNotFound
	}

	return proto.Unmarshal(data, pb)
}

type wrappedPluginContext struct {
	plugin.Context
}

var _ Context = &wrappedPluginContext{}

func (c *wrappedPluginContext) Get(key []byte, pb proto.Message) error {
	data := c.Context.Get(key)
	if len(data) == 0 {
		return ErrNotFound
	}

	return proto.Unmarshal(data, pb)
}

func (c *wrappedPluginContext) Set(key []byte, pb proto.Message) error {
	enc, err := proto.Marshal(pb)
	if err != nil {
		return err
	}
	c.Context.Set(key, enc)
	return nil
}

func MakePluginContract(c Contract) plugin.Contract {
	r, err := NewRequestDispatcher(c)
	if err != nil {
		panic(err)
	}

	return r
}

func Call(ctx Context, addr loom.Address, inpb proto.Message, outpb proto.Message) error {
	input, err := proto.Marshal(inpb)
	if err != nil {
		return err
	}

	output, err := ctx.Call(addr, input)
	if err != nil {
		return err
	}

	if outpb != nil {
		err = proto.Unmarshal(output, outpb)
		if err != nil {
			return err
		}
	}

	return nil
}

func WrapPluginContext(ctx plugin.Context) Context {
	return &wrappedPluginContext{ctx}
}
