package client

import (
	"encoding/hex"
	"errors"
	"strconv"

	"github.com/gogo/protobuf/proto"
	"github.com/loomnetwork/go-loom"
	"github.com/loomnetwork/go-loom/auth"
	"github.com/loomnetwork/go-loom/types"
	"github.com/loomnetwork/go-loom/vm"
)

type TxHandlerResult struct {
	Code  int32  `json:"code"`
	Error string `json:"log"`
	Data  []byte `json:"data"`
}

type BroadcastTxCommitResult struct {
	CheckTx   TxHandlerResult `json:"check_tx"`
	DeliverTx TxHandlerResult `json:"deliver_tx"`
	Hash      string          `json:"hash"`
	Height    int64           `json:"height"`
}

// Implements the DAppChainClient interface
type DAppChainRPCClient struct {
	chainID       string
	writeURI      string
	readURI       string
	txClient      *JSONRPCClient
	queryClient   *JSONRPCClient
	nextRequestID uint64
}

// NewDAppChainRPCClient creates a new dumb client that can be used to commit txs and query contract
// state via RPC.
// URI parameters should be specified as "tcp://<host>:<port>", writeURI the host that txs will be
// submitted to (port 46657 by default), readURI is the host that will be queried for current app
// state (47000 by default).
func NewDAppChainRPCClient(chainID, writeURI, readURI string) *DAppChainRPCClient {
	return &DAppChainRPCClient{
		chainID:       chainID,
		writeURI:      writeURI,
		readURI:       readURI,
		txClient:      NewJSONRPCClient(writeURI),
		queryClient:   NewJSONRPCClient(readURI),
		nextRequestID: 1,
	}
}

func (c *DAppChainRPCClient) getNextRequestID() string {
	id := strconv.FormatUint(c.nextRequestID, 10)
	c.nextRequestID++
	return id
}

func (c *DAppChainRPCClient) GetChainID() string {
	return c.chainID
}

func (c *DAppChainRPCClient) GetNonce(signer auth.Signer) (uint64, error) {
	params := map[string]interface{}{
		"key": hex.EncodeToString(signer.PublicKey()),
	}
	var r uint64
	err := c.queryClient.Call("nonce", params, c.getNextRequestID(), &r)
	return r, err
}

func (c *DAppChainRPCClient) CommitTx(signer auth.Signer, tx proto.Message) ([]byte, error) {
	// TODO: signing & noncing should be handled by middleware
	nonce, err := c.GetNonce(signer)
	if err != nil {
		return nil, err
	}
	txBytes, err := proto.Marshal(tx)
	if err != nil {
		return nil, err
	}
	nonceTxBytes, err := proto.Marshal(&auth.NonceTx{
		Inner:    txBytes,
		Sequence: nonce + 1,
	})
	if err != nil {
		return nil, err
	}
	signedTxBytes, err := proto.Marshal(auth.SignTx(signer, nonceTxBytes))
	if err != nil {
		return nil, err
	}
	params := map[string]interface{}{
		"tx": signedTxBytes,
	}
	var r BroadcastTxCommitResult
	if err = c.txClient.Call("broadcast_tx_commit", params, c.getNextRequestID(), &r); err != nil {
		return nil, err
	}
	if r.CheckTx.Code != 0 {
		if len(r.CheckTx.Error) != 0 {
			return nil, errors.New(r.CheckTx.Error)
		}
		return nil, errors.New("CheckTx failed")
	}
	if r.DeliverTx.Code != 0 {
		if len(r.DeliverTx.Error) != 0 {
			return nil, errors.New(r.DeliverTx.Error)
		}
		return nil, errors.New("DeliverTx failed")
	}
	return r.DeliverTx.Data, nil
}

func (c *DAppChainRPCClient) Query(contractAddr loom.LocalAddress, query proto.Message) ([]byte, error) {
	queryBytes, err := proto.Marshal(query)
	if err != nil {
		return nil, err
	}
	params := map[string]interface{}{
		"contract": contractAddr.String(),
		"query":    queryBytes,
		"vmType":   vm.VMType_PLUGIN,
	}
	var r []byte
	if err = c.queryClient.Call("query", params, c.getNextRequestID(), &r); err != nil {
		return nil, err
	}
	return r, nil
}

func (c *DAppChainRPCClient) Resolve(name string) (loom.Address, error) {
	params := map[string]interface{}{
		"name": name,
	}
	var addrStr string
	if err := c.queryClient.Call("resolve", params, c.getNextRequestID(), &addrStr); err != nil {
		return loom.Address{}, err
	}
	return loom.ParseAddress(addrStr)
}

func (c *DAppChainRPCClient) QueryEvm(contractAddr loom.LocalAddress, query []byte) ([]byte, error) {
	params := map[string]interface{}{
		"contract": contractAddr.String(),
		"query":    query,
		"vmType":   vm.VMType_EVM,
	}
	var r []byte
	if err := c.queryClient.Call("query", params, c.getNextRequestID(), &r); err != nil {
		return nil, err
	}
	return r, nil
}

func (c *DAppChainRPCClient) GetEvmTxReceipt(txHash []byte) (vm.EvmTxReceipt, error) {
	params := map[string]interface{}{
		"txHash": txHash,
	}
	var r []byte
	if err := c.queryClient.Call("txreceipt", params, c.getNextRequestID(), &r); err != nil {
		return vm.EvmTxReceipt{}, err
	}
	var receipt vm.EvmTxReceipt
	err := proto.Unmarshal(r, &receipt)

	return receipt, err
}

func (c *DAppChainRPCClient) CommitDeployTx(
	from loom.Address,
	signer auth.Signer,
	vmType vm.VMType,
	code []byte,
) ([]byte, error) {
	deployTxBytes, err := proto.Marshal(&vm.DeployTx{
		VmType: vmType,
		Code:   code,
	})
	if err != nil {
		return nil, err
	}
	msgBytes, err := proto.Marshal(&vm.MessageTx{
		From: from.MarshalPB(),
		To:   loom.Address{}.MarshalPB(), // not used
		Data: deployTxBytes,
	})
	if err != nil {
		return nil, err
	}
	tx := &types.Transaction{
		Id:   1,
		Data: msgBytes,
	}
	return c.CommitTx(signer, tx)
}

func (c *DAppChainRPCClient) CommitCallTx(
	caller loom.Address,
	contract loom.Address,
	signer auth.Signer,
	vmType vm.VMType,
	input []byte,
) ([]byte, error) {
	callTxBytes, err := proto.Marshal(&vm.CallTx{
		VmType: vm.VMType(vmType),
		Input:  input,
	})
	if err != nil {
		return nil, err
	}
	msgTx := &vm.MessageTx{
		From: caller.MarshalPB(),
		To:   contract.MarshalPB(),
		Data: callTxBytes,
	}
	msgBytes, err := proto.Marshal(msgTx)
	if err != nil {
		return nil, err
	}
	// tx ids associated with handlers in loadApp()
	tx := &types.Transaction{
		Id:   2,
		Data: msgBytes,
	}
	return c.CommitTx(signer, tx)
}
