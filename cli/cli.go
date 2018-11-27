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

var TxFlags struct {
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
	pflags.StringVarP(&TxFlags.WriteURI, "write", "w", "http://localhost:46658/rpc", "URI for sending txs")
	pflags.StringVarP(&TxFlags.ReadURI, "read", "r", "http://localhost:46658/query", "URI for quering app state")
	pflags.StringVarP(&TxFlags.ContractAddr, "contract", "", "", "contract address")
	pflags.StringVarP(&TxFlags.ChainID, "chain", "", "default", "chain ID")
	pflags.StringVarP(&TxFlags.PrivFile, "private-key", "p", "", "private key file")
	return cmd
}

func ContractResolveCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "resolve",
		Short: "resolve a contract method",
	}
	pflags := cmd.PersistentFlags()
	pflags.StringVarP(&TxFlags.WriteURI, "write", "w", "http://localhost:46658/rpc", "URI for sending txs")
	pflags.StringVarP(&TxFlags.ReadURI, "read", "r", "http://localhost:46658/query", "URI for quering app state")
	pflags.StringVarP(&TxFlags.ContractAddr, "contract", "", "", "contract name")
	pflags.StringVarP(&TxFlags.ChainID, "chain", "", "default", "chain ID")
	pflags.StringVarP(&TxFlags.PrivFile, "private-key", "p", "", "private key file")
	return cmd
}

func contract(defaultAddr string) (*client.Contract, error) {
	contractAddrStr := TxFlags.ContractAddr
	if contractAddrStr == "" {
		contractAddrStr = defaultAddr
	}

	if contractAddrStr == "" {
		return nil, errors.New("contract address or name required")
	}

	contractAddr, err := ResolveAddress(contractAddrStr)
	if err != nil {
		return nil, err
	}

	// create rpc client
	rpcClient := client.NewDAppChainRPCClient(TxFlags.ChainID, TxFlags.WriteURI, TxFlags.ReadURI)
	// create contract
	contract := client.NewContract(rpcClient, contractAddr.Local)
	return contract, nil
}

func CallContract(defaultAddr string, method string, params proto.Message, result interface{}) error {
	if TxFlags.PrivFile == "" {
		return errors.New("private key required to call contract")
	}

	privKeyB64, err := ioutil.ReadFile(TxFlags.PrivFile)
	if err != nil {
		return err
	}

	privKey, err := base64.StdEncoding.DecodeString(string(privKeyB64))
	if err != nil {
		return err
	}

	signer := auth.NewEd25519Signer(privKey)

	contract, err := contract(defaultAddr)
	if err != nil {
		return err
	}
	_, err = contract.Call(method, params, signer, result)
	return err
}

func StaticCallContract(defaultAddr string, method string, params proto.Message, result interface{}) error {
	contract, err := contract(defaultAddr)
	if err != nil {
		return err
	}

	_, err = contract.StaticCall(method, params, loom.RootAddress(TxFlags.ChainID), result)
	return err
}
