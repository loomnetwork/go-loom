// +build evm

package plasma_cash

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	pctypes "github.com/loomnetwork/go-loom/builtin/types/plasma_cash"
)

// Plasma Block
type PbBlock struct {
	block *pctypes.PlasmaBlock
}

func NewClientBlock(pb *pctypes.PlasmaBlock) Block {
	return &PbBlock{pb}
}

func (p *PbBlock) MerkleHash() []byte {
	return p.block.MerkleHash
}

func (p *PbBlock) TxFromSlot(slot uint64) (Tx, error) {
	var tx *pctypes.PlasmaTx

	if p.block.Transactions == nil {
		return nil, nil
	}
	for _, v := range p.block.Transactions {
		if v.Slot == slot {
			tx = v
			break
		}
	}
	if tx == nil {
		return nil, fmt.Errorf("can't find transaction at slot %d. We had %d Transactions\n", slot, len(p.block.Transactions))
	}

	prevBlock := big.NewInt(0)
	if tx.GetPreviousBlock() != nil {
		prevBlock = tx.GetPreviousBlock().Value.Int
	}
	address := tx.NewOwner.Local.String()
	ethAddress := common.HexToAddress(address)

	return &LoomTx{Slot: slot,
		PrevBlock:    prevBlock,
		Denomination: tx.Denomination.Value.Int,
		Owner:        ethAddress,
		Signature:    tx.Signature,
		TXProof:      tx.Proof}, nil
}
