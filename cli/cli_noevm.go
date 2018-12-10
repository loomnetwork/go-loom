// +build !evm

package cli

import (
	"encoding/base64"
	"errors"
	"io/ioutil"

	"github.com/gogo/protobuf/proto"
	"github.com/loomnetwork/go-loom/auth"
)

func CallContract(defaultAddr string, method string, params proto.Message, result interface{}) error {
	if TxFlags.PrivFile == "" {
		return errors.New("private key required to call contract")
	}

	privKeyB64, err := ioutil.ReadFile(TxFlags.PrivFile)
	if err != nil {
		return err
	}

	privKey, err := base64.StdEncoding.DecodeString(string(privKeyB64))
	if err != nil {
		return err
	}

	signer := auth.NewSigner(auth.SignerTypeEd25519, privKey)

	contract, err := contract(defaultAddr)
	if err != nil {
		return err
	}
	_, err = contract.Call(method, params, signer, result)
	return err
}
