package common

import (
	"math/big"
)

// Wraps BigInt and gives protobuf unmarshaling
type BigUint struct {
	big.Int
}

func NewBigUint(b big.Int) BigUint {
	return BigUint{b}
}
