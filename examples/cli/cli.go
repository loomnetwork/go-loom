package main

import (
	"log"

	loom "github.com/loomnetwork/go-loom"
	"github.com/loomnetwork/go-loom/auth"
	"github.com/loomnetwork/go-loom/client"
	"github.com/loomnetwork/go-loom/examples/types"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ed25519"
)

func main() {
	var writeURI, readURI, contractHexAddr, chainID string
	rootCmd := &cobra.Command{
		Use:   "cli",
		Short: "CLI example",
	}
	rootCmd.PersistentFlags().StringVarP(&writeURI, "write", "w", "http://localhost:46657", "URI for sending txs")
	rootCmd.PersistentFlags().StringVarP(&readURI, "read", "r", "http://localhost:47000", "URI for quering app state")
	rootCmd.PersistentFlags().StringVarP(&contractHexAddr, "contract", "", "0x005B17864f3adbF53b1384F2E6f2120c6652F779", "contract address")
	rootCmd.PersistentFlags().StringVarP(&chainID, "chain", "", "default", "chain ID")

	rpcClient := client.NewDAppChainRPCClient(chainID, writeURI, readURI)

	setMsgCmd := &cobra.Command{
		Use:   "set",
		Short: "Set message stored in the Helloworld contract",
		RunE: func(cmd *cobra.Command, args []string) error {
			contractAddr, err := loom.LocalAddressFromHexString(contractHexAddr)
			if err != nil {
				return err
			}
			contract := client.NewContract(rpcClient, contractAddr, "helloworld")
			// NOTE: usually you shouldn't generate a new key pair for every tx, but this is just an example...
			_, priv, err := ed25519.GenerateKey(nil)
			if err != nil {
				return err
			}
			signer := auth.NewEd25519Signer(priv)
			payload := &types.Dummy{
				Key:   "123",
				Value: "thefoxjumped",
			}
			if _, err := contract.Call("SetMsg", payload, signer, nil); err != nil {
				return err
			}
			return nil
		},
	}
	getMsgCmd := &cobra.Command{
		Use:   "get",
		Short: "Get message stored in the Helloworld contract",
		RunE: func(cmd *cobra.Command, args []string) error {
			contractAddr, err := loom.LocalAddressFromHexString(contractHexAddr)
			if err != nil {
				return err
			}
			contract := client.NewContract(rpcClient, contractAddr, "helloworld")
			query := &types.Dummy{
				Key: "123",
			}
			var result types.Dummy
			if _, err := contract.StaticCall("GetMsg", query, &result); err != nil {
				return err
			}
			log.Printf("{ Key: '%s', Value: '%s' }", result.Key, result.Value)
			return nil
		},
	}
	rootCmd.AddCommand(setMsgCmd)
	rootCmd.AddCommand(getMsgCmd)
	rootCmd.Execute()
}
