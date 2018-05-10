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

var writeURI, readURI, chainID string

func getContract(contractHexAddr, contractName string) (*client.Contract, error) {
	rpcClient := client.NewDAppChainRPCClient(chainID, writeURI, readURI)
	contractAddr, err := loom.LocalAddressFromHexString(contractHexAddr)
	if err != nil {
		return nil, err
	}
	return client.NewContract(rpcClient, contractAddr, contractName), nil
}

func main() {
	var contractHexAddr, contractName, methodName string
	rootCmd := &cobra.Command{
		Use:   "cli",
		Short: "CLI example",
	}
	rootCmd.PersistentFlags().StringVarP(&writeURI, "write", "w", "http://localhost:46657", "URI for sending txs")
	rootCmd.PersistentFlags().StringVarP(&readURI, "read", "r", "http://localhost:47000", "URI for quering app state")
	rootCmd.PersistentFlags().StringVarP(&contractHexAddr, "contract", "", "0x005B17864f3adbF53b1384F2E6f2120c6652F779", "contract address")
	rootCmd.PersistentFlags().StringVarP(&contractName, "name", "n", "helloworld", "smart contract name")
	rootCmd.PersistentFlags().StringVarP(&chainID, "chain", "", "default", "chain ID")
	rootCmd.PersistentFlags().StringVarP(&methodName, "method", "m", "", "smart contract method name")

	var key, value string

	callCmd := &cobra.Command{
		Use:   "call",
		Short: "Calls a method on a smart contract",
		RunE: func(cmd *cobra.Command, args []string) error {
			// NOTE: usually you shouldn't generate a new key pair for every tx, but this is just an example...
			_, privateKey, err := ed25519.GenerateKey(nil)
			if err != nil {
				return err
			}
			signer := auth.NewEd25519Signer(privateKey)
			contract, err := getContract(contractHexAddr, contractName)
			if err != nil {
				return err
			}
			params := &types.Dummy{
				Key:   key,
				Value: value,
			}
			if _, err := contract.Call(methodName, params, signer, nil); err != nil {
				return err
			}
			return nil
		},
	}
	callCmd.Flags().StringVarP(&key, "key", "k", "", "")
	callCmd.Flags().StringVarP(&value, "value", "v", "", "value to associate with the key")

	staticCallCmd := &cobra.Command{
		Use:   "static-call",
		Short: "Calls a read-only method on a smart contract",
		RunE: func(cmd *cobra.Command, args []string) error {
			contract, err := getContract(contractHexAddr, contractName)
			if err != nil {
				return err
			}
			params := &types.Dummy{
				Key: key,
			}
			var result types.Dummy
			if _, err := contract.StaticCall(methodName, params, &result); err != nil {
				return err
			}
			log.Printf("{ key: \"%s\", value: \"%s\" }\n", result.Key, result.Value)
			return nil
		},
	}
	staticCallCmd.Flags().StringVarP(&key, "key", "k", "", "")

	rootCmd.AddCommand(callCmd)
	rootCmd.AddCommand(staticCallCmd)
	rootCmd.Execute()
}
