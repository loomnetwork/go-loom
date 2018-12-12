// +build evm

package crypto

import (
	"crypto/sha256"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

func TestGenYubiSecp256k1Key(t *testing.T) {
	// check if yubiHsm config is given
	cfg := os.Getenv("YUBIHSM_CFG_FILE")
	if len(cfg) == 0 {
		t.Log("YubiHsm crypto testing disabled")
		return
	}

	t.Log("Generating YubiHSM private key")
	yubiPrivKey, err := GenYubiHsmPrivKey("ed25519", cfg)
	if err != nil {
		t.Fatal(err)
	}
	defer yubiPrivKey.UnloadYubiHsmPrivKey()

	yubiPrivKey.SaveYubiHsmPrivKey(cfg)
}

func TestSignYubiSecp256k1(t *testing.T) {
	// check if yubiHsm config is given
	cfg := os.Getenv("YUBIHSM_CFG_FILE")
	if len(cfg) == 0 {
		t.Log("YubiHsm crypto testing disabled")
		return
	}

	t.Log("Loading YubiHSM private key")
	yubiPrivKey, err := LoadYubiHsmPrivKey("ed25519", cfg)
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

	t.Logf("Hash length is %d", len(hash))
	if !secp256k1.VerifySignature(yubiPrivKey.pubKeyBytes[:], hash[:], sig) {
		t.Fatalf("Verification of signature has failed")
	}

	t.Logf("Sign/Verify using YubiHSM has been succeeded")
}
