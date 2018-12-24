// +build evm

package client

import (
	"io/ioutil"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

func LoadDAppChainContractABI(contractName string) (*abi.ABI, error) {
	_, filename, _, _ := runtime.Caller(0)
	abiPath := filepath.Join(filepath.Dir(filename), contractName+".abi")
	abiBytes, err := ioutil.ReadFile(abiPath)
	if err != nil {
		return nil, err
	}
	contractABI, err := abi.JSON(strings.NewReader(string(abiBytes)))
	if err != nil {
		return nil, err
	}
	return &contractABI, nil
}

func LoadDAppChainContractCode(contractName string) ([]byte, error) {
	_, filename, _, _ := runtime.Caller(0)
	binPath := filepath.Join(filepath.Dir(filename), contractName+".bin")
	hexByteCode, err := ioutil.ReadFile(binPath)
	if err != nil {
		return nil, err
	}
	return common.FromHex(string(hexByteCode)), nil
}
