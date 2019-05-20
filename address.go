package loom

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/ripemd160"

	"github.com/loomnetwork/go-loom/common"
	"github.com/loomnetwork/go-loom/types"
	"github.com/loomnetwork/go-loom/util"
)

var (
	ErrInvalidAddress = errors.New("invalid address")
)

type LocalAddress = common.LocalAddress

func LocalAddressFromHexString(hexAddr string) (LocalAddress, error) {
	if !strings.HasPrefix(hexAddr, "0x") {
		return nil, errors.New("hexAddr string has no 0x prefix")
	}
	bytes, err := hex.DecodeString(hexAddr[2:])
	if err != nil {
		return nil, err
	}
	if len(bytes) != 20 {
		return nil, fmt.Errorf("invalid local address %v", bytes)
	}
	return LocalAddress(bytes), nil
}

func LocalAddressFromPublicKey(pubKey []byte) LocalAddress {
	hasher := ripemd160.New()
	hasher.Write(pubKey[:]) // does not error
	return LocalAddress(hasher.Sum(nil))
}

// recent versions of Tendermint > 0.22.0
func LocalAddressFromPublicKeyV2(pubKey []byte) LocalAddress {
	hash := sha256.Sum256(pubKey)
	return LocalAddress(hash[:20])
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
		Local:   a.Local,
	}
}

func UnmarshalAddressPB(pb *types.Address) Address {
	return Address{
		ChainID: pb.ChainId,
		Local:   pb.Local,
	}
}

func RootAddress(chainID string) Address {
	return Address{
		ChainID: chainID,
		Local:   make([]byte, 20),
	}
}

// ParseAddress parses an address generated from String()
func ParseAddress(s string) (Address, error) {
	parts := strings.SplitN(s, ":", 2)
	if len(parts) != 2 {
		return Address{}, ErrInvalidAddress
	}

	local, err := LocalAddressFromHexString(parts[1])
	if err != nil {
		return Address{}, ErrInvalidAddress
	}

	return Address{ChainID: parts[0], Local: local}, nil
}

func MustParseAddress(s string) Address {
	addr, err := ParseAddress(s)
	if err != nil {
		panic(err)
	}
	return addr
}
