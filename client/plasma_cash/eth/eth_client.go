// +build evm

package eth

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	etypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	loom "github.com/loomnetwork/go-loom"
	pctypes "github.com/loomnetwork/go-loom/builtin/types/plasma_cash"
	"github.com/loomnetwork/go-loom/client/plasma_cash/eth/ethcontract"
	ltypes "github.com/loomnetwork/go-loom/types"
	"github.com/pkg/errors"
)

type EthPlasmaClientConfig struct {
	// URI of an Ethereum node
	EthereumURI string
	// Plasma contract address on Ethereum
	PlasmaHexAddress string
	// Private key that should be used to sign txs sent to Ethereum
	PrivateKey *ecdsa.PrivateKey
	// Override default gas computation when sending txs to Ethereum
	OverrideGas bool
	// How often Ethereum should be polled for mined txs (defaults to 10 secs).
	TxPollInterval time.Duration
	// Maximum amount of time to way for a tx to be mined by Ethereum (defaults to 5 mins).
	TxTimeout time.Duration
}

type EthPlasmaClient interface {
	Init() error
	CurrentPlasmaBlockNum() (*big.Int, error)
	LatestEthBlockNum() (uint64, error)
	// SubmitPlasmaBlock will submit a Plasma block to Ethereum and wait until the tx is confirmed.
	// The maximum wait time can be specified via the TxTimeout option in the client config.
	SubmitPlasmaBlock(plasmaBlockNum *big.Int, merkleRoot [32]byte) error

	FetchDeposits(startBlock, endBlock uint64) ([]*pctypes.PlasmaDepositEvent, error)
	FetchCoinReset(startBlock, endBlock uint64) ([]*pctypes.PlasmaCashCoinResetEvent, error)
	FetchWithdrews(startBlock, endBlock uint64) ([]*pctypes.PlasmaCashWithdrewEvent, error)
	FetchFinalizedExit(startBlock, endBlock uint64) ([]*pctypes.PlasmaCashFinalizedExitEvent, error)
	FetchStartedExit(startBlock, endBlock uint64) ([]*pctypes.PlasmaCashStartedExitEvent, error)
}

type EthPlasmaClientImpl struct {
	EthPlasmaClientConfig
	ethClient      *ethclient.Client
	plasmaContract *ethcontract.RootChain
}

func (c *EthPlasmaClientImpl) Init() error {
	var err error
	c.ethClient, err = ethclient.Dial(c.EthereumURI)
	if err != nil {
		return errors.Wrap(err, "failed to connect to Ethereum")
	}

	c.plasmaContract, err = ethcontract.NewRootChain(common.HexToAddress(c.PlasmaHexAddress), c.ethClient)
	if err != nil {
		return errors.Wrap(err, "failed to bind Plasma Solidity contract")
	}
	return nil
}

func (c *EthPlasmaClientImpl) CurrentPlasmaBlockNum() (*big.Int, error) {
	curEthPlasmaBlockNum, err := c.plasmaContract.CurrentBlock(nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to obtain current plasma block from Ethereum")
	}
	return curEthPlasmaBlockNum, nil
}

func (c *EthPlasmaClientImpl) LatestEthBlockNum() (uint64, error) {
	blockHeader, err := c.ethClient.HeaderByNumber(context.TODO(), nil)
	if err != nil {
		return 0, err
	}
	return blockHeader.Number.Uint64(), nil
}

// SubmitPlasmaBlock will submit a Plasma block to Ethereum and wait until the tx is confirmed.
func (c *EthPlasmaClientImpl) SubmitPlasmaBlock(blockNum *big.Int, merkleRoot [32]byte) error {
	failMsg := "failed to submit plasma block to Ethereum"
	auth := bind.NewKeyedTransactor(c.PrivateKey)
	if c.OverrideGas {
		auth.GasPrice = big.NewInt(20000)
		auth.GasLimit = uint64(3141592)
	}
	tx, err := c.plasmaContract.SubmitBlock(auth, merkleRoot)
	if err != nil {
		return errors.Wrap(err, failMsg)
	}
	receipt, err := c.waitForTxReceipt(context.TODO(), tx)
	if err != nil {
		return err
	}
	if receipt.Status == 0 {
		return errors.New(failMsg)
	}
	return nil
}

func (c *EthPlasmaClientImpl) FetchWithdrews(startBlock, endBlock uint64) ([]*pctypes.PlasmaCashWithdrewEvent, error) {
	filterOpts := &bind.FilterOpts{
		Start: startBlock,
		End:   &endBlock,
	}

	withdrawCoinEvents := []*pctypes.PlasmaCashWithdrewEvent{}

	iterator, err := c.plasmaContract.FilterWithdrew(filterOpts, nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get Plasma coin withdraw logs")
	}
	defer iterator.Close()

	for iterator.Next() {
		event := iterator.Event

		localOwnerAddress, err := loom.LocalAddressFromHexString(event.Owner.Hex())
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse plasma coin withdraw owner's address")
		}
		localContractAddr, err := loom.LocalAddressFromHexString(event.ContractAddress.Hex())
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse Plasma deposit contract address")
		}

		withdrawCoinEvents = append(withdrawCoinEvents, &pctypes.PlasmaCashWithdrewEvent{
			Slot:         event.Slot,
			Owner:        loom.Address{ChainID: "eth", Local: localOwnerAddress}.MarshalPB(),
			Mode:         uint32(event.Mode),
			Uid:          &ltypes.BigUInt{Value: *loom.NewBigUInt(event.Uid)},
			Denomination: &ltypes.BigUInt{Value: *loom.NewBigUInt(event.Denomination)},
			Contract:     loom.Address{ChainID: "eth", Local: localContractAddr}.MarshalPB(),
		})
	}

	if err := iterator.Error(); err != nil {
		return nil, errors.Wrapf(err, "failed to iterate event data for plasma coin withdraw")
	}

	return withdrawCoinEvents, nil
}

func (c *EthPlasmaClientImpl) FetchFinalizedExit(startBlock, endBlock uint64) ([]*pctypes.PlasmaCashFinalizedExitEvent, error) {
	filterOpts := &bind.FilterOpts{
		Start: startBlock,
		End:   &endBlock,
	}

	finalizedExitEvents := []*pctypes.PlasmaCashFinalizedExitEvent{}

	iterator, err := c.plasmaContract.FilterFinalizedExit(filterOpts, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get Plasma finalized exit logs")
	}
	defer iterator.Close()

	for iterator.Next() {
		event := iterator.Event
		localOwnerAddress, err := loom.LocalAddressFromHexString(event.Owner.Hex())
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse plasma finalized exit owner's address")
		}

		finalizedExitEvents = append(finalizedExitEvents, &pctypes.PlasmaCashFinalizedExitEvent{
			Owner: loom.Address{ChainID: "eth", Local: localOwnerAddress}.MarshalPB(),
			Slot:  event.Slot,
		})
	}

	if err := iterator.Error(); err != nil {
		return nil, errors.Wrapf(err, "failed to iterate event data for plasma finalized exit")
	}

	return finalizedExitEvents, nil
}

func (c *EthPlasmaClientImpl) FetchCoinReset(startBlock, endBlock uint64) ([]*pctypes.PlasmaCashCoinResetEvent, error) {
	filterOpts := &bind.FilterOpts{
		Start: startBlock,
		End:   &endBlock,
	}

	coinResetEvents := []*pctypes.PlasmaCashCoinResetEvent{}

	iterator, err := c.plasmaContract.FilterCoinReset(filterOpts, nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get Plasma start exit logs")
	}
	defer iterator.Close()

	for iterator.Next() {
		event := iterator.Event
		localOwnerAddress, err := loom.LocalAddressFromHexString(event.Owner.Hex())
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse plasma start exit owner's address")
		}

		coinResetEvents = append(coinResetEvents, &pctypes.PlasmaCashCoinResetEvent{
			Owner: loom.Address{ChainID: "eth", Local: localOwnerAddress}.MarshalPB(),
			Slot:  event.Slot,
		})
	}

	if err := iterator.Error(); err != nil {
		return nil, errors.Wrapf(err, "failed to iterate event data for plasma start exit")
	}

	return coinResetEvents, nil
}

func (c *EthPlasmaClientImpl) FetchStartedExit(startBlock, endBlock uint64) ([]*pctypes.PlasmaCashStartedExitEvent, error) {
	filterOpts := &bind.FilterOpts{
		Start: startBlock,
		End:   &endBlock,
	}

	startedExitEvents := []*pctypes.PlasmaCashStartedExitEvent{}

	iterator, err := c.plasmaContract.FilterStartedExit(filterOpts, nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get Plasma start exit logs")
	}
	defer iterator.Close()

	for iterator.Next() {
		event := iterator.Event
		localOwnerAddress, err := loom.LocalAddressFromHexString(event.Owner.Hex())
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse plasma start exit owner's address")
		}

		startedExitEvents = append(startedExitEvents, &pctypes.PlasmaCashStartedExitEvent{
			Owner: loom.Address{ChainID: "eth", Local: localOwnerAddress}.MarshalPB(),
			Slot:  event.Slot,
		})
	}

	if err := iterator.Error(); err != nil {
		return nil, errors.Wrapf(err, "failed to iterate event data for plasma start exit")
	}

	return startedExitEvents, nil
}

// FetchDeposits fetches all deposit events from an Ethereum node from startBlock to endBlock (inclusive).
func (c *EthPlasmaClientImpl) FetchDeposits(startBlock, endBlock uint64) ([]*pctypes.PlasmaDepositEvent, error) {
	// NOTE: Currently either all blocks from w.StartBlock are processed successfully or none are.
	filterOpts := &bind.FilterOpts{
		Start: startBlock,
		End:   &endBlock,
	}
	depositEvents := []*pctypes.PlasmaDepositEvent{}

	iterator, err := c.plasmaContract.FilterDeposit(filterOpts, nil, nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get Plasma deposit logs")
	}
	defer iterator.Close()

	for iterator.Next() {
		event := iterator.Event
		localFromAddr, err := loom.LocalAddressFromHexString(event.From.Hex())
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse Plasma deposit 'from' address")
		}
		localContractAddr, err := loom.LocalAddressFromHexString(event.ContractAddress.Hex())
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse Plasma deposit contract address")
		}

		depositEvents = append(depositEvents, &pctypes.PlasmaDepositEvent{
			Slot:         event.Slot,
			DepositBlock: &ltypes.BigUInt{Value: *loom.NewBigUInt(event.BlockNumber)},
			Denomination: &ltypes.BigUInt{Value: *loom.NewBigUIntFromInt(1)}, // TODO: ev.Denomination
			From:         loom.Address{ChainID: "eth", Local: localFromAddr}.MarshalPB(),
			Contract:     loom.Address{ChainID: "eth", Local: localContractAddr}.MarshalPB(),
			// TODO: store ev.Hash... it's always a hash of ev.Slot, so a bit redundant
		})
	}

	if err := iterator.Error(); err != nil {
		return nil, errors.Wrap(err, "failed to iterate event data for Plasma deposit")
	}

	return depositEvents, nil
}

// waitForTxReceipt waits for a tx to be confirmed.
// It stops waiting if the context is canceled, or the tx hasn't been confirmed after TxTimeout.
func (c *EthPlasmaClientImpl) waitForTxReceipt(ctx context.Context, tx *etypes.Transaction) (*etypes.Receipt, error) {
	timeout := c.TxTimeout
	if timeout == 0 {
		timeout = 5 * time.Minute
	}

	interval := c.TxPollInterval
	if interval == 0 {
		interval = 10 * time.Second
	}

	timer := time.NewTimer(timeout)
	ticker := time.NewTicker(interval)

	defer timer.Stop()
	defer ticker.Stop()

	txHash := tx.Hash()
	for {
		receipt, err := c.ethClient.TransactionReceipt(ctx, txHash)
		if err != nil {
			return nil, errors.Wrap(err, "failed to retrieve tx receipt")
		}
		if receipt != nil {
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

func NewEthPlasmaClient(ethCfg EthPlasmaClientConfig) EthPlasmaClient {
	return &EthPlasmaClientImpl{EthPlasmaClientConfig: ethCfg}
}
