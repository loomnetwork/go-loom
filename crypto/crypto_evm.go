// +build evm

package crypto

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/crypto"
)

func LoadECDSA(privKeyType string, filePath string) (PrivateKey, error) {
	switch privKeyType {
	case PrivKeyTypeFile:
		return LoadSecp256k1PrivKey(filePath)
	case PrivKeyTypeYubiHsm:
		return LoadYubiHsmPrivKey(filePath)
	default:
		panic("Unknow ECDSA private key type")
	}
}

func SoliditySign(hash []byte, prv PrivateKey) (sig []byte, err error) {
	switch prv.(type) {
	case Secp256k1PrivateKey:
		sig, err = crypto.Sign(hash, prv.(*ecdsa.PrivateKey))
	case YubiHsmPrivateKey:
		sig, err = YubiHsmSign(hash, prv.(*YubiHsmPrivateKey))
	}

	if err != nil {
		return nil, err
	}

	v := sig[len(sig)-1]
	sig[len(sig)-1] = v + 27

	return sig, nil
}
