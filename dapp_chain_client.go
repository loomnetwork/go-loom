package loomplugin

type VMType int32

const (
	VMType_PLUGIN VMType = 0
	VMType_EVM    VMType = 1
)

type DAppChainClient interface {
	GetNonce(signer Signer) (uint64, error)
	CommitTx(signer Signer, txBytes []byte) ([]byte, error)
	CommitDeployTx(from Address, signer Signer, vm VMType, code []byte) ([]byte, error)
	CommitCallTx(from Address, to Address, signer Signer, vm VMType, input []byte) ([]byte, error)
}
