// +build evm

package client

import (
	"fmt"
	"log"

	loom "github.com/loomnetwork/go-loom"
	"github.com/loomnetwork/go-loom/auth"
	pctypes "github.com/loomnetwork/go-loom/builtin/types/plasma_cash"
	"github.com/loomnetwork/go-loom/client/plasma_cash"
	"github.com/loomnetwork/go-loom/types"
)

// PlasmaCashClient client to loom plasma cash server
type PlasmaCashClient struct {
	url          string
	ChainID      string
	WriteURI     string
	ReadURI      string
	contractAddr string
	loomcontract *Contract
	caller       loom.Address
	signer       auth.Signer
}

// CurrentBlock gets the highest block of plasma cash
// does not grab pending transactions yet
func (c *PlasmaCashClient) CurrentBlock() (plasma_cash.Block, error) {
	return c.Block(0) //asking for block zero gives latest
}

// BlockNumber gets the current plasma cash block height
func (c *PlasmaCashClient) BlockNumber() (int64, error) {
	request := &pctypes.GetCurrentBlockRequest{}
	var result pctypes.GetCurrentBlockResponse

	if _, err := c.loomcontract.StaticCall("GetCurrentBlockRequest", request, c.caller, &result); err != nil {
		log.Fatalf("failed getting Block number - %v\n", err)

		return 0, err
	}

	return result.BlockHeight.Value.Int64(), nil
}

// Block get the block, transactions and proofs for a given block height
func (c *PlasmaCashClient) Block(blknum int64) (plasma_cash.Block, error) {
	blk := &types.BigUInt{*loom.NewBigUIntFromInt(blknum)}

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

// Sends a plasma cash transaciton to be added to the current plasma cash block
func (c *PlasmaCashClient) SendTransaction(slot uint64, prevBlock int64, denomination int64, newOwner string, sig []byte) error {
	loomAddress := fmt.Sprintf("chain:%s", newOwner)

	address := loom.MustParseAddress(loomAddress)
	tx := &pctypes.PlasmaTx{
		Slot:          uint64(slot),
		PreviousBlock: &types.BigUInt{*loom.NewBigUIntFromInt(prevBlock)},
		Denomination:  &types.BigUInt{*loom.NewBigUIntFromInt(denomination)},
		NewOwner:      address.MarshalPB(),
		Signature:     sig,
	}

	params := &pctypes.PlasmaTxRequest{
		Plasmatx: tx,
	}

	if _, err := c.loomcontract.Call("PlasmaTxRequest", params, c.signer, nil); err != nil {
		log.Fatalf("failed trying to send transaction - %v\n", err)

		return err
	}

	return nil
}

func NewPlasmaCashClient(signer auth.Signer, chainID, readuri, writeuri string) (plasma_cash.ChainServiceClient, error) {
	//for now assume plasmacash
	s := "plasmacash"

	// create rpc client
	rpcClient := NewDAppChainRPCClient(chainID, writeuri, readuri)

	// try to resolve it using registry
	contractAddr, err := rpcClient.Resolve(s)
	if err != nil {
		return nil, err
	}

	// create contract
	contract := NewContract(rpcClient, contractAddr.Local)
	caller := loom.RootAddress(chainID)

	return &PlasmaCashClient{loomcontract: contract, signer: signer, caller: caller}, nil
}
