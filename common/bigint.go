package common

import (
	"math/big"
)

// Wraps BigInt and gives protobuf unmarshaling
type BigUint = big.Int
