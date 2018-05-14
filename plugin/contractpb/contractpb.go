package contractpb

import (
	"errors"
	"fmt"
	"time"

	"github.com/gogo/protobuf/proto"

	"github.com/loomnetwork/go-loom"
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
	HasPermission(token []byte, roles []string) (bool, []string)
	HasPermissionFor(addr loom.Address, token []byte, roles []string) (bool, []string)
	GrantPermissionTo(addr loom.Address, token []byte, role string)
	GrantPermission(token []byte, roles []string)
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

// HasPermission checks whether the sender of the tx has any of the permission given in `roles` on `token`
func (c *wrappedPluginContext) HasPermission(token []byte, roles []string) (bool, []string) {
	addr := c.Message().Sender
	return c.HasPermissionFor(addr, token, roles)
}

// HasPermissionFor checks whether the given `addr` has any of the permission given in `roles` on `token`
func (c *wrappedPluginContext) HasPermissionFor(addr loom.Address, token []byte, roles []string) (bool, []string) {
	found := false
	foundRoles := []string{}
	for _, role := range roles {
		v := c.Context.Get(c.rolePermKey(addr, token, role))
		if v != nil && string(v) == "true" {
			found = true
			foundRoles = append(foundRoles, role)
		}
	}
	return found, foundRoles
}

// GrantPermissionTo sets a given `role` permission on `token` for the given `addr`
func (c *wrappedPluginContext) GrantPermissionTo(addr loom.Address, token []byte, role string) {
	c.Context.Set(c.rolePermKey(addr, token, role), []byte("true"))
}

func (c *wrappedPluginContext) rolePermKey(addr loom.Address, token []byte, role string) []byte {
	return []byte(fmt.Sprintf("%stoken:%s:role:%s", loom.PermPrefix(addr), token, []byte(role)))
}

// GrantPermission sets a given `role` permission on `token` for the sender of the tx
func (c *wrappedPluginContext) GrantPermission(token []byte, roles []string) {
	for _, r := range roles {
		c.GrantPermissionTo(c.Message().Sender, token, r)
	}
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

func CallEVM(ctx Context, addr loom.Address, input []byte, output *[]byte) error {
	resp, err := ctx.CallEVM(addr, input)
	*output = resp
	return err
}

func StaticCall(ctx Context, addr loom.Address, inpb proto.Message, outpb proto.Message) error {
	input, err := proto.Marshal(inpb)
	if err != nil {
		return err
	}

	output, err := ctx.StaticCall(addr, input)
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

func StaticCallEVM(ctx Context, addr loom.Address, input []byte, output *[]byte) error {
	resp, err := ctx.StaticCallEVM(addr, input)
	*output = resp
	return err
}

func WrapPluginContext(ctx plugin.Context) Context {
	return &wrappedPluginContext{ctx}
}
