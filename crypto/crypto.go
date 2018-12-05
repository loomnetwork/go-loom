package crypto

const (
	PrivKeyTypeFile    = "file"
	PrivKeyTypeYubiHsm = "yubihsm"
)

type PrivateKey interface{}
