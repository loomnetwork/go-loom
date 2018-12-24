// +build evm

package client

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

var ErrTxFailed = errors.New("tx failed")

func DefaultTransactOptsForIdentity(ident *Identity) *bind.TransactOpts {
	opts := bind.NewKeyedTransactor(ident.MainnetPrivKey)
	ethNet := os.Getenv("ETHEREUM_NETWORK")
	if ethNet == "" || ethNet == "ganache" {
		// If gas price isn't set explicitely then go-ethereum will attempt to query the suggested gas
		// price, unfortunatley ganache-cli v6.1.2 seems to encode the gas price in a format go-ethereum
		// can't decode correctly, so this error is returned whenver you attempt to call a contract:
		// failed to suggest gas price: json: cannot unmarshal hex number with leading zero digits into Go value of type *hexutil.Big
		//
		// Earlier versions of ganache-cli don't seem to exhibit this issue, but they're broken in other
		// ways (logs aren't hex-encoded correctly).
		opts.GasPrice = big.NewInt(20000)
		opts.GasLimit = uint64(3141592)
	}
	return opts
}

// waitForTxReceipt waits for a tx to be confirmed.
// It stops waiting if the context is canceled, or the tx hasn't been confirmed after the specified timeout.
func WaitForTxReceipt(ctx context.Context, ethClient *ethclient.Client, tx *types.Transaction, timeout time.Duration) (*types.Receipt, error) {
	if timeout == 0 {
		timeout = 5 * time.Minute
	}

	interval := 30 * time.Second

	timer := time.NewTimer(timeout)
	ticker := time.NewTicker(interval)

	defer timer.Stop()
	defer ticker.Stop()

	txHash := tx.Hash()
	for {
		receipt, err := ethClient.TransactionReceipt(ctx, txHash)
		if err != nil && err != ethereum.NotFound {
			return nil, errors.Wrap(err, "failed to retrieve tx receipt")
		}
		if err == nil {
			return receipt, nil
		}

		// Wait for the next round.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-timer.C:
			return nil, errors.New("timed out waiting for tx receipt")
		case <-ticker.C:
		}
	}
}

// waitForTxConfirmation waits for a tx to be confirmed as successful.
func WaitForTxConfirmation(ctx context.Context, ethClient *ethclient.Client, tx *types.Transaction, timeout time.Duration) error {
	r, err := WaitForTxReceipt(ctx, ethClient, tx, timeout)
	if err != nil {
		gasPrice := new(big.Int)
		if tx.GasPrice() != nil {
			gasPrice = tx.GasPrice()
		}
		cost := new(big.Int)
		if tx.Cost() != nil {
			cost = tx.Cost()
		}
		return errors.Wrap(err,
			fmt.Sprintf(
				"tx failed (gas: %v, gasPrice: %s, cost: %s)",
				tx.Gas(), gasPrice.String(), cost.String(),
			),
		)
	}
	if r.Status != types.ReceiptStatusSuccessful {
		return ErrTxFailed
	}
	return nil
}

// waitForTxConfirmationAndFee waits for a tx to be confirmed as successful, and returns the fee paid for the tx.
func WaitForTxConfirmationAndFee(ctx context.Context, ethClient *ethclient.Client, tx *types.Transaction, timeout time.Duration) (*big.Int, error) {
	r, err := WaitForTxReceipt(ctx, ethClient, tx, timeout)
	if err != nil {
		return nil, err
	}
	if r.Status != types.ReceiptStatusSuccessful {
		return nil, ErrTxFailed
	}
	return new(big.Int).Mul(tx.GasPrice(), big.NewInt(0).SetUint64(r.GasUsed)), nil
}
