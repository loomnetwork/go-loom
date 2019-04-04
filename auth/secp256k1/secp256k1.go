package secp256k1

import (
	"crypto/ecdsa"
)

// Secp256k1Signer implements the Signer interface using secp256k1 keys
type Secp256k1Signer struct {
	//nolint:unused
	privateKey *ecdsa.PrivateKey
}
