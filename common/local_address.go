package common

import (
	"bytes"
	"encoding/hex"

	"golang.org/x/crypto/sha3"
)

type LocalAddress []byte

// From ethereum with finalized sha3
// Note: only works with addresses up to 256 bit
func (a LocalAddress) Hex() string {
	unchecksummed := hex.EncodeToString(a)
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

// Unmarshal unmarshals protobuf data
func (a *LocalAddress) Unmarshal(data []byte) error {
	if len(data) == 0 {
		a = nil
		return nil
	}
	id := LocalAddress(make([]byte, len(data)))
	copy(id, data)
	*a = id
	return nil
}

// Marshal converts to a byte buffer for protobufs
func (a LocalAddress) Marshal() ([]byte, error) {
	if len(a) == 0 {
		return nil, nil
	}
	return []byte(a), nil
}

// Size returns the number of bytes for deterministic marshaling
func (a LocalAddress) Size() int {
	return len(a)
}
