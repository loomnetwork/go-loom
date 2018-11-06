// +build evm

package client

import (
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	loom "github.com/loomnetwork/go-loom"
	"github.com/loomnetwork/go-loom/auth"
	pctypes "github.com/loomnetwork/go-loom/builtin/types/plasma_cash"
	"github.com/loomnetwork/go-loom/client/plasma_cash"
	"github.com/loomnetwork/go-loom/types"
	"github.com/pkg/errors"
)

// PlasmaCashClient client to loom plasma cash server
type PlasmaCashClient struct {
	ChainID      string
	WriteURI     string
	ReadURI      string
	loomcontract *Contract
	caller       loom.Address
	signer       auth.Signer
}

// CurrentBlock gets the highest block of plasma cash
// does not grab pending transactions yet
func (c *PlasmaCashClient) CurrentBlock() (plasma_cash.Block, error) {
	return c.Block(big.NewInt(0)) //asking for block zero gives latest
}

// BlockNumber gets the current plasma cash block height
func (c *PlasmaCashClient) BlockNumber() (*big.Int, error) {
	request := &pctypes.GetCurrentBlockRequest{}
	var result pctypes.GetCurrentBlockResponse

	if _, err := c.loomcontract.StaticCall("GetCurrentBlockRequest", request, c.caller, &result); err != nil {
		log.Fatalf("failed getting Block number - %v\n", err)

		return big.NewInt(0), err
	}

	return result.BlockHeight.Value.Int, nil
}

// Get the transaction given slot and block height
func (c *PlasmaCashClient) PlasmaTx(blknum *big.Int, slot uint64) (plasma_cash.Tx, error) {
	blk := &types.BigUInt{Value: *loom.NewBigUInt(blknum)}
	var result pctypes.GetPlasmaTxResponse
	params := &pctypes.GetPlasmaTxRequest{
		BlockHeight: blk,
		Slot:        slot,
	}

	if _, err := c.loomcontract.StaticCall("GetPlasmaTxRequest", params, c.caller, &result); err != nil {
		return &plasma_cash.LoomTx{}, nil
	}

	tx := result.Plasmatx
	prevBlock := big.NewInt(0)
	if tx.GetPreviousBlock() != nil {
		prevBlock = tx.GetPreviousBlock().Value.Int
	}
	address := tx.NewOwner.Local.String()
	ethAddress := common.HexToAddress(address)

	return &plasma_cash.LoomTx{Slot: slot,
		PrevBlock:    prevBlock,
		Denomination: tx.Denomination.Value.Int,
		Owner:        ethAddress,
		Signature:    tx.Signature,
		TXProof:      tx.Proof}, nil
}

// Block get the block, transactions and proofs for a given block height
func (c *PlasmaCashClient) Block(blknum *big.Int) (plasma_cash.Block, error) {
	blk := &types.BigUInt{Value: *loom.NewBigUInt(blknum)}
	var result pctypes.GetBlockResponse
	params := &pctypes.GetBlockRequest{
		BlockHeight: blk,
	}

	if _, err := c.loomcontract.StaticCall("GetBlockRequest", params, c.caller, &result); err != nil {
		return &plasma_cash.PbBlock{}, nil
	}

	return plasma_cash.NewClientBlock(result.Block), nil
}

// SubmitBlock submits current plasma cash block to mainnet, useful for debugging
// ** note that only validators, or test clients use this method
// normal clients should not use it as it will be rejected by the server
func (c *PlasmaCashClient) SubmitBlock() error {
	request := &pctypes.SubmitBlockToMainnetRequest{}

	if _, err := c.loomcontract.Call("SubmitBlockToMainnet", request, c.signer, nil); err != nil {
		log.Fatalf("failed submitting block - %v\n", err)

		return err
	}

	log.Println("succeeded submitting a block ")

	return nil
}

// Deposit , submits a deposit from the Ethereum Blockchain onto the DAppChain
// ** note that only validators, or test clients use this method
// normal clients should not use it as it will be rejected by the server
func (c *PlasmaCashClient) Deposit(deposit *pctypes.DepositRequest) error {
	if _, err := c.loomcontract.Call("DepositRequest", deposit, c.signer, nil); err != nil {
		return errors.Wrap(err, "failed to commit DepositRequest tx")
	}
	return nil
}

// Sends a plasma cash transaciton to be added to the current plasma cash block
func (c *PlasmaCashClient) SendTransaction(slot uint64, prevBlock *big.Int, denomination *big.Int,
	newOwner, prevOwner string, sig []byte, hash []byte, replayProtectionSignature []byte) error {
	receiverAddr := loom.MustParseAddress(fmt.Sprintf("eth:%s", newOwner))
	senderAddr := loom.MustParseAddress(fmt.Sprintf("eth:%s", prevOwner))
	tx := &pctypes.PlasmaTx{
		Slot:          uint64(slot),
		PreviousBlock: &types.BigUInt{Value: *loom.NewBigUInt(prevBlock)},
		Denomination:  &types.BigUInt{Value: *loom.NewBigUInt(denomination)},
		NewOwner:      receiverAddr.MarshalPB(),
		Sender:        senderAddr.MarshalPB(),
		Signature:     sig,
		Hash:          hash,
	}

	params := &pctypes.PlasmaTxRequest{
		Plasmatx:                  tx,
		ReplayProtectionSignature: replayProtectionSignature,
	}

	if _, err := c.loomcontract.Call("PlasmaTxRequest", params, c.signer, nil); err != nil {
		log.Fatalf("failed trying to send transaction - %v\n", err)

		return err
	}

	return nil
}

func NewPlasmaCashClient(contractName string, signer auth.Signer, chainID, writeuri, readuri string) (plasma_cash.ChainServiceClient, error) {
	// create rpc client
	rpcClient := NewDAppChainRPCClient(chainID, writeuri, readuri)

	// try to resolve it using registry
	contractAddr, err := rpcClient.Resolve(contractName)
	if err != nil {
		return nil, err
	}

	// create contract
	contract := NewContract(rpcClient, contractAddr.Local)
	caller := loom.RootAddress(chainID)

	return &PlasmaCashClient{loomcontract: contract, signer: signer, caller: caller}, nil
}
