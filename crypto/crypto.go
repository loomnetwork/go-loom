package crypto

import (
	"crypto/ecdsa"
	"encoding/base64"
	"io"
	"io/ioutil"
	"os"

	"github.com/btcsuite/btcd/btcec"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
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

// LoadBtecSecp256k1PrivKey converts private key from btec secp256k1 to ecdsa
func LoadBtecSecp256k1PrivKey(file string) (*ecdsa.PrivateKey, error) {
	buf, err := LoadBtecSecp256k1PrivKeyByte(file)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load btec secp256k1 private key")
	}
	privKey, _ := btcec.PrivKeyFromBytes(btcec.S256(), common.FromHex(string(buf)))
	return privKey.ToECDSA(), nil
}

// LoadBtecSecp256k1PrivKeyByte reads 64 byte from private key file
func LoadBtecSecp256k1PrivKeyByte(file string) ([]byte, error) {
	fd, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer fd.Close()
	buf := make([]byte, 64)
	if _, err := io.ReadFull(fd, buf); err != nil {
		return nil, err
	}
	return buf, nil
}
