// +build evm

package validator_manager

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/loomnetwork/go-loom/client"
)

type MainnetVMCClient struct {
	contract  *VMC
	ethClient *ethclient.Client

	TxTimeout time.Duration
	Address   common.Address
}

func (c *MainnetVMCClient) BalanceOf(owner common.Address) (*big.Int, error) {
	bal, err := c.contract.BalanceOf(nil, owner)
	if err != nil {
		return nil, err
	}
	return bal, nil
}

func (c *MainnetVMCClient) TokenOfOwnerByIndex(owner common.Address, index int) (*big.Int, error) {
	tokenID, err := c.contract.TokenOfOwnerByIndex(nil, owner, new(big.Int).SetInt64(int64(index)))
	if err != nil {
		return nil, err
	}
	return tokenID, nil
}

func (c *MainnetVMCClient) OwnerOf(tokenID *big.Int) (common.Address, error) {
	return c.contract.OwnerOf(nil, tokenID)
}

func (c *MainnetVMCClient) SafeTransferFrom(caller *client.Identity, from common.Address, to common.Address, tokenId *big.Int, amount *big.Int, data []byte) error {
	tx, err := c.contract.SafeTransferFrom(client.DefaultTransactOptsForIdentity(caller), from, to, tokenId, amount, data)
	if err != nil {
		return err
	}
	return client.WaitForTxConfirmation(context.TODO(), c.ethClient, tx, c.TxTimeout)
}

func (c *MainnetVMCClient) ToggleToken(caller *client.Identity, tokenContract common.Address) error {
	opts := bind.NewKeyedTransactor(validatorKey)
	ethNet := os.Getenv("ETHEREUM_NETWORK")
	if ethNet == "" || ethNet == "ganache" {
		// hack to get around Ganache hex-encoding bug, see client.DefaultTransactOptsForIdentity for info
		opts.GasPrice = big.NewInt(20000)
		opts.GasLimit = uint64(3141592)
	}
	tx, err := c.contract.ToggleToken(opts, tokenContract)
	if err != nil {
		return err
	}
	return client.WaitForTxConfirmation(context.TODO(), c.ethClient, tx, c.TxTimeout)
}

func ConnectToMainnetVMCClient(ethClient *ethclient.Client, contractAddr string) (*MainnetVMCClient, error) {
	contractAddress := common.HexToAddress(contractAddr)
	contract, err := NewVMC(contractAddress, ethClient)
	if err != nil {
		return nil, err
	}
	return &MainnetVMCClient{
		contract:  contract,
		ethClient: ethClient,
		Address:   contractAddress,
	}, nil
}
