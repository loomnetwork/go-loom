// +build evm

package plasma_cash

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/crypto/sha3"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/loomnetwork/go-loom/common/evmcompat"
)

func Sha3WithRlpEncoding(value interface{}) ([]byte, error) {
	encodedBytes, err := rlp.EncodeToBytes(value)
	if err != nil {
		return nil, err
	}

	return Sha3(encodedBytes), nil
}

func Sha3(input []byte) []byte {
	hasher := sha3.NewKeccak256()
	hasher.Write(input)
	return hasher.Sum(nil)
}

func SolidityTypedSign(hash []byte, key *ecdsa.PrivateKey) ([]byte, error) {
	sig, err := evmcompat.SoliditySign(hash, key)
	if err != nil {
		return nil, err
	}

	// The first byte should be the signature mode, for details about the signature format refer to
	// https://github.com/loomnetwork/plasma-erc721/blob/master/server/contracts/Libraries/ECVerify.sol
	return append(make([]byte, 1, 66), sig...), nil
}
