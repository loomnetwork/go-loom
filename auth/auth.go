package auth

import (
	"fmt"

	"github.com/loomnetwork/go-loom/auth/secp256k1"
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
		return secp256k1.NewSecp256k1Signer(privKey)
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
