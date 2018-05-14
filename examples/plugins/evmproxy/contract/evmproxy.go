// +build evm

package main

import (
	"encoding/hex"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/loomnetwork/go-loom/examples/plugins/evmproxy/types"
	"github.com/loomnetwork/go-loom/plugin"
	"github.com/loomnetwork/go-loom/plugin/contractpb"
)

type EvmProxy struct {
}

func (c *EvmProxy) Meta() (plugin.Meta, error) {
	return plugin.Meta{
		Name:    "evmproxy",
		Version: "1.0.0",
	}, nil
}

func (c *EvmProxy) EthTransaction(ctx contractpb.Context, tx *types.EthTransaction) error {
	simpleStoreAddr, err := ctx.Resolve("SimpleStore")
	if err != nil {
		return err
	}

	input, err := hex.DecodeString(strings.TrimPrefix(tx.Data, "0x"))
	if err != nil {
		return err
	}

	evmOut := []byte{}
	err = contractpb.CallEVM(ctx, simpleStoreAddr, input, &evmOut)

	return err
}

func (c *EvmProxy) EthCall(ctx contractpb.Context, tx *types.EthCall) (*types.EthCallResult, error) {
	simpleStoreAddr, err := ctx.Resolve("SimpleStore")
	if err != nil {
		return nil, err
	}

	input, err := hex.DecodeString(strings.TrimPrefix(tx.Data, "0x"))
	if err != nil {
		return nil, err
	}

	evmOut := []byte{}
	err = contractpb.CallEVM(ctx, simpleStoreAddr, input, &evmOut)
	if err != nil {
		return nil, err
	}

	value := common.Bytes2Hex(evmOut)

	return &types.EthCallResult{
		Data: value,
	}, err
}

var Contract = contractpb.MakePluginContract(&EvmProxy{})

func main() {
	plugin.Serve(Contract)
}
