// +build evm

package auth

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/crypto"
)

type TronSigner struct {
	PrivateKey *ecdsa.PrivateKey
}

func (k *TronSigner) Sign(txBytes []byte) []byte {
	hash, err := GetTxHash(txBytes)
	if err != nil {
		panic(err)
	}
	signature, err := crypto.Sign(hash, k.PrivateKey)
	if err != nil {
		panic(err)
	}
	return signature
}

func (k *TronSigner) PublicKey() []byte {
	return crypto.FromECDSAPub(&k.PrivateKey.PublicKey)
}

