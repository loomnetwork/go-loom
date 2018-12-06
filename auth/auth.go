package auth

import (
	"fmt"
)

const (
	SignerTypeEd25519   = "ed25519"
	SignerTypeSecp256k1 = "secp256k1"
)

// Signer interface is used to sign transactions.
type Signer interface {
	Sign(msg []byte) []byte
	PublicKey() []byte
}

func NewSigner(signerType string, privKey []byte) Signer {
	switch signerType {
	case SignerTypeEd25519:
		return NewEd25519Signer(privKey)
	case SignerTypeSecp256k1:
		return NewSecp256k1Signer(privKey)
	default:
		panic(fmt.Errorf("Unknown signer type %s", signerType))
	}
	return nil
}

// SignTx generates a signed tx containing the given bytes.
func SignTx(signer Signer, txBytes []byte) *SignedTx {
	return &SignedTx{
		Inner:     txBytes,
		Signature: signer.Sign(txBytes),
		PublicKey: signer.PublicKey(),
	}
}
