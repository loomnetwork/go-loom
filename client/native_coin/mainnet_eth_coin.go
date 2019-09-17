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

type MainnetEthCoin struct {
	ethClient *ethclient.Client
	TxTimeout time.Duration
}

func (ec *MainnetEthCoin) BalanceOf(address common.Address) (*big.Int, error) {
	return ec.ethClient.BalanceAt(context.Background(), address, nil)
}

func (ec *MainnetEthCoin) Transfer(caller *client.Identity, from common.Address, to common.Address, amount *big.Int) error {
	opts := client.DefaultTransactOptsForIdentity(caller)

	nonce, err := ec.ethClient.PendingNonceAt(context.Background(), from)
	if err != nil {
		return err
	}

	tx := types.NewTransaction(nonce, to, amount, opts.GasLimit, opts.GasPrice, []byte{})

	chainID, err := ec.ethClient.NetworkID(context.Background())
	if err != nil {
		return err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), caller.MainnetPrivKey)
	if err != nil {
		return err
	}

	err = ec.ethClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return err
	}

	return nil
}

func ConnectToMainnetEthCoin(ethClient *ethclient.Client) (*MainnetEthCoin, error) {
	return &MainnetEthCoin{
		ethClient: ethClient,
	}, nil
}
