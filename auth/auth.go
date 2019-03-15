package auth

import (
	"fmt"

	"github.com/loomnetwork/go-loom"
	"github.com/loomnetwork/go-loom/vm"
	"github.com/loomnetwork/go-loom/auth/secp256k1"
	"github.com/loomnetwork/go-loom/auth/yubihsm"
	"github.com/loomnetwork/loomchain"
	"github.com/pkg/errors"
	"github.com/gogo/protobuf/proto"
)

const (
	SignerTypeEd25519   = "ed25519"
	SignerTypeSecp256k1 = "secp256k1"
	SignerTypeYubiHsm   = "yubihsm"
)

// Signer interface is used to sign transactions.
type Signer interface {
	Sign(msg []byte) []byte
	PublicKey() []byte
}

func NewSigner(signerType string, privKey interface{}) Signer {
	switch signerType {
	case SignerTypeEd25519:
		return NewEd25519Signer(privKey.([]byte))
	case SignerTypeSecp256k1:
		return secp256k1.NewSecp256k1Signer(privKey.([]byte))
	case SignerTypeYubiHsm:
		return yubihsm.NewYubiHsmSigner(privKey)
	default:
		panic(fmt.Errorf("Unknown signer type %s", signerType))
	}
	return nil
}

// SignTx generates a signed tx containing the given bytes.
func SignTx(signer Signer, txBytes []byte) *SignedTx {
	return &SignedTx{
		Inner:     txBytes,
		Signature: signer.Sign(txBytes),
		PublicKey: signer.PublicKey(),
	}
}

func GetFromToNonce(nonceBytes []byte) (loom.Address, loom.Address, uint64, error) {
	var nonceTx NonceTx
	if err := proto.Unmarshal(nonceBytes, &nonceTx); err != nil {
		return loom.Address{}, loom.Address{}, 0, errors.Wrap(err, "unwrap nonce Tx")
	}

	var tx loomchain.Transaction
	if err := proto.Unmarshal(nonceTx.Inner, &tx); err != nil {
		return loom.Address{}, loom.Address{}, 0, errors.New("unmarshal tx")
	}

	var msg vm.MessageTx
	if err := proto.Unmarshal(tx.Data, &msg); err != nil {
		return loom.Address{}, loom.Address{}, 0, errors.Wrapf(err, "unmarshal message tx %v", tx.Data)
	}

	if msg.From == nil {
		return loom.Address{}, loom.Address{}, 0, errors.Errorf("nil from address")
	}

	return loom.UnmarshalAddressPB(msg.From), loom.UnmarshalAddressPB(msg.To), nonceTx.Sequence, nil
}