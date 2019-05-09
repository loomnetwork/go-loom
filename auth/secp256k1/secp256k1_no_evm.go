// +build !evm

package secp256k1

type Secp256k1Signer struct {
}

func NewSecp256k1Signer(privateKey []byte) *Secp256k1Signer {
	panic("EVM build isn't activated")
}

func (s *Secp256k1Signer) Sign(msg []byte) []byte {
	panic("EVM build isn't activated")
}

func (s *Secp256k1Signer) PublicKey() []byte {
	panic("EVM build isn't activated")
}
