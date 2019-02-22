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

func GetStateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "get_dpos_state",
		Short: "Gets dpos state",
		RunE: func(cmd *cobra.Command, args []string) error {
			var resp dposv2.GetStateResponse
			err := cli.StaticCallContract(DPOSV2ContractName, "GetState", &dposv2.GetStateRequest{}, &resp)
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

func ChangeFeeCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "change_fee [new validator fee (in basis points)]",
		Short: "Changes a validator's fee after (with a 2 election delay)",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			candidateFee, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}
			if candidateFee > 10000 {
				errors.New("candidateFee is expressed in basis point (hundredths of a percent) and must be between 10000 (100%) and 0 (0%).")
			}
			return cli.CallContract(DPOSV2ContractName, "ChangeFee", &dposv2.ChangeCandidateFeeRequest{
				Fee: candidateFee,
			}, nil)
		},
	}
}

func RegisterCandidateCmdV2() *cobra.Command {
	return &cobra.Command{
		Use:   "register_candidateV2 [public key] [validator fee (in basis points)] [locktime tier]",
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

			tier := uint64(0)
			if len(args) == 3 {
				tier, err = strconv.ParseUint(args[2], 10, 64)
				if err != nil {
					return err
				}

				if tier > 3 {
					errors.New("Tier value must be integer 0 - 4")
				}
			}

			return cli.CallContract(DPOSV2ContractName, "RegisterCandidate2", &dposv2.RegisterCandidateRequestV2{
				PubKey:       pubKey,
				Fee:          candidateFee,
				Name:         candidateName,
				Description:  candidateDescription,
				Website:      candidateWebsite,
				LocktimeTier: tier,
			}, nil)
		},
	}
}

func UpdateCandidateInfoCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "update_candidate_info [name] [description] [website]",
		Short: "Update candidate information for a validator",
		Args:  cobra.MinimumNArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			candidateName := args[0]
			candidateDescription := args[1]
			candidateWebsite := args[2]

			return cli.CallContract(DPOSV2ContractName, "UpdateCandidateInfo", &dposv2.UpdateCandidateInfoRequest{
				Name:         candidateName,
				Description:  candidateDescription,
				Website:      candidateWebsite,
			}, nil)
		},
	}
}


func DelegateCmdV2() *cobra.Command {
	return &cobra.Command{
		Use:   "delegateV2 [validator address] [amount] [locktime tier]",
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

			var req dposv2.DelegateRequestV2
			req.Amount = &types.BigUInt{Value: *amount}
			req.ValidatorAddress = addr.MarshalPB()

			if len(args) == 3 {
				tier, err := strconv.ParseUint(args[2], 10, 64)
				if err != nil {
					return err
				}

				if tier > 3 {
					errors.New("Tier value must be integer 0 - 4")
				}

				req.LocktimeTier = tier
			}

			return cli.CallContract(DPOSV2ContractName, "Delegate2", &req, nil)
		},
	}
}

func RedelegateCmdV2() *cobra.Command {
	return &cobra.Command{
		Use:   "redelegateV2 [new validator address] [former validator address] [amount]",
		Short: "Redelegate tokens from one validator to another",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			validatorAddress, err := cli.ResolveAddress(args[0])
			if err != nil {
				return err
			}
			formerValidatorAddress, err := cli.ResolveAddress(args[1])
			if err != nil {
				return err
			}

			var req dposv2.RedelegateRequestV2
			req.ValidatorAddress = validatorAddress.MarshalPB()
			req.FormerValidatorAddress = formerValidatorAddress.MarshalPB()

			if len(args) == 3 {
				amount, err := cli.ParseAmount(args[2])
				if err != nil {
					return err
				}
				req.Amount = &types.BigUInt{Value: *amount}
			}

			return cli.CallContract(DPOSV2ContractName, "Redelegate", &req, nil)
		},
	}
}

func WhitelistCandidateCmdV2() *cobra.Command {
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

			return cli.CallContract(DPOSV2ContractName, "WhitelistCandidate", &dposv2.WhitelistCandidateRequestV2{
				CandidateAddress: candidateAddress.MarshalPB(),
				Amount: &types.BigUInt{
					Value: *amount,
				},
				LockTime: locktime,
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

func ChangeWhitelistAmountCmdV2() *cobra.Command {
	return &cobra.Command{
		Use:   "change_whitelist_amount [candidate address] [amount]",
		Short: "Changes a whitelisted candidate's whitelist amount",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			candidateAddress, err := cli.ResolveAddress(args[0])
			if err != nil {
				return err
			}
			amount, err := cli.ParseAmount(args[1])
			if err != nil {
				return err
			}
			return cli.CallContract(DPOSV2ContractName, "ChangeWhitelistAmount", &dposv2.ChangeWhitelistAmountRequestV2{
				CandidateAddress: candidateAddress.MarshalPB(),
				Amount: &types.BigUInt{
					Value: *amount,
				},
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

func CheckRewardsCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "check_rewards",
		Short: "check rewards statistics",
		Args:  cobra.MinimumNArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			var resp dposv2.CheckRewardsResponse
			err := cli.StaticCallContract(DPOSV2ContractName, "CheckRewards", &dposv2.CheckRewardsRequest{}, &resp)
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

func CheckDistributionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "check_distribution",
		Short: "check rewards distribution",
		Args:  cobra.MinimumNArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			var resp dposv2.CheckDistributionResponse
			err := cli.StaticCallContract(DPOSV2ContractName, "CheckDistribution", &dposv2.CheckDistributionRequest{}, &resp)
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

func TotalDelegationCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "total_delegation",
		Short: "check how much a delegator has delegated in total (to all validators)",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			addr, err := cli.ResolveAddress(args[0])
			if err != nil {
				return err
			}

			var resp dposv2.TotalDelegationResponse
			err = cli.StaticCallContract(DPOSV2ContractName, "TotalDelegation", &dposv2.TotalDelegationRequest{DelegatorAddress: addr.MarshalPB()}, &resp)
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

func TimeUntilElectionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "time_until_election",
		Short: "check how many seconds remain until the next election",
		Args:  cobra.MinimumNArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			var resp dposv2.TimeUntilElectionResponse
			err := cli.StaticCallContract(DPOSV2ContractName, "TimeUntilElection", &dposv2.TimeUntilElectionRequest{}, &resp)
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

func ListDelegationsCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list_delegations",
		Short: "list a candidate's delegations & delegation total",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			addr, err := cli.ResolveAddress(args[0])
			if err != nil {
				return err
			}

			var resp dposv2.ListDelegationsResponse
			err = cli.StaticCallContract(DPOSV2ContractName, "ListDelegations", &dposv2.ListDelegationsRequest{Candidate: addr.MarshalPB()}, &resp)
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

func ListAllDelegationsCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list_all_delegations",
		Short: "display the results of calling list_delegations for all candidates",
		Args:  cobra.MinimumNArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			var resp dposv2.ListAllDelegationsResponse
			err := cli.StaticCallContract(DPOSV2ContractName, "ListAllDelegations", &dposv2.ListAllDelegationsRequest{}, &resp)
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

// Oracle Commands for setting parameters

func SetElectionCycleCmdV2() *cobra.Command {
	return &cobra.Command{
		Use:   "set_election_cycle [election duration]",
		Short: "Set election cycle duration (in seconds)",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			electionCycleDuration, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			err = cli.CallContract(DPOSV2ContractName, "SetElectionCycle", &dposv2.SetElectionCycleRequestV2{
				ElectionCycle: int64(electionCycleDuration),
			}, nil)
			if err != nil {
				return err
			}
			return nil
		},
	}
}

func SetValidatorCountCmdV2() *cobra.Command {
	return &cobra.Command{
		Use:   "set_validator_count [validator count]",
		Short: "Set maximum number of validators",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			validatorCount, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			err = cli.CallContract(DPOSV2ContractName, "SetValidatorCount", &dposv2.SetValidatorCountRequestV2{
				ValidatorCount: int64(validatorCount),
			}, nil)
			if err != nil {
				return err
			}
			return nil
		},
	}
}

func SetMaxYearlyRewardCmdV2() *cobra.Command {
	return &cobra.Command{
		Use:   "set_max_yearly_reward [max yearly rewward amount]",
		Short: "Set maximum yearly reward",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			maxYearlyReward, err := cli.ParseAmount(args[0])
			if err != nil {
				return err
			}

			err = cli.CallContract(DPOSV2ContractName, "SetMaxYearlyReward", &dposv2.SetMaxYearlyRewardRequestV2{
				MaxYearlyReward: &types.BigUInt{
					Value: *maxYearlyReward,
				},
			}, nil)
			if err != nil {
				return err
			}
			return nil
		},
	}
}

func SetRegistrationRequirementCmdV2() *cobra.Command {
	return &cobra.Command{
		Use:   "set_registration_requirement [registration_requirement]",
		Short: "Set minimum self-delegation required of a new Candidate",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			registrationRequirement, err := cli.ParseAmount(args[0])
			if err != nil {
				return err
			}

			err = cli.CallContract(DPOSV2ContractName, "SetRegistrationRequirement", &dposv2.SetRegistrationRequirementRequestV2{
				RegistrationRequirement: &types.BigUInt{
					Value: *registrationRequirement,
				},
			}, nil)
			if err != nil {
				return err
			}
			return nil
		},
	}
}

func SetOracleAddressCmdV2() *cobra.Command {
	return &cobra.Command{
		Use:   "set_oracle_address [oracle address]",
		Short: "Set oracle address",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			oracleAddress, err := cli.ParseAddress(args[0])
			if err != nil {
				return err
			}
			err = cli.CallContract(DPOSV2ContractName, "SetOracleAddress", &dposv2.SetOracleAddressRequestV2{OracleAddress: oracleAddress.MarshalPB()}, nil)
			if err != nil {
				return err
			}
			return nil
		},
	}
}

func SetSlashingPercentagesCmdV2() *cobra.Command {
	return &cobra.Command{
		Use:   "set_slashing_percentages [crash fault slashing percentage] [byzantine fault slashing percentage",
		Short: "Set crash and byzantine fualt slashing percentages expressed in basis points",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			registrationRequirement, err := cli.ParseAmount(args[1])
			if err != nil {
				return err
			}

			err = cli.CallContract(DPOSV2ContractName, "SetRegistrationRequirement", &dposv2.SetRegistrationRequirementRequestV2{
				RegistrationRequirement: &types.BigUInt{
					Value: *registrationRequirement,
				},
			}, nil)
			if err != nil {
				return err
			}
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
		registercmd,
		ListCandidatesCmdV2(),
		ListValidatorsCmdV2(),
		ListDelegationsCmd(),
		ListAllDelegationsCmd(),
		UnregisterCandidateCmdV2(),
		UpdateCandidateInfoCmd(),
		DelegateCmdV2(),
		RedelegateCmdV2(),
		WhitelistCandidateCmdV2(),
		RemoveWhitelistedCandidateCmdV2(),
		ChangeWhitelistAmountCmdV2(),
		CheckDelegationCmdV2(),
		CheckDistributionCmd(),
		CheckRewardsCmd(),
		UnbondCmdV2(),
		ClaimDistributionCmdV2(),
		SetElectionCycleCmdV2(),
		SetValidatorCountCmdV2(),
		SetMaxYearlyRewardCmdV2(),
		SetRegistrationRequirementCmdV2(),
		SetOracleAddressCmdV2(),
		SetSlashingPercentagesCmdV2(),
		ChangeFeeCmd(),
		TimeUntilElectionCmd(),
		TotalDelegationCmd(),
		GetStateCmd(),
	)
}
