// +build evm

package secp256k1

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

const (
	Secp256k1PubKeyBytes  = 33
	Secp256k1PrivKeyBytes = 32
	Secp256k1SigBytes     = 65
)

// Secp256k1Signer implements the Signer interface using secp256k1 keys
type Secp256k1Signer struct {
	privateKey *ecdsa.PrivateKey
}

func NewSecp256k1Signer(privateKey []byte) *Secp256k1Signer {
	var err error

	secp256k1Signer := &Secp256k1Signer{}
	if privateKey == nil {
		secp256k1Signer.privateKey, err = ecdsa.GenerateKey(secp256k1.S256(), rand.Reader)
		if err != nil {
			panic(err)
		}
	} else {
		if len(privateKey) != Secp256k1PrivKeyBytes {
			panic(errors.New("Invalid private key length"))
		}

		hexPrivKey := hex.EncodeToString(privateKey)
		secp256k1Signer.privateKey, err = crypto.HexToECDSA(hexPrivKey)
		if err != nil {
			panic(err)
		}
	}

	return secp256k1Signer
}

func GenSecp256k1Key() ([]byte, []byte) {
	var pubKey []byte

	key, err := ecdsa.GenerateKey(secp256k1.S256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	privKeyBytes := ecdsaToBytes(key)
	pubKey = secp256k1.CompressPubkey(key.X, key.Y)

	return pubKey, privKeyBytes[:]
}

func VerifyBytes(pubKey []byte, msg []byte, sig []byte) bool {
	var sigBytes [Secp256k1SigBytes - 1]byte

	if len(sig) != Secp256k1SigBytes {
		panic("Invalid signature Secp256k1SigBytes length")
	}

	copy(sigBytes[:], sig[:])
	hash := sha256.Sum256(msg)

	return secp256k1.VerifySignature(pubKey, hash[:], sigBytes[:])
}

func (s *Secp256k1Signer) Sign(msg []byte) []byte {
	privKeyBytes := ecdsaToBytes(s.privateKey)

	hash := sha256.Sum256(msg)
	sigBytes, err := secp256k1.Sign(hash[:], privKeyBytes[:])
	if err != nil {
		panic(err)
	}

	return sigBytes
}

func (s *Secp256k1Signer) PublicKey() []byte {
	return secp256k1.CompressPubkey(s.privateKey.X, s.privateKey.Y)
}

func (s *Secp256k1Signer) verifyBytes(msg []byte, sig []byte) bool {
	return VerifyBytes(s.PublicKey(), msg, sig)
}

func ecdsaToBytes(privKey *ecdsa.PrivateKey) [Secp256k1PrivKeyBytes]byte {
	var privKeyBytes [Secp256k1PrivKeyBytes]byte

	blob := privKey.D.Bytes()
	copy(privKeyBytes[Secp256k1PrivKeyBytes-len(blob):], blob)

	return privKeyBytes
}
