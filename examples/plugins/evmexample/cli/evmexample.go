// +build evm

package main

import (
	"fmt"
	"github.com/loomnetwork/go-loom"
	"github.com/loomnetwork/go-loom/auth"
	"github.com/loomnetwork/go-loom/client"
	"github.com/loomnetwork/go-loom/examples/plugins/evmexample/types"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ed25519"
	"log"
	"os"
)

var RootCmd = &cobra.Command{
	Use:   "simplestore",
	Short: "SimpeStore DAppChain",
}

type persistFlags struct {
	WriteUri        string `json:"wrjiteuri"`
	ReadUri         string `json:"readuri"`
	ContractHexAddr string `json:"contractaddr"`
	ChainId         string `json:"chainid"`
}

func main() {
	var flags persistFlags
	RootCmd.PersistentFlags().StringVarP(&flags.WriteUri, "write", "w", "http://localhost:46657", "URI for sending txs")
	RootCmd.PersistentFlags().StringVarP(&flags.ReadUri, "read", "r", "http://localhost:9999", "URI for quering app state")
	RootCmd.PersistentFlags().StringVarP(&flags.ContractHexAddr, "contract", "c", "0x005B17864f3adbF53b1384F2E6f2120c6652F779", "contract address")
	RootCmd.PersistentFlags().StringVarP(&flags.ChainId, "chainId", "i", "default", "chain ID")

	getCmd := &cobra.Command{
		Use:   "get",
		Short: "get the value from the store",
		RunE: func(cmd *cobra.Command, args []string) error {
			result, err := GetValueCmd(flags.ChainId, flags.WriteUri, flags.ReadUri, flags.ContractHexAddr)
			if err == nil {
				log.Printf("{ Out: '%d' }", result)
				fmt.Println("output:", result)
			}
			return err
		},
	}

	var value int
	setCmd := &cobra.Command{
		Use:   "set",
		Short: "set the value in the store",
		RunE: func(cmd *cobra.Command, args []string) error {
			return SetValueCmd(flags.ChainId, flags.WriteUri, flags.ReadUri, flags.ContractHexAddr, value)
		},
	}
	setCmd.Flags().IntVarP(&value, "value", "v", 0, "value to set in store")
	RootCmd.AddCommand(
		getCmd,
		setCmd,
	)
	err := RootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func GetValueCmd(chainId, writeUri, readUri, contractHexAddr string) (int64, error) {
	rpcClient := client.NewDAppChainRPCClient(chainId, writeUri, readUri)
	contractAddr, err := loom.LocalAddressFromHexString(contractHexAddr)
	if err != nil {
		return 0, err
	}
	contract := client.NewContract(rpcClient, contractAddr, "EvmExample")

	dummy := &types.Dummy{}
	result := &types.WrapValue{}
	// NOTE: usually you shouldn't generate a new key pair for every tx, but this is just an example...
	_, priv, err := ed25519.GenerateKey(nil)
	if err != nil {
		return 0, err
	}
	signer := auth.NewEd25519Signer(priv)
	_, err = contract.Call("GetValue", dummy, signer, result)

	return result.Value, err
}

func SetValueCmd(chainId, writeUri, readUri, contractHexAddr string, value int) error {
	rpcClient := client.NewDAppChainRPCClient(chainId, writeUri, readUri)
	contractAddr, err := loom.LocalAddressFromHexString(contractHexAddr)
	if err != nil {
		return err
	}
	contract := client.NewContract(rpcClient, contractAddr, "EvmExample")

	// NOTE: usually you shouldn't generate a new key pair for every tx,
	// but this is just an example...
	_, priv, err := ed25519.GenerateKey(nil)
	if err != nil {
		return err
	}
	signer := auth.NewEd25519Signer(priv)

	payload := &types.WrapValue{
		Value: int64(value),
	}
	_, err = contract.Call("SetValue", payload, signer, nil)
	return err
}
