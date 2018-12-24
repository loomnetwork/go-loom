// +build evm

package timelock

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/loomnetwork/go-loom/client"
)

type MainnetTimelockFactoryClient struct {
	ethClient *ethclient.Client
	contract  *LoomTimelockFactory
	// Mainnet TimelockFactory contract address
	Address   common.Address
	TxTimeout time.Duration
}

func (c *MainnetTimelockFactoryClient) DeployTimelock(caller *client.Identity, beneficiary common.Address, validatorName string, pubKey string, amount *big.Int, duration *big.Int) (common.Address, error) {
	tx, err := c.contract.DeployTimeLock(client.DefaultTransactOptsForIdentity(caller), beneficiary, validatorName, pubKey, amount, duration)
	if err != nil {
		return common.HexToAddress("0x0"), err
	}
	err = client.WaitForTxConfirmation(context.TODO(), c.ethClient, tx, c.TxTimeout)
	if err != nil {
		return common.HexToAddress("0x0"), err
	}

	filterOpts := &bind.FilterOpts{
		Start: 0,
	}
	it, err := c.contract.FilterLoomTimeLockCreated(filterOpts)
	if err != nil {
		return common.HexToAddress("0x0"), nil
	}
	it.Next()
	timelockAddress := it.Event.TimelockContractAddress

	return timelockAddress, nil
}

func ConnectToMainnetTimelockFactory(ethClient *ethclient.Client, address string) (*MainnetTimelockFactoryClient, error) {
	contractAddr := common.HexToAddress(address)
	contract, err := NewLoomTimelockFactory(contractAddr, ethClient)
	if err != nil {
		return nil, err
	}
	return &MainnetTimelockFactoryClient{
		ethClient: ethClient,
		contract:  contract,
		Address:   contractAddr,
	}, nil
}
