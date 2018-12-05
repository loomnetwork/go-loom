// +build evm

package secp256k1

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

	sig := signer.Sign(testMsg)
	if len(sig) != Secp256k1SigBytes {
		return nil, errors.New("Invalid params for VerifySignature")
	}

	if signer.verifyBytes(testMsg, sig) == false {
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

func TestCompareAddress(t *testing.T) {
	// get ethereum address
	privKey, err := crypto.HexToECDSA(TestEthereumPrivKey)
	if err != nil {
		t.Fatal(err)
	}
	ethAddr1 := crypto.PubkeyToAddress(privKey.PublicKey)

	// get pubkey address by secp256k1 signer
	key, _ := hex.DecodeString(TestEthereumPrivKey)
	signer := NewSecp256k1Signer(key)
	ethAddr2 := crypto.PubkeyToAddress(signer.privateKey.PublicKey)

	if !bytes.Equal(ethAddr1.Bytes(), ethAddr2.Bytes()) {
		t.Fatal("public key address isn't mismatched")
	}
}
