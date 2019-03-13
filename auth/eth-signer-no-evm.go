// +build !evm

package auth

type EthSigner66Byte struct {
	PrivateKey interface{}
}

func NewEthSigner66Byte(_ []byte) Signer {
	panic("EVM build isn't activated")
}

func (k *EthSigner66Byte) Sign(_ []byte) []byte {
	return nil
}

func (k *EthSigner66Byte) PublicKey() []byte {
	return nil
}
