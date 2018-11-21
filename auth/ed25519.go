package auth

import "golang.org/x/crypto/ed25519"

// Ed25519Signer implements the Signer interface using ed25519 keys.
type Ed25519Signer struct {
	privateKey ed25519.PrivateKey
}

func NewEd25519Signer(privateKey []byte) *Ed25519Signer {
	var err error
	if privateKey == nil {
		_, privateKey, err = ed25519.GenerateKey(nil)
		if err != nil {
			panic(err)
		}
	}

	return &Ed25519Signer{privateKey}
}

func (s *Ed25519Signer) Sign(msg []byte) []byte {
	return ed25519.Sign(s.privateKey, msg)
}

func (s *Ed25519Signer) PublicKey() []byte {
	return []byte(s.privateKey.Public().(ed25519.PublicKey))
}
