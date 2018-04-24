package loomplugin

import (
	"time"

	"github.com/loomnetwork/loom-plugin/types"
)

type StaticAPI interface {
	StaticCall(addr Address, input []byte) ([]byte, error)
}

type VolatileAPI interface {
	Call(addr Address, input []byte) ([]byte, error)
}

type API interface {
	StaticAPI
	VolatileAPI
}

type StaticContext interface {
	StaticAPI
	Get(key []byte) []byte
	// Has checks if a key exists.
	Has(key []byte) bool
	Block() types.BlockHeader
	Now() time.Time
	Message() types.Message
	ContractAddress() Address
}

type Context interface {
	StaticContext
	VolatileAPI
	// Set sets the key. Panics on nil key.
	Set(key, value []byte)

	// Delete deletes the key. Panics on nil key.
	Delete(key []byte)
	Emit(event []byte)
}

type Contract interface {
	Meta() (types.ContractMeta, error)
	Init(ctx Context, req *types.Request) error
	Call(ctx Context, req *types.Request) (*types.Response, error)
	StaticCall(ctx StaticContext, req *types.Request) (*types.Response, error)
}
