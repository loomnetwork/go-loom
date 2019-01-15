// +build evm

package crypto

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

func TestGenYubiSecp256k1Key(t *testing.T) {
	// check if yubiHsm config is given
	cfg := os.Getenv("YUBIHSM_CFG_FILE")
	if len(cfg) == 0 {
		t.Log("YUBIHSM_CFG_FILE is not set")
		return
	}

	t.Log("Generating YubiHSM private key")
	yubiPrivKey, err := GenYubiHsmPrivKey(cfg)
	if err != nil {
		t.Fatal(err)
	}
	defer yubiPrivKey.UnloadYubiHsmPrivKey()
}

func TestSignYubiSecp256k1(t *testing.T) {
	// check if yubiHsm config is given
	cfg := os.Getenv("YUBIHSM_CFG_FILE")
	if len(cfg) == 0 {
		t.Log("YUBIHSM_CFG_FILE is not set")
		return
	}

	t.Log("Loading YubiHSM private key")
	yubiPrivKey, err := LoadYubiHsmPrivKey(cfg)
	if err != nil {
		t.Fatal(err)
	}
	defer yubiPrivKey.UnloadYubiHsmPrivKey()

	t.Logf("LoadYubiHsmPrivKey succeeded")

	testMsg := []byte{'t', 'e', 's', 't'}
	hash := sha256.Sum256(testMsg)
	sig, err := YubiHsmSign(hash[:], yubiPrivKey)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("hash:%s, len:%d", hex.EncodeToString(hash[:]), len(hash))
	t.Logf("sig:%s, len:%d", hex.EncodeToString(sig), len(sig))
	t.Logf("pubkey: %s, len:%d", hex.EncodeToString(yubiPrivKey.GetPubKeyBytes()), len(yubiPrivKey.GetPubKeyBytes()))
	t.Logf("pubkey uncompressed: %s, len:%d", hex.EncodeToString(yubiPrivKey.pubKeyUncompressed), len(yubiPrivKey.pubKeyUncompressed))

	// try to recover pubkey
	recPubKeyBytes, err := secp256k1.RecoverPubkey(hash[:], sig)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("recovered pubkey:%s, len:%d", hex.EncodeToString(recPubKeyBytes), len(recPubKeyBytes))

	if !bytes.Equal(yubiPrivKey.pubKeyUncompressed, recPubKeyBytes) {
		t.Fatal("pubkey is mismatch")
	}

	sig1 := sig[:len(sig)-1]
	if !secp256k1.VerifySignature(yubiPrivKey.pubKeyBytes, hash[:], sig1) {
		t.Fatal("Verification of signature has failed")
	}
}
