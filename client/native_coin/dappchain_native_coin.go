package native_coin

import (
	"math/big"

	"github.com/loomnetwork/go-loom"
	"github.com/loomnetwork/go-loom/builtin/types/coin"
	"github.com/loomnetwork/go-loom/client"
	"github.com/loomnetwork/go-loom/types"
)

// DAppChainNativeCoin is a client-side binding for the builtin coin Go contracts.
type DAppChainNativeCoin struct {
	contract *client.Contract
	chainID  string

	Address loom.Address
}

func (ec *DAppChainNativeCoin) BalanceOf(identity *client.Identity) (*big.Int, error) {
	ownerAddr := loom.Address{
		ChainID: ec.chainID,
		Local:   identity.LoomAddr.Local,
	}
	req := &coin.BalanceOfRequest{
		Owner: ownerAddr.MarshalPB(),
	}
	var resp coin.BalanceOfResponse
	_, err := ec.contract.StaticCall("BalanceOf", req, ownerAddr, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Balance != nil {
		return resp.Balance.Value.Int, nil
	}
	return nil, nil
}

func (ec *DAppChainNativeCoin) Approve(owner *client.Identity, spender loom.Address, amount *big.Int) error {
	req := &coin.ApproveRequest{
		Spender: spender.MarshalPB(),
		Amount:  &types.BigUInt{Value: *loom.NewBigUInt(amount)},
	}
	_, err := ec.contract.Call("Approve", req, owner.LoomSigner, nil)
	return err
}

func (ec *DAppChainNativeCoin) Transfer(owner *client.Identity, to loom.Address, amount *big.Int) error {
	req := &coin.TransferRequest{
		To:     to.MarshalPB(),
		Amount: &types.BigUInt{Value: *loom.NewBigUInt(amount)},
	}
	_, err := ec.contract.Call("Transfer", req, owner.LoomSigner, nil)
	return err
}

/** Connectors */

func ConnectToDAppChainLoomContract(loomClient *client.DAppChainRPCClient) (*DAppChainNativeCoin, error) {
	return connectToDAppChainNativeCoin(loomClient, "coin")
}

func ConnectToDAppChainETHContract(loomClient *client.DAppChainRPCClient) (*DAppChainNativeCoin, error) {
	return connectToDAppChainNativeCoin(loomClient, "ethcoin")
}

func connectToDAppChainNativeCoin(loomClient *client.DAppChainRPCClient, name string) (*DAppChainNativeCoin, error) {
	contractAddr, err := loomClient.Resolve(name)
	if err != nil {
		return nil, err
	}

	return &DAppChainNativeCoin{
		contract: client.NewContract(loomClient, contractAddr.Local),
		chainID:  loomClient.GetChainID(),
		Address:  contractAddr,
	}, nil
}
