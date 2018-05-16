package loom

import (
	"math/big"

	"github.com/loomnetwork/go-loom/common"
	"github.com/loomnetwork/go-loom/types"
)

type (
	BigUInt     = common.BigUInt
	BlockHeader = types.BlockHeader
)

// NewBigUint creates a biguint from a bigint
func NewBigUInt(i *big.Int) *BigUInt {
	return &BigUInt{i}
}

// NewBigUintFromInt creates a biguint from a int64
func NewBigUIntFromInt(i int64) *BigUInt {
	return &BigUInt{big.NewInt(i)}
}
