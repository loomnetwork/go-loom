// +build evm

package auth

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/crypto"
	sha3 "github.com/miguelmota/go-solidity-sha3"
)

type TronSigner struct {
	PrivateKey *ecdsa.PrivateKey
}

func (k *TronSigner) Sign(txBytes []byte) []byte {
	signature, err := crypto.Sign(sha3.SoliditySHA3(txBytes), k.PrivateKey)
	if err != nil {
		panic(err)
	}
	return signature
}

func (k *TronSigner) PublicKey() []byte {
	return crypto.FromECDSAPub(&k.PrivateKey.PublicKey)
}
