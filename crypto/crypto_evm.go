// +build evm

package crypto

import (
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
)

func SoliditySign(hash []byte, prv PrivateKey) (sig []byte, err error) {
	switch prv.(type) {
	case *Secp256k1PrivateKey:
		sig, err = crypto.Sign(hash, prv.(*Secp256k1PrivateKey).ToECDSAPrivKey())
	case YubiHsmPrivateKey:
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
