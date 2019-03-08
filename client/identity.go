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
	"golang.org/x/crypto/ed25519"
)

type Identity struct {
	MainnetPrivKey *ecdsa.PrivateKey
	MainnetAddr    common.Address
	LoomSigner     auth.Signer
	LoomAddr       loom.Address
}

func CreateRandomIdentity(chainID string) (*Identity, error) {
    return CreateIdentity(nil, nil, chainID)
}

func CreateIdentity(hexKeyOrECDSA interface{}, signerOrKey interface{}, chainID string) (*Identity, error) {
	var mainnetPrivKey *ecdsa.PrivateKey
	var signer auth.Signer
	var err error
	var loomAddr loom.Address
	var ethAddr common.Address

	// Convert hex key to crypto.ECDSA key
	if hexKeyOrECDSA == nil {
		mainnetPrivKey, err = crypto.GenerateKey()
		if err != nil {
			return nil, err
		}
	} else {
		switch hexKeyOrECDSA.(type) {
		case *ecdsa.PrivateKey:
			mainnetPrivKey = hexKeyOrECDSA.(*ecdsa.PrivateKey)
			break
		case string:
			mainnetPrivKey, err = crypto.HexToECDSA(strings.TrimPrefix(hexKeyOrECDSA.(string), "0x"))
			if err != nil {
				return nil, err
			}
			break
		default:
			return nil, errors.New("Invalid mainnet key/signer type")
		}
	}
	ethAddr = crypto.PubkeyToAddress(mainnetPrivKey.PublicKey)

	// Convert dappchain key to signer
	if signerOrKey == nil {
		_, priv, err := ed25519.GenerateKey(nil)
		if err != nil {
			return nil, err
		}
		signer = auth.NewEd25519Signer(priv)
		if err != nil {
			return nil, err
		}
	} else {
		switch signerOrKey.(type) {
		case auth.Signer:
			signer = signerOrKey.(auth.Signer)
			break
		case string:
			privKey, err := base64.StdEncoding.DecodeString(signerOrKey.(string))
			if err != nil {
				return nil, err
			}
			signer = auth.NewEd25519Signer(privKey)

			break
		default:
			return nil, errors.New("Invalid dappchain key/signer type")
		}
	}
	loomAddr = loom.Address{
		ChainID: chainID,
		Local:   loom.LocalAddressFromPublicKey(signer.PublicKey()),
	}

	identity := &Identity{
		MainnetPrivKey: mainnetPrivKey,
		MainnetAddr:    ethAddr,
		LoomSigner:     signer,
		LoomAddr:       loomAddr,
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
