// +build evm

package yubihsm

import (
	"crypto/sha256"
	"errors"

	"github.com/loomnetwork/go-loom/crypto"
)

func NewYubiHsmSigner(privateKey crypto.PrivateKey) *YubiHsmSigner {
	yubiHsmSigner := &YubiHsmSigner{}
	if privateKey == nil {
		panic(errors.New("The private key should be given for YubiHSM signer"))
	}

	yubiHsmSigner.PrivateKey = privateKey
	return yubiHsmSigner
}

func (s *YubiHsmSigner) PublicKey() []byte {
	return s.PrivateKey.(*crypto.YubiHsmPrivateKey).GetPubKeyBytes()
}

func (s *YubiHsmSigner) Sign(msg []byte) []byte {
	hash := sha256.Sum256(msg)
	sigBytes, err := crypto.YubiHsmSign(hash[:], s.PrivateKey.(*crypto.YubiHsmPrivateKey))
	if err != nil {
		panic(err)
	}

	return sigBytes
}
