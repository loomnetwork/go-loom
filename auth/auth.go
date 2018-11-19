package auth

import (
	"os"

	"github.com/tendermint/tendermint/crypto/secp256k1"
	"golang.org/x/crypto/ed25519"
)

const (
	EnableSecp256k1EnvVarName = "GOLOOM_ENABLE_SECP256K1"
	Secp256k1PrivKeySize      = 32
)

// Signer interface is used to sign transactions.
type Signer interface {
	Sign(msg []byte) []byte
	PublicKey() []byte
}

func NewAuthKey() ([]byte, []byte, error) {
	if os.Getenv(EnableSecp256k1EnvVarName) == "1" {
		privKey := secp256k1.GenPrivKey()
		return privKey.PubKey().Bytes(), privKey.Bytes(), nil
	}

	return ed25519.GenerateKey(nil)
}

func NewSigner(privateKey []byte) Signer {
	if os.Getenv(EnableSecp256k1EnvVarName) == "1" {
		return NewSecp256k1Signer(privateKey)
	}

	return NewEd25519Signer(privateKey)
}

// Ed25519Signer implements the Signer interface using ed25519 keys.
type Ed25519Signer struct {
	privateKey ed25519.PrivateKey
}

func NewEd25519Signer(privateKey []byte) *Ed25519Signer {
	return &Ed25519Signer{privateKey}
}

func (s *Ed25519Signer) Sign(msg []byte) []byte {
	return ed25519.Sign(s.privateKey, msg)
}

func (s *Ed25519Signer) PublicKey() []byte {
	return []byte(s.privateKey.Public().(ed25519.PublicKey))
}

// Secp256k1Signer implements the Signer interface using secp256k1 keys
type Secp256k1Signer struct {
	privateKey secp256k1.PrivKeySecp256k1
}

func NewSecp256k1Signer(privateKey []byte) *Secp256k1Signer {
	secp256k1Signer := &Secp256k1Signer{}

	copy(secp256k1Signer.privateKey[:], privateKey)
	return secp256k1Signer
}

func (s *Secp256k1Signer) Sign(msg []byte) []byte {
	sig, err := s.privateKey.Sign(msg)
	if err != nil {
		return nil
	}
	return sig.Bytes()
}

func (s *Secp256k1Signer) PublicKey() []byte {
	return s.privateKey.PubKey().Bytes()
}

// SignTx generates a signed tx containing the given bytes.
func SignTx(signer Signer, txBytes []byte) *SignedTx {
	return &SignedTx{
		Inner:     txBytes,
		Signature: signer.Sign(txBytes),
		PublicKey: signer.PublicKey(),
	}
}
