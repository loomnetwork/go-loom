package main

import (
	"fmt"
	"os"

	lcrypto "github.com/loomnetwork/go-loom/crypto"
	"github.com/spf13/cobra"
)

const (
	YUBIHSM_DEFAULT_CFG_FILE = "yubihsm.cfg"
	YUBIHSM_DEFAULT_KEY_TYPE = lcrypto.PrivateKeyTypeEd25519
)

var RootCmd = &cobra.Command{
	Use:   "yubihsm-key-util",
	Short: "YubiHSM key management utility",
}

type yubiKeyUtilFlags struct {
	ConfigFile string
	KeyType    string
}

var yubiKeyUtilCmdFlags yubiKeyUtilFlags

func gen_yubihsm_key() error {
	yubiPrivKey, err := lcrypto.GenYubiHsmPrivKey(yubiKeyUtilCmdFlags.ConfigFile)
	if err != nil {
		return err
	}
	defer yubiPrivKey.UnloadYubiHsmPrivKey()

	// print public key and address
	pubKeyAddr := yubiPrivKey.GetPubKeyAddr()

	fmt.Printf("Private Key Type:   %s\n", yubiPrivKey.GetKeyType())
	fmt.Printf("Private Key ID:     %d\n", yubiPrivKey.GetPrivKeyID())
	fmt.Printf("Public Key address: %s\n", pubKeyAddr)

	return nil
}

func load_yubihsm_key() error {
	yubiPrivKey, err := lcrypto.LoadYubiHsmPrivKey(yubiKeyUtilCmdFlags.ConfigFile)
	if err != nil {
		return err
	}
	defer yubiPrivKey.UnloadYubiHsmPrivKey()

	// print public key and address
	pubKeyAddr := yubiPrivKey.GetPubKeyAddr()

	fmt.Printf("Private Key Type:   %s\n", yubiPrivKey.GetKeyType())
	fmt.Printf("Private Key ID:     %d\n", yubiPrivKey.GetPrivKeyID())
	fmt.Printf("Public Key address: %s\n", pubKeyAddr)

	return nil
}

func genKeyCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "genkey",
		Short: "Generate ed25519/secp256k1 asymmetric key",
		RunE: func(cmd *cobra.Command, args []string) error {
			return gen_yubihsm_key()
		},
	}
}

func loadKeyCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "loadkey",
		Short: "Load ed25519/secp256k1 asymmetric key",
		RunE: func(cmd *cobra.Command, args []string) error {
			return load_yubihsm_key()
		},
	}
}

func main() {
	RootCmd.AddCommand(
		genKeyCommand(),
		loadKeyCommand(),
	)
	pflags := RootCmd.PersistentFlags()
	pflags.StringVarP(&yubiKeyUtilCmdFlags.ConfigFile, "config", "c", YUBIHSM_DEFAULT_CFG_FILE, "YubiHSM config file")

	err := RootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
