// +build evm

package dposv3

import (
	"github.com/loomnetwork/go-loom"
	dpostypes "github.com/loomnetwork/go-loom/builtin/types/dposv3"
	"github.com/loomnetwork/go-loom/client"
	"math/big"
)

// DAppChainDPOSContract is a client-side binding for the builtin coin Go contract.
type DAppChainDPOSContract struct {
	contract *client.Contract
	chainID  string

	Address loom.Address
}

func ConnectToDAppChainDPOSContract(loomClient *client.DAppChainRPCClient) (*DAppChainDPOSContract, error) {
	contractAddr, err := loomClient.Resolve("dposV3")
	if err != nil {
		return nil, err
	}

	return &DAppChainDPOSContract{
		contract: client.NewContract(loomClient, contractAddr.Local),
		chainID:  loomClient.GetChainID(),
		Address:  contractAddr,
	}, nil
}

// Check and Claim rewards client. TODO: Implement the rest.
func (dpos *DAppChainDPOSContract) CheckRewardsFromAllValidators(identity *client.Identity, address loom.Address) (*big.Int, error) {
	req := &dpostypes.CheckDelegatorRewardsRequest{
		Delegator: address.MarshalPB(),
	}
	var resp dpostypes.CheckDelegatorRewardsResponse
	_, err := dpos.contract.StaticCall("CheckRewardsFromAllValidators", req, identity.LoomAddr, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Amount.Value.Int, err
}

func (dpos *DAppChainDPOSContract) ClaimRewardsFromAllValidators(identity *client.Identity) (*big.Int, error) {
	req := &dpostypes.ClaimDelegatorRewardsRequest{}
	var resp dpostypes.ClaimDelegatorRewardsResponse
	_, err := dpos.contract.Call("ClaimRewardsFromAllValidators", req, identity.LoomSigner, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Amount.Value.Int, err
}

func (dpos *DAppChainDPOSContract) TimeUntilElections(identity *client.Identity) (int64, error) {
	req := &dpostypes.TimeUntilElectionRequest{}
	var resp dpostypes.TimeUntilElectionResponse
	_, err := dpos.contract.StaticCall("TimeUntilElection", req, identity.LoomAddr, &resp)
	if err != nil {
		return -1, err
	}
	return resp.TimeUntilElection, nil
}

func (dpos *DAppChainDPOSContract) GetRewardsDelegation(identity *client.Identity, address loom.Address) (*dpostypes.Delegation, error) {
	req := &dpostypes.CheckRewardDelegationRequest{
		ValidatorAddress: address.MarshalPB(),
	}
	var resp dpostypes.CheckRewardDelegationResponse
	_, err := dpos.contract.StaticCall("CheckRewardDelegation", req, identity.LoomAddr, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Delegation, nil
}
