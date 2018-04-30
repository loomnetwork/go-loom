package loomplugin

import (
	"fmt"
	"strings"

	"github.com/loomnetwork/go-loom/common"
	"github.com/loomnetwork/go-loom/types"
	"github.com/loomnetwork/go-loom/util"
)

type Address struct {
	ChainID string
	Local   common.LocalAddress
}

func (a Address) String() string {
	return fmt.Sprintf("%s:%s", a.ChainID, a.Local.String())
}

func (a Address) Bytes() []byte {
	return util.PrefixKey([]byte(a.ChainID), a.Local.Bytes())
}

func (a Address) Compare(other Address) int {
	ret := strings.Compare(a.ChainID, other.ChainID)
	if ret == 0 {
		ret = a.Local.Compare(other.Local)
	}
	return ret
}

func (a Address) IsEmpty() bool {
	return a.ChainID == "" && len(a.Local.Bytes()) == 0
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
		Local:   common.LocalAddress(pb.Local),
	}
}

func RootAddress(chainID string) Address {
	return Address{
		ChainID: chainID,
		Local:   common.LocalAddressFromBytes(make([]byte, 20, 20)),
	}
}
