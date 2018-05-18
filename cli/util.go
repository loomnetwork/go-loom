package cli

import (
	"encoding/base64"
	"encoding/hex"
	"strings"

	loom "github.com/loomnetwork/go-loom"
)

func ParseBytes(s string) ([]byte, error) {
	if strings.HasPrefix(s, "0x") {
		return hex.DecodeString(s[2:])
	}

	b, err := hex.DecodeString(s)
	if err != nil {
		b, err = base64.StdEncoding.DecodeString(s)
	}

	return b, err
}

func ParseAddress(s string) (loom.Address, error) {
	addr, err := loom.ParseAddress(s)
	if err == nil {
		return addr, nil
	}

	b, err := ParseBytes(s)
	if len(b) != 20 {
		return loom.Address{}, loom.ErrInvalidAddress
	}

	return loom.Address{ChainID: txFlags.ChainID, Local: loom.LocalAddress(b)}, nil
}
