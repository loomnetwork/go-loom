// +build evm

package dposv2

import (
	"github.com/loomnetwork/go-loom"
	dpostypes "github.com/loomnetwork/go-loom/builtin/types/dposv2"
	"github.com/loomnetwork/go-loom/client"
	"github.com/pkg/errors"
	"math/big"
)

// DAppChainDPOSContract is a client-side binding for the builtin coin Go contract.
type DAppChainDPOSContract struct {
	contract *client.Contract
	chainID  string

	Address loom.Address
}

func ConnectToDAppChainDPOSContract(loomClient *client.DAppChainRPCClient) (*DAppChainDPOSContract, error) {
	contractAddr, err := loomClient.Resolve("dposV2")
	if err != nil {
		return nil, err
	}

	return &DAppChainDPOSContract{
		contract: client.NewContract(loomClient, contractAddr.Local),
		chainID:  loomClient.GetChainID(),
		Address:  contractAddr,
	}, nil
}

func (dpos *DAppChainDPOSContract) CheckDistributions(identity *client.Identity) (*big.Int, error) {
	req := &dpostypes.CheckDistributionRequest{}
	var resp dpostypes.CheckDistributionResponse
	_, err := dpos.contract.StaticCall("CheckDistribution", req, identity.LoomAddr, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Amount.Value.Int, err
}

func (dpos *DAppChainDPOSContract) ListCandidates(identity *client.Identity) ([]*dpostypes.CandidateV2, error) {
	req := &dpostypes.ListCandidateRequestV2{}
	var resp dpostypes.ListCandidateResponseV2
	_, err := dpos.contract.StaticCall("ListCandidates", req, identity.LoomAddr, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Candidates, err
}

func (dpos *DAppChainDPOSContract) ListValidators(identity *client.Identity) ([]*dpostypes.ValidatorStatisticV2, error) {
	req := &dpostypes.ListValidatorsRequestV2{}
	var resp dpostypes.ListValidatorsResponseV2
	_, err := dpos.contract.StaticCall("ListValidators", req, identity.LoomAddr, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Statistics, err
}

func (dpos *DAppChainDPOSContract) ProcessRequestBatch(identity *client.Identity, req *dpostypes.RequestBatchV2) error {
	_, err := dpos.contract.Call("ProcessRequestBatch", req, identity.LoomSigner, nil)
	return err
}

func (dpos *DAppChainDPOSContract) GetRequestBatchTally(identity *client.Identity) (*dpostypes.RequestBatchTallyV2, error) {
	req := &dpostypes.GetRequestBatchTallyRequestV2{}
	resp := &dpostypes.RequestBatchTallyV2{}
	if _, err := dpos.contract.StaticCall("GetRequestBatchTally", req, identity.LoomAddr, resp); err != nil {
		return nil, errors.Wrap(err, "failed to get request batch tally")
	}

	return resp, nil
}

func (dpos *DAppChainDPOSContract) ChangeFee(identity *client.Identity, candidateFee uint64) error {
	req := &dpostypes.ChangeCandidateFeeRequest{
		Fee: candidateFee,
	}
	_, err := dpos.contract.Call("ChangeFee", req, identity.LoomSigner, nil)
	return err
}

func (dpos *DAppChainDPOSContract) ClaimRewards(identity *client.Identity, addr loom.Address) (*dpostypes.ClaimDistributionResponseV2, error) {
	req := &dpostypes.ClaimDistributionRequestV2{
		WithdrawalAddress: addr.MarshalPB(),
	}
	resp := &dpostypes.ClaimDistributionResponseV2{}

	_, err := dpos.contract.Call("ClaimDistribution", req, identity.LoomSigner, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (dpos *DAppChainDPOSContract) RegisterCandidate(identity *client.Identity, pubKey []byte, candidateFee uint64, candidateName string, candidateDescription string, candidateWebsite string) error {
	req := &dpostypes.RegisterCandidateRequestV2{
		PubKey:      pubKey,
		Fee:         candidateFee,
		Name:        candidateName,
		Description: candidateDescription,
		Website:     candidateWebsite,
	}
	_, err := dpos.contract.Call("RegisterCandidate", req, identity.LoomSigner, nil)
	return err
}

func (dpos *DAppChainDPOSContract) UnregisterCandidate(identity *client.Identity) error {
	req := &dpostypes.UnregisterCandidateRequestV2{}
	_, err := dpos.contract.Call("UnregisterCandidate", req, identity.LoomSigner, nil)
	return err
}
