// +build !evm

package auth

func NewSecp256k1Signer(privateKey []byte) Signer {
	panic("EVM build isn't activated")
}
