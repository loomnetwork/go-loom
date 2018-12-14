// +build !evm

package crypto

func Sign(hash []byte, prv PrivateKey) (sig []byte, err error) {
	panic("EVM build isn't activated")
}

func LoadSecp256k1PrivKey(filePath string) (interface{}, error) {
	panic("EVM build isn't activated")
}
