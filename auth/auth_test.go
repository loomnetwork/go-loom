// +build evm

package auth

import (
	"crypto/sha256"
	"testing"

	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

func TestSecp256k1Sign(t *testing.T) {
	testMsg := []byte{'t', 'e', 's', 't'}

	signer := NewSecp256k1Signer(nil)
	if len(signer.privateKey) != 32 {
		t.Fatalf("Invalid private key length:%d", len(signer.privateKey))
	}

	sig := signer.Sign(testMsg)

	msg := sha256.Sum256(testMsg)
	if len(msg) != 32 || len(sig) != 64 || len(signer.publicKey) == 0 {
		t.Fatalf("Invalid params(msg_len:%d, sig_len:%d, pubkey_len:%d) for VerifySignature",
			len(msg), len(sig), len(signer.publicKey))
	}

	if secp256k1.VerifySignature(signer.publicKey[:], msg[:], sig) == false {
		t.Fatal("Signature is invalid")
	}
}
