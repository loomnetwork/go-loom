// +build evm

package plasma_cash

import (
	"crypto/ecdsa"

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
