// +build evm

package address_mapper

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/loomnetwork/go-loom"
	"github.com/loomnetwork/go-loom/builtin/types/address_mapper"
	"github.com/loomnetwork/go-loom/client"
	"github.com/loomnetwork/go-loom/common/evmcompat"
	ssha "github.com/miguelmota/go-solidity-sha3"
)

var (
	SupportedChainID = []string{"eth", "tron", "binance"}
)

type DAppChainAddressMapper struct {
	contract *client.Contract
	chainID  string
}

func ConnectToDAppChainAddressMapper(loomClient *client.DAppChainRPCClient) (*DAppChainAddressMapper, error) {
	mapperAddr, err := loomClient.Resolve("addressmapper")
	if err != nil {
		return nil, err
	}
	return &DAppChainAddressMapper{
		contract: client.NewContract(loomClient, mapperAddr.Local),
		chainID:  loomClient.GetChainID(),
	}, nil
}

func (am *DAppChainAddressMapper) GetMappedAccount(account loom.Address) (loom.Address, error) {
	req := &address_mapper.AddressMapperGetMappingRequest{
		From: account.MarshalPB(),
	}
	resp := &address_mapper.AddressMapperGetMappingResponse{}
	_, err := am.contract.StaticCall("GetMapping", req, account, resp)
	if err != nil {
		return loom.Address{}, err
	}
	return loom.UnmarshalAddressPB(resp.To), nil
}

func (am *DAppChainAddressMapper) HasIdentityMapping(account loom.Address) (bool, error) {
	req := &address_mapper.AddressMapperHasMappingRequest{
		From: account.MarshalPB(),
	}
	resp := &address_mapper.AddressMapperHasMappingResponse{}
	_, err := am.contract.StaticCall("HasMapping", req, account, resp)
	if err != nil {
		return false, err
	}
	return resp.HasMapping, nil
}

func (am *DAppChainAddressMapper) GetNonce(account loom.Address) (uint64, error) {
	fmt.Println("GetNONCE CALLED")
	req := &address_mapper.AddressMapperGetNonceRequest{
		Address: account.MarshalPB(),
	}
	resp := &address_mapper.AddressMapperGetNonceResponse{}
	fmt.Printf("CALL GETNONCE with REQUEST %+v \n", req.Address.String())
	_, err := am.contract.StaticCall("GetNonce", req, account, resp)
	if err != nil {
		fmt.Println("GETNONCE ERROR IN go-loom")
		return uint64(0), err
	}
	fmt.Println("GetNONCE PASS IN go-loom")
	return resp.Nonce, nil
}

// AddIdentityMapping creates a bi-directional mapping between a Mainnet & DAppChain account.
func (am *DAppChainAddressMapper) AddIdentityMapping(identity *client.Identity) error {
	mainnetAddrBytes, err := loom.LocalAddressFromHexString(identity.MainnetAddr.Hex())
	if err != nil {
		return err
	}
	from := loom.Address{
		ChainID: "eth",
		Local:   mainnetAddrBytes,
	}
	to := loom.Address{
		ChainID: am.chainID,
		Local:   identity.LoomAddr.Local,
	}

	mappedAccount, err := am.GetMappedAccount(from)
	if err == nil {
		if mappedAccount.Compare(to) != 0 {
			return fmt.Errorf("Account %v is mapped to %v", from, mappedAccount)
		}
		return nil
	}

	nonce, err := am.GetNonce(to)
	if err != nil {
		return err
	}
	fmt.Printf("Mapping account %v to %v\n", from, to)
	sig, err := signIdentityMapping(from, to, identity.MainnetPrivKey, nonce)
	if err != nil {
		return err
	}
	req := &address_mapper.AddressMapperAddIdentityMappingRequest{
		From:      from.MarshalPB(),
		To:        to.MarshalPB(),
		Signature: sig,
	}
	_, err = am.contract.Call("AddIdentityMapping", req, identity.LoomSigner, nil)
	return err
}

func signIdentityMapping(from, to loom.Address, key *ecdsa.PrivateKey, nonce uint64) ([]byte, error) {
	var foreignChainID string
	for _, c := range SupportedChainID {
		if from.ChainID == c {
			foreignChainID = c
			break
		}
		if to.ChainID == c {
			foreignChainID = c
			break
		}
	}
	hash := ssha.SoliditySHA3(
		ssha.Address(common.BytesToAddress(from.Local)),
		ssha.Address(common.BytesToAddress(to.Local)),
		// ssha.Uint64(nonce),
		ssha.String(foreignChainID),
	)
	sig, err := evmcompat.SoliditySign(hash, key)
	if err != nil {
		return nil, err
	}
	// Prefix the sig with a single byte indicating the sig type, in this case EIP712
	return append(make([]byte, 1, 66), sig...), nil
}
