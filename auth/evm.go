// +build evm

package auth

import (
	"github.com/ethereum/go-ethereum/common"
	sha3 "github.com/miguelmota/go-solidity-sha3"
	"github.com/pkg/errors"
)

func GetTxHash(nonceBytes []byte) ([]byte, error) {
	from, to, nonce, err := GetFromToNonce(nonceBytes)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshalling tx details")
	}
	return sha3.SoliditySHA3(
		sha3.Address(common.BytesToAddress(from.Local)),
		sha3.Address(common.BytesToAddress(to.Local)),
		sha3.Uint64(nonce),
		nonceBytes,
	), nil
}