// +build evm

package client

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/loomnetwork/go-loom"
	"github.com/loomnetwork/go-loom/auth"
)

type MirroredTokenContract struct {
	Contract    *EvmContract
	ContractABI *abi.ABI
	ChainID     string

	Address loom.Address
}

func DeployTokenToDAppChain(loomClient *DAppChainRPCClient, contractName string, directory string,
	gatewayAddr loom.Address, creator auth.Signer) (*MirroredTokenContract, error) {
	contractPath := strings.Join([]string{directory, contractName}, "/")
	ContractABI, err := LoadDAppChainContractABI(contractPath)
	if err != nil {
		return nil, err
	}
	byteCode, err := LoadDAppChainContractCode(contractPath)
	if err != nil {
		return nil, err
	}
	// append constructor args to bytecode
	// TODO: DeployContract() really should take care of this itself
	input, err := ContractABI.Pack("", common.BytesToAddress(gatewayAddr.Local))
	if err != nil {
		return nil, err
	}
	byteCode = append(byteCode, input...)
	contract, _, err := DeployContract(loomClient, byteCode, creator, contractName)
	if err != nil {
		return nil, err
	}
	return &MirroredTokenContract{
		Contract:    contract,
		ContractABI: ContractABI,
		ChainID:     loomClient.GetChainID(),
		Address:     contract.Address,
	}, nil
}

func ConnectToMirroredTokenContract(
	loomClient *DAppChainRPCClient, contractName string, directory string,
) (*MirroredTokenContract, error) {
	contractPath := strings.Join([]string{directory, contractName}, "/")
	ContractABI, err := LoadDAppChainContractABI(contractPath)
	if err != nil {
		return nil, err
	}

	contractAddr, err := loomClient.Resolve(contractName)
	if err != nil {
		return nil, err
	}

	return &MirroredTokenContract{
		Contract:    NewEvmContract(loomClient, contractAddr.Local),
		ContractABI: ContractABI,
		ChainID:     loomClient.GetChainID(),
		Address:     contractAddr,
	}, nil
}

func (c *MirroredTokenContract) StaticCallEVM(method string, result interface{}, params ...interface{}) error {
	input, err := c.ContractABI.Pack(method, params...)
	if err != nil {
		return err
	}
	output, err := c.Contract.StaticCall(input, c.Address)
	if err != nil {
		return err
	}
	return c.ContractABI.Unpack(result, method, output)
}

func (c *MirroredTokenContract) CallEVM(method string, signer auth.Signer, params ...interface{}) error {
	input, err := c.ContractABI.Pack(method, params...)
	if err != nil {
		return err
	}
	_, err = c.Contract.Call(input, signer)
	return err
}
