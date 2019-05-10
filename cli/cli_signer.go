package cli

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/loomnetwork/go-loom/auth"
	"github.com/loomnetwork/go-loom/crypto"
)

func GetSigner(privFile, hsmConfigFile, algo string) (auth.Signer, error) {
	var signerType string

	if privFile == "" && hsmConfigFile == "" {
		return nil, errors.New("private key required to call contract")
	}

	var privKey crypto.PrivateKey
	var signer auth.Signer
	var err error
	if hsmConfigFile != "" {
		privKey, err = crypto.LoadYubiHsmPrivKey(hsmConfigFile)
		if err != nil {
			return nil, err
		}
		signerType = auth.SignerTypeYubiHsm
		signer = auth.NewSigner(signerType, privKey)
	} else {
		switch algo {
		case auth.SignerTypeSecp256k1:
			privKey, err = crypto.LoadSecp256k1PrivKey(privFile)
			if err != nil {
				return nil, err
			}
			signer = auth.NewSigner(algo, privKey)
		case auth.SignerTypeEd25519:
			privKeyB64, err := ioutil.ReadFile(privFile)
			if err != nil {
				return nil, err
			}

			privKey, err := base64.StdEncoding.DecodeString(string(privKeyB64))
			if err != nil {
				return nil, err
			}

			signer = auth.NewSigner(algo, privKey)
		case auth.SignerTypeTron:
			privKey, err = crypto.LoadBtecSecp256k1PrivKey(privFile)
			if err != nil {
				return nil, err
			}
			signer = auth.NewSigner(algo, privKey)
		default:
			return nil, fmt.Errorf("Unknown signer type %s", algo)
		}
	}

	return signer, err
}
