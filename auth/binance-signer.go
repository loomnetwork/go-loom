// +build evm

package auth

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"errors"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/loomnetwork/go-loom/common/evmcompat"
	sha3 "github.com/miguelmota/go-solidity-sha3"
)

type BinanceSigner struct {
	privateKey *ecdsa.PrivateKey
}

func NewBinanceSigner(privateKey []byte) *BinanceSigner {
	var err error

	binanceSigner := &BinanceSigner{}
	if privateKey == nil {
		binanceSigner.privateKey, err = ecdsa.GenerateKey(secp256k1.S256(), rand.Reader)
		if err != nil {
			panic(err)
		}
	} else {
		if len(privateKey) != Secp256k1PrivKeyBytes {
			panic(errors.New("Invalid private key length"))
		}

		hexPrivKey := hex.EncodeToString(privateKey)
		binanceSigner.privateKey, err = crypto.HexToECDSA(hexPrivKey)
		if err != nil {
			panic(err)
		}
	}
	return binanceSigner
}

func (s *BinanceSigner) Sign(txBytes []byte) []byte {
	signature, err := evmcompat.GenerateTypedSig(
		sha3.SoliditySHA3(txBytes),
		s.privateKey,
		evmcompat.SignatureType_BINANCE,
	)
	if err != nil {
		panic(err)
	}
	return signature
}

func (s *BinanceSigner) PublicKey() []byte {
	return secp256k1.CompressPubkey(s.privateKey.X, s.privateKey.Y)
}
