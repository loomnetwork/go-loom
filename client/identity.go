// +build evm

package client

import (
	"crypto/ecdsa"
	"encoding/base64"
	"github.com/pkg/errors"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/loomnetwork/go-loom"
	"github.com/loomnetwork/go-loom/auth"
)

type Identity struct {
	MainnetPrivKey *ecdsa.PrivateKey
	MainnetAddr    common.Address
	LoomSigner     auth.Signer
	LoomAddr       loom.Address
}

func CreateIdentity(hexKeyOrECDSA interface{}, signerOrKey interface{}, chainID string) (*Identity, error) {
	var mainnetPrivKey *ecdsa.PrivateKey
	var signer auth.Signer
	var err error

	// Convert hex key to crypto.ECDSA key
	switch signerOrKey.(type) {
	case *ecdsa.PrivateKey:
		mainnetPrivKey = hexKeyOrECDSA.(*ecdsa.PrivateKey)
	case string:
		mainnetPrivKey, err = crypto.HexToECDSA(strings.TrimPrefix(hexKeyOrECDSA.(string), "0x"))
		if err != nil {
			return nil, err
		}
		break
	default:
		return nil, errors.New("Invalid mainnet key/signer type")
	}

	// Convert dappchain key to crypto.ECDSA key
	switch signerOrKey.(type) {
	case auth.Signer:
		signer = signerOrKey.(auth.Signer)
	case string:
		privKey, err := base64.StdEncoding.DecodeString(signerOrKey.(string))
		if err != nil {
			return nil, err
		}

		signer = auth.NewEd25519Signer(privKey)
		if err != nil {
			return nil, err
		}
		break
	default:
		return nil, errors.New("Invalid dappchain key/signer type")
	}

	identity := &Identity{
		MainnetPrivKey: mainnetPrivKey,
		MainnetAddr:    crypto.PubkeyToAddress(mainnetPrivKey.PublicKey),
		LoomSigner:     signer,
		LoomAddr: loom.Address{
			ChainID: chainID,
			Local:   loom.LocalAddressFromPublicKey(signer.PublicKey()),
		},
	}
	return identity, nil
}

func LoomAddressFromEthereumAddress(ethAddr common.Address) (loom.Address, error) {
	addrBytes, err := loom.LocalAddressFromHexString(ethAddr.Hex())
	if err != nil {
		return loom.Address{}, err
	}
	return loom.Address{
		ChainID: "eth",
		Local:   addrBytes,
	}, nil
}
