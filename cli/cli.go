package cli

import (
	"errors"

	"github.com/gogo/protobuf/proto"
	"github.com/spf13/cobra"

	"github.com/loomnetwork/go-loom"
	"github.com/loomnetwork/go-loom/client"
)

var TxFlags struct {
	WriteURI      string
	ReadURI       string
	MainnetURI    string
	ContractAddr  string
	ChainID       string
	PrivFile      string
	HsmConfigFile string
	Algo          string
	ChainType     string
}

func ContractCallCommand(name string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "call",
		Short: "call a contract method",
	}
	if name != "" {
		cmd = &cobra.Command{
			Use:   name,
			Short: "call a method of the " + name + " contract",
		}
	}

	pflags := cmd.PersistentFlags()
	pflags.StringVarP(&TxFlags.WriteURI, "write", "w", "http://localhost:46658/rpc", "URI for sending txs")
	pflags.StringVarP(&TxFlags.ReadURI, "read", "r", "http://localhost:46658/query", "URI for quering app state")
	pflags.StringVarP(&TxFlags.MainnetURI, "ethereum", "e", "http://localhost:8545", "URI for talking to Ethereum")
	pflags.StringVarP(&TxFlags.ContractAddr, "contract", "", "", "contract address")
	pflags.StringVarP(&TxFlags.ChainID, "chain", "", "default", "chain ID")
	pflags.StringVarP(&TxFlags.PrivFile, "key", "k", "", "private key file")
	pflags.StringVarP(&TxFlags.HsmConfigFile, "hsmconfig", "", "", "hsm config file")
	pflags.StringVarP(&TxFlags.Algo, "algo", "", "ed25519", "crypto algo Ed25519 or Secp256k1")
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
	pflags.StringVarP(&TxFlags.MainnetURI, "ethereum", "e", "http://localhost:8545", "URI for talking to Ethereum")
	pflags.StringVarP(&TxFlags.ContractAddr, "contract", "", "", "contract name")
	pflags.StringVarP(&TxFlags.ChainID, "chain", "", "default", "chain ID")
	pflags.StringVarP(&TxFlags.PrivFile, "key", "k", "", "private key file")
	pflags.StringVarP(&TxFlags.HsmConfigFile, "hsmconfig", "", "", "hsm config file")
	pflags.StringVarP(&TxFlags.Algo, "algo", "", "ed25519", "crypto algo for the key- default is Ed25519 or Secp256k1")

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
	signer, err := GetSigner(TxFlags.PrivFile, TxFlags.HsmConfigFile, TxFlags.Algo)
	if err != nil {
		return err
	}
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
