package loomplugin

type DAppChainClient interface {
	CommitTx(signer Signer, txBytes []byte) error
}
