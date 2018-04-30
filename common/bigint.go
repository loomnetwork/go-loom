package common

import (
	"fmt"
	"math/big"
)

// BigUint is a simple wrapper on bigint to support marshaling to protobufs
type BigUint struct {
	*big.Int
}

// Unmarshal unmarshals protobuf data
func (b *BigUint) Unmarshal(by []byte) error {
	fmt.Printf("unmarshal bytes-%v\n", by)
	b.Int = big.NewInt(0).SetBytes(by)
	return nil
}

// Marshal converts to a byte buffer for protobufs
func (b *BigUint) Marshal() ([]byte, error) {
	return b.Int.Bytes(), nil
}

// Cmp compares x and y and returns:
//
//   -1 if x <  y
//    0 if x == y
//   +1 if x >  y
//
func (b *BigUint) Cmp(i *BigUint) int {
	return b.Int.Cmp(i.Int)
}

// Sub sets z to the difference x-y and returns z.
func (b *BigUint) Sub(x, y *BigUint) *BigUint {
	return NewBigUint(b.Int.Sub(x.Int, y.Int))
}

// Add sets z to the sum x+y and returns z.
func (b *BigUint) Add(x, y *BigUint) *BigUint {
	return NewBigUint(b.Int.Add(x.Int, y.Int))
}

// Mul sets z to the product x*y and returns z.
func (b *BigUint) Mul(x, y *BigUint) *BigUint {
	return NewBigUint(b.Int.Mul(x.Int, y.Int))
}

// NewBigUint creates a biguint from a bigint
func NewBigUint(i *big.Int) (b *BigUint) {
	return &BigUint{i}
}

// NewBigUintFromInt creates a biguint from a int64
func NewBigUintFromInt(i int64) (b *BigUint) {
	return &BigUint{big.NewInt(i)}
}
