// +build evm

package dposv2

import (
	"github.com/loomnetwork/go-loom"
	"github.com/loomnetwork/go-loom/auth"
	dpostypes "github.com/loomnetwork/go-loom/builtin/types/dposv2"
	"github.com/loomnetwork/go-loom/client"
	"github.com/pkg/errors"
)

// DAppChainDPOSContract is a client-side binding for the builtin coin Go contract.
type DAppChainDPOSContract struct {
	contract      *client.Contract
	chainID       string
	signer        auth.Signer
	SignerAddress loom.Address
	Address       loom.Address
}

func ConnectToDAppChainDPOSContract(loomClient *client.DAppChainRPCClient, signer auth.Signer) (*DAppChainDPOSContract, error) {
	contractAddr, err := loomClient.Resolve("dposV2")
	if err != nil {
		return nil, err
	}

	localAddr := loom.LocalAddressFromPublicKey(signer.PublicKey())
	signerAddress := loom.Address{
		ChainID: loomClient.GetChainID(),
		Local:   localAddr,
	}

	return &DAppChainDPOSContract{
		contract:      client.NewContract(loomClient, contractAddr.Local),
		chainID:       loomClient.GetChainID(),
		signer:        signer,
		SignerAddress: signerAddress,
		Address:       contractAddr,
	}, nil
}

func (dpos *DAppChainDPOSContract) ListCandidates() ([]*dpostypes.CandidateV2, error) {
	req := &dpostypes.ListCandidateRequestV2{}
	var resp dpostypes.ListCandidateResponseV2
	_, err := dpos.contract.StaticCall("ListCandidates", req, dpos.SignerAddress, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Candidates, err
}

func (dpos *DAppChainDPOSContract) ListValidators() ([]*dpostypes.ValidatorStatisticV2, error) {
	req := &dpostypes.ListValidatorsRequestV2{}
	var resp dpostypes.ListValidatorsResponseV2
	_, err := dpos.contract.StaticCall("ListValidators", req, dpos.SignerAddress, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Statistics, err
}

func (dpos *DAppChainDPOSContract) ClaimRewards(withdrawalAddress string) (*dpostypes.ClaimDistributionResponseV2, error) {
	local, err := loom.LocalAddressFromHexString(withdrawalAddress)
	if err != nil {
		return nil, err
	}
	addr := loom.Address{
		ChainID: dpos.chainID,
		Local:   local,
	}

	req := &dpostypes.ClaimDistributionRequestV2{
		WithdrawalAddress: addr.MarshalPB(),
	}
	resp := &dpostypes.ClaimDistributionResponseV2{}

	_, err = dpos.contract.Call("ClaimDistribution", req, dpos.signer, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (dpos *DAppChainDPOSContract) ProcessRequestBatch(req *dpostypes.RequestBatchV2) error {
	_, err := dpos.contract.Call("ProcessRequestBatch", req, dpos.signer, nil)
	return err
}

func (dpos *DAppChainDPOSContract) GetRequestBatchTally() (*dpostypes.RequestBatchTallyV2, error) {
	req := &dpostypes.GetRequestBatchTallyRequestV2{}
	resp := &dpostypes.RequestBatchTallyV2{}
	if _, err := dpos.contract.StaticCall("GetRequestBatchTally", req, dpos.SignerAddress, resp); err != nil {
		return nil, errors.Wrap(err, "failed to get request batch tally")
	}

	return resp, nil
}

func (dpos *DAppChainDPOSContract) ChangeFee(candidateFee uint64) error {
	req := &dpostypes.ChangeCandidateFeeRequest{
		Fee: candidateFee,
	}
	_, err := dpos.contract.Call("ChangeFee", req, dpos.signer, nil)
	return err
}

func (dpos *DAppChainDPOSContract) RegisterCandidate(pubKey []byte, candidateFee uint64, candidateName string, candidateDescription string, candidateWebsite string) error {
	req := &dpostypes.RegisterCandidateRequestV2{
		PubKey:      pubKey,
		Fee:         candidateFee,
		Name:        candidateName,
		Description: candidateDescription,
		Website:     candidateWebsite,
	}
	_, err := dpos.contract.Call("RegisterCandidate", req, dpos.signer, nil)
	return err
}

func (dpos *DAppChainDPOSContract) UnregisterCandidate() error {
	req := &dpostypes.UnregisterCandidateRequestV2{}
	_, err := dpos.contract.Call("UnregisterCandidate", req, dpos.signer, nil)
	return err
}
