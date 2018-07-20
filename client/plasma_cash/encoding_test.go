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
		Denomination: 1,
		Owner:        ownerAddr,
	}
	txBytes, err := tx.RlpEncode()
	if err != nil {
		c.Fatal(err)
	}
	hexStr := common.Bytes2Hex(txBytes)
	c.Assert(hexStr, Equals, "d8058001945194b63f10691e46635b27925100cfc0a5ceca62")

	tx.PrevBlock = big.NewInt(85478557858583)
	txBytes, err = tx.RlpEncode()
	if err != nil {
		c.Fatal(err)
	}
	hexStr = common.Bytes2Hex(txBytes)
	c.Assert(hexStr, Equals, "de05864dbe0713bb1701945194b63f10691e46635b27925100cfc0a5ceca62")
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
		Denomination: 1,
		Owner:        ownerAddr,
	}
	hexStr := common.Bytes2Hex(tx.Hash())
	c.Assert(hexStr, Equals, "b10da41825f94bd447ebce74913e82ceae90c6ba27aa6781d611f8530f78ec4c")
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
		Denomination: 1,
		Owner:        ownerAddr,
	}
	sig, err := tx.Sign(privKey)
	if err != nil {
		c.Fatal(err)
	}

	hexStr := common.Bytes2Hex(sig)
	c.Assert(hexStr, Equals, "00b0e4901dc74b9851dba3c52406e1325c2ac9c4fe9f4d0379099a3357b763c96c104d3fffb78e99515db2e583568588d740b743ad3105d63fb252014f806fd06b1b")

	signer, err := evmcompat.SolidityRecover(tx.Hash(), sig[1:])
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
		Denomination: 1,
		Owner:        ownerAddr,
	}
	txBytes, err := tx.RlpEncode()
	if err != nil {
		c.Fatal(err)
	}
	hexStr := common.Bytes2Hex(txBytes)
	c.Assert(hexStr, Equals, "da028203e8019450bce46ff7f6b92e4d383e4ada3ecba9e86d1292")
}
