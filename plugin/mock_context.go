package plugin

import (
	"bytes"
	"encoding/binary"
	"time"

	"golang.org/x/crypto/sha3"

	"github.com/gogo/protobuf/proto"
	loom "github.com/loomnetwork/go-loom"
	"github.com/loomnetwork/go-loom/types"
	"github.com/loomnetwork/go-loom/util"
)

type FakeContext struct {
	caller        loom.Address
	address       loom.Address
	block         loom.BlockHeader
	data          map[string][]byte
	contractNonce uint64
	contracts     map[string]Contract
	validators    loom.ValidatorSet
}

var _ Context = &FakeContext{}

func createAddress(parent loom.Address, nonce uint64) loom.Address {
	var nonceBuf bytes.Buffer
	binary.Write(&nonceBuf, binary.BigEndian, nonce)
	data := util.PrefixKey(parent.Bytes(), nonceBuf.Bytes())
	hash := sha3.Sum256(data)
	return loom.Address{
		ChainID: parent.ChainID,
		Local:   hash[12:],
	}
}

func CreateFakeContext(caller, address loom.Address) *FakeContext {
	return &FakeContext{
		caller:     caller,
		address:    address,
		data:       make(map[string][]byte),
		contracts:  make(map[string]Contract),
		validators: loom.NewValidatorSet(),
	}
}

func (c *FakeContext) WithBlock(header loom.BlockHeader) *FakeContext {
	return &FakeContext{
		caller:     c.caller,
		address:    c.address,
		data:       c.data,
		contracts:  c.contracts,
		validators: c.validators,
		block:      header,
	}
}

func (c *FakeContext) WithSender(caller loom.Address) *FakeContext {
	return &FakeContext{
		caller:     caller,
		address:    c.address,
		data:       c.data,
		contracts:  c.contracts,
		validators: c.validators,
		block:      c.block,
	}
}

func (c *FakeContext) WithAddress(addr loom.Address) *FakeContext {
	return &FakeContext{
		caller:     c.caller,
		address:    addr,
		data:       c.data,
		contracts:  c.contracts,
		validators: c.validators,
		block:      c.block,
	}
}

func (c *FakeContext) CreateContract(contract Contract) loom.Address {
	addr := createAddress(c.address, c.contractNonce)
	c.contractNonce++
	c.contracts[addr.String()] = contract
	return addr
}

func (c *FakeContext) Call(addr loom.Address, input []byte) ([]byte, error) {
	contract := c.contracts[addr.String()]

	ctx := &FakeContext{
		caller:     c.address,
		address:    addr,
		data:       c.data,
		contracts:  c.contracts,
		validators: c.validators,
	}

	var req Request
	err := proto.Unmarshal(input, &req)
	if err != nil {
		return nil, err
	}

	resp, err := contract.Call(ctx, &req)
	if err != nil {
		return nil, err
	}

	return proto.Marshal(resp)
}

func (c *FakeContext) StaticCall(addr loom.Address, input []byte) ([]byte, error) {
	contract := c.contracts[addr.String()]

	ctx := &FakeContext{
		caller:    c.address,
		address:   addr,
		data:      c.data,
		contracts: c.contracts,
	}

	var req Request
	err := proto.Unmarshal(input, &req)
	if err != nil {
		return nil, err
	}

	resp, err := contract.StaticCall(ctx, &req)
	if err != nil {
		return nil, err
	}

	return proto.Marshal(resp)
}

func (c *FakeContext) CallEVM(addr loom.Address, input []byte) ([]byte, error) {
	return nil, nil
}

func (c *FakeContext) StaticCallEVM(addr loom.Address, input []byte) ([]byte, error) {
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
	return time.Unix(c.block.Time, 0)
}

func (c *FakeContext) Emit(event []byte) {
}

func (c *FakeContext) makeKey(key []byte) string {
	return string(util.PrefixKey(c.address.Bytes(), key))
}

func (c *FakeContext) Get(key []byte) []byte {
	v, _ := c.data[c.makeKey(key)]
	return v
}

func (c *FakeContext) Has(key []byte) bool {
	_, ok := c.data[c.makeKey(key)]
	return ok
}

func (c *FakeContext) Set(key []byte, value []byte) {
	c.data[c.makeKey(key)] = value
}

func (c *FakeContext) Delete(key []byte) {
	delete(c.data, c.makeKey(key))
}

func (c *FakeContext) SetValidatorPower(pubKey []byte, power int64) {
	c.validators.Set(&loom.Validator{PubKey: pubKey, Power: power})
}

func (c *FakeContext) Validators() []*loom.Validator {
	return c.validators.Slice()
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
