// +build evm

package dposv2

import (
	"github.com/loomnetwork/go-loom"
	dpostypes "github.com/loomnetwork/go-loom/builtin/types/dposv2"
	"github.com/loomnetwork/go-loom/client"
	"github.com/pkg/errors"
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

func (dpos *DAppChainDPOSContract) ListCandidates(identity *client.Identity) ([]*dpostypes.CandidateV2, error) {
	owner := loom.Address{
		ChainID: dpos.chainID,
		Local:   identity.LoomAddr.Local,
	}
	req := &dpostypes.ListCandidateRequestV2{}
	var resp dpostypes.ListCandidateResponseV2
	_, err := dpos.contract.StaticCall("ListCandidates", req, owner, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Candidates, err
}

func (dpos *DAppChainDPOSContract) ListValidators(identity *client.Identity) ([]*dpostypes.ValidatorStatisticV2, error) {
	owner := loom.Address{
		ChainID: dpos.chainID,
		Local:   identity.LoomAddr.Local,
	}
	req := &dpostypes.ListValidatorsRequestV2{}
	var resp dpostypes.ListValidatorsResponseV2
	_, err := dpos.contract.StaticCall("ListValidators", req, owner, &resp)
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
	caller := loom.Address{
		ChainID: dpos.chainID,
		Local:   identity.LoomAddr.Local,
	}
	req := &dpostypes.GetRequestBatchTallyRequestV2{}
	resp := &dpostypes.RequestBatchTallyV2{}
	if _, err := dpos.contract.StaticCall("GetRequestBatchTally", req, caller, resp); err != nil {
		return nil, errors.Wrap(err, "failed to get request batch tally")
	}

	return resp, nil
}

func (dpos *DAppChainDPOSContract) ChangeFee(identity *client.Identity, candidateFee uint64) error {
	req := &dpostypes.ChangeCandidateFeeRequest{
		Fee:         candidateFee,
	}
	_, err := dpos.contract.Call("ChangeFee", req, identity.LoomSigner, nil)
	return err
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
