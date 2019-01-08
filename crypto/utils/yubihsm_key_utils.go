package main

import (
	"fmt"

	lcrypto "github.com/loomnetwork/go-loom/crypto"
)

const (
	YUBIHSM_CFG_FILE = "yubihsm.cfg"
)

func gen_yubihsm_secp256k1() {
	yubiPrivKey, err := lcrypto.GenYubiHsmPrivKey(YUBIHSM_CFG_FILE)
	if err != nil {
		panic(err)
	}
	defer yubiPrivKey.UnloadYubiHsmPrivKey()

	fmt.Printf("Private Key ID: %v\n", yubiPrivKey.GetPrivKeyID())

	// print public key and address
	pubKeyAddr := yubiPrivKey.GetPubKeyAddr()

	fmt.Printf("Public Key address: %v\n", pubKeyAddr)
}

func main() {
	// generate yubihsm secp256 key and print details
	gen_yubihsm_secp256k1()
}
