package commands

import (
	"errors"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	tgtypes "github.com/loomnetwork/go-loom/builtin/types/transfer_gateway"
	"github.com/loomnetwork/go-loom/cli"
)

const GatewayName = "gateway"
const LoomGatewayName = "loomcoin-gateway"

func UnsafeResetBlockCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "unsafe_reset_block [blockNumber] [gatewayType]",
		Short: "Resets the block mainnet",
		Args:  cobra.MinimumNArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			var block uint64
			var err error
			if len(args) == 0 {
				block = uint64(0)
			} else {
				block, err = strconv.ParseUint(args[0], 10, 64)
				if err != nil {
					return err
				}
			}

			var name string
			if len(args) <= 1 || (strings.Compare(args[1], GatewayName) == 0) {
				name = GatewayName
			} else if strings.Compare(args[1], LoomGatewayName) == 0 {
				name = LoomGatewayName
			} else {
				errors.New("Invalid gateway name")
			}

			return cli.CallContract(name, "ResetMainnetBlock", &tgtypes.TransferGatewayResetMainnetBlockRequest{
				LastMainnetBlockNum: block,
			}, nil)
		},
	}
}

func AddUnsafeCommands(root *cobra.Command) {
	root.AddCommand(
		UnsafeResetBlockCmd(),
	)
}
