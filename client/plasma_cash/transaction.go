// +build evm

package plasma_cash

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/crypto/sha3"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/loomnetwork/go-loom/common/evmcompat"
)

type LoomTx struct {
	Slot         uint64
	Denomination *big.Int
	Owner        common.Address
	PrevBlock    *big.Int
	Signature    []byte
	TXProof      []byte
}

func (l *LoomTx) Sig() []byte {
	return l.Signature
}

func (l *LoomTx) Proof() Proof {
	return l.TXProof
}

func (l *LoomTx) NewOwner() common.Address {
	return l.Owner
}

func (l *LoomTx) Sign(key *ecdsa.PrivateKey) ([]byte, error) {
	hash, err := l.Hash()
	if err != nil {
		return nil, err
	}

	sig, err := evmcompat.SoliditySign(hash, key)
	if err != nil {
		return nil, err
	}

	// The first byte should be the signature mode, for details about the signature format refer to
	// https://github.com/loomnetwork/plasma-erc721/blob/master/server/contracts/Libraries/ECVerify.sol
	return append(make([]byte, 1, 66), sig...), nil
}

func (l *LoomTx) RlpEncode() ([]byte, error) {
	return rlp.EncodeToBytes([]interface{}{
		uint64(l.Slot),
		l.PrevBlock,
		l.Denomination,
		l.Owner,
	})
}

func (l *LoomTx) Hash() ([]byte, error) {
	if l.PrevBlock.Cmp(big.NewInt(0)) != 0 {
		ret, err := l.rlpEncodeWithSha3()
		if err != nil {
			return nil, err
		}
		return ret, nil
	}

	data, err := soliditySha3(l.Slot)
	if err != nil {
		return nil, err
	}
	if len(data) != 32 {
		return nil, fmt.Errorf("wrong hash size! expected 32, got %v", len(data))
	}
	return data, nil
}

func soliditySha3(data uint64) ([]byte, error) {
	pairs := []*evmcompat.Pair{&evmcompat.Pair{Type: "uint64", Value: strconv.FormatUint(data, 10)}}
	hash, err := evmcompat.SoliditySHA3(pairs)
	if err != nil {
		return []byte{}, err
	}
	return hash, err
}

func (l *LoomTx) rlpEncodeWithSha3() ([]byte, error) {
	hash, err := l.RlpEncode()
	if err != nil {
		return []byte{}, err
	}
	d := sha3.NewLegacyKeccak256()
	d.Write(hash)
	return d.Sum(nil), nil
}
