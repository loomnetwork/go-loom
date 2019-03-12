// +build evm

package auth

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gogo/protobuf/proto"
	"github.com/loomnetwork/go-loom"
	"github.com/loomnetwork/go-loom/common/evmcompat"
	"github.com/loomnetwork/go-loom/vm"
	"github.com/loomnetwork/loomchain"
	sha3 "github.com/miguelmota/go-solidity-sha3"
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

func (k *EthSigner66Byte) Sign(txBytes []byte) []byte {
	var nonceTx NonceTx
	if err := proto.Unmarshal(txBytes, &nonceTx); err != nil {
		panic(err)
	}

	var tx loomchain.Transaction
	if err := proto.Unmarshal(nonceTx.Inner, &tx); err != nil {
		panic("throttle: unmarshal tx")
	}
	var msg vm.MessageTx
	if err := proto.Unmarshal(tx.Data, &msg); err != nil {
		panic("unmarshal message tx")
	}
	from := loom.UnmarshalAddressPB(msg.From)
	to := loom.UnmarshalAddressPB(msg.To)

	hash := sha3.SoliditySHA3(
		sha3.Address(common.BytesToAddress(from.Local)),
		sha3.Address(common.BytesToAddress(to.Local)),
		sha3.Uint64(nonceTx.Sequence),
		txBytes,
	)
	signature, err := evmcompat.GenerateTypedSig(hash, k.PrivateKey, evmcompat.SignatureType_EIP712)
	if err != nil {
		panic(err)
	}
	return signature
}

func (k *EthSigner66Byte) PublicKey() []byte {
	return crypto.FromECDSAPub(&k.PrivateKey.PublicKey)
}

type EthSigner65Byte struct {
	PrivateKey *ecdsa.PrivateKey
}

func NewEthSigner65Byte(_ []byte) Signer {
	panic("EVM build isn't activated")
}

func (k *EthSigner65Byte) Sign(_ []byte) []byte {
	sigBytes, err := evmcompat.SoliditySign(k.PublicKey(), k.PrivateKey)
	if err != nil {
		panic(err)
	}
	return sigBytes
}

func (k *EthSigner65Byte) PublicKey() []byte {
	ethLocalAdr, err := loom.LocalAddressFromHexString(crypto.PubkeyToAddress(k.PrivateKey.PublicKey).Hex())
	if err != nil {
		panic(err)
	}
	ethPublicAddr := loom.Address{ChainID: EthChainId, Local: ethLocalAdr}
	hash := crypto.Keccak256(ethPublicAddr.Bytes())

	return hash
}
