// +build evm

package client

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/loomnetwork/go-loom/common/evmcompat"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

var ErrTxFailed = errors.New("tx failed")
var ErrValnotFound = errors.New("validator not found")

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

// Parses a serialized signatures array into a list of (v,r,s) triples plus their corresponding validator indexes, in order. Refer to https://github.com/loomnetwork/transfer-gateway-v2/pull/83/files#diff-0aada7672d303fc5bbdeb252dc7ff653R208 for more information.
func ParseSigs(sigs []byte, hash []byte, validators []common.Address) ([]uint8, [][32]byte, [][32]byte, []*big.Int, error) {
	var vs []uint8
	var rs [][32]byte
	var ss [][32]byte
	var validatorIndexes []*big.Int


    // don't try splitting if 65 or 66
    var splitSigs [][]byte
    if len(sigs) == 65 {
        splitSigs = [][]byte{sigs}
    } else if len(sigs) == 66 {
        splitSigs = [][]byte{sigs[1:]} // remove the mode flag
    } else {
	    splitSigs = split(sigs, 65) // assume we receive unprefixed if more than 1 element
    }

	for _, sig := range splitSigs {
		validator, err := evmcompat.SolidityRecover(hash, sig)
		if err != nil {
			return nil, nil, nil, nil, err
		}

		var r [32]byte
		copy(r[:], sig[0:32])

		var s [32]byte
		copy(s[:], sig[32:64])

		v := uint8(sig[64])

		index, err := indexOfValidator(validator, validators)
		if err != nil {
			continue
		}

		vs = append(vs, v)
		rs = append(rs, r)
		ss = append(ss, s)
		validatorIndexes = append(validatorIndexes, index)
	}
	return vs, rs, ss, validatorIndexes, nil
}

func indexOfValidator(v common.Address, validators []common.Address) (*big.Int, error) {
	for key, value := range validators {
		if v.Hex() == value.Hex() {
			return big.NewInt(int64(key)), nil
		}
	}
	return nil, ErrValnotFound
}

func split(buf []byte, lim int) [][]byte {
	var chunk []byte
	chunks := make([][]byte, 0, len(buf)/lim+1)
	for len(buf) >= lim {
		chunk, buf = buf[:lim], buf[lim:]
		chunks = append(chunks, chunk)
	}
	if len(buf) > 0 {
		chunks = append(chunks, buf[:len(buf)])
	}
	return chunks
}
