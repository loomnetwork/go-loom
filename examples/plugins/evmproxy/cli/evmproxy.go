// +build evm

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/loomnetwork/go-loom"
	"github.com/loomnetwork/go-loom/auth"
	"github.com/loomnetwork/go-loom/client"
	"github.com/loomnetwork/go-loom/examples/plugins/evmproxy/types"
	"github.com/spf13/cobra"
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
				log.Printf("{ Out: '%v' }", result)
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

func GetValueCmd(chainId, writeUri, readUri, contractHexAddr string) (string, error) {
	rpcClient := client.NewDAppChainRPCClient(chainId, writeUri, readUri)
	contractAddr, err := loom.LocalAddressFromHexString(contractHexAddr)
	if err != nil {
		return "", err
	}
	contract := client.NewContract(rpcClient, contractAddr)

	ethCall := &types.EthCall{
		Data: "0x6d4ce63c",
	}

	ethCallResult := &types.EthCallResult{}
	// NOTE: usually you shouldn't generate a new key pair for every tx, but this is just an example...
	signer := auth.NewSecp256k1Signer(nil)
	_, err = contract.Call("EthCall", ethCall, signer, ethCallResult)

	return ethCallResult.GetData(), err
}

func SetValueCmd(chainId, writeUri, readUri, contractHexAddr string, value int) error {
	rpcClient := client.NewDAppChainRPCClient(chainId, writeUri, readUri)
	contractAddr, err := loom.LocalAddressFromHexString(contractHexAddr)
	if err != nil {
		return err
	}
	contract := client.NewContract(rpcClient, contractAddr)

	// NOTE: usually you shouldn't generate a new key pair for every tx,
	// but this is just an example...
	signer := auth.NewSecp256k1Signer(nil)

	// set(uint256) = 60fe47b1 (4 bytes)
	// 0000000000000000000000000000000000000000000000000000000000000001 = 1 (64 bytes)
	payload := &types.EthCall{
		Data: "0x60fe47b10000000000000000000000000000000000000000000000000000000000000002",
	}
	_, err = contract.Call("EthTransaction", payload, signer, nil)
	return err
}
