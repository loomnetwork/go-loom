// +build !evm

package auth

import (
	"crypto/ecdsa"
	"fmt"
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

func NewTronSigner(privateKey []byte) *TronSigner {
	panic("EVM build isn't activated")
}

func (k *TronSigner) Sign(_ []byte) []byte {
	return nil
}

func (k *TronSigner) PublicKey() []byte {
	return nil
}

type BinanceSigner struct {
	PrivateKey *ecdsa.PrivateKey
}

func NewBinanceSigner(privateKey []byte) *BinanceSigner {
	panic("EVM build isn't activated")
}

func (k *BinanceSigner) Sign(_ []byte) []byte {
	return nil
}

func (k *BinanceSigner) PublicKey() []byte {
	return nil
}
