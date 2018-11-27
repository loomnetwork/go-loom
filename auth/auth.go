package auth

// Signer interface is used to sign transactions.
type Signer interface {
	Sign(msg []byte) []byte
	PublicKey() []byte
}

// SignTx generates a signed tx containing the given bytes.
func SignTx(signer Signer, txBytes []byte) *SignedTx {
	return &SignedTx{
		Inner:     txBytes,
		Signature: signer.Sign(txBytes),
		PublicKey: signer.PublicKey(),
	}
}
