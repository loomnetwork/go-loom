package commands

import (
	"encoding/base64"
	"fmt"

	loom "github.com/loomnetwork/go-loom"
	"github.com/spf13/cobra"
)

func PubkeyToAddress() *cobra.Command {
	return &cobra.Command{
		Use:   "pubkey [public key]",
		Short: "Converts a public key to an address",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			pubKey, err := base64.StdEncoding.DecodeString(args[0])
			if err != nil {
				return err
			}
			fmt.Println("Address", loom.LocalAddressFromPublicKey(pubKey))
			return nil
		},
	}
}

func AddGeneralCommands(root *cobra.Command) {
	root.AddCommand(
		PubkeyToAddress(),
	)
}
