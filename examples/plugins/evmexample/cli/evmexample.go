// +build evm

package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/loomnetwork/go-loom"
	"github.com/loomnetwork/go-loom/auth"
	"github.com/loomnetwork/go-loom/client"
	"github.com/loomnetwork/go-loom/examples/plugins/evmexample/types"
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
	ContractName    string `json:"contract-name"`
}

func main() {
	var flags persistFlags
	RootCmd.PersistentFlags().StringVarP(&flags.WriteUri, "write", "w", "http://localhost:46657", "URI for sending txs")
	RootCmd.PersistentFlags().StringVarP(&flags.ReadUri, "read", "r", "http://localhost:9999", "URI for quering app state")
	RootCmd.PersistentFlags().StringVarP(&flags.ContractHexAddr, "contract", "c", "", "contract address")
	RootCmd.PersistentFlags().StringVarP(&flags.ChainId, "chainId", "i", "default", "chain ID")
	RootCmd.PersistentFlags().StringVarP(&flags.ContractName, "contract-name", "n", "evmexample", "contract name")

	var binFile string
	deployCmd := &cobra.Command{
		Use:   "deploy",
		Short: "deploy SimpleStore Solidity contract",
		RunE: func(cmd *cobra.Command, args []string) error {
			contract, txHash, err := deployTx(flags.ChainId, flags.WriteUri, flags.ReadUri, binFile, flags.ContractName)
			if err == nil {
				fmt.Println("address ", contract.Address.String())
				fmt.Println("transaction hash", txHash)
			}
			return err
		},
	}
	deployCmd.Flags().StringVarP(
		&binFile,
		"bytecode file",
		"b",
		"./contracts/SimpleStore.bin",
		"file containg solidty bytcode",
	)

	getCmd := &cobra.Command{
		Use:   "get",
		Short: "get the value from the store",
		RunE: func(cmd *cobra.Command, args []string) error {
			result, err := GetValueCmd(flags.ChainId, flags.WriteUri, flags.ReadUri, flags.ContractHexAddr, flags.ContractName)
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
			return SetValueCmd(flags.ChainId, flags.WriteUri, flags.ReadUri, flags.ContractHexAddr, flags.ContractName, value)
		},
	}
	setCmd.Flags().IntVarP(&value, "value", "v", 0, "value to set in store")

	RootCmd.AddCommand(
		deployCmd,
		getCmd,
		setCmd,
	)
	err := RootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func deployTx(chainId, writeUri, readUri, binFile, name string) (*client.EvmContract, []byte, error) {
	bytetext, err := ioutil.ReadFile(binFile)

	if err != nil {
		return nil, nil, err
	}
	if string(bytetext[0:2]) == "0x" {
		bytetext = bytetext[2:]
	}
	bytecode, err := hex.DecodeString(string(bytetext))
	if err != nil {
		return nil, nil, err
	}

	// NOTE: usually you shouldn't generate a new key pair for every tx,
	// but this is just an example...
	signer := auth.NewSigner(auth.SignerTypeEd25519, nil)

	rpcClient := client.NewDAppChainRPCClient(chainId, writeUri, readUri)
	return client.DeployContract(rpcClient, bytecode, signer, name)
}

func GetValueCmd(chainId, writeUri, readUri, contractHexAddr, name string) (int64, error) {
	rpcClient := client.NewDAppChainRPCClient(chainId, writeUri, readUri)

	var contractLocalAddr loom.LocalAddress
	var err error
	if contractHexAddr != "" {
		if name != "" {
			fmt.Println("Both name and address entered, using address ", contractHexAddr)
		}
		contractLocalAddr, err = loom.LocalAddressFromHexString(contractHexAddr)
		if err != nil {
			return 0, err
		}
	} else {
		contractAddr, err := rpcClient.Resolve(name)
		if err != nil {
			return 0, err
		}
		contractLocalAddr = contractAddr.Local
	}
	contract := client.NewContract(rpcClient, contractLocalAddr)

	dummy := &types.Dummy{}
	result := &types.WrapValue{}
	// NOTE: usually you shouldn't generate a new key pair for every tx, but this is just an example...
	signer := auth.NewSigner(auth.SignerTypeEd25519, nil)
	_, err = contract.Call("GetValue", dummy, signer, result)

	return result.Value, err
}

func SetValueCmd(chainId, writeUri, readUri, contractHexAddr, name string, value int) error {
	rpcClient := client.NewDAppChainRPCClient(chainId, writeUri, readUri)
	var contractLocalAddr loom.LocalAddress
	var err error
	if contractHexAddr != "" {
		if name != "" {
			fmt.Println("Both name and address entered, using address ", contractHexAddr)
		}
		contractLocalAddr, err = loom.LocalAddressFromHexString(contractHexAddr)
		if err != nil {
			return err
		}
	} else {
		contractAddr, err := rpcClient.Resolve(name)
		if err != nil {
			return err
		}
		contractLocalAddr = contractAddr.Local
	}
	contract := client.NewContract(rpcClient, contractLocalAddr)

	// NOTE: usually you shouldn't generate a new key pair for every tx,
	// but this is just an example...
	signer := auth.NewSigner(auth.SignerTypeEd25519, nil)

	payload := &types.WrapValue{
		Value: int64(value),
	}
	_, err = contract.Call("SetValue", payload, signer, nil)
	return err
}
