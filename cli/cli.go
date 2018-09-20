package cli

import (
	"encoding/base64"
	"errors"
	"io/ioutil"

	"github.com/gogo/protobuf/proto"
	"github.com/spf13/cobra"

	"github.com/loomnetwork/go-loom"
	"github.com/loomnetwork/go-loom/auth"
	"github.com/loomnetwork/go-loom/client"
)

var txFlags struct {
	WriteURI     string
	ReadURI      string
	ContractAddr string
	ChainID      string
	PrivFile     string
}

func ContractCallCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "call",
		Short: "call a contract method",
	}
	pflags := cmd.PersistentFlags()
	pflags.StringVarP(&txFlags.WriteURI, "write", "w", "http://localhost:46658/rpc", "URI for sending txs")
	pflags.StringVarP(&txFlags.ReadURI, "read", "r", "http://localhost:46658/query", "URI for quering app state")
	pflags.StringVarP(&txFlags.ContractAddr, "contract", "", "", "contract address")
	pflags.StringVarP(&txFlags.ChainID, "chain", "", "default", "chain ID")
	pflags.StringVarP(&txFlags.PrivFile, "private-key", "p", "", "private key file")
	return cmd
}

func contract(defaultAddr, defaultVersion string) (*client.Contract, error) {
	contractAddrStr := txFlags.ContractAddr
	if contractAddrStr == "" {
		contractAddrStr = defaultAddr
	}

	if contractAddrStr == "" {
		return nil, errors.New("contract address or name required")
	}

	contractAddr, err := ResolveAddress(contractAddrStr, defaultVersion)
	if err != nil {
		return nil, err
	}

	// create rpc client
	rpcClient := client.NewDAppChainRPCClient(txFlags.ChainID, txFlags.WriteURI, txFlags.ReadURI)
	// create contract
	contract := client.NewContract(rpcClient, contractAddr.Local)
	return contract, nil
}

func CallContract(defaultAddr string, version string, method string, params proto.Message, result interface{}) error {
	if txFlags.PrivFile == "" {
		return errors.New("private key required to call contract")
	}

	privKeyB64, err := ioutil.ReadFile(txFlags.PrivFile)
	if err != nil {
		return err
	}

	privKey, err := base64.StdEncoding.DecodeString(string(privKeyB64))
	if err != nil {
		return err
	}

	signer := auth.NewEd25519Signer(privKey)

	contract, err := contract(defaultAddr, version)
	if err != nil {
		return err
	}
	_, err = contract.Call(method, params, signer, result)
	return err
}

func StaticCallContract(defaultAddr string, defaultVersion string, method string, params proto.Message, result interface{}) error {
	contract, err := contract(defaultAddr, defaultVersion)
	if err != nil {
		return err
	}

	_, err = contract.StaticCall(method, params, loom.RootAddress(txFlags.ChainID), result)
	return err
}
