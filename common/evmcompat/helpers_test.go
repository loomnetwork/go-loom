// +build evm

package evmcompat

import (
	"testing"
)

/*
describe('solidity tight packing multiple arguments', function () {
  it('should equal', function () {
    var a = abi.solidityPack(
      [ 'bytes32', 'uint32', 'uint32', 'uint32', 'uint32' ],
      [ new Buffer('123456', 'hex'), 6, 7, 8, 9 ]
    )
    var b = '123456000000000000000000000000000000000000000000000000000000000000000006000000070000000800000009'
    assert.equal(a.toString('hex'), b.toString('hex'))
  })
})
*/

// uint32 8
// uint256 64
// func TestSolidityPackedBytes(t *testing.T) {
// 	want := "0000000843989fb883ba8111221e8912389753847589383700000007"

// 	pairs := []*Pair{
// 		&Pair{"uint32", "8"},
// 		&Pair{"Address", "43989fb883ba8111221e89123897538475893837"},
// 		&Pair{"uint32", "7"},
// 	}

// 	g, err := SolidityPackedBytes(pairs)
// 	if err != nil {
// 		t.Errorf("TestSolidityPackedBytes failed got error %q", err)
// 	}
// 	got := hex.EncodeToString(g)

// 	if got != want {
// 		t.Errorf("TestSolidityPackedBytes failed got %q, want %q", got, want)
// 	}

// 	g = ssha.ConcatByteSlices(
// 		ssha.Uint32(8),
// 		ssha.Address("43989fb883ba8111221e89123897538475893837"),
// 		ssha.Uint32(7),
// 	)
// 	got = hex.EncodeToString(g)

// 	if got != want {
// 		t.Errorf("TestSolidityPackedBytes failed got %q, want %q", got, want)
// 	}

// 	wantsha3 := "5611aae8648e01a2e4721917fd1706014b8f4d387928e3cad536be41e5af4f77"

// 	g2, err := SoliditySHA3(pairs)
// 	if err != nil {
// 		t.Errorf("TestSolidityPackedBytes failed got error %q", err)
// 	}
// 	gotsha3 := hex.EncodeToString(g2)

// 	if gotsha3 != wantsha3 {
// 		t.Errorf("TestSolidityPackedBytes failed got %q, want %q", gotsha3, wantsha3)
// 	}

// 	g2 = ssha.SoliditySHA3(g)
// 	gotsha3 = hex.EncodeToString(g2)
// 	if gotsha3 != wantsha3 {
// 		t.Errorf("TestSolidityPackedBytes failed got %q, want %q", gotsha3, wantsha3)
// 	}
// }

// func TestSoliditySha3(t *testing.T) {
// 	want := "43989fb883ba8111221e8912389753847589383700000000000000000000000000000000000000002710564fe203"

// 	pairs := []*Pair{
// 		&Pair{"Address", "43989fb883ba8111221e89123897538475893837"},
// 		&Pair{"Address", "0000000000000000000000000000000000000000"},
// 		&Pair{"uint16", "10000"},
// 		&Pair{"uint32", "1448075779"},
// 	}

// 	g, err := SolidityPackedBytes(pairs)
// 	if err != nil {
// 		t.Errorf("TestSoliditySha3 failed got error %q", err)
// 	}
// 	got := hex.EncodeToString(g)

// 	if got != want {
// 		t.Errorf("TestSoliditySha3 failed got %q, want %q", got, want)
// 	}

// 	slices := [][]byte{
// 		ssha.Address("43989fb883ba8111221e89123897538475893837"),
// 		ssha.Address("0000000000000000000000000000000000000000"),
// 		ssha.Uint16(uint16(10000)),
// 		ssha.Uint32(uint32(1448075779)),
// 	}
// 	g = ssha.ConcatByteSlices(slices...)
// 	got = hex.EncodeToString(g)

// 	if got != want {
// 		t.Errorf("TestSoliditySha3 failed got %q, want %q", got, want)
// 	}

// 	wantsha3 := "7221df1d75e4baccbccd8a1fb33dbc5fca5f3c543e4acbb37c1b9edf990d3e1e"

// 	g2, err := SoliditySHA3(pairs)
// 	if err != nil {
// 		t.Errorf("TestSoliditySha3 failed got error %q", err)
// 	}
// 	gotsha3 := hex.EncodeToString(g2)

// 	if gotsha3 != wantsha3 {
// 		t.Errorf("TestSoliditySha3 failed got %q, want %q", gotsha3, wantsha3)
// 	}

// 	g2 = ssha.SoliditySHA3(slices...)
// 	gotsha3 = hex.EncodeToString(g2)

// 	if gotsha3 != wantsha3 {
// 		t.Errorf("TestSoliditySha3 failed got %q, want %q", gotsha3, wantsha3)
// 	}
// }

// func TestSolidityPackedBytesTypeAddress(t *testing.T) {
// 	want := "43989fb883ba8111221e89123897538475893837"
// 	pairs := []*Pair{
// 		&Pair{"Address", "43989fb883ba8111221e89123897538475893837"},
// 	}

// 	g, err := SolidityPackedBytes(pairs)
// 	if err != nil {
// 		t.Errorf("TestSolidityPackedBytesTypeAddress failed got error %q", err)
// 	}
// 	got := hex.EncodeToString([]byte(g))

// 	if got != want {
// 		t.Errorf("TestSolidityPackedBytesTypeAddress failed got %q, want %q", got, want)
// 	}

// 	g = ssha.Address("43989fb883ba8111221e89123897538475893837")
// 	got = hex.EncodeToString(g)

// 	if got != want {
// 		t.Errorf("TestSolidityPackedBytesTypeAddress failed got %q, want %q", got, want)
// 	}
// }

// func TestSolidityPackedUint16(t *testing.T) {
// 	want := "002a"

// 	pairs := []*Pair{&Pair{"uint16", "42"}}

// 	g, err := SolidityPackedBytes(pairs)
// 	if err != nil {
// 		t.Errorf("TestSolidityPackedBytes failed got error %q", err)
// 	}
// 	got := hex.EncodeToString([]byte(g))

// 	if got != want {
// 		t.Errorf("TestSolidityPackedBytes failed got %q, want %q", got, want)
// 	}

// 	g = ssha.Uint16(uint16(42))
// 	got = hex.EncodeToString(g)

// 	if got != want {
// 		t.Errorf("TestSolidityPackedBytes failed got %q, want %q", got, want)
// 	}
// }

// func TestSolidityPackedUint256(t *testing.T) {
// 	want := "000000000000000000000000000000000000000000000000000000000000002a"

// 	pairs := []*Pair{&Pair{"uint256", "42"}}

// 	g, err := SolidityPackedBytes(pairs)
// 	if err != nil {
// 		t.Errorf("TestSolidityPackedBytes failed got error %q", err)
// 	}
// 	got := hex.EncodeToString([]byte(g))

// 	if got != want {
// 		t.Errorf("TestSolidityPackedBytes failed got %q, want %q", got, want)
// 	}

// 	g = ssha.Uint256(big.NewInt(42))
// 	got = hex.EncodeToString(g)

// 	if got != want {
// 		t.Errorf("TestSolidityPackedBytes failed got %q, want %q", got, want)
// 	}
// }

// func TestSoliditySha3With256(t *testing.T) {
// 	want := "9f022fbbf24efa13621bbc6c2fc2f3b1f742d3320123acde9a25a9b5e25d81a9"

// 	pairs := []*Pair{
// 		&Pair{"uint256", "42"},
// 		&Pair{"Address", "32be343b94f860124dc4fee278fdcbd38c102d88"},
// 		&Pair{"Address", "74ff65739a88fdaf9675ff33405f760b53832ad7"},
// 		&Pair{"uint256", "52"},
// 	}

// 	g, err := SolidityPackedBytes(pairs)
// 	if err != nil {
// 		t.Errorf("TestSoliditySha3With256 failed got error %q", err)
// 	}
// 	if len(g) != 104 {
// 		t.Errorf("length unexpected")
// 	}

// 	d := sha3.NewKeccak256()
// 	d.Write(g)
// 	hash := d.Sum(nil)
// 	if hex.EncodeToString(hash) != want {
// 		t.Errorf("hashes don't match")
// 	}

// 	g = ssha.ConcatByteSlices(
// 		ssha.Uint256(big.NewInt(42)),
// 		ssha.Address("32be343b94f860124dc4fee278fdcbd38c102d88"),
// 		ssha.Address("74ff65739a88fdaf9675ff33405f760b53832ad7"),
// 		ssha.Uint256(big.NewInt(52)),
// 	)
// 	if len(g) != 104 {
// 		t.Errorf("length unexpected")
// 	}

// 	d = sha3.NewKeccak256()
// 	d.Write(g)
// 	hash = d.Sum(nil)
// 	if hex.EncodeToString(hash) != want {
// 		t.Errorf("hashes don't match")
// 	}
// }

// func TestAnotherSoliditySha3With256(t *testing.T) {
// 	want := "00000000000000000000000000000000000000000000000000000000564fe20343989fb883ba8111221e8912389753847589383766989fb883ba8111221e8912389753847589386700000000000000000000000000000000000000000000000000000000564fe203"

// 	pairs := []*Pair{
// 		&Pair{"uint256", "1448075779"},
// 		&Pair{"Address", "43989fb883ba8111221e89123897538475893837"},
// 		&Pair{"Address", "66989fb883ba8111221e89123897538475893867"},
// 		&Pair{"uint256", "1448075779"},
// 	}

// 	g, err := SolidityPackedBytes(pairs)
// 	if err != nil {
// 		t.Errorf("TestSoliditySha3With256 failed got error %q", err)
// 	}
// 	if len(g) != 104 {
// 		t.Errorf("length unexpected")
// 	}
// 	got := hex.EncodeToString([]byte(g))
// 	if want != got {
// 		t.Errorf("hashes don't match -\n%s\n%s", want, g)
// 	}

// 	g = ssha.ConcatByteSlices(
// 		ssha.Uint256(big.NewInt(1448075779)),
// 		ssha.Address("43989fb883ba8111221e89123897538475893837"),
// 		ssha.Address("66989fb883ba8111221e89123897538475893867"),
// 		ssha.Uint256(big.NewInt(1448075779)),
// 	)
// 	if len(g) != 104 {
// 		t.Errorf("length unexpected")
// 	}
// 	got = hex.EncodeToString(g)
// 	if want != got {
// 		t.Errorf("hashes don't match -\n%s\n%s", want, g)
// 	}
// }

// func TestAnotherSoliditySha3WithUnit64(t *testing.T) {
// 	want := "920ae4155769cd69c30626f054134b5f003772473f57f84837402df6d166e663"

// 	pairs := []*Pair{
// 		&Pair{"uint32", "5"},
// 	}

// 	g, err := SoliditySHA3(pairs)
// 	if err != nil {
// 		t.Errorf("TestSoliditySha3With256 failed got error %q", err)
// 	}
// 	got := hex.EncodeToString([]byte(g))
// 	if want != got {
// 		t.Errorf("hashes don't match -\n%s\n%s", want, got)
// 	}

// 	g = ssha.SoliditySHA3(ssha.Uint32(uint32(5)))
// 	got = hex.EncodeToString(g)
// 	if want != got {
// 		t.Errorf("hashes don't match -\n%s\n%s", want, got)
// 	}

// 	want2 := "fe07a98784cd1850eae35ede546d7028e6bf9569108995fc410868db775e5e6a"

// 	pairs2 := []*Pair{
// 		&Pair{"uint64", "5"},
// 	}

// 	g2, err := SoliditySHA3(pairs2)
// 	if err != nil {
// 		t.Errorf("TestSoliditySha3With256 failed got error %q", err)
// 	}
// 	got2 := hex.EncodeToString([]byte(g2))
// 	if want2 != got2 {
// 		t.Errorf("hashes don't match -\n%s\n%s", want2, got2)
// 	}

// 	g2 = ssha.SoliditySHA3(ssha.Uint64(uint64(5)))
// 	got2 = hex.EncodeToString(g2)
// 	if want2 != got2 {
// 		t.Errorf("hashes don't match -\n%s\n%s", want2, got2)
// 	}
// }

func TestSolidityUnpackBytes(t *testing.T) {
	SolidityUnpackBytes("nope")
}
