package commands

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/loomnetwork/go-loom/builtin/types/dposv3"
	"github.com/loomnetwork/go-loom/cli"
	"github.com/loomnetwork/go-loom/types"
)

const DPOSV3ContractName = "dposV3"

var (
	candidateName        string
	candidateDescription string
	candidateWebsite     string
)

func UnregisterCandidateCmdV3() *cobra.Command {
	return &cobra.Command{
		Use:   "unregister_candidate_v3",
		Short: "Unregisters the candidate (only called if previously registered)",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cli.CallContract(DPOSV3ContractName, "UnregisterCandidate", &dposv3.UnregisterCandidateRequest{}, nil)
		},
	}
}

func GetStateCmdV3() *cobra.Command {
	return &cobra.Command{
		Use:   "get_dpos_state_v3",
		Short: "Gets dpos state",
		RunE: func(cmd *cobra.Command, args []string) error {
			var resp dposv3.GetStateResponse
			err := cli.StaticCallContract(DPOSV3ContractName, "GetState", &dposv3.GetStateRequest{}, &resp)
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

func ListValidatorsCmdV3() *cobra.Command {
	return &cobra.Command{
		Use:   "list_validators_v3",
		Short: "List the current validators",
		RunE: func(cmd *cobra.Command, args []string) error {
			var resp dposv3.ListValidatorsResponse
			err := cli.StaticCallContract(DPOSV3ContractName, "ListValidators", &dposv3.ListValidatorsRequest{}, &resp)
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

func ListCandidatesCmdV3() *cobra.Command {
	return &cobra.Command{
		Use:   "list_candidates_v3",
		Short: "List the registered candidates",
		RunE: func(cmd *cobra.Command, args []string) error {
			var resp dposv3.ListCandidatesResponse
			err := cli.StaticCallContract(DPOSV3ContractName, "ListCandidates", &dposv3.ListCandidatesRequest{}, &resp)
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

func ChangeFeeCmdV3() *cobra.Command {
	return &cobra.Command{
		Use:   "change_fee_v3 [new validator fee (in basis points)]",
		Short: "Changes a validator's fee after (with a 2 election delay)",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			candidateFee, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}
			if candidateFee > 10000 {
				errors.New("candidateFee is expressed in basis points (hundredths of a percent) and must be between 10000 (100%) and 0 (0%).")
			}
			return cli.CallContract(DPOSV3ContractName, "ChangeFee", &dposv3.ChangeCandidateFeeRequest{
				Fee: candidateFee,
			}, nil)
		},
	}
}

func RegisterCandidateCmdV3() *cobra.Command {
	return &cobra.Command{
		Use:   "register_candidate_v3 [public key] [validator fee (in basis points)] [locktime tier] [maximum referral percentage]",
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
			if len(args) >= 3 {
				tier, err = strconv.ParseUint(args[2], 10, 64)
				if err != nil {
					return err
				}

				if tier > 3 {
					errors.New("Tier value must be integer 0 - 4")
				}
			}
			maxReferralPercentage := uint64(0)
			if len(args) >= 4 {
				maxReferralPercentage, err = strconv.ParseUint(args[3], 10, 64)
				if err != nil {
					return err
				}
			}

			return cli.CallContract(DPOSV3ContractName, "RegisterCandidate", &dposv3.RegisterCandidateRequest{
				PubKey:                pubKey,
				Fee:                   candidateFee,
				Name:                  candidateName,
				Description:           candidateDescription,
				Website:               candidateWebsite,
				LocktimeTier:          tier,
				MaxReferralPercentage: maxReferralPercentage,
			}, nil)
		},
	}
}

func UpdateCandidateInfoCmdV3() *cobra.Command {
	return &cobra.Command{
		Use:   "update_candidate_info_v3 [name] [description] [website] [maximum referral percentage]",
		Short: "Update candidate information for a validator",
		Args:  cobra.MinimumNArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			candidateName := args[0]
			candidateDescription := args[1]
			candidateWebsite := args[2]
			maxReferralPercentage := uint64(0)
			if len(args) >= 4 {
				percentage, err := strconv.ParseUint(args[3], 10, 64)
				if err != nil {
					return err
				}
				if percentage > 10000 {
					errors.New("maxReferralFee is expressed in basis points (hundredths of a percent) and must be between 10000 (100%) and 0 (0%).")
				}
				maxReferralPercentage = percentage
			}

			return cli.CallContract(DPOSV3ContractName, "UpdateCandidateInfo", &dposv3.UpdateCandidateInfoRequest{
				Name:                  candidateName,
				Description:           candidateDescription,
				Website:               candidateWebsite,
				MaxReferralPercentage: maxReferralPercentage,
			}, nil)
		},
	}
}

func DelegateCmdV3() *cobra.Command {
	return &cobra.Command{
		Use:   "delegate_v3 [validator address] [amount] [locktime tier] [referrer]",
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

			var req dposv3.DelegateRequest
			req.Amount = &types.BigUInt{Value: *amount}
			req.ValidatorAddress = addr.MarshalPB()

			if len(args) >= 3 {
				tier, err := strconv.ParseUint(args[2], 10, 64)
				if err != nil {
					return err
				}

				if tier > 3 {
					errors.New("Tier value must be integer 0 - 4")
				}

				req.LocktimeTier = tier
			}

			if len(args) >= 4 {
				req.Referrer = args[3]
			}

			return cli.CallContract(DPOSV3ContractName, "Delegate", &req, nil)
		},
	}
}

func RedelegateCmdV3() *cobra.Command {
	return &cobra.Command{
		Use:   "redelegate_v3 [new validator address] [former validator address] [index] [amount] [referrer]",
		Short: "Redelegate tokens from one validator to another",
		Args:  cobra.MinimumNArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			validatorAddress, err := cli.ResolveAddress(args[0])
			if err != nil {
				return err
			}
			formerValidatorAddress, err := cli.ResolveAddress(args[1])
			if err != nil {
				return err
			}

			index, err := strconv.ParseUint(args[2], 10, 64)
			if err != nil {
				return err
			}

			var req dposv3.RedelegateRequest
			req.ValidatorAddress = validatorAddress.MarshalPB()
			req.FormerValidatorAddress = formerValidatorAddress.MarshalPB()
			req.Index = index

			if len(args) >= 4 {
				amount, err := cli.ParseAmount(args[3])
				if err != nil {
					return err
				}
				req.Amount = &types.BigUInt{Value: *amount}
			}

			if len(args) >= 5 {
				req.Referrer = args[4]
			}

			return cli.CallContract(DPOSV3ContractName, "Redelegate", &req, nil)
		},
	}
}

func WhitelistCandidateCmdV3() *cobra.Command {
	return &cobra.Command{
		Use:   "whitelist_candidate_v3 [candidate address] [amount] [locktime tier]",
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

			tier := uint64(0)
			if len(args) >= 3 {
				tier, err = strconv.ParseUint(args[2], 10, 64)
				if err != nil {
					return err
				}

				if tier > 3 {
					errors.New("Tier value must be integer 0 - 4")
				}
			}

			return cli.CallContract(DPOSV3ContractName, "WhitelistCandidate", &dposv3.WhitelistCandidateRequest{
				CandidateAddress: candidateAddress.MarshalPB(),
				Amount: &types.BigUInt{
					Value: *amount,
				},
				LocktimeTier: dposv3.LocktimeTier(tier),
			}, nil)
		},
	}
}

func RemoveWhitelistedCandidateCmdV3() *cobra.Command {
	return &cobra.Command{
		Use:   "remove_whitelisted_candidate_v3 [candidate address]",
		Short: "remove a candidate's whitelist entry",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			candidateAddress, err := cli.ResolveAddress(args[0])
			if err != nil {
				return err
			}

			return cli.CallContract(DPOSV3ContractName, "RemoveWhitelistedCandidate", &dposv3.RemoveWhitelistedCandidateRequest{
				CandidateAddress: candidateAddress.MarshalPB(),
			}, nil)
		},
	}
}

func ChangeWhitelistInfoCmdV3() *cobra.Command {
	return &cobra.Command{
		Use:   "change_whitelist_info_v3 [candidate address] [amount] [locktime tier]",
		Short: "Changes a whitelisted candidate's whitelist amount and tier",
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

			tier := uint64(0)
			if len(args) >= 3 {
				tier, err = strconv.ParseUint(args[2], 10, 64)
				if err != nil {
					return err
				}

				if tier > 3 {
					errors.New("Tier value must be integer 0 - 4")
				}
			}

			return cli.CallContract(DPOSV3ContractName, "ChangeWhitelistInfo", &dposv3.ChangeWhitelistInfoRequest{
				CandidateAddress: candidateAddress.MarshalPB(),
				Amount: &types.BigUInt{
					Value: *amount,
				},
				LocktimeTier: dposv3.LocktimeTier(tier),
			}, nil)
		},
	}
}

func CheckDelegationCmdV3() *cobra.Command {
	return &cobra.Command{
		Use:   "check_delegation_v3 [validator address] [delegator address]",
		Short: "check delegation to a particular validator",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			var resp dposv3.CheckDelegationResponse
			validatorAddress, err := cli.ParseAddress(args[0])
			if err != nil {
				return err
			}
			delegatorAddress, err := cli.ParseAddress(args[1])
			if err != nil {
				return err
			}
			err = cli.StaticCallContract(DPOSV3ContractName, "CheckDelegation", &dposv3.CheckDelegationRequest{ValidatorAddress: validatorAddress.MarshalPB(), DelegatorAddress: delegatorAddress.MarshalPB()}, &resp)
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

func UnbondCmdV3() *cobra.Command {
	return &cobra.Command{
		Use:   "unbond_v3 [validator address] [amount] [index]",
		Short: "De-allocate tokens from a validator",
		Args:  cobra.MinimumNArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			addr, err := cli.ResolveAddress(args[0])
			if err != nil {
				return err
			}

			amount, err := cli.ParseAmount(args[1])
			if err != nil {
				return err
			}

			index, err := strconv.ParseUint(args[2], 10, 64)
			if err != nil {
				return err
			}

			return cli.CallContract(DPOSV3ContractName, "Unbond", &dposv3.UnbondRequest{
				ValidatorAddress: addr.MarshalPB(),
				Amount: &types.BigUInt{
					Value: *amount,
				},
				Index: index,
			}, nil)
		},
	}
}

func CheckRewardsCmdV3() *cobra.Command {
	return &cobra.Command{
		Use:   "check_rewards_v3",
		Short: "check rewards statistics",
		Args:  cobra.MinimumNArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			var resp dposv3.CheckRewardsResponse
			err := cli.StaticCallContract(DPOSV3ContractName, "CheckRewards", &dposv3.CheckRewardsRequest{}, &resp)
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

func TotalDelegationCmdV3() *cobra.Command {
	return &cobra.Command{
		Use:   "total_delegation_v3 [delegator]",
		Short: "check how much a delegator has delegated in total (to all validators)",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			addr, err := cli.ResolveAddress(args[0])
			if err != nil {
				return err
			}

			var resp dposv3.TotalDelegationResponse
			err = cli.StaticCallContract(DPOSV3ContractName, "TotalDelegation", &dposv3.TotalDelegationRequest{DelegatorAddress: addr.MarshalPB()}, &resp)
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

func CheckAllDelegationsCmdV3() *cobra.Command {
	return &cobra.Command{
		Use:   "check_all_delegations_v3 [delegator]",
		Short: "display all of a particular delegator's delegations",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			addr, err := cli.ResolveAddress(args[0])
			if err != nil {
				return err
			}

			var resp dposv3.CheckAllDelegationsResponse
			err = cli.StaticCallContract(DPOSV3ContractName, "CheckAllDelegations", &dposv3.CheckAllDelegationsRequest{DelegatorAddress: addr.MarshalPB()}, &resp)
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

func TimeUntilElectionCmdV3() *cobra.Command {
	return &cobra.Command{
		Use:   "time_until_election_v3",
		Short: "check how many seconds remain until the next election",
		Args:  cobra.MinimumNArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			var resp dposv3.TimeUntilElectionResponse
			err := cli.StaticCallContract(DPOSV3ContractName, "TimeUntilElection", &dposv3.TimeUntilElectionRequest{}, &resp)
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

func ListDelegationsCmdV3() *cobra.Command {
	return &cobra.Command{
		Use:   "list_delegations_v3",
		Short: "list a candidate's delegations & delegation total",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			addr, err := cli.ResolveAddress(args[0])
			if err != nil {
				return err
			}

			var resp dposv3.ListDelegationsResponse
			err = cli.StaticCallContract(DPOSV3ContractName, "ListDelegations", &dposv3.ListDelegationsRequest{Candidate: addr.MarshalPB()}, &resp)
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

func ListAllDelegationsCmdV3() *cobra.Command {
	return &cobra.Command{
		Use:   "list_all_delegations_v3",
		Short: "display the results of calling list_delegations for all candidates",
		Args:  cobra.MinimumNArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			var resp dposv3.ListAllDelegationsResponse
			err := cli.StaticCallContract(DPOSV3ContractName, "ListAllDelegations", &dposv3.ListAllDelegationsRequest{}, &resp)
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

func SetElectionCycleCmdV3() *cobra.Command {
	return &cobra.Command{
		Use:   "set_election_cycle_v3 [election duration]",
		Short: "Set election cycle duration (in seconds)",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			electionCycleDuration, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			err = cli.CallContract(DPOSV3ContractName, "SetElectionCycle", &dposv3.SetElectionCycleRequest{
				ElectionCycle: int64(electionCycleDuration),
			}, nil)
			if err != nil {
				return err
			}
			return nil
		},
	}
}

func SetValidatorCountCmdV3() *cobra.Command {
	return &cobra.Command{
		Use:   "set_validator_count_v3 [validator count]",
		Short: "Set maximum number of validators",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			validatorCount, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			err = cli.CallContract(DPOSV3ContractName, "SetValidatorCount", &dposv3.SetValidatorCountRequest{
				ValidatorCount: int64(validatorCount),
			}, nil)
			if err != nil {
				return err
			}
			return nil
		},
	}
}

func SetMaxYearlyRewardCmdV3() *cobra.Command {
	return &cobra.Command{
		Use:   "set_max_yearly_reward_v3 [max yearly rewward amount]",
		Short: "Set maximum yearly reward",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			maxYearlyReward, err := cli.ParseAmount(args[0])
			if err != nil {
				return err
			}

			err = cli.CallContract(DPOSV3ContractName, "SetMaxYearlyReward", &dposv3.SetMaxYearlyRewardRequest{
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

func SetRegistrationRequirementCmdV3() *cobra.Command {
	return &cobra.Command{
		Use:   "set_registration_requirement_v3 [registration_requirement]",
		Short: "Set minimum self-delegation required of a new Candidate",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			registrationRequirement, err := cli.ParseAmount(args[0])
			if err != nil {
				return err
			}

			err = cli.CallContract(DPOSV3ContractName, "SetRegistrationRequirement", &dposv3.SetRegistrationRequirementRequest{
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

func SetOracleAddressCmdV3() *cobra.Command {
	return &cobra.Command{
		Use:   "set_oracle_address_v3 [oracle address]",
		Short: "Set oracle address",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			oracleAddress, err := cli.ParseAddress(args[0])
			if err != nil {
				return err
			}
			err = cli.CallContract(DPOSV3ContractName, "SetOracleAddress", &dposv3.SetOracleAddressRequest{OracleAddress: oracleAddress.MarshalPB()}, nil)
			if err != nil {
				return err
			}
			return nil
		},
	}
}

func SetSlashingPercentagesCmdV3() *cobra.Command {
	return &cobra.Command{
		Use:   "set_slashing_percentages_v3 [crash fault slashing percentage] [byzantine fault slashing percentage",
		Short: "Set crash and byzantine fualt slashing percentages expressed in basis points",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			registrationRequirement, err := cli.ParseAmount(args[1])
			if err != nil {
				return err
			}

			err = cli.CallContract(DPOSV3ContractName, "SetRegistrationRequirement", &dposv3.SetRegistrationRequirementRequest{
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

func AddDPOSV3(root *cobra.Command) {
	registercmd := RegisterCandidateCmdV3()
	registercmd.Flags().StringVarP(&candidateName, "name", "", "", "candidate name")
	registercmd.Flags().StringVarP(&candidateDescription, "description", "", "", "candidate description")
	registercmd.Flags().StringVarP(&candidateWebsite, "website", "", "", "candidate website")
	root.AddCommand(
		registercmd,
		ListCandidatesCmdV3(),
		ListValidatorsCmdV3(),
		ListDelegationsCmdV3(),
		ListAllDelegationsCmdV3(),
		UnregisterCandidateCmdV3(),
		UpdateCandidateInfoCmdV3(),
		DelegateCmdV3(),
		RedelegateCmdV3(),
		WhitelistCandidateCmdV3(),
		RemoveWhitelistedCandidateCmdV3(),
		ChangeWhitelistInfoCmdV3(),
		CheckDelegationCmdV3(),
		CheckAllDelegationsCmdV3(),
		CheckRewardsCmdV3(),
		UnbondCmdV3(),
		SetElectionCycleCmdV3(),
		SetValidatorCountCmdV3(),
		SetMaxYearlyRewardCmdV3(),
		SetRegistrationRequirementCmdV3(),
		SetOracleAddressCmdV3(),
		SetSlashingPercentagesCmdV3(),
		ChangeFeeCmdV3(),
		TimeUntilElectionCmdV3(),
		TotalDelegationCmdV3(),
		GetStateCmdV3(),
	)
}
