// +build evm

package auth

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/loomnetwork/go-loom"
	"github.com/loomnetwork/go-loom/common/evmcompat"
)

const (
	EthChainId = "eth"
)

type EthSigner struct {
	PrivateKey *ecdsa.PrivateKey
}

func NewEthSigner(_ []byte) Signer {
	panic("EVM build isn't activated")
}

func (k *EthSigner) Sign(_ []byte) []byte {
	sigBytes, err := evmcompat.SoliditySign(k.PublicKey(), k.PrivateKey)
	if err != nil {
		panic(err)
	}
	return sigBytes
}

func (k *EthSigner) PublicKey() []byte {
	ethLocalAdr, err := loom.LocalAddressFromHexString(crypto.PubkeyToAddress(k.PrivateKey.PublicKey).Hex())
	if err != nil {
		panic(err)
	}
	ethPublicAddr := loom.Address{ChainID: EthChainId, Local: ethLocalAdr}
	hash := crypto.Keccak256(ethPublicAddr.Bytes())

	return hash
}
