// +build !evm

package auth

func NewEthSigner66Byte(privateKey []byte) Signer {
	panic("EVM build isn't activated")
}
