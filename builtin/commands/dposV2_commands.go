package commands

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/loomnetwork/go-loom/builtin/types/dposv2"
	"github.com/loomnetwork/go-loom/cli"
	"github.com/loomnetwork/go-loom/types"
)

const DPOSV2ContractName = "dposV2"

var (
	candidateName        string
	candidateDescription string
	candidateWebsite     string
)

func UnregisterCandidateCmdV2() *cobra.Command {
	return &cobra.Command{
		Use:   "unregister_candidateV2",
		Short: "Unregisters the candidate (only called if previously registered)",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cli.CallContract(DPOSV2ContractName, "UnregisterCandidate", &dposv2.UnregisterCandidateRequestV2{}, nil)
		},
	}
}

func ListValidatorsCmdV2() *cobra.Command {
	return &cobra.Command{
		Use:   "list_validatorsV2",
		Short: "List the current validators",
		RunE: func(cmd *cobra.Command, args []string) error {
			var resp dposv2.ListValidatorsResponseV2
			err := cli.StaticCallContract(DPOSV2ContractName, "ListValidators", &dposv2.ListValidatorsRequestV2{}, &resp)
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

func ListCandidatesCmdV2() *cobra.Command {
	return &cobra.Command{
		Use:   "list_candidatesV2",
		Short: "List the registered candidates",
		RunE: func(cmd *cobra.Command, args []string) error {
			var resp dposv2.ListCandidateResponseV2
			err := cli.StaticCallContract(DPOSV2ContractName, "ListCandidates", &dposv2.ListCandidateRequestV2{}, &resp)
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

func RegisterCandidateCmdV2() *cobra.Command {
	return &cobra.Command{
		Use:   "register_candidateV2 [public key] [validator fee (in basis points)]",
		Short: "Register a candidate for validator",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			pubKey, err := base64.StdEncoding.DecodeString(args[0])
			candidateFee, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}
			if candidateFee > 10000 {
				errors.New("candidateFee is expressed in basis point (hundredths of a percent) and must be between 10000 (100%) and 0 (0%).")
			}
			return cli.CallContract(DPOSV2ContractName, "RegisterCandidate", &dposv2.RegisterCandidateRequestV2{
				PubKey:      pubKey,
				Fee:         candidateFee,
				Name:        candidateName,
				Description: candidateDescription,
				Website:     candidateWebsite,
			}, nil)
		},
	}
}

func DelegateCmdV2() *cobra.Command {
	return &cobra.Command{
		Use:   "delegateV2 [validator address] [amount]",
		Short: "delegate tokens to a validator",
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

			return cli.CallContract(DPOSV2ContractName, "Delegate", &dposv2.DelegateRequestV2{
				ValidatorAddress: addr.MarshalPB(),
				Amount: &types.BigUInt{
					Value: *amount,
				},
			}, nil)
		},
	}
}

func WhitelistCandidateCmdV2() *cobra.Command {
	// Keep increamenting blocknumber to make sure ProcessRequestBatch is successful
	var currentBlockNumber uint64
	return &cobra.Command{
		Use:   "whitelist_candidate [candidate address] [amount] [lock time]",
		Short: "Whitelist candidate & credit candidate's self delegation without token deposit",
		Args:  cobra.MinimumNArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			candidateAddress, err := cli.ResolveAddress(args[0])
			if err != nil {
				return err
			}
			amount, err := cli.ParseAmount(args[1])
			if err != nil {
				return err
			}
			locktime, err := strconv.ParseUint(args[2], 10, 64)
			if err != nil {
				return err
			}

			currentBlockNumber++

			return cli.CallContract(DPOSV2ContractName, "ProcessRequestBatch", &dposv2.RequestBatchV2{
				Batch: []*dposv2.BatchRequestV2{
					&dposv2.BatchRequestV2{
						Payload: &dposv2.BatchRequestV2_WhitelistCandidate{
							&dposv2.WhitelistCandidateRequestV2{
								CandidateAddress: candidateAddress.MarshalPB(),
								Amount: &types.BigUInt{
									Value: *amount,
								},
								LockTime: locktime,
							},
						},
						Meta: &dposv2.BatchRequestMetaV2{
							BlockNumber: currentBlockNumber,
							LogIndex:    0,
							TxIndex:     0,
						},
					},
				},
			}, nil)
		},
	}
}

func RemoveWhitelistedCandidateCmdV2() *cobra.Command {
	return &cobra.Command{
		Use:   "remove_whitelisted_candidate [candidate address]",
		Short: "remove a candidate's whitelist entry",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			candidateAddress, err := cli.ResolveAddress(args[0])
			if err != nil {
				return err
			}

			return cli.CallContract(DPOSV2ContractName, "RemoveWhitelistedCandidate", &dposv2.RemoveWhitelistedCandidateRequestV2{
				CandidateAddress: candidateAddress.MarshalPB(),
			}, nil)
		},
	}
}

func CheckDelegationCmdV2() *cobra.Command {
	return &cobra.Command{
		Use:   "check_delegationV2 [validator address] [delegator address]",
		Short: "check delegation to a particular validator",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			var resp dposv2.CheckDelegationResponseV2
			validatorAddress, err := cli.ParseAddress(args[0])
			if err != nil {
				return err
			}
			delegatorAddress, err := cli.ParseAddress(args[1])
			if err != nil {
				return err
			}
			err = cli.StaticCallContract(DPOSV2ContractName, "CheckDelegation", &dposv2.CheckDelegationRequestV2{ValidatorAddress: validatorAddress.MarshalPB(), DelegatorAddress: delegatorAddress.MarshalPB()}, &resp)
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

func UnbondCmdV2() *cobra.Command {
	return &cobra.Command{
		Use:   "unbondV2 [validator address] [amount]",
		Short: "De-allocate tokens from a validator",
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
			return cli.CallContract(DPOSV2ContractName, "Unbond", &dposv2.UnbondRequestV2{
				ValidatorAddress: addr.MarshalPB(),
				Amount: &types.BigUInt{
					Value: *amount,
				},
			}, nil)
		},
	}
}

func ClaimDistributionCmdV2() *cobra.Command {
	return &cobra.Command{
		Use:   "claim_distributionV2 [withdrawal address]",
		Short: "claim dpos distributions due to a validator or delegator",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			addr, err := cli.ResolveAddress(args[0])
			if err != nil {
				return err
			}

			var resp dposv2.ClaimDistributionResponseV2
			err = cli.CallContract(DPOSV2ContractName, "ClaimDistribution", &dposv2.ClaimDistributionRequestV2{
				WithdrawalAddress: addr.MarshalPB(),
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

func AddDPOSV2(root *cobra.Command) {
	registercmd := RegisterCandidateCmdV2()
	registercmd.Flags().StringVarP(&candidateName, "name", "", "", "candidate name")
	registercmd.Flags().StringVarP(&candidateDescription, "description", "", "", "candidate description")
	registercmd.Flags().StringVarP(&candidateWebsite, "website", "", "", "candidate website")
	root.AddCommand(
		ListValidatorsCmdV2(),
		registercmd,
		ListCandidatesCmdV2(),
		UnregisterCandidateCmdV2(),
		DelegateCmdV2(),
		WhitelistCandidateCmdV2(),
		RemoveWhitelistedCandidateCmdV2(),
		CheckDelegationCmdV2(),
		UnbondCmdV2(),
		ClaimDistributionCmdV2(),
	)
}
