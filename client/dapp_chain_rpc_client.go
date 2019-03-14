package client

import (
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"net/http"
	"strconv"

	"github.com/gogo/protobuf/proto"
	"github.com/loomnetwork/go-loom"
	"github.com/loomnetwork/go-loom/auth"
	ptypes "github.com/loomnetwork/go-loom/plugin/types"
	"github.com/loomnetwork/go-loom/types"
	"github.com/loomnetwork/go-loom/vm"
	"github.com/pkg/errors"
)

const (
	ShortPollLimit = 10
	ShortPollDelay = 1 * time.Second

	CodeTypeOK = 0
)

type TxHandlerResult struct {
	Code  int32  `json:"code"`
	Error string `json:"log"`
	Data  []byte `json:"data"`
}

type BoradcastTxSyncResult struct {
	Code  int32  `json:"code"`
	Data  []byte `json:"data"`
	Hash  string `json:"hash"`
	Error string `json:"log"`
}

type TxQueryResult struct {
	TxResult TxHandlerResult `json:"tx_result"`
}

type BroadcastTxCommitResult struct {
	CheckTx   TxHandlerResult `json:"check_tx"`
	DeliverTx TxHandlerResult `json:"deliver_tx"`
	Hash      string          `json:"hash"`
	Height    string          `json:"height"`
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

func NewDAppChainRPCClientShareTransport(chainID, writeURI, readURI string, transport *http.Transport) *DAppChainRPCClient {
	return &DAppChainRPCClient{
		chainID:       chainID,
		writeURI:      writeURI,
		readURI:       readURI,
		txClient:      NewJSONRPCClientShareTransport(writeURI, transport),
		queryClient:   NewJSONRPCClientShareTransport(readURI, transport),
		nextRequestID: 1,
	}
}

func (c *DAppChainRPCClient) CloseIdleConnections() {
	c.txClient.client.Transport.(*http.Transport).CloseIdleConnections()
	c.queryClient.client.Transport.(*http.Transport).CloseIdleConnections()
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
	var rRes string
	err := c.queryClient.Call("nonce", params, c.getNextRequestID(), &rRes)
	if err != nil {
		return 0, err
	}
	return strconv.ParseUint(rRes, 10, 64)
}

func (c *DAppChainRPCClient) GetNonce2(caller loom.Address, isAddressMapped bool) (uint64, error) {
	var accountType string
	if isAddressMapped {
		accountType = "2"
	} else {
		accountType = "1"
	}
	accountType = accountType
	params := map[string]interface{}{
		"chainId":     caller.ChainID,
		"local":       caller.Local,
		"accountType": accountType,
	}
	var rRes string
	err := c.queryClient.Call("nonce2", params, c.getNextRequestID(), &rRes)
	if err != nil {
		return 0, err
	}
	return strconv.ParseUint(rRes, 10, 64)
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

	var r BoradcastTxSyncResult
	if err = c.txClient.Call("broadcast_tx_sync", params, c.getNextRequestID(), &r); err != nil {
		return nil, err
	}
	if r.Code != CodeTypeOK {
		if len(r.Error) != 0 {
			return nil, errors.New(r.Error)
		}
		return nil, fmt.Errorf("CheckTx failed")
	}

	txResult, err := c.pollTx(r.Hash, ShortPollLimit, ShortPollDelay)
	if err != nil {
		return nil, err
	}

	return txResult.Data, nil
}

func (c *DAppChainRPCClient) CommitTx2(signer auth.Signer, tx proto.Message, caller loom.Address, chainName string) ([]byte, error) {
	// TODO: signing & noncing should be handled by middleware
	nonce, err := c.GetNonce2(caller, len(chainName) > 0)
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
	signedTx := auth.SignTx(signer, nonceTxBytes)
	signedTx.ChainName = chainName
	signedTxBytes, err := proto.Marshal(signedTx)
	if err != nil {
		return nil, err
	}
	params := map[string]interface{}{
		"tx": signedTxBytes,
	}

	var r BoradcastTxSyncResult
	if err = c.txClient.Call("broadcast_tx_sync", params, c.getNextRequestID(), &r); err != nil {
		return nil, err
	}
	if r.Code != CodeTypeOK {
		if len(r.Error) != 0 {
			return nil, errors.New(r.Error)
		}
		return nil, fmt.Errorf("CheckTx failed")
	}

	txResult, err := c.pollTx(r.Hash, ShortPollLimit, ShortPollDelay)
	if err != nil {
		return nil, err
	}

	return txResult.Data, nil
}

func (c *DAppChainRPCClient) pollTx(hash string, shortPollLimit int, shortPollDelay time.Duration) (*TxHandlerResult, error) {
	var result TxQueryResult
	var err error

	decodedHash, err := hex.DecodeString(hash)
	if err != nil {
		return nil, errors.Wrapf(err, "error while polling for tx")
	}
	params := map[string]interface{}{
		"hash": decodedHash,
	}

	for i := 0; i < shortPollLimit; i++ {
		// Delaying in beginning of the loop, as immediate poll will likely result in "not found"
		time.Sleep(shortPollDelay)

		if err = c.txClient.Call("tx", params, c.getNextRequestID(), &result); err != nil {
			if !strings.Contains(err.Error(), "not found") {
				// Bailing early if error is due to something other than pending tx.
				return nil, errors.Wrap(err, "error while polling for tx")
			}
		} else {
			if result.TxResult.Code != CodeTypeOK {
				if len(result.TxResult.Error) != 0 {
					return nil, errors.New(result.TxResult.Error)
				}
				return nil, fmt.Errorf("DeliverTx failed")
			}
			break
		}
	}

	if err != nil {
		return nil, errors.Wrap(err, "max retry exceeded while polling for tx")
	}

	return &result.TxResult, nil
}

func (c *DAppChainRPCClient) Query(caller loom.Address, contractAddr loom.LocalAddress, query proto.Message) ([]byte, error) {
	queryBytes, err := proto.Marshal(query)
	if err != nil {
		return nil, err
	}
	params := map[string]interface{}{
		"caller":   caller.String(),
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

// GetCode returns the runtime byte-code of a contract running on a DAppChain's EVM.
// Gives an error for non-EVM contracts.
// contract - address of the contract in the form of a string. (Use loom.Address.String() to convert)
// return []byte - runtime bytecode of the contract.
func (c *DAppChainRPCClient) GetEvmCode(contract string) ([]byte, error) {
	params := map[string]interface{}{
		"contract": contract,
	}

	var bytecode []byte
	if err := c.queryClient.Call("getevmcode", params, c.getNextRequestID(), &bytecode); err != nil {
		return []byte{}, err
	}
	return bytecode, nil
}

func (c *DAppChainRPCClient) GetEvmLogs(filter string) (ptypes.EthFilterLogList, error) {
	params := map[string]interface{}{
		"filter": filter,
	}

	var r []byte
	if err := c.queryClient.Call("getevmlogs", params, c.getNextRequestID(), &r); err != nil {
		return ptypes.EthFilterLogList{}, err
	}
	var logs ptypes.EthFilterLogList
	if err := proto.Unmarshal(r, &logs); err != nil {
		return ptypes.EthFilterLogList{}, err
	}

	return logs, nil
}

// Sets up new filter for polling
// https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_newfilter
func (c *DAppChainRPCClient) NewEvmFilter(filter string) (string, error) {
	params := map[string]interface{}{
		"filter": filter,
	}

	var id string
	if err := c.queryClient.Call("newevmfilter", params, c.getNextRequestID(), &id); err != nil {
		return "", err
	}
	return id, nil
}

// https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_newblockfilter
func (c *DAppChainRPCClient) NewBlockEvmFilter() (string, error) {
	params := map[string]interface{}{}
	var id string
	if err := c.queryClient.Call("newblockevmfilter", params, c.getNextRequestID(), &id); err != nil {
		return "", err
	}
	return id, nil
}

// https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_newpendingtransactionfilter
func (c *DAppChainRPCClient) NewPendingTransactionEvmFilter() (string, error) {
	params := map[string]interface{}{}

	var id string
	if err := c.queryClient.Call("newpendingtransactionevmfilter", params, c.getNextRequestID(), &id); err != nil {
		return "", err
	}
	return id, nil
}

// Get logs since last poll
// https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_getfilterchanges
// could return protbuf of EthFilterLogList, EthBlockHashList or EthTxHashList
func (c *DAppChainRPCClient) GetEvmFilterChanges(id string) ([]byte, error) {
	params := map[string]interface{}{
		"id": id,
	}

	var r []byte
	if err := c.queryClient.Call("getevmfilterchanges", params, c.getNextRequestID(), &r); err != nil {
		return nil, err
	}

	return r, nil
}

// Forget filter
// https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_uninstallfilter
func (c *DAppChainRPCClient) UninstallEvmFilter(id string) (bool, error) {
	params := map[string]interface{}{
		"id": id,
	}

	var ok bool
	if err := c.queryClient.Call("uninstallevmfilter", params, c.getNextRequestID(), &ok); err != nil {
		return ok, err
	}
	return ok, nil
}

func (c *DAppChainRPCClient) GetEvmBlockByNumber(number string, full bool) (ptypes.EthBlockInfo, error) {
	params := map[string]interface{}{
		"number": number,
		"full":   full,
	}

	var r []byte
	if err := c.queryClient.Call("getevmblockbynumber", params, c.getNextRequestID(), &r); err != nil {
		return ptypes.EthBlockInfo{}, err
	}
	var blockInfo ptypes.EthBlockInfo
	if err := proto.Unmarshal(r, &blockInfo); err != nil {
		return ptypes.EthBlockInfo{}, err
	}

	return blockInfo, nil
}

func (c *DAppChainRPCClient) GetEvmBlockByHash(hash []byte, full bool) (ptypes.EthBlockInfo, error) {
	params := map[string]interface{}{
		"hash": hash,
		"full": full,
	}

	var r []byte
	if err := c.queryClient.Call("getevmblockbyhash", params, c.getNextRequestID(), &r); err != nil {
		return ptypes.EthBlockInfo{}, err
	}
	var blockInfo ptypes.EthBlockInfo
	if err := proto.Unmarshal(r, &blockInfo); err != nil {
		return ptypes.EthBlockInfo{}, err
	}

	return blockInfo, nil
}

func (c *DAppChainRPCClient) GetEvmTransactionByHash(hash []byte) (ptypes.EvmTxObject, error) {
	params := map[string]interface{}{
		"hash": hash,
	}

	var r []byte
	if err := c.queryClient.Call("getevmtransactionbyhash", params, c.getNextRequestID(), &r); err != nil {
		return ptypes.EvmTxObject{}, err
	}
	var txInfo ptypes.EvmTxObject
	if err := proto.Unmarshal(r, &txInfo); err != nil {
		return ptypes.EvmTxObject{}, err
	}

	return txInfo, nil
}

func (c *DAppChainRPCClient) QueryEvm(caller loom.Address, contractAddr loom.LocalAddress, query []byte) ([]byte, error) {
	params := map[string]interface{}{
		"caller":   caller.String(),
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

func (c *DAppChainRPCClient) GetEvmTxReceipt(txHash []byte) (ptypes.EvmTxReceipt, error) {
	params := map[string]interface{}{
		"txHash": txHash,
	}
	var r []byte
	if err := c.queryClient.Call("evmtxreceipt", params, c.getNextRequestID(), &r); err != nil {
		return ptypes.EvmTxReceipt{}, err
	}
	var receipt ptypes.EvmTxReceipt
	err := proto.Unmarshal(r, &receipt)

	return receipt, err
}

func (c *DAppChainRPCClient) CommitDeployTx(
	from loom.Address,
	signer auth.Signer,
	vmType vm.VMType,
	code []byte,
	name string,
) ([]byte, error) {
	deployTxBytes, err := proto.Marshal(&vm.DeployTx{
		VmType: vmType,
		Code:   code,
		Name:   name,
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
	return c.CommitTx2(signer, tx, from, "")
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
	return c.CommitTx2(signer, tx, caller, "")
}

func (c *DAppChainRPCClient) CommitDeployTx2(
	from loom.Address,
	signer auth.Signer,
	vmType vm.VMType,
	code []byte,
	name, chainName string,
) ([]byte, error) {
	deployTxBytes, err := proto.Marshal(&vm.DeployTx{
		VmType: vmType,
		Code:   code,
		Name:   name,
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
	return c.CommitTx2(signer, tx, from, chainName)
}

func (c *DAppChainRPCClient) CommitCallTx2(
	caller loom.Address,
	contract loom.Address,
	signer auth.Signer,
	vmType vm.VMType,
	input []byte,
	chainName string,
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
	return c.CommitTx2(signer, tx, caller, chainName)
}
