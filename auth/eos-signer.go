// +build evm

package auth

import (
	"crypto/ecdsa"
	"github.com/eosspark/eos-go/crypto/ecc"
	sha3 "github.com/miguelmota/go-solidity-sha3"
	"github.com/ethereum/go-ethereum/crypto"
)

type EosSigner struct {
	PrivateKey *ecc.PrivateKey
}

func (e *EosSigner) Sign(txBytes []byte) []byte {
	signature, err := e.PrivateKey.Sign(sha3.SoliditySHA3(txBytes))
	if err != nil {
		panic(err)
	}
	sigByes, err := signature.Pack()
	if err != nil {
		panic(err)
	}
	return sigByes
}

func (e *EosSigner) PublicKey() []byte {
	btcecPubKey, err := e.PrivateKey.PublicKey().Key()
	if err != nil {
		panic(err)
	}
	ecdsaPubKey := ecdsa.PublicKey(*btcecPubKey)
	return crypto.FromECDSAPub(&ecdsaPubKey)
}
