package plugin

import (
	"time"

	loom "github.com/loomnetwork/go-loom"
	cctypes "github.com/loomnetwork/go-loom/builtin/types/chainconfig"
	ptypes "github.com/loomnetwork/go-loom/plugin/types"
	"github.com/loomnetwork/go-loom/types"
)

type (
	Request            = ptypes.Request
	Response           = ptypes.Response
	Meta               = ptypes.ContractMeta
	EncodingType       = ptypes.EncodingType
	ContractMethodCall = ptypes.ContractMethodCall
	Code               = ptypes.PluginCode
)

var (
	EncodingType_JSON      = ptypes.EncodingType_JSON
	EncodingType_PROTOBUF3 = ptypes.EncodingType_PROTOBUF3
)

type Message struct {
	Sender loom.Address
}

type StaticAPI interface {
	StaticCall(addr loom.Address, input []byte) ([]byte, error)
	StaticCallEVM(addr loom.Address, input []byte) ([]byte, error)
	Resolve(name string) (loom.Address, error)
	EmitTopics(event []byte, topics ...string)
	Emit(event []byte)
}

type VolatileAPI interface {
	Call(addr loom.Address, input []byte) ([]byte, error)
	CallEVM(addr loom.Address, input []byte, value *loom.BigUInt) ([]byte, error)
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
	Get(key []byte) []byte
	Has(key []byte) bool
	Range(prefix []byte) RangeData
	Block() loom.BlockHeader
	Now() time.Time
	Message() Message
	GetEvmTxReceipt([]byte) (ptypes.EvmTxReceipt, error)
	ContractAddress() loom.Address
	FeatureEnabled(name string, defaultVal bool) bool
	Config() *cctypes.Config
	EnabledFeatures() []string
	Validators() []*types.Validator
	// ContractRecord retrieves the contract meta data stored in the Registry.
	// NOTE: This method requires Registry v2.
	ContractRecord(contractAddr loom.Address) (*ContractRecord, error)
}

// Context is the low-level context provided to RequestDispatcher.StaticCall().
type Context interface {
	StaticContext
	VolatileAPI
	Set(key, value []byte)
	Delete(key []byte)
}

type Contract interface {
	Meta() (Meta, error)
	Init(ctx Context, req *Request) error
	Call(ctx Context, req *Request) (*Response, error)
	StaticCall(ctx StaticContext, req *Request) (*Response, error)
}
