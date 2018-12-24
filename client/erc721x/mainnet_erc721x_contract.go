// +build evm

package erc721x

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/loomnetwork/go-loom/client"
)

type MainnetERC721XClient struct {
	contract  *ERC721X
	ethClient *ethclient.Client

	TxTimeout time.Duration
	Address   common.Address
}

func (c *MainnetERC721XClient) BalanceOf(owner common.Address) (*big.Int, error) {
	bal, err := c.contract.BalanceOf(nil, owner)
	if err != nil {
		return nil, err
	}
	return bal, nil
}

func (c *MainnetERC721XClient) TokenOfOwnerByIndex(owner common.Address, index int) (*big.Int, error) {
	tokenID, err := c.contract.TokenOfOwnerByIndex(nil, owner, new(big.Int).SetInt64(int64(index)))
	if err != nil {
		return nil, err
	}
	return tokenID, nil
}

func (c *MainnetERC721XClient) OwnerOf(tokenID *big.Int) (common.Address, error) {
	return c.contract.OwnerOf(nil, tokenID)
}

func (c *MainnetERC721XClient) SafeTransferFrom(caller *client.Identity, from common.Address, to common.Address, tokenId *big.Int, amount *big.Int, data []byte) error {
	tx, err := c.contract.SafeTransferFrom(client.DefaultTransactOptsForIdentity(caller), from, to, tokenId, amount, data)
	if err != nil {
		return err
	}
	return client.WaitForTxConfirmation(context.TODO(), c.ethClient, tx, c.TxTimeout)
}

func ConnectToMainnetERC721XClient(ethClient *ethclient.Client, contractAddr string) (*MainnetERC721XClient, error) {
	contractAddress := common.HexToAddress(contractAddr)
	contract, err := NewERC721X(contractAddress, ethClient)
	if err != nil {
		return nil, err
	}
	return &MainnetERC721XClient{
		contract:  contract,
		ethClient: ethClient,
		Address:   contractAddress,
	}, nil
}
