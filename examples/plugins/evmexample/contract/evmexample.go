// +build evm

package main

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/loomnetwork/go-loom/examples/plugins/evmexample/types"
	"github.com/loomnetwork/go-loom/plugin"
	"github.com/loomnetwork/go-loom/plugin/contractpb"
	"math/big"
	"strconv"
	"strings"
)

var (
	SimpleStoreABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"set\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"
)

type EvmExample struct {
}

func (c *EvmExample) Meta() (plugin.Meta, error) {
	return plugin.Meta{
		Name:    "EvmExample",
		Version: "1.0.0",
	}, nil
}

func (c *EvmExample) SetValue(ctx contractpb.Context, value *types.WrapValue) error {
	ssAddr, err := ctx.Resolve("SimpleStore")
	if err != nil {
		return err
	}
	abiSS, err := abi.JSON(strings.NewReader(SimpleStoreABI))
	if err != nil {
		return err
	}
	input, err := abiSS.Pack("set", big.NewInt(value.Value))
	if err != nil {
		return err
	}
	evmOut := []byte{}
	err = contractpb.CallEVM(ctx, ssAddr, input, &evmOut)
	return err
}

func (c *EvmExample) GetValue(ctx contractpb.Context, req *types.Dummy) (*types.WrapValue, error) {
	ssAddr, err := ctx.Resolve("SimpleStore")
	if err != nil {
		return nil, err
	}
	abiSS, err := abi.JSON(strings.NewReader(SimpleStoreABI))
	input, err := abiSS.Pack("get")

	evmOut := []byte{}
	err = contractpb.CallEVM(ctx, ssAddr, input, &evmOut)
	if err != nil {
		return nil, err
	}
	value, err := strconv.ParseInt(common.Bytes2Hex(evmOut), 16, 64)
	if err != nil {
		return nil, err
	}
	return &types.WrapValue{
		Value: value,
	}, err
}

var Contract = contractpb.MakePluginContract(&EvmExample{})

func main() {
	plugin.Serve(Contract)
}
