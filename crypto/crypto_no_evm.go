// +build !evm

package crypto

func Sign(hash []byte, prv PrivateKey) (sig []byte, err error) {
	panic("EVM build isn't activated")
}
