// +build !evm

package yubihsm

func NewYubiHsmSigner(privateKey interface{}) *YubiHsmSigner {
	panic("EVM build isn't activated")
}

func (s *YubiHsmSigner) Sign(msg []byte) []byte {
	panic("EVM build isn't activated")
}

func (s *YubiHsmSigner) PublicKey() []byte {
	panic("EVM build isn't activated")
}
