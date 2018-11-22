package auth

import (
	"crypto/ecdsa"
	"crypto/rand"

	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

// Secp256k1Signer implements the Signer interface using secp256k1 keys
type Secp256k1Signer struct {
	privateKey [32]byte
	publicKey  [33]byte
}

func NewSecp256k1Signer(privateKey []byte) *Secp256k1Signer {
	secp256k1Signer := &Secp256k1Signer{}

	if privateKey == nil {
		return genKeyPair()
	}

	bitCurve := secp256k1.S256()
	x, y := bitCurve.ScalarBaseMult(privateKey)

	pubKey := secp256k1.CompressPubkey(x, y)
	copy(secp256k1Signer.publicKey[:], pubKey[:])
	copy(secp256k1Signer.privateKey[:], privateKey[:])

	return secp256k1Signer
}

func genKeyPair() *Secp256k1Signer {
	secp256k1Signer := &Secp256k1Signer{}

	key, err := ecdsa.GenerateKey(secp256k1.S256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	pubKey := secp256k1.CompressPubkey(key.X, key.Y)
	copy(secp256k1Signer.publicKey[:], pubKey[:])

	blob := key.D.Bytes()
	copy(secp256k1Signer.privateKey[32-len(blob):], blob)

	return secp256k1Signer
}

func (s *Secp256k1Signer) Sign(msg []byte) []byte {
	sig, err := secp256k1.Sign(msg, s.privateKey[:])
	if err != nil {
		return nil
	}
	return sig
}

func (s *Secp256k1Signer) PublicKey() []byte {
	return s.publicKey[:]
}
