package commands

import (
	`github.com/loomnetwork/go-loom/builtin/types/karma`
	`github.com/loomnetwork/go-loom/cli`
	`github.com/spf13/cobra`
	`strconv`
)

const KarmaContractName = "karma"

func UpdateSourcesForUserCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "update [user] [oracle] [name] [count]",
		Short: "Update sources for a user",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			user, err := cli.ResolveAddress(args[0])
			if err != nil {
				return err
			}
			oracle, err := cli.ResolveAddress(args[1])
			if err != nil {
				return err
			}
			name := args[2]
			count, err := strconv.ParseInt(args[3], 10, 64)
			if err != nil {
				return err
			}
			return cli.CallContract(KarmaContractName, "UpdateSourcesForUser", &karma.KarmaStateUser{
				User: user.MarshalPB(),
				Oracle: oracle.MarshalPB(),
				SourceStates: []*karma.KarmaSource{{name, count}},
			}, nil)
		},
	}
}

func AddKarma(root *cobra.Command) {
	root.AddCommand(
		UpdateSourcesForUserCmd(),
	)
}
