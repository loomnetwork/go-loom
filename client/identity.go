// +build evm

package client

import (
	"crypto/ecdsa"
	"encoding/base64"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	loom "github.com/loomnetwork/go-loom"
	"github.com/loomnetwork/go-loom/auth"
)

type Identity struct {
	MainnetPrivKey *ecdsa.PrivateKey
	MainnetAddr    common.Address
	LoomSigner     auth.Signer
	LoomAddr       loom.Address
}

func CreateIdentity(mainnetPrivKey *ecdsa.PrivateKey, loomSigner auth.Signer, chainID string) (*Identity, error) {
	var mainnetAddr common.Address
	if mainnetPrivKey != nil {
		mainnetAddr = crypto.PubkeyToAddress(mainnetPrivKey.PublicKey)
	} else {
		mainnetAddr = common.HexToAddress("0x0")
	}
	identity := &Identity{
		MainnetPrivKey: mainnetPrivKey,
		MainnetAddr:    mainnetAddr,
		LoomSigner:     loomSigner,
		LoomAddr: loom.Address{
			ChainID: chainID,
			Local:   loom.LocalAddressFromPublicKey(loomSigner.PublicKey()),
		},
	}
	return identity, nil
}

func CreateIdentityStr(ethKey string, dappchainKey string, chainID string) (*Identity, error) {
	mainnetPrivKey, err := crypto.HexToECDSA(strings.TrimPrefix(ethKey, "0x"))
	if err != nil {
		return nil, err
	}
	keyBytes, err := base64.StdEncoding.DecodeString(dappchainKey)
	if err != nil {
		return nil, err
	}
	loomSigner := auth.NewEd25519Signer(keyBytes)
	identity := &Identity{
		MainnetPrivKey: mainnetPrivKey,
		MainnetAddr:    crypto.PubkeyToAddress(mainnetPrivKey.PublicKey),
		LoomSigner:     loomSigner,
		LoomAddr: loom.Address{
			ChainID: chainID,
			Local:   loom.LocalAddressFromPublicKey(loomSigner.PublicKey()),
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

func LoomAddressFromTronAddress(tronAddr common.Address) (loom.Address, error) {
	addrBytes, err := loom.LocalAddressFromHexString(tronAddr.Hex())
	if err != nil {
		return loom.Address{}, err
	}
	return loom.Address{
		ChainID: "tron",
		Local:   addrBytes,
	}, nil
}

func LoomAddressFromBinanceAddress(binanceAddr common.Address) loom.Address {
	return loom.Address{ChainID: "binance", Local: binanceAddr.Bytes()}
}
