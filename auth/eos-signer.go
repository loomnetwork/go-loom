// +build evm

package auth

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"strings"

	"github.com/eosspark/eos-go/crypto/ecc"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gogo/protobuf/proto"
	sha3 "github.com/miguelmota/go-solidity-sha3"

	"github.com/loomnetwork/go-loom/common/evmcompat"
)

type EosSigner struct {
	PrivateKey *ecc.PrivateKey
}

func (e *EosSigner) Sign(txBytes []byte) []byte {
	signature, err := e.PrivateKey.Sign(sha3.SoliditySHA3(txBytes))
	if err != nil {
		panic(err)
	}
	sigByes, err := signature.Pack()
	if err != nil {
		panic(err)
	}
	return sigByes
}

func (e *EosSigner) PublicKey() []byte {
	btcecPubKey, err := e.PrivateKey.PublicKey().Key()
	if err != nil {
		panic(err)
	}
	ecdsaPubKey := ecdsa.PublicKey(*btcecPubKey)
	return crypto.FromECDSAPub(&ecdsaPubKey)
}

type EosScatterSigner struct {
	PrivateKey *ecc.PrivateKey
}

func (e *EosScatterSigner) Sign(txBytes []byte) []byte {
	hash := sha256.Sum256([]byte(strings.ToUpper(hex.EncodeToString(txBytes))))
	var nonceTx NonceTx
	if err := proto.Unmarshal(txBytes, &nonceTx); err != nil {
		panic(err)
	}
	typedSignature := []byte{byte(evmcompat.SignatureType_EOS_SCATTER)}

	nonceBytes := []byte(strconv.FormatUint(nonceTx.Sequence, 10))[:6]
	nonceSha := sha256.Sum256(nonceBytes)
	typedSignature = append(typedSignature, nonceBytes...)

	hash_1 := sha256.Sum256([]byte("0x" + hex.EncodeToString(hash[:])))
	hash_2 := sha256.Sum256([]byte(hex.EncodeToString(nonceSha[:6])))
	scatterMsgHash := sha256.Sum256([]byte(hex.EncodeToString(hash_1[:]) + hex.EncodeToString(hash_2[:])))

	signature, err := e.PrivateKey.Sign(scatterMsgHash[:])
	if err != nil {
		panic(err)
	}
	typedSignature = append(typedSignature, []byte(signature.String())...)

	return typedSignature
}

func (e *EosScatterSigner) PublicKey() []byte {
	btcecPubKey, err := e.PrivateKey.PublicKey().Key()
	if err != nil {
		panic(err)
	}
	ecdsaPubKey := ecdsa.PublicKey(*btcecPubKey)
	return crypto.FromECDSAPub(&ecdsaPubKey)
}
