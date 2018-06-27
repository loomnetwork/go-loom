// +build evm

package evmcompat

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/crypto/sha3"
)

//TODO in future all interfaces and not do conversions from strings
type Pair struct {
	Type  string
	Value string
}

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
