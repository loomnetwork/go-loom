// +build evm

package timelock

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/loomnetwork/go-loom/client"
)

type MainnetTokenTimelockClient struct {
	ethClient *ethclient.Client
	contract  *TokenTimelock
	// Mainnet TimelockFactory contract address
	Address   common.Address
	TxTimeout time.Duration
}

func (c *MainnetTokenTimelockClient) Release(caller *client.Identity) error {
	tx, err := c.contract.Release(client.DefaultTransactOptsForIdentity(caller))
	if err != nil {
		return err
	}
	return client.WaitForTxConfirmation(context.TODO(), c.ethClient, tx, c.TxTimeout)
}

func ConnectToMainnetTimelock(ethClient *ethclient.Client, address string) (*MainnetTokenTimelockClient, error) {
	contractAddr := common.HexToAddress(address)
	contract, err := NewTokenTimelock(contractAddr, ethClient)
	if err != nil {
		return nil, err
	}
	return &MainnetTokenTimelockClient{
		ethClient: ethClient,
		contract:  contract,
		Address:   contractAddr,
	}, nil
}
