package loomplugin

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"strings"

	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"

	"github.com/loomnetwork/go-loom/types"
	"github.com/loomnetwork/go-loom/util"
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

func LocalAddressFromPublicKey(pubKey []byte) LocalAddress {
	hasher := ripemd160.New()
	hasher.Write(pubKey[:]) // does not error
	return LocalAddress(hasher.Sum(nil))
}

type Address struct {
	ChainID string
	Local   LocalAddress
}

func (a Address) String() string {
	return fmt.Sprintf("%s:%s", a.ChainID, a.Local.String())
}

func (a Address) Bytes() []byte {
	return util.PrefixKey([]byte(a.ChainID), a.Local)
}

func (a Address) Compare(other Address) int {
	ret := strings.Compare(a.ChainID, other.ChainID)
	if ret == 0 {
		ret = a.Local.Compare(other.Local)
	}
	return ret
}

func (a Address) IsEmpty() bool {
	return a.ChainID == "" && len(a.Local) == 0
}

func (a Address) MarshalPB() *types.Address {
	return &types.Address{
		ChainId: a.ChainID,
		Local:   []byte(a.Local),
	}
}

func UnmarshalAddressPB(pb *types.Address) Address {
	return Address{
		ChainID: pb.ChainId,
		Local:   LocalAddress(pb.Local),
	}
}

func RootAddress(chainID string) Address {
	return Address{
		ChainID: chainID,
		Local:   make([]byte, 20, 20),
	}
}
