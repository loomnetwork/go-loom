package plugin

import (
	"time"

	loom "github.com/loomnetwork/go-loom"
	"github.com/loomnetwork/go-loom/plugin/types"
)

type (
	Request            = types.Request
	Response           = types.Response
	Meta               = types.ContractMeta
	EncodingType       = types.EncodingType
	ContractMethodCall = types.ContractMethodCall
	Code               = types.PluginCode
)

var (
	EncodingType_JSON      = types.EncodingType_JSON
	EncodingType_PROTOBUF3 = types.EncodingType_PROTOBUF3
)

type Message struct {
	Sender loom.Address
}

type StaticAPI interface {
	StaticCall(addr loom.Address, input []byte) ([]byte, error)
	StaticCallEVM(addr loom.Address, input []byte) ([]byte, error)
	Resolve(name string) (loom.Address, error)
	ValidatorPower(pubKey []byte) int64
	EmitTopics(event []byte, topics ...string) error
	Emit(event []byte) error
}

type VolatileAPI interface {
	Call(addr loom.Address, input []byte) ([]byte, error)
	CallEVM(addr loom.Address, input []byte, value *loom.BigUInt) ([]byte, error)
	// Privileged API
	SetValidatorPower(pubKey []byte, power int64) (error)
}

type API interface {
	StaticAPI
	VolatileAPI
}

// RangeEntry a single entry in a range
type RangeEntry struct {
	Key   []byte
	Value []byte
}

// RangeData an array of key value pairs for a range of data
type RangeData []*RangeEntry

type ContractRecord struct {
	ContractName    string
	ContractAddress loom.Address
	CreatorAddress  loom.Address
}

// StaticContext is the low-level context provided to RequestDispatcher.Call().
// The primary implementation of this interface is plugin.contractContext (loomchain/plugin package).
// For external GRPC contracts plugin.contractContext is wrapped by GRPCContext (go-loom/plugin package).
type StaticContext interface {
	StaticAPI
	Get(key []byte) ([]byte,error)
	Has(key []byte) bool
	Range(prefix []byte) (RangeData,error)
	Block() loom.BlockHeader
	Now() time.Time
	Message() Message
	GetEvmTxReceipt([]byte) (types.EvmTxReceipt, error)
	ContractAddress() loom.Address
	FeatureEnabled(name string, defaultVal bool) bool
	// ContractRecord retrieves the contract meta data stored in the Registry.
	// NOTE: This method requires Registry v2.
	ContractRecord(contractAddr loom.Address) (*ContractRecord, error)
}

// Context is the low-level context provided to RequestDispatcher.StaticCall().
type Context interface {
	StaticContext
	VolatileAPI
	Set(key, value []byte) error
	Delete(key []byte) error
}

type Contract interface {
	Meta() (Meta, error)
	Init(ctx Context, req *Request) error
	Call(ctx Context, req *Request) (*Response, error)
	StaticCall(ctx StaticContext, req *Request) (*Response, error)
}
