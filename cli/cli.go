package cli

import (
	"errors"
	"io/ioutil"

	"github.com/gogo/protobuf/proto"
	"github.com/spf13/cobra"

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
	pflags.StringVarP(&txFlags.PrivFile, "key", "k", "", "private key file")
	return cmd
}

func contract() (*client.Contract, error) {
	// create rpc client
	rpcClient := client.NewDAppChainRPCClient(txFlags.ChainID, txFlags.WriteURI, txFlags.ReadURI)

	contractAddr, err := ParseAddress(txFlags.ContractAddr)
	if err != nil {
		// if address invalid, try to resolve it using registry
		contractAddr, err = rpcClient.Resolve(txFlags.ContractAddr)
		if err != nil {
			return nil, err
		}
	}
	// create contract
	contract := client.NewContract(rpcClient, contractAddr.Local)
	return contract, nil
}

func CallContract(method string, params proto.Message, result interface{}) error {
	if txFlags.PrivFile == "" {
		return errors.New("private key required to call contract")
	}
	privKey, err := ioutil.ReadFile(txFlags.PrivFile)
	if err != nil {
		return err
	}

	signer := auth.NewEd25519Signer(privKey)

	contract, err := contract()
	if err != nil {
		return err
	}
	_, err = contract.Call(method, params, signer, result)
	return err
}

func StaticCallContract(method string, params proto.Message, result interface{}) error {
	contract, err := contract()
	if err != nil {
		return err
	}

	_, err = contract.StaticCall(method, params, result)
	return err
}
