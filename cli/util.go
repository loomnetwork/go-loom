package cli

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	loom "github.com/loomnetwork/go-loom"
	amtypes "github.com/loomnetwork/go-loom/builtin/types/address_mapper"
	"github.com/loomnetwork/go-loom/client"
	"github.com/pkg/errors"
)

func ParseBytes(s string) ([]byte, error) {
	if strings.HasPrefix(s, "0x") {
		return hex.DecodeString(s[2:])
	}

	b, err := hex.DecodeString(s)
	if err != nil {
		b, err = base64.StdEncoding.DecodeString(s)
	}

	return b, err
}

func parseAddress(s string) (loom.Address, error) {
	addr, err := loom.ParseAddress(s)
	if err == nil {
		return addr, nil
	}

	b, err := ParseBytes(s)
	if err != nil {
		return loom.Address{}, err
	}
	if len(b) != 20 {
		return loom.Address{}, loom.ErrInvalidAddress
	}

	return loom.Address{ChainID: TxFlags.ChainID, Local: loom.LocalAddress(b)}, nil
}

func ResolveAddress(s, chainID, URI string) (loom.Address, error) {
	rpcClient := client.NewDAppChainRPCClient(chainID, URI+"/rpc", URI+"/query")
	contractAddr, err := parseAddress(s)
	if err != nil {
		// if address invalid, try to resolve it using registry
		contractAddr, err = rpcClient.Resolve(s)
		if err != nil {
			return loom.Address{}, err
		}
	}

	return contractAddr, nil
}

func sciNot(m, n int64) *loom.BigUInt {
	ret := loom.NewBigUIntFromInt(10)
	ret.Exp(ret, loom.NewBigUIntFromInt(n), nil)
	ret.Mul(ret, loom.NewBigUIntFromInt(m))
	return ret
}

func ParseAmount(s string) (*loom.BigUInt, error) {
	// TODO: allow more precision
	val, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return nil, err
	}
	return sciNot(val, 18), nil
}

func getMappedAccount(mapper *client.Contract, account loom.Address) (loom.Address, error) {
	req := &amtypes.AddressMapperGetMappingRequest{
		From: account.MarshalPB(),
	}
	resp := &amtypes.AddressMapperGetMappingResponse{}
	_, err := mapper.StaticCall("GetMapping", req, account, resp)
	if err != nil {
		return loom.Address{}, err
	}
	return loom.UnmarshalAddressPB(resp.To), nil
}

func ParseAddress(address string, callFlags *ContractCallFlags) (loom.Address, error) {
	var addr loom.Address
	addr, err := parseAddress(address)
	if err != nil {
		return addr, errors.Wrap(err, "failed to parse address")
	}
	//Resolve address if chainID does not match prefix
	if addr.ChainID != callFlags.ChainID {
		rpcClient := client.NewDAppChainRPCClient(callFlags.ChainID, callFlags.URI+"/rpc",
			callFlags.URI+"/query")
		mapperAddr, err := rpcClient.Resolve("addressmapper")
		if err != nil {
			return addr, errors.Wrap(err, "failed to resolve DAppChain Address Mapper address")
		}
		mapper := client.NewContract(rpcClient, mapperAddr.Local)
		mappedAccount, err := getMappedAccount(mapper, addr)
		if err != nil {
			return addr, fmt.Errorf("No account information found for %v", addr)
		}
		addr = mappedAccount
	}
	return addr, nil
}
