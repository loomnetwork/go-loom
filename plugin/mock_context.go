package plugin

import (
	"time"

	loom "github.com/loomnetwork/go-loom"
	"github.com/loomnetwork/go-loom/types"
)

type FakeContext struct {
	caller  loom.Address
	address loom.Address
	data    map[string][]byte
}

var _ Context = &FakeContext{}

func CreateFakeContext(caller, address loom.Address) *FakeContext {
	return &FakeContext{
		caller:  caller,
		address: address,
		data:    make(map[string][]byte),
	}
}

func (c *FakeContext) WithSender(caller loom.Address) Context {
	return &FakeContext{
		caller:  caller,
		address: c.address,
		data:    c.data,
	}
}

func (c *FakeContext) Call(addr loom.Address, input []byte) ([]byte, error) {
	return nil, nil
}

func (c *FakeContext) CallEVM(addr loom.Address, input []byte) ([]byte, error) {
	return nil, nil
}

func (c *FakeContext) StaticCall(addr loom.Address, input []byte) ([]byte, error) {
	return nil, nil
}

func (c *FakeContext) Resolve(name string) (loom.Address, error) {
	return loom.Address{}, nil
}

func (c *FakeContext) ValidatorPower(pubKey []byte) int64 {
	return 0
}

func (c *FakeContext) Message() Message {
	return Message{
		Sender: c.caller,
	}
}

func (c *FakeContext) Block() types.BlockHeader {
	return types.BlockHeader{}
}

func (c *FakeContext) ContractAddress() loom.Address {
	return c.address
}

func (c *FakeContext) Now() time.Time {
	return time.Unix(0, 0)
}

func (c *FakeContext) Emit(event []byte) {
}

func (c *FakeContext) Get(key []byte) []byte {
	v, _ := c.data[string(key)]
	return v
}

func (c *FakeContext) Has(key []byte) bool {
	_, ok := c.data[string(key)]
	return ok
}

func (c *FakeContext) Set(key []byte, value []byte) {
	c.data[string(key)] = value
}

func (c *FakeContext) Delete(key []byte) {
}

func (c *FakeContext) SetValidatorPower(pubKey []byte, power int64) {
}

func (c *FakeContext) HasPermission(token []byte, roles []string) (bool, []string) {
	addr := c.Message().Sender
	return c.HasPermissionFor(addr, token, roles)
}

// HasPermissionFor checks whether the given `addr` has any of the permission given in `roles` on `token`
func (c *FakeContext) HasPermissionFor(addr loom.Address, token []byte, roles []string) (bool, []string) {
	return false, nil
}

// GrantPermissionTo sets a given `role` permission on `token` for the given `addr`
func (c *FakeContext) GrantPermissionTo(addr loom.Address, token []byte, role string) {
}

// GrantPermission sets a given `role` permission on `token` for the sender of the tx
func (c *FakeContext) GrantPermission(token []byte, roles []string) {
}
