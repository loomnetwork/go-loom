package crypto

import (
	"encoding/base64"
	"io/ioutil"
)

const (
	PrivateKeyTypeEd25519   = "ed25519"
	PrivateKeyTypeSecp256k1 = "secp256k1"
)

type PrivateKey interface{}

func LoadEd25519PrivKey(path string) ([]byte, error) {
	privKeyB64, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	privKey, err := base64.StdEncoding.DecodeString(string(privKeyB64))
	if err != nil {
		return nil, err
	}

	return privKey, nil
}
