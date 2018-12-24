// +build evm

package evm_eth

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/loomnetwork/go-loom"
	"github.com/loomnetwork/go-loom/auth"
	"github.com/loomnetwork/go-loom/client"
	"strings"
)

// DAppChainEthEvmTestContract is a client-side binding for the EthCoinIntegrationTest.sol contract
// that's deployed to the Loom EVM.
type DAppChainEthEvmTestContract struct {
	*client.MirroredTokenContract
}

func (c *DAppChainEthEvmTestContract) Deposit(caller *client.Identity, amount *big.Int) error {
	// TODO: doesn't work yet because there's no way to pass the msg.value via EvmContract, need to expand API
	return c.CallEVM("deposit", caller.LoomSigner)
}

func (c *DAppChainEthEvmTestContract) Withdraw(caller *client.Identity, amount *big.Int) error {
	return c.CallEVM("withdraw", caller.LoomSigner, amount)
}

func (c *DAppChainEthEvmTestContract) WithdrawThenFail(caller *client.Identity, amount *big.Int) error {
	return c.CallEVM("withdrawThenFail", caller.LoomSigner, amount)
}

func (c *DAppChainEthEvmTestContract) Transfer(caller *client.Identity, recipient loom.Address, amount *big.Int) error {
	recipientAddr := common.BytesToAddress(recipient.Local)
	// TODO: doesn't work yet because there's no way to pass the msg.value via EvmContract, need to expand API
	return c.CallEVM("transfer", caller.LoomSigner, recipientAddr, amount)
}

func (c *DAppChainEthEvmTestContract) Balance(account loom.Address) (*big.Int, error) {
	var result *big.Int
	if err := c.StaticCallEVM("getBalance", &result, common.BytesToAddress(account.Local)); err != nil {
		return common.Big0, err
	}
	return result, nil
}

/** Connectors

 */

func DeployEthEvmTestContractToDAppChain(
	loomClient *client.DAppChainRPCClient,
	creator auth.Signer,
) (*DAppChainEthEvmTestContract, error) {
	contractName := "EthCoinIntegrationTest"
	contractPath := strings.Join([]string{"evm_eth", contractName}, "/")
	contractABI, err := client.LoadDAppChainContractABI(contractPath)
	if err != nil {
		return nil, err
	}
	byteCode, err := client.LoadDAppChainContractCode(contractPath)
	if err != nil {
		return nil, err
	}
	input, err := contractABI.Pack("")
	if err != nil {
		return nil, err
	}
	byteCode = append(byteCode, input...)
	contract, _, err := client.DeployContract(loomClient, byteCode, creator, contractName)
	if err != nil {
		return nil, err
	}

	c := &client.MirroredTokenContract{
		Contract:    contract,
		ContractABI: contractABI,
		ChainID:     loomClient.GetChainID(),
		Address:     contract.Address,
	}

	return &DAppChainEthEvmTestContract{c}, nil
}
