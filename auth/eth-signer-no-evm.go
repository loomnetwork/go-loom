// +build !evm

package auth

import (
	"crypto/ecdsa"
)

type EthSigner66Byte struct {
	PrivateKey *ecdsa.PrivateKey
}

func (k *EthSigner66Byte) Sign(_ []byte) []byte {
	return nil
}

func (k *EthSigner66Byte) PublicKey() []byte {
	return nil
}
