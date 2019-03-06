// +build evm

package native_coin

import (
	"math/big"

	"github.com/loomnetwork/go-loom"
	"github.com/loomnetwork/go-loom/auth"
	"github.com/loomnetwork/go-loom/builtin/types/coin"
	"github.com/loomnetwork/go-loom/client"
	"github.com/loomnetwork/go-loom/types"
)

// DAppChainNativeCoin is a client-side binding for the builtin coin Go contracts.
type DAppChainNativeCoin struct {
	contract *client.Contract
	chainID  string
	Signer   auth.Signer

	SignerAddress loom.Address
	Address       loom.Address
}

func (ec *DAppChainNativeCoin) toLoomAddr(addr string) (loom.Address, error) {
	local, err := loom.LocalAddressFromHexString(addr)
	if err != nil {
		return loom.RootAddress(ec.chainID), err
	}
	ownerAddr := loom.Address{
		ChainID: ec.chainID,
		Local:   local,
	}
	return ownerAddr, nil
}

func (ec *DAppChainNativeCoin) BalanceOf(ownerAddrStr string) (*big.Int, error) {

	ownerAddr, err := ec.toLoomAddr(ownerAddrStr)
	if err != nil {
		return nil, err
	}

	req := &coin.BalanceOfRequest{
		Owner: ownerAddr.MarshalPB(),
	}
	var resp coin.BalanceOfResponse
	_, err = ec.contract.StaticCall("BalanceOf", req, ownerAddr, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Balance != nil {
		return resp.Balance.Value.Int, nil
	}
	return nil, nil
}

func (ec *DAppChainNativeCoin) Approve(spenderAddrStr string, amount *big.Int) error {
	spenderAddr, err := ec.toLoomAddr(spenderAddrStr)
	if err != nil {
		return err
	}

	req := &coin.ApproveRequest{
		Spender: spenderAddr.MarshalPB(),
		Amount:  &types.BigUInt{Value: *loom.NewBigUInt(amount)},
	}
	_, err = ec.contract.Call("Approve", req, ec.Signer, nil)
	return err
}

func (ec *DAppChainNativeCoin) Transfer(toAddrStr string, amount *big.Int) error {
	toAddr, err := ec.toLoomAddr(toAddrStr)
	if err != nil {
		return err
	}

	req := &coin.TransferRequest{
		To:     toAddr.MarshalPB(),
		Amount: &types.BigUInt{Value: *loom.NewBigUInt(amount)},
	}
	_, err = ec.contract.Call("Transfer", req, ec.Signer, nil)
	return err
}

/** Connectors */

func ConnectToDAppChainLoomContract(loomClient *client.DAppChainRPCClient, signer auth.Signer) (*DAppChainNativeCoin, error) {
	return connectToDAppChainNativeCoin(loomClient, signer, "coin")
}

func ConnectToDAppChainETHContract(loomClient *client.DAppChainRPCClient, signer auth.Signer) (*DAppChainNativeCoin, error) {
	return connectToDAppChainNativeCoin(loomClient, signer, "ethcoin")
}

func connectToDAppChainNativeCoin(loomClient *client.DAppChainRPCClient, signer auth.Signer, name string) (*DAppChainNativeCoin, error) {
	contractAddr, err := loomClient.Resolve(name)
	if err != nil {
		return nil, err
	}

	localAddr := loom.LocalAddressFromPublicKey(signer.PublicKey())
	signerAddress := loom.Address{
		ChainID: loomClient.GetChainID(),
		Local:   localAddr,
	}

	return &DAppChainNativeCoin{
		contract:      client.NewContract(loomClient, contractAddr.Local),
		Signer:        signer,
		SignerAddress: signerAddress,
		chainID:       loomClient.GetChainID(),
		Address:       contractAddr,
	}, nil
}
