package contract

import (
	"time"

	loom "github.com/loomnetwork/loom-plugin"
	"github.com/loomnetwork/loom-plugin/store"
)

type Meta struct {
	Name    string
	Version string
}

type StaticAPI interface {
	StaticCall(addr loom.Address, input []byte) ([]byte, error)
}

type VolatileAPI interface {
	Call(addr loom.Address, input []byte) ([]byte, error)
}

type Message struct {
	Sender loom.Address
}

// BlockHeader interface wraps an abci.Header
type BlockHeader interface {
	// TODO
}

type ReadOnlyState interface {
	store.KVReader
	Block() BlockHeader
}

type StaticContext interface {
	StaticAPI
	ReadOnlyState
	Now() time.Time
	Message() Message
	ContractAddress() loom.Address
}

type Context interface {
	StaticContext
	VolatileAPI
	store.KVWriter
	Emit(event []byte)
}

type Contract interface {
	Meta() Meta
	Init(ctx Context, input []byte) ([]byte, error)
	Call(ctx Context, input []byte) ([]byte, error)
	StaticCall(ctx StaticContext, input []byte) ([]byte, error)
}
