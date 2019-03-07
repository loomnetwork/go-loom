// +build !evm

package auth

func Eth256Signer(privateKey []byte) Signer {
	panic("EVM build isn't activated")
}
