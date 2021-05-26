// +build evm

package native_coin

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/loomnetwork/go-loom/client"
)

// MainnetEthCoin is a client-side binding for the ether on ethereum
type MainnetEthCoin struct {
	ethClient *ethclient.Client
	TxTimeout time.Duration
}

func (ec *MainnetEthCoin) BalanceOf(address common.Address) (*big.Int, error) {
	return ec.ethClient.BalanceAt(context.Background(), address, nil)
}

func (ec *MainnetEthCoin) Transfer(caller *client.Identity, from common.Address, to common.Address, amount *big.Int) error {
	nonce, err := ec.ethClient.PendingNonceAt(context.Background(), caller.MainnetAddr)
	if err != nil {
		return err
	}

	opts := client.DefaultTransactOptsForIdentity(caller)

	tx := types.NewTransaction(nonce, to, amount, opts.GasLimit, opts.GasPrice, []byte{})

	chainID, err := ec.ethClient.NetworkID(context.Background())
	if err != nil {
		return err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), caller.MainnetPrivKey)
	if err != nil {
		return err
	}
	return ec.ethClient.SendTransaction(context.Background(), signedTx)
}

func ConnectToMainnetEthCoin(ethClient *ethclient.Client) (*MainnetEthCoin, error) {
	return &MainnetEthCoin{
		ethClient: ethClient,
	}, nil
}
