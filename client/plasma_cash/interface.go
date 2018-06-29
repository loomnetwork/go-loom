// +build evm

package plasma_cash

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Proof []byte

type Block interface {
	MerkleHash() []byte
	TxFromSlot(slot uint64) (Tx, error)
}

type Tx interface {
	RlpEncode() ([]byte, error)
	Sign(key *ecdsa.PrivateKey) ([]byte, error)
	Sig() []byte
	NewOwner() common.Address
	Proof() Proof
}

type ChainServiceClient interface {
	CurrentBlock() (Block, error)
	BlockNumber() (int64, error)

	Block(blknum int64) (Block, error)
	//Proof(blknum int64, slot uint64) (Proof, error)

	SubmitBlock() error

	SendTransaction(slot uint64, prevBlock int64, denomination int64, newOwner string, sig []byte) error
}

type Account struct {
	Address    string
	PrivateKey *ecdsa.PrivateKey
}

type TokenContract interface {
	Register() error
	Deposit(int64) (common.Hash, error)
	BalanceOf() (int64, error)

	Account() (*Account, error)
}

type PlasmaCoinState uint8

const (
	PlasmaCoinDeposited PlasmaCoinState = iota
	PlasmaCoinExiting
	PlasmaCoinChallenged
	PlasmaCoinResponded
	PlasmaCoinExited
)

type PlasmaCoin struct {
	UID             uint64
	DepositBlockNum int64
	Denomination    uint32
	Owner           string
	State           PlasmaCoinState
}

type DepositEventData struct {
	// Plasma slot, a unique identifier, assigned to the deposit.
	Slot uint64
	// Index of the Plasma block in which the deposit transaction was included.
	BlockNum *big.Int
}

type RootChainClient interface {
	FinalizeExits() error
	Withdraw(slot uint64) error
	WithdrawBonds() error
	PlasmaCoin(slot uint64) (*PlasmaCoin, error)
	StartExit(slot uint64, prevTx Tx, exitingTx Tx, prevTxProof Proof,
		exitingTxProof Proof, sigs []byte, prevTxBlkNum int64, txBlkNum int64) ([]byte, error)

	ChallengeBefore(slot uint64, prevTx Tx, exitingTx Tx,
		prevTxInclusionProof Proof, exitingTxInclusionProof Proof,
		sig []byte, prevTxBlockNum int64, exitingTxBlockNum int64) ([]byte, error)

	RespondChallengeBefore(slot uint64, challengingBlockNumber int64,
		challengingTransaction Tx, proof Proof, sig []byte) ([]byte, error)

	ChallengeBetween(slot uint64, challengingBlockNumber int64,
		challengingTransaction Tx, proof Proof, sig []byte) ([]byte, error)

	ChallengeAfter(slot uint64, challengingBlockNumber int64,
		challengingTransaction Tx, proof Proof, sig []byte) ([]byte, error)

	SubmitBlock(blockNum *big.Int, merkleRoot [32]byte) error

	DebugCoinMetaData(slots []uint64)
	DepositEventData(txHash common.Hash) (*DepositEventData, error)
}
