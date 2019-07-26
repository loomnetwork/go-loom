// +build !evm

package evmcompat

import "github.com/loomnetwork/go-loom/plugin/contractpb"

type SignatureType uint8

const (
	SignatureType_EIP712  SignatureType = 0
	SignatureType_GETH    SignatureType = 1
	SignatureType_TREZOR  SignatureType = 2
	SignatureType_TRON    SignatureType = 3
	SignatureType_BINANCE SignatureType = 4
)

func GetAllowedSignatureTypes(ctx contractpb.StaticContext) []SignatureType {
	var allowedSigTypes []SignatureType
	// AuthSigTxFeature is in the form 'auth:sigtx:..' e.g. auth:sigtx:default, auth:sigtx:eth
	if ctx.FeatureEnabled("auth:sigtx:default", false) {
		allowedSigTypes = append(allowedSigTypes, SignatureType_EIP712)
	}
	if ctx.FeatureEnabled("auth:sigtx:eth", false) {
		allowedSigTypes = append(allowedSigTypes, SignatureType_GETH)
	}
	if ctx.FeatureEnabled("auth:sigtx:tron", false) {
		allowedSigTypes = append(allowedSigTypes, SignatureType_TRON)
	}
	if ctx.FeatureEnabled("auth:sigtx:binance", false) {
		allowedSigTypes = append(allowedSigTypes, SignatureType_BINANCE)
	}
	return allowedSigTypes
}
