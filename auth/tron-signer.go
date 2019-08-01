// +build evm

package auth

import (
	"crypto/ecdsa"

	"github.com/btcsuite/btcd/btcec"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/loomnetwork/go-loom/common/evmcompat"
	sha3 "github.com/miguelmota/go-solidity-sha3"
)

type TronSigner struct {
	PrivateKey *ecdsa.PrivateKey
}

func NewTronSigner(privateKey []byte) *TronSigner {
	if privateKey == nil {
		privKey, err := btcec.NewPrivateKey(btcec.S256())
		if err != nil {
			panic(err)
		}
		return &TronSigner{PrivateKey: privKey.ToECDSA()}
	}

	privKey, _ := btcec.PrivKeyFromBytes(btcec.S256(), privateKey)
	return &TronSigner{PrivateKey: privKey.ToECDSA()}
}

func (k *TronSigner) Sign(txBytes []byte) []byte {
	signature, err := evmcompat.GenerateTypedSig(
		sha3.SoliditySHA3(
			sha3.String("\x19TRON Signed Message:\n32"),
			sha3.SoliditySHA3(txBytes),
		),
		k.PrivateKey,
		evmcompat.SignatureType_TRON,
	)
	if err != nil {
		panic(err)
	}
	return signature
}

func (k *TronSigner) PublicKey() []byte {
	return crypto.FromECDSAPub(&k.PrivateKey.PublicKey)
}
