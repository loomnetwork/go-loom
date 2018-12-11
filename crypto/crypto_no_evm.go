// +build !evm

package crypto

func LoadECDSA(filePath string) (*PrivateKey, error) {
	panic("EVM build isn't activated")
}

func Sign(hash []byte, prv PrivateKey) (sig []byte, err error) {
	panic("EVM build isn't activated")
}
