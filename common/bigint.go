package common

import (
	"math/big"
)

// BigUint is a simple wrapper on bigint to support marshaling to protobufs
type BigUInt struct {
	*big.Int
}

// Unmarshal unmarshals protobuf data
func (b *BigUInt) Unmarshal(by []byte) error {
	b.Int = big.NewInt(0).SetBytes(by)
	return nil
}

// Marshal converts to a byte buffer for protobufs
func (b *BigUInt) Marshal() ([]byte, error) {
	return b.Int.Bytes(), nil
}

// Cmp compares x and y and returns:
//
//   -1 if x <  y
//    0 if x == y
//   +1 if x >  y
//
func (b *BigUInt) Cmp(i *BigUInt) int {
	return b.Int.Cmp(i.Int)
}

// Sub sets z to the difference x-y and returns z.
func (b *BigUInt) Sub(x, y *BigUInt) *BigUInt {
	return &BigUInt{b.Int.Sub(x.Int, y.Int)}
}

// Add sets z to the sum x+y and returns z.
func (b *BigUInt) Add(x, y *BigUInt) *BigUInt {
	return &BigUInt{b.Int.Add(x.Int, y.Int)}
}

// Mul sets z to the product x*y and returns z.
func (b *BigUInt) Mul(x, y *BigUInt) *BigUInt {
	return &BigUInt{b.Int.Mul(x.Int, y.Int)}
}
