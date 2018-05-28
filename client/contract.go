package client

import (
	"errors"
	"reflect"

	"github.com/gogo/protobuf/proto"
	"github.com/loomnetwork/go-loom"
	"github.com/loomnetwork/go-loom/auth"
	"github.com/loomnetwork/go-loom/plugin"
	ptypes "github.com/loomnetwork/go-loom/plugin/types"
	"github.com/loomnetwork/go-loom/types"
	"github.com/loomnetwork/go-loom/vm"
)

var (
	typeOfPBMessage = reflect.TypeOf((*proto.Message)(nil)).Elem()
)

// Contract provides a thin abstraction over DAppChainClient that makes it easier to perform
// read & write operations on a contract running on a Loom DAppChain.
type Contract struct {
	client  *DAppChainRPCClient
	Address loom.Address
	Name    string
}

func NewContract(client *DAppChainRPCClient, contractAddr loom.LocalAddress) *Contract {
	return &Contract{
		client: client,
		Address: loom.Address{
			ChainID: client.GetChainID(),
			Local:   contractAddr,
		},
	}
}

func NewNamedContract(client *DAppChainRPCClient, name string) *Contract {
	addr, err := client.Resolve(name)
	if err != nil {
		return nil
	}
	return &Contract{
		client:  client,
		Address: addr,
		Name:    name,
	}
}

func NewDeployEvmContract(client *DAppChainRPCClient, signer auth.Signer, byteCode []byte, name string) *Contract {
	callerAddr := loom.Address{
		ChainID: client.GetChainID(),
		Local:   loom.LocalAddressFromPublicKey(signer.PublicKey()),
	}
	resp, err := client.CommitDeployTx(callerAddr, signer, vm.VMType_EVM, byteCode, name)
	if err != nil {
		return nil
	}
	response := vm.DeployResponse{}
	err = proto.Unmarshal(resp, &response)
	if err != nil {
		return nil
	}
	return &Contract{
		client:  client,
		Address: loom.UnmarshalAddressPB(response.Contract),
		Name:    name,
	}
}

func (c *Contract) Call(method string, args proto.Message, signer auth.Signer, result interface{}) (interface{}, error) {
	if result != nil && !reflect.TypeOf(result).Implements(typeOfPBMessage) {
		return nil, errors.New("Contract.Call result parameter must be a protobuf")
	}

	argsBytes, err := proto.Marshal(args)
	if err != nil {
		return nil, err
	}
	methodCallBytes, err := proto.Marshal(&plugin.ContractMethodCall{
		Method: method,
		Args:   argsBytes,
	})
	if err != nil {
		return nil, err
	}
	requestBytes, err := proto.Marshal(&plugin.Request{
		ContentType: plugin.EncodingType_PROTOBUF3,
		Accept:      plugin.EncodingType_PROTOBUF3,
		Body:        methodCallBytes,
	})
	if err != nil {
		return nil, err
	}
	callTxBytes, err := proto.Marshal(&vm.CallTx{
		VmType: vm.VMType_PLUGIN,
		Input:  requestBytes,
	})
	if err != nil {
		return nil, err
	}
	callerAddr := loom.Address{
		ChainID: c.client.GetChainID(),
		Local:   loom.LocalAddressFromPublicKey(signer.PublicKey()),
	}
	msgTxBytes, err := proto.Marshal(&vm.MessageTx{
		From: callerAddr.MarshalPB(),
		To:   c.Address.MarshalPB(),
		Data: callTxBytes,
	})
	if err != nil {
		return nil, err
	}
	resultBytes, err := c.client.CommitTx(signer, &types.Transaction{
		Id:   2,
		Data: msgTxBytes,
	})
	if err != nil {
		return nil, err
	}
	if result != nil && len(resultBytes) > 0 {
		response := &ptypes.Response{}
		err = proto.Unmarshal(resultBytes, response)
		if err != nil {
			return nil, nil
		}
		if err := proto.Unmarshal(response.Body, result.(proto.Message)); err != nil {
			return result, err
		}
	}
	return nil, nil
}

func (c *Contract) StaticCall(method string, args proto.Message, result interface{}) (interface{}, error) {
	if result == nil || !reflect.TypeOf(result).Implements(typeOfPBMessage) {
		return nil, errors.New("Contract.StaticCall result parameter must be a protobuf")
	}
	argsBytes, err := proto.Marshal(args)
	if err != nil {
		return nil, err
	}
	methodCall := &plugin.ContractMethodCall{
		Method: method,
		Args:   argsBytes,
	}
	resultBytes, err := c.client.Query(c.Address.Local, methodCall)
	if err != nil {
		return nil, err
	}
	if len(resultBytes) > 0 {
		if err := proto.Unmarshal(resultBytes, result.(proto.Message)); err != nil {
			return result, err
		}
	}
	return nil, nil
}

func (c *Contract) CallEvm(input []byte, signer auth.Signer) ([]byte, error) {
	callerAddr := loom.Address{
		ChainID: c.client.GetChainID(),
		Local:   loom.LocalAddressFromPublicKey(signer.PublicKey()),
	}
	return c.client.CommitCallTx(callerAddr, c.Address, signer, vm.VMType_EVM, input)
}

func (c *Contract) StaticCallEvm(input []byte) ([]byte, error) {
	return c.client.QueryEvm(c.Address.Local, input)
}
