package cli

import (
	"errors"

	"github.com/gogo/protobuf/proto"
	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"

	"github.com/loomnetwork/go-loom"
	"github.com/loomnetwork/go-loom/client"
)

type ContractCallFlags struct {
	URI           string
	ContractAddr  string
	ChainID       string
	PrivFile      string
	HsmConfigFile string
	Algo          string
}

var TxFlags struct {
	URI           string
	MainnetURI    string
	ContractAddr  string
	ChainID       string
	PrivFile      string
	HsmConfigFile string
	Algo          string
	CallerChainID string
}

// ContractCallCommand is deprecated in favor of CallContractWithFlags / StaticCallContractWithFlags
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
	pflags.StringVarP(&TxFlags.URI, "uri", "u", "http://localhost:46658", "DAppChain base URI")
	pflags.StringVarP(&TxFlags.MainnetURI, "ethereum", "e", "http://localhost:8545", "URI for talking to Ethereum")
	pflags.StringVarP(&TxFlags.ContractAddr, "contract", "", "", "contract address")
	pflags.StringVarP(&TxFlags.ChainID, "chain", "", "default", "chain ID")
	pflags.StringVarP(&TxFlags.PrivFile, "key", "k", "", "private key file")
	pflags.StringVarP(&TxFlags.HsmConfigFile, "hsmconfig", "", "", "hsm config file")
	pflags.StringVar(&TxFlags.Algo, "algo", "ed25519", "Signing algo: ed25519, secp256k1, tron")
	pflags.StringVar(&TxFlags.CallerChainID, "caller-chain", "", "Overrides chain ID of caller")

	return cmd
}

func ContractResolveCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "resolve",
		Short: "resolve a contract method",
	}
	pflags := cmd.PersistentFlags()
	pflags.StringVarP(&TxFlags.URI, "uri", "u", "http://localhost:46658", "DAppChain base URI")
	pflags.StringVarP(&TxFlags.MainnetURI, "ethereum", "e", "http://localhost:8545", "URI for talking to Ethereum")
	pflags.StringVarP(&TxFlags.ContractAddr, "contract", "", "", "contract name")
	pflags.StringVarP(&TxFlags.ChainID, "chain", "", "default", "chain ID")
	pflags.StringVarP(&TxFlags.PrivFile, "key", "k", "", "private key file")
	pflags.StringVarP(&TxFlags.HsmConfigFile, "hsmconfig", "", "", "hsm config file")
	pflags.StringVarP(&TxFlags.Algo, "algo", "", "ed25519", "crypto algo for the key- default is Ed25519 or Secp256k1")

	return cmd
}

func ResolveContract(callFlags *ContractCallFlags, defaultAddr string) (*client.Contract, error) {
	contractAddrStr := callFlags.ContractAddr
	if contractAddrStr == "" {
		contractAddrStr = defaultAddr
	}

	if contractAddrStr == "" {
		return nil, errors.New("contract address or name required")
	}

	contractAddr, err := ResolveAddress(contractAddrStr, callFlags.ChainID, callFlags.URI)
	if err != nil {
		return nil, err
	}
	// create rpc client
	rpcClient := client.NewDAppChainRPCClient(callFlags.ChainID, callFlags.URI+"/rpc", callFlags.URI+"/query")
	// create contract
	contract := client.NewContract(rpcClient, contractAddr.Local)
	return contract, nil
}

func CallContractWithFlags(callFlags *ContractCallFlags, defaultAddr string, method string, params proto.Message, result interface{}) error {
	signer, err := GetSigner(callFlags.PrivFile, callFlags.HsmConfigFile, callFlags.Algo)
	if err != nil {
		return err
	}
	contract, err := ResolveContract(callFlags, defaultAddr)
	if err != nil {
		return err
	}
	_, err = contract.Call(method, params, signer, result)
	return err
}

func StaticCallContractWithFlags(callFlags *ContractCallFlags, defaultAddr string, method string, params proto.Message, result interface{}) error {
	contract, err := ResolveContract(callFlags, defaultAddr)
	if err != nil {
		return err
	}
	_, err = contract.StaticCall(method, params, loom.RootAddress(callFlags.ChainID), result)
	return err
}

// AddContractCallFlags adds flags to the given flag set that are used when calling a mutable method
// on a Go contract.
func AddContractCallFlags(flagSet *flag.FlagSet, callFlags *ContractCallFlags) {
	flagSet.StringVarP(&callFlags.URI, "uri", "u", "http://localhost:46658", "DAppChain base URI")
	flagSet.StringVar(&callFlags.ContractAddr, "contract", "", "contract address")
	flagSet.StringVarP(&callFlags.ChainID, "chain", "c", "default", "chain ID")
	flagSet.StringVarP(&callFlags.PrivFile, "key", "k", "", "private key file")
	flagSet.StringVar(&callFlags.HsmConfigFile, "hsm", "", "hsm config file")
	flagSet.StringVar(&callFlags.Algo, "algo", "ed25519", "Signing algo: ed25519, secp256k1, tron")
}

// AddContractStaticCallFlags adds flags to the given flag set that are used when calling a static
// method on a Go contract.
func AddContractStaticCallFlags(flagSet *flag.FlagSet, callFlags *ContractCallFlags) {
	flagSet.StringVarP(&callFlags.URI, "uri", "u", "http://localhost:46658", "DAppChain base URI")
	flagSet.StringVar(&callFlags.ContractAddr, "contract", "", "contract address")
	flagSet.StringVarP(&callFlags.ChainID, "chain", "c", "default", "chain ID")
}

func contract(defaultAddr string) (*client.Contract, error) {
	contractAddrStr := TxFlags.ContractAddr
	if contractAddrStr == "" {
		contractAddrStr = defaultAddr
	}

	if contractAddrStr == "" {
		return nil, errors.New("contract address or name required")
	}

	contractAddr, err := ResolveAddress(contractAddrStr, TxFlags.ChainID, TxFlags.URI)
	if err != nil {
		return nil, err
	}
	// create rpc client
	rpcClient := client.NewDAppChainRPCClient(TxFlags.ChainID, TxFlags.URI+"/rpc", TxFlags.URI+"/query")
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
