// +build evm

package crypto

import (
	"github.com/ethereum/go-ethereum/crypto"
)

func LoadECDSA(hsmEnabled bool, filePath string) (PrivateKey, error) {
	if hsmEnabled {
		return LoadSecp256k1PrivKey(filePath)
	}

	return LoadYubiHsmPrivKey(filePath)
}

func SoliditySign(hash []byte, prv PrivateKey) (sig []byte, err error) {
	switch prv.(type) {
	case Secp256k1PrivateKey:
		sig, err = crypto.Sign(hash, prv.(*Secp256k1PrivateKey).ToECDSAPrivKey())
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
