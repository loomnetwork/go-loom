// +build evm

package erc721

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/loomnetwork/go-loom"
	"github.com/loomnetwork/go-loom/auth"
	"github.com/loomnetwork/go-loom/client"
)

type DAppChainERC721Contract struct {
	*client.MirroredTokenContract
}

func (c *DAppChainERC721Contract) OwnerOf(tokenID *big.Int) (loom.LocalAddress, error) {
	var result common.Address
	if err := c.StaticCallEVM("ownerOf", &result, tokenID); err != nil {
		return nil, err
	}
	addr, err := loom.LocalAddressFromHexString(result.Hex())
	if err != nil {
		return nil, err
	}
	return addr, nil
}

// Approve grants authorization to an entity to transfer the given token at a later time
func (c *DAppChainERC721Contract) Approve(identity *client.Identity, to loom.Address, tokenID *big.Int) error {
	toAddr := common.BytesToAddress(to.Local)
	return c.CallEVM("approve", identity.LoomSigner, toAddr, tokenID)
}

func (c *DAppChainERC721Contract) BalanceOf(identity *client.Identity) (uint64, error) {
	result := new(big.Int)
	addr := common.BytesToAddress(identity.LoomAddr.Local)
	if err := c.StaticCallEVM("balanceOf", &result, addr); err != nil {
		return 0, err
	}
	return result.Uint64(), nil
}

func (c *DAppChainERC721Contract) TransferFrom(from *client.Identity, to *client.Identity, tokenID *big.Int) error {
	fromAddr := common.BytesToAddress(from.LoomAddr.Local)
	toAddr := common.BytesToAddress(to.LoomAddr.Local)
	return c.CallEVM("transferFrom", from.LoomSigner, fromAddr, toAddr, tokenID)
}

func (c *DAppChainERC721Contract) MintTo(identity *client.Identity, to loom.Address, tokenID *big.Int) error {
	toAddr := common.BytesToAddress(to.Local)
	return c.CallEVM("mintTo", identity.LoomSigner, toAddr, tokenID)
}

/**
  Connectors
*/

func DeployERC721ToDAppChain(loomClient *client.DAppChainRPCClient, contractName string,
	gatewayAddr loom.Address, creator auth.Signer) (*DAppChainERC721Contract, error) {
	contract, err := client.DeployTokenToDAppChain(loomClient, contractName, "erc721", gatewayAddr, creator)
	if err != nil {
		return nil, err
	}
	return &DAppChainERC721Contract{contract}, nil
}

func ConnectERC721ToDAppChain(
	loomClient *client.DAppChainRPCClient, contractName string,
) (*DAppChainERC721Contract, error) {
	contract, err := client.ConnectToMirroredTokenContract(loomClient, contractName, "erc721")
	if err != nil {
		return nil, err
	}
	return &DAppChainERC721Contract{contract}, nil
}
