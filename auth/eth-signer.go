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

type EthSigner66Byte struct {
	PrivateKey *ecdsa.PrivateKey
}

func NewEthSigner66Byte(_ []byte) Signer {
	panic("EVM build isn't activated")
}

func (k *EthSigner66Byte) Sign(_ []byte) []byte {
	sigBytes, err := evmcompat.GenerateTypedSig(k.PublicKey(), k.PrivateKey, evmcompat.SignatureType_EIP712)
	if err != nil {
		panic(err)
	}
	return sigBytes
}

func (k *EthSigner66Byte) PublicKey() []byte {
	ethLocalAdr, err := loom.LocalAddressFromHexString(crypto.PubkeyToAddress(k.PrivateKey.PublicKey).Hex())
	if err != nil {
		panic(err)
	}
	ethPublicAddr := loom.Address{ChainID: EthChainId, Local: ethLocalAdr}
	hash := crypto.Keccak256(ethPublicAddr.Bytes())

	return hash
}
