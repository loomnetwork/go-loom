// +build !evm

package crypto

func LoadECDSA(privKeyType string, filePath string) (*PrivateKey, error) {
	panic("EVM build isn't activated")
}

func Sign(hash []byte, prv PrivateKey) (sig []byte, err error) {
	panic("EVM build isn't activated")
}
