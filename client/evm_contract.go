package client

import (
	"github.com/gogo/protobuf/proto"
	"github.com/loomnetwork/go-loom"
	"github.com/loomnetwork/go-loom/auth"
	"github.com/loomnetwork/go-loom/vm"
)

// Contract provides a thin abstraction over DAppChainClient that makes it easier to perform
// read & write operations on a contract running the EVM of a Loom DAppChain.
type EvmContract struct {
	client  *DAppChainRPCClient
	Address loom.Address
	Name    string
}

func NewEvmContract(client *DAppChainRPCClient, contractAddr loom.LocalAddress) *EvmContract {
	return &EvmContract{
		client: client,
		Address: loom.Address{
			ChainID: client.GetChainID(),
			Local:   contractAddr,
		},
	}
}

func DeployContract(client *DAppChainRPCClient, byteCode []byte, signer auth.Signer, name string) (*EvmContract, []byte, error) {
	callerAddr := loom.Address{
		ChainID: client.GetChainID(),
		Local:   loom.LocalAddressFromPublicKey(signer.PublicKey()),
	}
	resp, err := client.CommitDeployTx(callerAddr, signer, vm.VMType_EVM, byteCode, name)
	if err != nil {
		return nil, []byte{}, err
	}
	response := vm.DeployResponse{}
	err = proto.Unmarshal(resp, &response)
	if err != nil {
		return nil, []byte{}, err
	}
	data := vm.DeployResponseData{}
	err = proto.Unmarshal(response.Output, &data)
	if err != nil {
		return nil, []byte{}, err
	}
	return &EvmContract{
		client:  client,
		Address: loom.UnmarshalAddressPB(response.Contract),
		Name:    name,
	}, data.GetTxHash(), nil
}

func (c *EvmContract) Call(input []byte, signer auth.Signer) ([]byte, error) {
	callerAddr := loom.Address{
		ChainID: c.client.GetChainID(),
		Local:   loom.LocalAddressFromPublicKey(signer.PublicKey()),
	}
	return c.client.CommitCallTx(callerAddr, c.Address, signer, vm.VMType_EVM, input)
}

func (c *EvmContract) StaticCall(input []byte, caller loom.Address) ([]byte, error) {
	return c.client.QueryEvm(caller, c.Address.Local, input)
}
