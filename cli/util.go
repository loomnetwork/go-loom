package cli

import (
	"encoding/base64"
	"encoding/hex"
	"strconv"
	"strings"

	loom "github.com/loomnetwork/go-loom"
	"github.com/loomnetwork/go-loom/client"
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

func ParseAddress(s string) (loom.Address, error) {
	addr, err := loom.ParseAddress(s)
	if err == nil {
		return addr, nil
	}

	b, err := ParseBytes(s)
	if err != nil {
		return loom.Address{}, nil
	}
	if len(b) != 20 {
		return loom.Address{}, loom.ErrInvalidAddress
	}

	return loom.Address{ChainID: TxFlags.ChainID, Local: loom.LocalAddress(b)}, nil
}

func ResolveAddress(s string) (loom.Address, error) {
	rpcClient := client.NewDAppChainRPCClient(TxFlags.ChainID, TxFlags.WriteURI, TxFlags.ReadURI)
	contractAddr, err := ParseAddress(s)
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
