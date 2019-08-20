package plugin

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strings"
	"time"

	"github.com/gogo/protobuf/proto"
	loom "github.com/loomnetwork/go-loom"
	cctypes "github.com/loomnetwork/go-loom/builtin/types/chainconfig"
	"github.com/loomnetwork/go-loom/config"
	ptypes "github.com/loomnetwork/go-loom/plugin/types"
	"github.com/loomnetwork/go-loom/types"
	"github.com/loomnetwork/go-loom/util"
	"github.com/pkg/errors"
	"golang.org/x/crypto/sha3"
)

type FEvent struct {
	Event  []byte
	Topics []string
}

type FakeContext struct {
	caller        loom.Address
	address       loom.Address
	block         loom.BlockHeader
	data          map[string][]byte
	contractNonce uint64
	contracts     map[string]Contract
	registry      map[string]*ContractRecord
	validators    loom.ValidatorSet
	Events        []FEvent
	ethBalances   map[string]*loom.BigUInt
	features      map[string]bool
	config        *cctypes.Config
}

var _ Context = &FakeContext{}

func createAddress(parent loom.Address, nonce uint64) loom.Address {
	var nonceBuf bytes.Buffer
	err := binary.Write(&nonceBuf, binary.BigEndian, nonce)
	if err != nil {
		panic(err)
	}
	data := util.PrefixKey(parent.Bytes(), nonceBuf.Bytes())
	hash := sha3.Sum256(data)
	return loom.Address{
		ChainID: parent.ChainID,
		Local:   hash[12:],
	}
}

func CreateFakeContext(caller, address loom.Address) *FakeContext {
	return &FakeContext{
		caller:      caller,
		address:     address,
		data:        make(map[string][]byte),
		contracts:   make(map[string]Contract),
		registry:    make(map[string]*ContractRecord),
		validators:  loom.NewValidatorSet(),
		Events:      make([]FEvent, 0),
		ethBalances: make(map[string]*loom.BigUInt),
		features:    make(map[string]bool),
		config:      config.DefaultConfig(),
	}
}

func (c *FakeContext) shallowClone() *FakeContext {
	return &FakeContext{
		caller:        c.caller,
		address:       c.address,
		block:         c.block,
		data:          c.data,
		contractNonce: c.contractNonce,
		contracts:     c.contracts,
		registry:      c.registry,
		validators:    c.validators,
		Events:        c.Events,
		ethBalances:   c.ethBalances,
		features:      c.features,
		config:        c.config,
	}
}

func (c *FakeContext) WithBlock(header loom.BlockHeader) *FakeContext {
	clone := c.shallowClone()
	clone.block = header
	return clone
}

func (c *FakeContext) WithSender(caller loom.Address) *FakeContext {
	clone := c.shallowClone()
	clone.caller = caller
	return clone
}

func (c *FakeContext) WithAddress(addr loom.Address) *FakeContext {
	clone := c.shallowClone()
	clone.address = addr
	return clone
}

func (c *FakeContext) WithValidators(validators []*types.Validator) *FakeContext {
	clone := c.shallowClone()
	clone.validators = loom.NewValidatorSet(validators...)
	return clone
}

func (c *FakeContext) CreateContract(contract Contract) loom.Address {
	addr := createAddress(c.address, c.contractNonce)
	c.contractNonce++
	address := addr.String()
	c.contracts[address] = contract
	return addr

}

func (c *FakeContext) RegisterContract(contractName string, contractAddr, creatorAddr loom.Address) {
	c.registry[contractAddr.String()] = &ContractRecord{
		ContractName:    contractName,
		ContractAddress: contractAddr,
		CreatorAddress:  creatorAddr,
	}
	c.registry[contractName] = &ContractRecord{
		ContractName:    contractName,
		ContractAddress: contractAddr,
		CreatorAddress:  creatorAddr,
	}
}

func (c *FakeContext) Call(addr loom.Address, input []byte) ([]byte, error) {
	contract := c.contracts[addr.String()]

	ctx := c.WithSender(c.address).WithAddress(addr)

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

	ctx := c.WithSender(c.address).WithAddress(addr)

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

func (c *FakeContext) CallEVM(addr loom.Address, input []byte, value *loom.BigUInt) ([]byte, error) {
	return nil, nil
}

func (c *FakeContext) StaticCallEVM(addr loom.Address, input []byte) ([]byte, error) {
	return nil, nil
}

func (c *FakeContext) Resolve(name string) (loom.Address, error) {
	for addrStr, contract := range c.contracts {
		meta, err := contract.Meta()
		if err != nil {
			return loom.Address{}, err
		}
		if meta.Name == name {
			return loom.MustParseAddress(addrStr), nil
		}
	}
	record, ok := c.registry[name]
	if ok {
		return record.ContractAddress, nil
	}

	return loom.Address{}, fmt.Errorf("failed  to resolve address of contract '%s'", name)
}

func (c *FakeContext) SetFeature(name string, val bool) {
	c.features[name] = val
}

func (c *FakeContext) FeatureEnabled(name string, defaultVal bool) bool {
	if val, ok := c.features[name]; ok {
		return val
	}
	return defaultVal
}

func (c *FakeContext) SetConfigSetting(name, value string) error {
	if err := config.SetConfigSetting(c.config, name, value); err != nil {
		return err
	}
	return nil
}

func (c *FakeContext) Config() *cctypes.Config {
	return c.config
}

func (c *FakeContext) EnabledFeatures() []string {
	return nil
}

func (c *FakeContext) Message() Message {
	return Message{
		Sender: c.caller,
	}
}

func (c *FakeContext) Block() types.BlockHeader {
	return c.block
}

func (c *FakeContext) ContractAddress() loom.Address {
	return c.address
}

func (c *FakeContext) GetEvmTxReceipt([]byte) (ptypes.EvmTxReceipt, error) {
	return ptypes.EvmTxReceipt{}, nil
}

func (c *FakeContext) SetTime(t time.Time) {
	c.block.Time = t.Unix()
}

func (c *FakeContext) Now() time.Time {
	return time.Unix(c.block.Time, 0)
}

func (c *FakeContext) EmitTopics(event []byte, topics ...string) {
	//Store last emitted strings, to make it testable
	c.Events = append(c.Events, FEvent{event, topics})
}

func (c *FakeContext) Emit(event []byte) {
}

// Prefix the given key with the contract address
func (c *FakeContext) makeKey(key []byte) string {
	return string(util.PrefixKey(c.address.Bytes(), key))
}

// Strip the contract address from the given key (i.e. inverse of makeKey)
func (c *FakeContext) recoverKey(key string, prefix []byte) ([]byte, error) {
	return util.UnprefixKey([]byte(key), util.PrefixKey(c.address.Bytes(), prefix))
}

func (c *FakeContext) Range(prefix []byte) RangeData {
	ret := make(RangeData, 0)
	keyedPrefix := c.makeKey(prefix)
	for key, value := range c.data {
		if strings.HasPrefix(key, keyedPrefix) {
			k, err := c.recoverKey(key, prefix)
			if err != nil {
				panic(err)
			}
			r := &RangeEntry{
				Key:   k,
				Value: value,
			}

			ret = append(ret, r)
		}
	}
	return ret
}

func (c *FakeContext) Get(key []byte) []byte {
	v := c.data[c.makeKey(key)]
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

func (c *FakeContext) Validators() []*loom.Validator {
	return c.validators.Slice()
}

func (c *FakeContext) ContractRecord(contractAddr loom.Address) (*ContractRecord, error) {
	rec := c.registry[contractAddr.String()]
	if rec == nil {
		return nil, errors.New("contract not found in registry")
	}
	return rec, nil
}
