package common

import (
	"bytes"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
)

type LocalAddress []byte

// Unmarshal unmarshals protobuf data
func (a LocalAddress) Unmarshal(b []byte) error {
	a = b
	return nil
}

// Marshal converts to a byte buffer for protobufs
func (a LocalAddress) Marshal() ([]byte, error) {
	return []byte(a.String()), nil
}

// From ethereum with finalized sha3
// Note: only works with addresses up to 256 bit
func (a LocalAddress) Hex() string {
	unchecksummed := hex.EncodeToString(a)
	fmt.Printf("unchecksummed-%s\n", unchecksummed)
	sha := sha3.New256()
	sha.Write([]byte(unchecksummed))
	hash := sha.Sum(nil)

	result := []byte(unchecksummed)
	for i := 0; i < len(result); i++ {
		hashByte := hash[i/2]
		if i%2 == 0 {
			hashByte = hashByte >> 4
		} else {
			hashByte &= 0xf
		}
		if result[i] > '9' && hashByte > 7 {
			result[i] -= 32
		}
	}
	return string(result)
}

func (a LocalAddress) String() string {
	return "0x" + a.Hex()
}

func (a LocalAddress) Compare(other LocalAddress) int {
	return bytes.Compare([]byte(a), []byte(other))
}

func LocalAddressFromPublicKey(pubKey []byte) LocalAddress {
	hasher := ripemd160.New()
	hasher.Write(pubKey[:]) // does not error
	return LocalAddress(hasher.Sum(nil))
}
