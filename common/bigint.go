package common

import (
	"math/big"
)

var (
	zeroInt = big.NewInt(0)
)

// BigUInt is a simple wrapper on bigint to support marshaling to protobufs
type BigUInt struct {
	*big.Int
}

func (b *BigUInt) int() *big.Int {
	if b == nil {
		return nil
	}

	return b.Int
}

// Unmarshal unmarshals protobuf data
func (b *BigUInt) Unmarshal(by []byte) error {
	b.Int = big.NewInt(0).SetBytes(by)
	return nil
}

// Size returns the length of the protobuf data
func (b *BigUInt) Size() int {
	return len(b.Bytes())
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

// Div sets z to the quotient x/y for y != 0 and returns z.
// If y == 0, a division-by-zero run-time panic occurs.
// Div implements Euclidean division (unlike Go); see DivMod for more details.
func (b *BigUInt) Div(x, y *BigUInt) *BigUInt {
	return &BigUInt{b.Int.Div(x.Int, y.Int)}
}

// Exp sets z = x**y mod |m| (i.e. the sign of m is ignored), and returns z.
// If y <= 0, the result is 1 mod |m|; if m == nil or m == 0, z = x**y.
//
// Modular exponentation of inputs of a particular size is not a
// cryptographically constant-time operation.
func (b *BigUInt) Exp(x, y, m *BigUInt) *BigUInt {
	return &BigUInt{b.Int.Exp(x.Int, y.Int, m.int())}
}

func (b *BigUInt) Uint64() uint64 {
	return b.Int.Uint64()
}

func BigZero() *BigUInt {
	return &BigUInt{new(big.Int).Set(zeroInt)}
}

func IsZero(b BigUInt) bool {
	return b.Int.Cmp(zeroInt) == 0
}

func IsPositive(b BigUInt) bool {
	return b.Int.Cmp(zeroInt) > 0
}
