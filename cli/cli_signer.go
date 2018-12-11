package cli

import (
	"encoding/base64"
	"errors"
	"io/ioutil"

	"github.com/loomnetwork/go-loom/crypto"

	"github.com/loomnetwork/go-loom/auth"
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
		privKey, err = crypto.LoadYubiHsmPrivKey(algo, hsmConfigFile)
		if err != nil {
			return nil, err
		}
		signerType = auth.SignerTypeYubiHsm
		signer = auth.NewSigner(signerType, privKey)
	} else {
		if algo == "secp256k1" {
			privKey, err = crypto.LoadSecp256k1PrivKey(privFile)
			if err != nil {
				return nil, err
			}
			signer = auth.NewSigner(algo, privKey)
		} else {
			//ed25519
			privKeyB64, err := ioutil.ReadFile(privFile)
			if err != nil {
				return nil, err
			}

			privKey, err := base64.StdEncoding.DecodeString(string(privKeyB64))
			if err != nil {
				return nil, err
			}

			signer = auth.NewSigner(algo, privKey)
		}
	}

	return signer, err
}
