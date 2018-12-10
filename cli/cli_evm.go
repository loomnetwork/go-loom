// +build evm

package cli

import (
	"errors"

	"github.com/loomnetwork/go-loom/crypto"

	"github.com/gogo/protobuf/proto"
	"github.com/loomnetwork/go-loom/auth"
)

func CallContract(defaultAddr string, method string, params proto.Message, result interface{}) error {
	var signerType string

	if TxFlags.PrivFile == "" {
		return errors.New("private key required to call contract")
	}

	privKey, err := crypto.LoadECDSA(TxFlags.HsmEnabled, TxFlags.PrivFile)
	if err != nil {
		return err
	}

	if TxFlags.HsmEnabled {
		signerType = auth.SignerTypeYubiHsm
	} else {
		signerType = auth.SignerTypeSecp256k1
	}

	signer := auth.NewSigner(signerType, privKey)
	contract, err := contract(defaultAddr)
	if err != nil {
		return err
	}
	_, err = contract.Call(method, params, signer, result)
	return err
}
