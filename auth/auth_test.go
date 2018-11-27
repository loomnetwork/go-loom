// +build evm

package auth

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
)

var (
	TestEthereumPrivKey = "b04df8f5492ef497f6202a34669a6ebbd8340c7a3f02f7f921c1b98d538e7947"
)

func testSign(privKey []byte, t *testing.T) ([]byte, error) {
	testMsg := []byte{'t', 'e', 's', 't'}

	signer := NewSecp256k1Signer(privKey)
	if len(signer.privateKey) != 32 {
		return nil, errors.New("Invalid private key length")
	}

	sig := signer.Sign(testMsg)

	if len(sig) != Secp256k1SigBytes || len(signer.publicKey) != Secp256k1PubKeyBytes {
		return nil, errors.New("Invalid params for VerifySignature")
	}

	if VerifyBytes(signer.publicKey[:], testMsg, sig) == false {
		return nil, errors.New("Signature is invalid")
	}

	return sig, nil
}

func TestSecp256k1Sign(t *testing.T) {
	if _, err := testSign(nil, t); err != nil {
		t.Fatal(err)
	}
}

func TestImportEthereumKey(t *testing.T) {
	key, _ := hex.DecodeString(TestEthereumPrivKey)
	if _, err := testSign(key, t); err != nil {
		t.Fatal(err)
	}
}

func TestImportSecp256k1Key(t *testing.T) {
	signer := NewSecp256k1Signer(nil)

	hexKey := hex.EncodeToString(signer.privateKey[:])
	_, err := crypto.HexToECDSA(hexKey)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCompareSig(t *testing.T) {
	var sig2 [Secp256k1SigBytes]byte
	testMsg := []byte{'t', 'e', 's', 't'}

	key, _ := hex.DecodeString(TestEthereumPrivKey)
	sig1, _ := testSign(key, t)

	privKey, err := crypto.HexToECDSA(TestEthereumPrivKey)
	if err != nil {
		t.Fatal(err)
	}

	hash := sha256.Sum256(testMsg)
	sig2Bytes, err := crypto.Sign(hash[:], privKey)
	if err != nil {
		t.Fatal(err)
	}
	copy(sig2[:], sig2Bytes[:])

	if !bytes.Equal(sig1, sig2[:]) {
		t.Fatal("the signature is mismatched")
	}
}
