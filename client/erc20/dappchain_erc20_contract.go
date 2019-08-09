// +build evm

package erc20

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	loom "github.com/loomnetwork/go-loom"
	"github.com/loomnetwork/go-loom/auth"
	"github.com/loomnetwork/go-loom/client"
)

type DAppChainERC20Contract struct {
	*client.MirroredTokenContract
}

func (c *DAppChainERC20Contract) BalanceOf(caller *client.Identity) (*big.Int, error) {
	ownerAddr := common.BytesToAddress(caller.LoomAddr.Local)
	var result *big.Int
	if err := c.StaticCallEVM("balanceOf", &result, ownerAddr); err != nil {
		return nil, err
	}
	return result, nil
}

// Approve grants authorization to an entity to transfer the given tokens at a later time
func (c *DAppChainERC20Contract) Approve(identity *client.Identity, to loom.Address, amount *big.Int) error {
	toAddr := common.BytesToAddress(to.Local)
	return c.CallEVM("approve", identity.LoomSigner, toAddr, amount)
}

func (c *DAppChainERC20Contract) Transfer(identity *client.Identity, to loom.Address, amount *big.Int) error {
	toAddr := common.BytesToAddress(to.Local)
	return c.CallEVM("transfer", identity.LoomSigner, toAddr, amount)
}

func (c *DAppChainERC20Contract) MintTo(identity *client.Identity, to loom.Address, amount *big.Int) error {
	toAddr := common.BytesToAddress(to.Local)
	return c.CallEVM("mintTo", identity.LoomSigner, toAddr, amount)
}

/**
  Functions to connect / deploy the ERC20 contracts
*/

func DeployERC20ToDAppChain(loomClient *client.DAppChainRPCClient, contractName string,
	gatewayAddr loom.Address, creator auth.Signer) (*DAppChainERC20Contract, error) {
	contract, err := client.DeployTokenToDAppChain(loomClient, contractName, "erc20", gatewayAddr, creator)
	if err != nil {
		return nil, err
	}
	return &DAppChainERC20Contract{contract}, nil
}

func ConnectERC20ToDAppChain(
	loomClient *client.DAppChainRPCClient, contractName string,
) (*DAppChainERC20Contract, error) {
	contract, err := client.ConnectToMirroredTokenContract(loomClient, contractName, "erc20")
	if err != nil {
		return nil, err
	}
	return &DAppChainERC20Contract{contract}, nil
}

