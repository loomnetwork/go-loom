// +build evm

package plasma_cash

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/loomnetwork/go-loom/common/evmcompat"
	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type EncodingTestSuite struct{}

var _ = Suite(&EncodingTestSuite{})

func (s *EncodingTestSuite) TestUnsignedTxRlpEncode(c *C) {
	privKey, err := crypto.HexToECDSA("bb63b692f9d8f21f0b978b596dc2b8611899f053d68aec6c1c20d1df4f5b6ee2")
	if err != nil {
		c.Fatal(err)
	}
	ownerAddr := crypto.PubkeyToAddress(privKey.PublicKey)
	tx := &LoomTx{
		Slot:         5,
		PrevBlock:    big.NewInt(0),
		Denomination: big.NewInt(1),
		Owner:        ownerAddr,
		Nonce:        0,
	}
	txBytes, err := tx.RlpEncode()
	if err != nil {
		c.Fatal(err)
	}
	hexStr := common.Bytes2Hex(txBytes)
	c.Assert(hexStr, Equals, "d9058001945194b63f10691e46635b27925100cfc0a5ceca6280")

	tx.PrevBlock = big.NewInt(85478557858583)
	txBytes, err = tx.RlpEncode()
	if err != nil {
		c.Fatal(err)
	}
	hexStr = common.Bytes2Hex(txBytes)
	c.Assert(hexStr, Equals, "df05864dbe0713bb1701945194b63f10691e46635b27925100cfc0a5ceca6280")
}

func (s *EncodingTestSuite) TestTxHash(c *C) {
	privKey, err := crypto.HexToECDSA("bb63b692f9d8f21f0b978b596dc2b8611899f053d68aec6c1c20d1df4f5b6ee2")
	if err != nil {
		c.Fatal(err)
	}
	ownerAddr := crypto.PubkeyToAddress(privKey.PublicKey)
	tx := &LoomTx{
		Slot:         5,
		PrevBlock:    big.NewInt(85478557858583),
		Denomination: big.NewInt(1),
		Owner:        ownerAddr,
		Nonce:        0,
	}

	hash, err := tx.Hash()
	if err != nil {
		c.Fatal(err)
	}

	hexStr := common.Bytes2Hex(hash)
	c.Assert(hexStr, Equals, "e973665de5d78b6f4ab345c5e9f9f11ba326d35a1bd324dc462177060f142d0a")
}

func (s *EncodingTestSuite) TestTxSignature(c *C) {
	privKey, err := crypto.HexToECDSA("bb63b692f9d8f21f0b978b596dc2b8611899f053d68aec6c1c20d1df4f5b6ee2")
	if err != nil {
		c.Fatal(err)
	}
	ownerAddr := crypto.PubkeyToAddress(privKey.PublicKey)
	tx := &LoomTx{
		Slot:         5,
		PrevBlock:    big.NewInt(85478557858583),
		Denomination: big.NewInt(1),
		Owner:        ownerAddr,
		Nonce:        0,
	}
	sig, err := tx.Sign(privKey)
	if err != nil {
		c.Fatal(err)
	}

	hexStr := common.Bytes2Hex(sig)
	c.Assert(hexStr, Equals, "00dcfe2f3aca758f0c35994da2af16ac686cd9cff6f2bb0bbdcfbdcb6f4139229271bac95b0a7ae94d3eaa506b77dd4b09c084dee9fbbe3c2d4843978d5e33ca011c")

	hash, err := tx.Hash()
	if err != nil {
		c.Fatal(err)
	}

	signer, err := evmcompat.SolidityRecover(hash, sig[1:])
	if err != nil {
		c.Fatal(err)
	}
	c.Assert(signer.Hex(), Equals, ownerAddr.Hex())
}

func (s *EncodingTestSuite) TestUnsignedTxRlpEncode2(c *C) {
	ownerStr := "0x50bce46ff7f6b92e4d383e4ada3ecba9e86d1292"
	ownerAddr := common.HexToAddress(ownerStr)

	tx := &LoomTx{
		Slot:         2,
		PrevBlock:    big.NewInt(1000),
		Denomination: big.NewInt(1),
		Owner:        ownerAddr,
		Nonce:        0,
	}
	txBytes, err := tx.RlpEncode()
	if err != nil {
		c.Fatal(err)
	}
	hexStr := common.Bytes2Hex(txBytes)
	c.Assert(hexStr, Equals, "db028203e8019450bce46ff7f6b92e4d383e4ada3ecba9e86d129280")
}
