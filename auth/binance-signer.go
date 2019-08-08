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
)

type BinanceSigner struct {
	privateKey *ecdsa.PrivateKey
}

// NewBinanceSigner creates a new signer that can be used to sign txs & messages using the given
// secp256k1 private key. If the given private key is nil a new random key will be generated.
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
		evmcompat.GenSHA256(txBytes),
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
