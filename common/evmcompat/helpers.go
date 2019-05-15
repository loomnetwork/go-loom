// +build evm

package evmcompat

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"github.com/eosspark/eos-go/crypto/ecc"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/sha3"
	ssha "github.com/miguelmota/go-solidity-sha3"
)

type SignatureType uint8

const (
	SignatureType_EIP712 SignatureType = iota
	SignatureType_GETH
	SignatureType_TREZOR
	SignatureType_EOS
)

// SoliditySign signs the given data with the specified private key and returns the 65-byte signature.
// The signature is in a format that's compatible with the ecverify() Solidity function.
func SoliditySign(data []byte, privKey *ecdsa.PrivateKey) ([]byte, error) {
	sig, err := crypto.Sign(data, privKey)
	if err != nil {
		return nil, err
	}

	v := sig[len(sig)-1]
	sig[len(sig)-1] = v + 27
	return sig, nil
}

// SolidityRecover recovers the Ethereum address from the signed hash and the 65-byte signature.
func SolidityRecover(hash []byte, sig []byte) (common.Address, error) {
	if len(sig) != 65 {
		return common.Address{}, fmt.Errorf("signature must be 65 bytes, got %d bytes", len(sig))
	}
	stdSig := make([]byte, 65)
	copy(stdSig[:], sig[:])
	stdSig[len(sig)-1] -= 27

	var signer common.Address
	pubKey, err := crypto.Ecrecover(hash, stdSig)
	if err != nil {
		return signer, err
	}

	copy(signer[:], crypto.Keccak256(pubKey[1:])[12:])
	return signer, nil
}

// GenerateTypedSig signs the given data with the specified private key and returns the 66-byte signature
// (the first byte of which is used to denote the SignatureType).
func GenerateTypedSig(data []byte, privKey *ecdsa.PrivateKey, sigType SignatureType) ([]byte, error) {
	if sigType != SignatureType_EIP712 {
		return nil, errors.New("signing failed, sig type not implemented")
	}

	sig, err := SoliditySign(data, privKey)
	if err != nil {
		return nil, err
	}
	// Prefix the sig with a single byte indicating the sig type, in this case EIP712
	typedSig := append(make([]byte, 0, 66), byte(SignatureType_EIP712))
	return append(typedSig, sig...), nil
}

// RecoverAddressFromTypedSig recovers the Ethereum address from a signed hash and a 66-byte signature
// (the first byte of which is expected to denote the SignatureType).
func RecoverAddressFromTypedSig(hash []byte, sig []byte) (common.Address, error) {
	var err error

	switch SignatureType(sig[0]) {
	case SignatureType_EIP712:
		err = recoverAddressFromEIP712(sig[1:])
	case SignatureType_GETH:
		hash, err = recoverAddressFromGeth(hash, sig[1:])
	case SignatureType_TREZOR:
		hash, err = recoverAddressFromTrezor(hash, sig[1:])
	case SignatureType_EOS:
		return recoverAddressFromEos(hash, sig[1:])
	default:
		err = fmt.Errorf("invalid signature type: %d", sig[0])
	}
	if err != nil {
		return common.Address{}, err
	}

	return SolidityRecover(hash, sig[1:])
}

func recoverAddressFromEIP712(sig []byte) error {
	if len(sig) != 65 {
		return fmt.Errorf("signature must be 66 bytes, not %d bytes", len(sig))
	}
	return nil
}

func recoverAddressFromGeth(hash []byte, sig []byte) ([]byte, error) {
	if len(sig) != 65 {
		return nil, fmt.Errorf("signature must be 66 bytes, not %d bytes", len(sig))
	}
	return ssha.SoliditySHA3(
		ssha.String("\x19Ethereum Signed Message:\n32"),
		ssha.Bytes32(hash),
	), nil
}

func recoverAddressFromTrezor(hash []byte, sig []byte) ([]byte, error) {
	if len(sig) != 65 {
		return nil, fmt.Errorf("signature must be 66 bytes, not %d bytes", len(sig))
	}
	return ssha.SoliditySHA3(
		ssha.String("\x19Ethereum Signed Message:\n\x20"),
		ssha.Bytes32(hash),
	), nil
}

func recoverAddressFromEos(hash []byte, sig []byte) (common.Address, error) {
	var signer common.Address

	if len(sig) != 107 {
		return signer, fmt.Errorf("eos signature must be 107 bytes, not %d bytes", len(sig))
	}
	signature, err := ecc.NewSignature(string(sig[6:]))
	if err != nil {
		return signer, fmt.Errorf("cannot unpack eos signature %v", string(sig))
	}

	nonceHash := sha256.Sum256([]byte(hex.EncodeToString(sig[:6])))
	scatterMsgHash := sha256.Sum256([]byte(hex.EncodeToString(hash) + hex.EncodeToString(nonceHash[:])))

	pubKey, err := signature.PublicKey(scatterMsgHash[:])
	if err != nil {
		return signer, fmt.Errorf("cannot get publlic key from hash %v", string(hash))
	}
	btcecPubKey, err := pubKey.Key()
	if err != nil {
		return signer, errors.Wrapf(err, "retrieve btcec key from eos key %v", pubKey)
	}
	local := crypto.PubkeyToAddress(ecdsa.PublicKey(*btcecPubKey))
	return local, nil
}

//TODO in future all interfaces and not do conversions from strings
type Pair struct {
	Type  string
	Value string
}

// NOTE: This function is deprecated, use the one in github.com/miguelmota/go-solidity-sha3 instead!
func SoliditySHA3(pairs []*Pair) ([]byte, error) {
	//convert to packed bytes like solidity
	data, err := SolidityPackedBytes(pairs)
	if err != nil {
		return nil, err
	}

	d := sha3.NewKeccak256()
	d.Write(data)
	return d.Sum(nil), nil
}

func SolidityPackedBytes(pairs []*Pair) ([]byte, error) {
	var b bytes.Buffer

	for _, pair := range pairs {
		fmt.Printf("%v\n", pair)
		switch strings.ToLower(pair.Type) {
		case "address":
			decoded, err := hex.DecodeString(pair.Value)
			if err != nil {
				return nil, err
			}
			if len(decoded) != 20 {
				return nil, fmt.Errorf("we don't support partial addresses, the len was %d we wanted 20", len(decoded))
			}
			b.Write(decoded)
		case "uint16": //"uint", "uint16", "uint64":
			//pack integers
			u, err := strconv.ParseUint(pair.Value, 10, 32)
			if err != nil {
				return nil, err
			}
			var bTest []byte = make([]byte, 2)
			//			binary.LittleEndian.PutUint32(bTest, uint32(u))
			//			fmt.Printf("little-%v\n", bTest)
			binary.BigEndian.PutUint16(bTest, uint16(u))
			b.Write(bTest)
		case "uint32": //"uint", "uint16", "uint64":
			//pack integers
			u, err := strconv.ParseUint(pair.Value, 10, 32)
			if err != nil {
				return nil, err
			}
			var bTest []byte = make([]byte, 4)
			//			binary.LittleEndian.PutUint32(bTest, uint32(u))
			//			fmt.Printf("little-%v\n", bTest)
			binary.BigEndian.PutUint32(bTest, uint32(u))
			b.Write(bTest)
		case "uint64": //"uint", "uint16", "uint64":
			//pack integers
			u, err := strconv.ParseUint(pair.Value, 10, 64)
			if err != nil {
				return nil, err
			}
			var bTest []byte = make([]byte, 8)
			//			binary.LittleEndian.PutUint32(bTest, uint32(u))
			//			fmt.Printf("little-%v\n", bTest)
			binary.BigEndian.PutUint64(bTest, u)
			b.Write(bTest)
		case "uint256":
			n := new(big.Int)
			_, valid := n.SetString(pair.Value, 10)
			if !valid {
				return nil, errors.New("invalid big int")
			}

			bytes := n.Bytes()
			padlen := 32 - len(bytes)
			if padlen < 0 {
				return nil, errors.New("big int byte length too large")
			}
			pad := make([]byte, padlen, padlen)
			b.Write(pad)
			b.Write(bytes)
		}
	}

	return b.Bytes(), nil
}

// SolidityUnpackString is a function to decode data string from solidity encoded event data into given types.
// Currently this function supports uint8, uint32, uint256, address, string data types.
func SolidityUnpackString(data string, types []string) ([]interface{}, error) {

	if data[0:2] == "0x" {
		data = data[2:]
	}
	var resp = make([]interface{}, len(types))
	var stringCount = 0
	for i := 0; i < len(types); i++ {
		partialData := data[i*64 : (i+1)*64]
		convertedData, count, err := parseNextValueFromSolidityHexStr(partialData, types[i], data[i*64:], len(types)-i, stringCount)
		if err != nil {
			return nil, err
		}
		stringCount = count
		resp[i] = convertedData
	}
	return resp, nil
}

// This internal function parses hexstring into given data types.
func parseNextValueFromSolidityHexStr(partialData, typeString, dataLeft string, chunkLeft, stringCount int) (interface{}, int, error) {
	switch typeString {
	case "uint8":
		theInt, err := strconv.ParseUint(partialData, 10, 8)
		if err != nil {
			return nil, stringCount, err
		}
		return uint8(theInt), stringCount, nil

	case "uint32":
		theInt, err := strconv.ParseUint(partialData, 10, 32)
		if err != nil {
			return nil, stringCount, err
		}
		return uint32(theInt), stringCount, nil

	case "uint256":
		i := new(big.Int)
		theInt, ok := i.SetString(partialData, 16)
		if !ok {
			return nil, stringCount, errors.New(fmt.Sprintf("Error parsing big.Int from %s", partialData))
		}
		return theInt, stringCount, nil

	case "address":
		sliced := "0x" + partialData[24:]
		strings.ToLower(sliced)
		return sliced, stringCount, nil

	case "string":
		// find len of string
		// chunkLeft*64 : have to skip all chunk left
		// stringCount*64*2 : have to skip all string chunk passed

		dataChunkIndex := chunkLeft*64 + stringCount*64*2
		lenChunk := dataLeft[dataChunkIndex : dataChunkIndex+64]

		// find chunk of string
		// +64 for skip len chunk
		stringChunk := dataLeft[dataChunkIndex+64 : dataChunkIndex+64*2]

		//decode string
		byteString, err := hex.DecodeString(stringChunk)
		if err != nil {
			return nil, stringCount, err
		}

		//substring equal to len
		i := new(big.Int)
		stringLen, ok := i.SetString(lenChunk, 16)
		if !ok {
			return nil, stringCount, errors.New(fmt.Sprintf("Error parsing big.Int from %s", partialData))
		}
		byteString = byteString[:stringLen.Int64()]
		stringConverted := string(byteString)

		return strings.ToLower(stringConverted), stringCount + 1, nil
	}
	return typeString, stringCount, nil
}
