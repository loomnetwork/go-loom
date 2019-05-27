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

// ParseAddress attempts to parse the given string into an address, if the resulting address doesn't
// have a chain ID the given chain ID will be used instead.
func ParseAddress(s, chainID string) (loom.Address, error) {
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

	return loom.Address{ChainID: chainID, Local: loom.LocalAddress(b)}, nil
}

// ResolveAddress attempts to parse the given string into an address, if that fails it assumes the
// string corresponds to a contract name and attempts to obtain the corresponding contract address.
func ResolveAddress(s, chainID, URI string) (loom.Address, error) {
	rpcClient := client.NewDAppChainRPCClient(chainID, URI+"/rpc", URI+"/query")
	contractAddr, err := ParseAddress(s, chainID)
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

// ResolveAccountAddress attempts to parse the given string into the address of a user account.
// If the chain ID on the parsed address doesn't match the chain ID specified in chainFlags then
// the address is resolved to an on-chain address via the address mapper.
func ResolveAccountAddress(address string, chainFlags *ContractCallFlags) (loom.Address, error) {
	addr, err := ParseAddress(address, chainFlags.ChainID)
	if err != nil {
		return addr, errors.Wrap(err, "failed to parse address")
	}
	// Resolve address if chainID doesn't match
	if addr.ChainID != chainFlags.ChainID {
		rpcClient := client.NewDAppChainRPCClient(chainFlags.ChainID, chainFlags.URI+"/rpc", chainFlags.URI+"/query")
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
