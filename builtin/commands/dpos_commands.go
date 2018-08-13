package commands

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/loomnetwork/go-loom/builtin/types/dpos"
	"github.com/loomnetwork/go-loom/cli"
)

const DPOSContractName = "dpos"

func ListWitnessesCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list_witnesses",
		Short: "List the current witnesses",
		RunE: func(cmd *cobra.Command, args []string) error {
			var resp dpos.ListWitnessesResponse
			err := cli.StaticCallContract(DPOSContractName, "ListWitnesses", &dpos.ListWitnessesRequest{}, &resp)
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

func ListCandidatesCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list_candidates",
		Short: "List the registered candidates",
		RunE: func(cmd *cobra.Command, args []string) error {
			var resp dpos.ListCandidateResponse
			err := cli.StaticCallContract(DPOSContractName, "ListCandidates", &dpos.ListCandidateRequest{}, &resp)
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

func RegisterCandidateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "register_candidate [public key]",
		Short: "Register a candidate for witness",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			pubKey, err := cli.ParseBytes(args[0])
			if err != nil {
				return err
			}
			return cli.CallContract(DPOSContractName, "RegisterCandidate", &dpos.RegisterCandidateRequest{
				PubKey: pubKey,
			}, nil)
		},
	}
}

func VoteCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "vote [candidate address] [amount]",
		Short: "Allocate votes to a candidate",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			addr, err := cli.ResolveAddress(args[0])
			if err != nil {
				return err
			}

			amount, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				return err
			}

			return cli.CallContract(DPOSContractName, "Vote", &dpos.VoteRequest{
				CandidateAddress: addr.MarshalPB(),
				Amount:           amount,
			}, nil)
		},
	}
}

func ElectCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "elect",
		Short: "Run an election",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cli.CallContract(DPOSContractName, "Elect", &dpos.ElectRequest{}, nil)
		},
	}
}

func AddDPOS(root *cobra.Command) {
	root.AddCommand(
		ListWitnessesCmd(),
		RegisterCandidateCmd(),
		VoteCmd(),
		ElectCmd(),
		ListCandidatesCmd(),
	)
}
