package yubihsm

import (
	"github.com/loomnetwork/go-loom/crypto"
)

// YubiHsmSigner implements signer by YubiHSM secp256k1
type YubiHsmSigner struct {
	PrivateKey crypto.PrivateKey
}
