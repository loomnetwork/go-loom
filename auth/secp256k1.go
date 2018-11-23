package auth

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"errors"

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
		pubKeyBytes, privKeyBytes := GenSecp256k1Key()

		copy(secp256k1Signer.publicKey[:], pubKeyBytes[:])
		copy(secp256k1Signer.privateKey[:], privKeyBytes[:])
	} else {
		if len(privateKey) != 32 {
			panic(errors.New("Invalid private key length"))
		}

		copy(secp256k1Signer.privateKey[:], privateKey[:])

		bitCurve := secp256k1.S256()
		x, y := bitCurve.ScalarBaseMult(privateKey)

		pubKeyBytes := secp256k1.CompressPubkey(x, y)
		copy(secp256k1Signer.publicKey[:], pubKeyBytes[:])
	}

	return secp256k1Signer
}

func GenSecp256k1Key() ([]byte, []byte) {
	var pubKey []byte

	key, err := ecdsa.GenerateKey(secp256k1.S256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	pubKey = secp256k1.CompressPubkey(key.X, key.Y)

	privKeyBytes := make([]byte, 32)
	blob := key.D.Bytes()
	copy(privKeyBytes[32-len(blob):], blob)

	return pubKey, privKeyBytes[:]
}

func (s *Secp256k1Signer) Sign(msg []byte) []byte {
	var sig [64]byte

	hash := sha256.Sum256(msg)
	sigBytes, err := secp256k1.Sign(hash[:], s.privateKey[:])
	if err != nil {
		panic(err)
	}
	copy(sig[:], sigBytes[:])
	return sig[:]
}

func (s *Secp256k1Signer) PublicKey() []byte {
	return s.publicKey[:]
}
