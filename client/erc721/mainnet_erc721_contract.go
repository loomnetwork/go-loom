// +build evm

package erc721

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/loomnetwork/go-loom/client"
)

type MainnetERC721Client struct {
	contract  *ERC721
	ethClient *ethclient.Client

	TxTimeout time.Duration
	Address   common.Address
}

func (c *MainnetERC721Client) BalanceOf(owner common.Address) (uint64, error) {
	bal, err := c.contract.BalanceOf(nil, owner)
	if err != nil {
		return 0, err
	}
	return bal.Uint64(), nil
}

func (c *MainnetERC721Client) TokenOfOwnerByIndex(owner common.Address, index int) (*big.Int, error) {
	tokenID, err := c.contract.TokenOfOwnerByIndex(nil, owner, new(big.Int).SetInt64(int64(index)))
	if err != nil {
		return nil, err
	}
	return tokenID, nil
}

func (c *MainnetERC721Client) SafeTransferFrom(caller *client.Identity, from common.Address, to common.Address, tokenId *big.Int, data []byte) error {
	tx, err := c.contract.SafeTransferFrom(client.DefaultTransactOptsForIdentity(caller), from, to, tokenId, data)
	if err != nil {
		return err
	}
	return client.WaitForTxConfirmation(context.TODO(), c.ethClient, tx, c.TxTimeout)
}

func (c *MainnetERC721Client) OwnerOf(tokenID *big.Int) (common.Address, error) {
	return c.contract.OwnerOf(nil, tokenID)
}

func (c *MainnetERC721Client) MintTo(caller *client.Identity, to common.Address, amount *big.Int) error {
	tx, err := c.contract.MintTo(client.DefaultTransactOptsForIdentity(caller), to, amount)
	if err != nil {
		return err
	}
	return client.WaitForTxConfirmation(context.TODO(), c.ethClient, tx, c.TxTimeout)
}

func ConnectToMainnetERC721(ethClient *ethclient.Client, contractAddr string) (*MainnetERC721Client, error) {
	contractAddress := common.HexToAddress(contractAddr)
	contract, err := NewERC721(contractAddress, ethClient)
	if err != nil {
		return nil, err
	}
	return &MainnetERC721Client{
		contract:  contract,
		ethClient: ethClient,
		Address:   contractAddress,
	}, nil
}
