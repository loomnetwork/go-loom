package auth

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"errors"

	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

const (
	Secp256k1PubKeyBytes  = 33
	Secp256k1PrivKeyBytes = 32
	Secp256k1SigBytes     = 64
)

// Secp256k1Signer implements the Signer interface using secp256k1 keys
type Secp256k1Signer struct {
	privateKey [Secp256k1PrivKeyBytes]byte
	publicKey  [Secp256k1PubKeyBytes]byte
}

func NewSecp256k1Signer(privateKey []byte) *Secp256k1Signer {
	secp256k1Signer := &Secp256k1Signer{}
	if privateKey == nil {
		pubKeyBytes, privKeyBytes := GenSecp256k1Key()

		copy(secp256k1Signer.publicKey[:], pubKeyBytes[:])
		copy(secp256k1Signer.privateKey[:], privKeyBytes[:])
	} else {
		if len(privateKey) != Secp256k1PrivKeyBytes {
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

	privKeyBytes := make([]byte, Secp256k1PrivKeyBytes)
	blob := key.D.Bytes()
	copy(privKeyBytes[Secp256k1PrivKeyBytes-len(blob):], blob)

	return pubKey, privKeyBytes[:]
}

func (s *Secp256k1Signer) Sign(msg []byte) []byte {
	var sig [Secp256k1SigBytes]byte

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
