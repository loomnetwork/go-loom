package commands

import (
	`fmt`
	`github.com/loomnetwork/go-loom/builtin/types/karma`
	`github.com/loomnetwork/go-loom/cli`
	"github.com/pkg/errors"
	`github.com/spf13/cobra`
	`strconv`
)

const KarmaContractName = "karma"

func UpdateSourcesForUserCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "update [user] [oracle] [name] [count]",
		Short: "Update sources for a user",
		Args:  cobra.MinimumNArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			user, err := cli.ResolveAddress(args[0])
			if err != nil {
				return errors.Wrap(err, "resolve urser address")
			}
			oracle, err := cli.ResolveAddress(args[1])
			if err != nil {
				return errors.Wrap(err, "resolve oracle address")
			}
			name := args[2]
			count, err := strconv.ParseInt(args[3], 10, 64)
			if err != nil {
				return errors.Wrap(err, "parse count")
			}
			return cli.CallContract(KarmaContractName, "UpdateSourcesForUser", &karma.KarmaStateUser{
				User: user.MarshalPB(),
				Oracle: oracle.MarshalPB(),
				SourceStates: []*karma.KarmaSource{{name, count}},
			}, nil)
		},
	}
}

func GetConfig() *cobra.Command {
	return &cobra.Command{
		Use:   "getconfig [user]",
		Short: "Get karma config parameters",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			oracle, err := cli.ResolveAddress(args[0])
			if err != nil {
				return errors.Wrap(err, "resolve oracle address")
			}
			var resp karma.KarmaConfig
			if err := cli.StaticCallContract(KarmaContractName, "GetConfig", oracle.MarshalPB(), &resp); err !=  nil {
				return errors.Wrap(err, "call contract, method: GetConfig")
			}
			out, err := formatJSON(&resp)
			if err != nil {
				return errors.Wrap(err, "format JSON returned by karma GetConfig")
			}
			fmt.Println(out)
			return nil
		},
	}
}

func UpdateConfig() *cobra.Command {
	cmdFlags := struct {
		Enabled bool
		SessionMaxAccessCount int64
		SessionDuration int64
		DeployEnabled bool
		CallEnabled bool
	} {}
	cmdFlags = cmdFlags
	cmd := &cobra.Command{
		Use:   "updateconfig [oracle]",
		Short: "Update karma config parameters",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			oracle, err := cli.ResolveAddress(args[0])
			if err != nil {
				return errors.Wrap(err, "resolve oracle address")
			}
			var kConfig karma.KarmaConfigValidator
			kConfig.Enabled = cmdFlags.Enabled
			kConfig.SessionMaxAccessCount = cmdFlags.SessionMaxAccessCount
			kConfig.SessionDuration = cmdFlags.SessionDuration
			kConfig.DeployEnabled = cmdFlags.DeployEnabled
			kConfig.CallEnabled = cmdFlags.CallEnabled
			kConfig.Oracle = oracle.MarshalPB()
			return cli.CallContract(KarmaContractName, "UpdateConfig", &kConfig, nil)
		},
	}
	cmd.Flags().BoolVar(&cmdFlags.Enabled, "enabled", true,  "enable karma")
	cmd.Flags().Int64VarP(&cmdFlags.SessionMaxAccessCount, "contract-name", "c", 0, "maxium Tx per session")
	cmd.Flags().Int64VarP(&cmdFlags.SessionDuration, "input", "d", 0, "session duration")
	cmd.Flags().BoolVar(&cmdFlags.DeployEnabled, "deploy-enable", true , "contract address")
	cmd.Flags().BoolVar(&cmdFlags.CallEnabled, "call-enable", true,  "contract address")
	
	return cmd
}

func AddKarma(root *cobra.Command) {
	root.AddCommand(
		UpdateSourcesForUserCmd(),
		GetConfig(),
		UpdateConfig(),
	)
}
