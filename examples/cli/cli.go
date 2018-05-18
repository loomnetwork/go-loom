package main

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/loomnetwork/go-loom/builtin/commands"
	"github.com/loomnetwork/go-loom/cli"
	"github.com/loomnetwork/go-loom/examples/types"
)

var writeURI, readURI, chainID string

func main() {
	rootCmd := &cobra.Command{
		Use:   "cli",
		Short: "CLI example",
	}
	callCmd := cli.ContractCallCommand()
	rootCmd.AddCommand(callCmd)
	commands.Add(callCmd)

	defaultContract := "helloworld"
	var key, value string

	setMsgCmd := &cobra.Command{
		Use:   "set_msg",
		Short: "Calls the SetMsg method of the helloworld contract",
		RunE: func(cmd *cobra.Command, args []string) error {
			params := &types.MapEntry{
				Key:   key,
				Value: value,
			}

			return cli.CallContract(defaultContract, "SetMsg", params, nil)
		},
	}
	setMsgCmd.Flags().StringVarP(&key, "key", "k", "", "")
	setMsgCmd.Flags().StringVarP(&value, "value", "v", "", "value to associate with the key")
	callCmd.AddCommand(setMsgCmd)

	getMsgCmd := &cobra.Command{
		Use:   "get_msg",
		Short: "Calls the GetMsg method of the helloworld contract",
		RunE: func(cmd *cobra.Command, args []string) error {
			params := &types.MapEntry{
				Key: key,
			}
			var result types.MapEntry
			err := cli.StaticCallContract(defaultContract, "GetMsg", params, &result)
			if err != nil {
				return err
			}
			log.Printf("{ key: \"%s\", value: \"%s\" }\n", result.Key, result.Value)
			return nil
		},
	}
	getMsgCmd.Flags().StringVarP(&key, "key", "k", "", "")
	callCmd.AddCommand(getMsgCmd)
	rootCmd.Execute()
}
