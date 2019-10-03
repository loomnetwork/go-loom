package commands

import (
	"encoding/base64"
	"fmt"

	loom "github.com/loomnetwork/go-loom"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func newAddressToB64Command() *cobra.Command {
	return &cobra.Command{
		Use:     "addr-to-b64",
		Short:   "convert hexstring address to base 64 address",
		Example: "loom resolve addr-to-b64 0x9F5137fF296469cdc3D137273fF9A4Df76044758",
		Args:    cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			addr, err := loom.LocalAddressFromHexString(args[0])
			if err != nil {
				return errors.Wrap(err, "invalid account address")
			}

			fmt.Printf("Address in base64: %s\n", base64.StdEncoding.EncodeToString([]byte(addr)))
			return nil
		},
	}
}

func newPubkeyToAddressCommand() *cobra.Command {
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
		newPubkeyToAddressCommand(),
		newAddressToB64Command(),
	)
}
