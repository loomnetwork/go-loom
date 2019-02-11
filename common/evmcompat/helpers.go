// +build evm

package evmcompat

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/sha3"
	ssha "github.com/miguelmota/go-solidity-sha3"
)

type SignatureType uint8

const (
	SignatureType_EIP712 SignatureType = 0
	SignatureType_GETH   SignatureType = 1
	SignatureType_TREZOR SignatureType = 2
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
	var signer common.Address

	if len(sig) != 66 {
		return signer, fmt.Errorf("signature must be 66 bytes, not %d bytes", len(sig))
	}

	switch SignatureType(sig[0]) {
	case SignatureType_EIP712:
	case SignatureType_GETH:
		hash = ssha.SoliditySHA3(
			ssha.String("\x19Ethereum Signed Message:\n32"),
			ssha.Bytes32(hash),
		)
	case SignatureType_TREZOR:
		hash = ssha.SoliditySHA3(
			ssha.String("\x19Ethereum Signed Message:\n\x20"),
			ssha.Bytes32(hash),
		)
	default:
		return signer, fmt.Errorf("invalid signature type: %d", sig[0])
	}

	signer, err := SolidityRecover(hash, sig[1:])
	return signer, err
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

// func main() {
// 	sum := 0
// 	for i := 0; i < 10; ixb++ {
// 		sum += i
// 	}
// 	fmt.Println(sum)
// }

func SolidityUnpackBytes(data string) {
	types := [5]string{"uint256", "string", "string", "address", "uint256"}
	data = "0x0000000000000000000000000000000000000000000000000000000000000003000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000c48cf958324a23f77044b63949df104cca6fce2000000000000000000000000000000000000000000000000010f15cb27f673d5c00000000000000000000000000000000000000000000000000000000000000033078300000000000000000000000000000000000000000000000000000000000"
	log.Println(data[0:2])
	if data[0:2] == "0x" {
		data = data[2:]
	}
	for i := 0; i < len(types); i++ {
		partialData := data[i*64 : (i+1)*64]
		log.Println("partialData", partialData)
		TypeConverter(partialData, types[i])

	}

}

func TypeConverter(partialData, typeString string) {
	switch typeString {
	case "uint256":
		i := new(big.Int)
		theInt, _ := i.SetString(partialData, 10)
		log.Println("theInt", theInt)
	}
}
