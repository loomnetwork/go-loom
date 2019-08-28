// +build evm

package validator_manager

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/loomnetwork/go-loom/client"
	ssha "github.com/miguelmota/go-solidity-sha3"
)

type MainnetVMCClient struct {
	contract  *ValidatorManagerContract
	ethClient *ethclient.Client

	TxTimeout time.Duration
	Address   common.Address
}

func (c *MainnetVMCClient) Contract() *ValidatorManagerContract {
	return c.contract
}

func (c *MainnetVMCClient) GetValidators() ([]common.Address, error) {
	return c.contract.GetValidators(nil)
}



func (c *MainnetVMCClient) RotateValidators(caller *client.Identity, newValidators []common.Address, newPowers []uint64, sigs []byte) error {
	// Calculate the msg hash
	hash, err := c.calculateHash(newValidators, newPowers)
	if err != nil {
		return err
	}

	// Get validator list to sort the sigs properly
	validators, err := c.GetValidators()
	if err != nil {
		return err
	}
	// Break down sigs
	v, r, s, valIndexes, err := client.ParseSigs(sigs, hash, validators)
	if err != nil {
		return err
	}

	tx, err := c.contract.RotateValidators(client.DefaultTransactOptsForIdentity(caller), newValidators, newPowers, valIndexes, v, r, s)
	if err != nil {
		return err
	}
	return client.WaitForTxConfirmation(context.TODO(), c.ethClient, tx, c.TxTimeout)
}

func ConnectToMainnetVMCClient(ethClient *ethclient.Client, contractAddr string) (*MainnetVMCClient, error) {
	contractAddress := common.HexToAddress(contractAddr)
	contract, err := NewValidatorManagerContract(contractAddress, ethClient)
	if err != nil {
		return nil, err
	}
	return &MainnetVMCClient{
		contract:  contract,
		ethClient: ethClient,
		Address:   contractAddress,
	}, nil
}

func (c *MainnetVMCClient) calculateHash(newValidators []common.Address, newPowers []uint64) ([]byte, error) {
	nonce, err := c.contract.Nonce(nil)
	if err != nil {
		return nil, err
	}

	hash := ssha.SoliditySHA3(
		ssha.AddressArray(newValidators),
		ssha.Uint8Array(newPowers),
	)

	return ssha.SoliditySHA3(
		ssha.Address(c.Address),
		ssha.Uint256(nonce),
		hash,
	), nil
}
