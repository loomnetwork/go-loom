// +build !evm

package evmcompat

type SignatureType uint8

const (
	SignatureType_EIP712  SignatureType = 0
	SignatureType_GETH    SignatureType = 1
	SignatureType_TREZOR  SignatureType = 2
	SignatureType_TRON    SignatureType = 3
	SignatureType_BINANCE SignatureType = 4
)
