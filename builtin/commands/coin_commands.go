package commands

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/loomnetwork/go-loom/builtin/types/coin"
	"github.com/loomnetwork/go-loom/cli"
	"github.com/loomnetwork/go-loom/types"
)

const CoinContractName = "coin"

func TransferCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "transfer [address] [amount]",
		Short: "Transfer coins to another account",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			addr, err := cli.ResolveAddress(args[0])
			if err != nil {
				return err
			}

			amount, err := cli.ParseAmount(args[1])
			if err != nil {
				return err
			}
			return cli.CallContract(CoinContractName, "Transfer", &coin.TransferRequest{
				To: addr.MarshalPB(),
				Amount: &types.BigUInt{
					Value: *amount,
				},
			}, nil)
		},
	}
}

func BalanceCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "balance [address]",
		Short: "Fetch the balance of a coin account",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			addr, err := cli.ResolveAddress(args[0])
			if err != nil {
				return err
			}
			var resp coin.BalanceOfResponse
			err = cli.StaticCallContract(CoinContractName, "BalanceOf", &coin.BalanceOfRequest{
				Owner: addr.MarshalPB(),
			}, &resp)
			if err != nil {
				return err
			}
			out, err := formatJSON(&resp)
			if err != nil {
				return err
			}
			fmt.Println(out)
			return nil
		},
	}
}

func AddCoin(root *cobra.Command) {
	root.AddCommand(
		BalanceCmd(),
		TransferCmd(),
	)
}
