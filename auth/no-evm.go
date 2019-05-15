// +build !evm

package auth

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/eosspark/eos-go/crypto/ecc"
)

func GetTxHash(_ []byte) ([]byte, error) {
	return nil, fmt.Errorf("EVM build isn't activated")
}

func NewSecp256k1Signer(_ []byte) Signer {
	panic("EVM build isn't activated")
}

type EthSigner66Byte struct {
	PrivateKey *ecdsa.PrivateKey
}

func (k *EthSigner66Byte) Sign(_ []byte) []byte {
	return nil
}

func (k *EthSigner66Byte) PublicKey() []byte {
	return nil
}

type TronSigner struct {
	PrivateKey *ecdsa.PrivateKey
}

func (k *TronSigner) Sign(_ []byte) []byte {
	return nil
}

func (k *TronSigner) PublicKey() []byte {
	return nil
}

type EosSigner struct {
	PrivateKey *ecc.PrivateKey
}

func (k *EosSigner) Sign(_ []byte) []byte {
	return nil
}

func (k *EosSigner) PublicKey() []byte {
	return nil
}

