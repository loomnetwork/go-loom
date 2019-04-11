// +build evm

package crypto

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
	ssha "github.com/miguelmota/go-solidity-sha3"
)

type Secp256k1PrivateKey ecdsa.PrivateKey

func (s *Secp256k1PrivateKey) ToECDSAPrivKey() *ecdsa.PrivateKey {
	ecdsaPrivKey := ecdsa.PrivateKey(*s)
	return &ecdsaPrivKey
}

func LoadSecp256k1PrivKey(filePath string) (*Secp256k1PrivateKey, error) {
	privKey, err := crypto.LoadECDSA(filePath)
	if err != nil {
		return nil, err
	}

	secpPrivKey := Secp256k1PrivateKey(*privKey)

	return &secpPrivKey, nil
}

func SoliditySign(hash []byte, prv PrivateKey) (sig []byte, err error) {
	switch prv.(type) {
	case *Secp256k1PrivateKey:
		sig, err = crypto.Sign(hash, prv.(*Secp256k1PrivateKey).ToECDSAPrivKey())
	case *YubiHsmPrivateKey:
		//TODO this feels out of place
		sig, err = YubiHsmSign(hash, prv.(*YubiHsmPrivateKey))
	default:
		return nil, fmt.Errorf("unknown private key type")
	}

	if err != nil {
		return nil, err
	}

	v := sig[len(sig)-1]
	sig[len(sig)-1] = v + 27

	return sig, nil
}

func SoliditySignPrefixed(hash []byte, prv PrivateKey) (sig []byte, err error) {
	// Need to prefix the hash with the Ethereum Signed Message
	hash = ssha.SoliditySHA3(
		[]string{"string", "bytes32"},
		"\x19Ethereum Signed Message:\n32",
		hash,
	)

	switch prv.(type) {
	case *Secp256k1PrivateKey:
		sig, err = crypto.Sign(hash, prv.(*Secp256k1PrivateKey).ToECDSAPrivKey())
	case *YubiHsmPrivateKey:
		//TODO this feels out of place
		sig, err = YubiHsmSign(hash, prv.(*YubiHsmPrivateKey))
	default:
		return nil, fmt.Errorf("unknown private key type")
	}

	if err != nil {
		return nil, err
	}

	v := sig[len(sig)-1]
	sig[len(sig)-1] = v + 27

	return sig, nil
}
