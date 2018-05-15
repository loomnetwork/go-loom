// +build evm

package main

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/loomnetwork/go-loom/examples/plugins/evmexample/types"
	"github.com/loomnetwork/go-loom/plugin"
	"github.com/loomnetwork/go-loom/plugin/contractpb"
	"io/ioutil"
	"math/big"
	"strconv"
	"strings"
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
	simpleStoreAddr, err := ctx.Resolve("SimpleStore")
	if err != nil {
		return err
	}
	simpleStoreData, err := ioutil.ReadFile("SimpleStore.abi")
	if err != nil {
		return err
	}
	abiSimpleStore, err := abi.JSON(strings.NewReader(string(simpleStoreData)))
	if err != nil {
		return err
	}
	input, err := abiSimpleStore.Pack("set", big.NewInt(value.Value))
	if err != nil {
		return err
	}
	evmOut := []byte{}
	err = contractpb.CallEVM(ctx, simpleStoreAddr, input, &evmOut)
	return err
}

func (c *EvmExample) GetValue(ctx contractpb.Context, req *types.Dummy) (*types.WrapValue, error) {
	simpleStoreAddr, err := ctx.Resolve("SimpleStore")
	if err != nil {
		return nil, err
	}
	simpleStoreData, err := ioutil.ReadFile("SimpleStore.abi")
	if err != nil {
		return nil, err
	}
	abiSimpleStore, err := abi.JSON(strings.NewReader(string(simpleStoreData)))
	if err != nil {
		return nil, err
	}
	input, err := abiSimpleStore.Pack("get")
	if err != nil {
		return nil, err
	}
	evmOut := []byte{}
	err = contractpb.StaticCallEVM(ctx, simpleStoreAddr, input, &evmOut)
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
