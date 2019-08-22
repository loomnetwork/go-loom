// +build evm

package erc20

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/loomnetwork/go-loom/client"
)

type MainnetERC20Client struct {
	contract  *ERC20
	ethClient *ethclient.Client

	TxTimeout time.Duration
	Address   common.Address
}

func (c *MainnetERC20Client) BalanceOf(addr common.Address) (*big.Int, error) {
	bal, err := c.contract.BalanceOf(nil, addr)
	if err != nil {
		return nil, err
	}
	return bal, nil
}

func (c *MainnetERC20Client) TransferFrom(caller *client.Identity, from common.Address, to common.Address, amount *big.Int) error {
	tx, err := c.contract.TransferFrom(client.DefaultTransactOptsForIdentity(caller), from, to, amount)
	if err != nil {
		return err
	}
	return client.WaitForTxConfirmation(context.TODO(), c.ethClient, tx, c.TxTimeout)
}

func (c *MainnetERC20Client) Approve(caller *client.Identity, spender common.Address, amount *big.Int) error {
	tx, err := c.contract.Approve(client.DefaultTransactOptsForIdentity(caller), spender, amount)
	if err != nil {
		return err
	}
	return client.WaitForTxConfirmation(context.TODO(), c.ethClient, tx, c.TxTimeout)
}

func (c *MainnetERC20Client) Transfer(caller *client.Identity, to common.Address, amount *big.Int) error {
	tx, err := c.contract.Transfer(client.DefaultTransactOptsForIdentity(caller), to, amount)
	if err != nil {
		return err
	}
	return client.WaitForTxConfirmation(context.TODO(), c.ethClient, tx, c.TxTimeout)
}

func (c *MainnetERC20Client) MintTo(caller *client.Identity, to common.Address, amount *big.Int) error {
	tx, err := c.contract.MintTo(client.DefaultTransactOptsForIdentity(caller), to, amount)
	if err != nil {
		return err
	}
	return client.WaitForTxConfirmation(context.TODO(), c.ethClient, tx, c.TxTimeout)
}

func ConnectToMainnetERC20(ethClient *ethclient.Client, contractAddr string) (*MainnetERC20Client, error) {
	contractAddress := common.HexToAddress(contractAddr)
	contract, err := NewERC20(contractAddress, ethClient)
	if err != nil {
		return nil, err
	}
	return &MainnetERC20Client{
		contract:  contract,
		ethClient: ethClient,
		Address:   contractAddress,
	}, nil
}
