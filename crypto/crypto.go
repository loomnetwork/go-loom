package crypto

const (
	PrivateKeyTypeEd25519   = "ed25519"
	PrivateKeyTypeSecp256k1 = "secp256k1"
)

type PrivateKey interface{}
