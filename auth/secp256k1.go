package auth

import "github.com/tendermint/tendermint/crypto/secp256k1"

// Secp256k1Signer implements the Signer interface using secp256k1 keys
type Secp256k1Signer struct {
	privateKey secp256k1.PrivKeySecp256k1
}

func NewSecp256k1Signer(privateKey []byte) *Secp256k1Signer {
	secp256k1Signer := &Secp256k1Signer{}

	if privateKey == nil {
		secp256k1Signer.privateKey = secp256k1.GenPrivKey()
	} else {
		copy(secp256k1Signer.privateKey[:], privateKey[:])
	}
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
