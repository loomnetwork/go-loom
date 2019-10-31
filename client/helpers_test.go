// +build evm

package client

import (
	"encoding/hex"
	"fmt"
	"testing"

	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	tgtypes "github.com/loomnetwork/go-loom/builtin/types/transfer_gateway"
	"github.com/stretchr/testify/assert"
)

func TestWithdrawalHash(t *testing.T) {
	t.Logf("Testing withdrawal hash")
	hash := WithdrawalHash(
		common.HexToAddress("0x8b7495748a4aa54d98660cb51f6fa7f244568b9d"), // withdrawer
		common.HexToAddress("0xE681fc277ef3Eed61f0d48FB87Aa2EEe9F67aBB4"), // token addr
		common.HexToAddress("0x5D442Ef71427Fb0F57D63605F6FAf34cA8e78341"), // gateway addr
		tgtypes.TransferGatewayTokenKind_ERC20,
		big.NewInt(0),
		big.NewInt(2000000000000000000),
		big.NewInt(0),
		true,
	)
	assert.Equal(t, hex.EncodeToString(hash), "ed49f376605734945e3b1ddd0ec6341a76fdcda5b47138a4afc6d4b9ea04ab05")

}

func TestParseSigWithoutMode(t *testing.T) {
	t.Logf("Testing parse sigs")
	sigs, _ := hex.DecodeString("cd7f07b4f35d2d2dee86bde44d765aef81673745aab5d5aaf4422dc73938237d2cbc5105bc0ceddbf4037b62003159903d35b834496a622ba4d9117008c164401c")
	validators := []common.Address{
		common.HexToAddress("0x0C0eaEC5552C93a22fB628De3bd18406E1e74989"),
		common.HexToAddress("0xB84B25a63BCEEB318FBC412203D6d70Fef8E8883"),
		common.HexToAddress("0x58e28D2cE00886dfd03Ec47543c15EA185922242"),
		common.HexToAddress("0x45c6971A31C15D8B4E11e22901014b2E5e37c1a8"),
		common.HexToAddress("0xce67056aD7C12bF52A1659FC9a474881ef17ab85"),
		common.HexToAddress("0xe5ac31f6890b4571F0acbE019d24F13E17Db428c"),
	}

	hash, _ := hex.DecodeString("9be6cc490c68327498647b5a846b34565b4358a806d8b7e25a64058cfec744a0")

	vs, rs, ss, ind, _ := ParseSigs(sigs, hash, validators)
	assert.Equal(t, len(vs), 1, "incorrect v length")
	assert.Equal(t, len(rs), 1, "incorrect v length")
	assert.Equal(t, len(ss), 1, "incorrect v length")
	assert.Equal(t, len(ind), 1, "incorrect v length")

	sig := make([]byte, 0, 65)
	sig = append(sig, rs[0][:]...)
	sig = append(sig, ss[0][:]...)
	sig = append(sig, vs[0]-27)

	validatorBytes, err := crypto.SigToPub(hash, sig)
	if err != nil {
		fmt.Println(err)
	}

	addr := crypto.PubkeyToAddress(*validatorBytes)
	assert.Equal(t, addr.String(), validators[ind[0].Int64()].String(), "Validator did not match")
}

func TestParseSigWithMode(t *testing.T) {
	t.Logf("Testing parse sigs")
	sigs, _ := hex.DecodeString("00cd7f07b4f35d2d2dee86bde44d765aef81673745aab5d5aaf4422dc73938237d2cbc5105bc0ceddbf4037b62003159903d35b834496a622ba4d9117008c164401c")
	validators := []common.Address{
		common.HexToAddress("0x0C0eaEC5552C93a22fB628De3bd18406E1e74989"),
		common.HexToAddress("0xB84B25a63BCEEB318FBC412203D6d70Fef8E8883"),
		common.HexToAddress("0x58e28D2cE00886dfd03Ec47543c15EA185922242"),
		common.HexToAddress("0x45c6971A31C15D8B4E11e22901014b2E5e37c1a8"),
		common.HexToAddress("0xce67056aD7C12bF52A1659FC9a474881ef17ab85"),
		common.HexToAddress("0xe5ac31f6890b4571F0acbE019d24F13E17Db428c"),
	}

	hash, _ := hex.DecodeString("9be6cc490c68327498647b5a846b34565b4358a806d8b7e25a64058cfec744a0")

	vs, rs, ss, ind, _ := ParseSigs(sigs, hash, validators)
	assert.Equal(t, len(vs), 1, "incorrect v length")
	assert.Equal(t, len(rs), 1, "incorrect v length")
	assert.Equal(t, len(ss), 1, "incorrect v length")
	assert.Equal(t, len(ind), 1, "incorrect v length")

	sig := make([]byte, 0, 65)
	sig = append(sig, rs[0][:]...)
	sig = append(sig, ss[0][:]...)
	sig = append(sig, vs[0]-27)

	validatorBytes, err := crypto.SigToPub(hash, sig)
	if err != nil {
		fmt.Println(err)
	}

	addr := crypto.PubkeyToAddress(*validatorBytes)
	assert.Equal(t, addr.String(), validators[ind[0].Int64()].String(), "Validator did not match")
}

func TestParseSigs(t *testing.T) {
	t.Logf("Testing parse sigs")
	sigs, _ := hex.DecodeString("cd7f07b4f35d2d2dee86bde44d765aef81673745aab5d5aaf4422dc73938237d2cbc5105bc0ceddbf4037b62003159903d35b834496a622ba4d9117008c164401c6b86642dc24d77f9dcb8da9ced39bd83bbc9c5d536879b377a15cd158296377b3b82da79ba147325977ec82307326fbb33764ac849b57fa4d0b0e7b6d10a762b1bd5574ca4729ee8ee8567b7cb4f13725d8f5143b82d609a66795253b57ec541b454960a8c7aff544c7cf20105fc85dbdc23aa3f8640bd17615dabf606eb2f2cfe1bf7d6a8480464a2323138470e068c031b39e4f3a96e177df800f8b50e9a460c1b6a4ceb40bb337e5932196c011cbbe268dd6786c94d2a669c8d4b0fd0dbc9ce6f1c3bb1bfa61a7c9a632e4bdd8eace611a9ea6b9ddafff53d109f2da16fc2cf14963650f4abe5811621b3a8f44cc99b0084d3c2c90d40e1ea68b1a456bcb368c4331b0817bfea874076384a41d1638d6a5dfae1ab0e05e8e32acf6d1d00ac71dea0fd5abef247260e79512601feb770cd4c85a5061da7db2c1937970b5de0e308ccfc1c")
	validators := []common.Address{
		common.HexToAddress("0x0C0eaEC5552C93a22fB628De3bd18406E1e74989"),
		common.HexToAddress("0xB84B25a63BCEEB318FBC412203D6d70Fef8E8883"),
		common.HexToAddress("0x58e28D2cE00886dfd03Ec47543c15EA185922242"),
		common.HexToAddress("0x45c6971A31C15D8B4E11e22901014b2E5e37c1a8"),
		common.HexToAddress("0xce67056aD7C12bF52A1659FC9a474881ef17ab85"),
		common.HexToAddress("0xe5ac31f6890b4571F0acbE019d24F13E17Db428c"),
	}

	hash, _ := hex.DecodeString("9be6cc490c68327498647b5a846b34565b4358a806d8b7e25a64058cfec744a0")

	vs, rs, ss, ind, _ := ParseSigs(sigs, hash, validators)

	for i, _ := range vs {
		r := rs[i]
		s := ss[i]
		v := vs[i]

		// r + s + v
		sig := make([]byte, 0, 65)
		sig = append(sig, r[:]...)
		sig = append(sig, s[:]...)
		sig = append(sig, v-27)

		validatorBytes, err := crypto.SigToPub(hash, sig)
		if err != nil {
			fmt.Println(err)
		}

		addr := crypto.PubkeyToAddress(*validatorBytes)
		assert.Equal(t, addr.String(), validators[ind[i].Int64()].String(), "Validator did not match")
	}
}

func TestConcatSigs(t *testing.T) {
	sigs, _ := hex.DecodeString("cd7f07b4f35d2d2dee86bde44d765aef81673745aab5d5aaf4422dc73938237d2cbc5105bc0ceddbf4037b62003159903d35b834496a622ba4d9117008c164401c6b86642dc24d77f9dcb8da9ced39bd83bbc9c5d536879b377a15cd158296377b3b82da79ba147325977ec82307326fbb33764ac849b57fa4d0b0e7b6d10a762b1bd5574ca4729ee8ee8567b7cb4f13725d8f5143b82d609a66795253b57ec541b454960a8c7aff544c7cf20105fc85dbdc23aa3f8640bd17615dabf606eb2f2cfe1bf7d6a8480464a2323138470e068c031b39e4f3a96e177df800f8b50e9a460c1b6a4ceb40bb337e5932196c011cbbe268dd6786c94d2a669c8d4b0fd0dbc9ce6f1c3bb1bfa61a7c9a632e4bdd8eace611a9ea6b9ddafff53d109f2da16fc2cf14963650f4abe5811621b3a8f44cc99b0084d3c2c90d40e1ea68b1a456bcb368c4331b0817bfea874076384a41d1638d6a5dfae1ab0e05e8e32acf6d1d00ac71dea0fd5abef247260e79512601feb770cd4c85a5061da7db2c1937970b5de0e308ccfc1c")
	sortedSigs, _ := hex.DecodeString("0817bfea874076384a41d1638d6a5dfae1ab0e05e8e32acf6d1d00ac71dea0fd5abef247260e79512601feb770cd4c85a5061da7db2c1937970b5de0e308ccfc1ccd7f07b4f35d2d2dee86bde44d765aef81673745aab5d5aaf4422dc73938237d2cbc5105bc0ceddbf4037b62003159903d35b834496a622ba4d9117008c164401c3bb1bfa61a7c9a632e4bdd8eace611a9ea6b9ddafff53d109f2da16fc2cf14963650f4abe5811621b3a8f44cc99b0084d3c2c90d40e1ea68b1a456bcb368c4331b6b86642dc24d77f9dcb8da9ced39bd83bbc9c5d536879b377a15cd158296377b3b82da79ba147325977ec82307326fbb33764ac849b57fa4d0b0e7b6d10a762b1bf7d6a8480464a2323138470e068c031b39e4f3a96e177df800f8b50e9a460c1b6a4ceb40bb337e5932196c011cbbe268dd6786c94d2a669c8d4b0fd0dbc9ce6f1cd5574ca4729ee8ee8567b7cb4f13725d8f5143b82d609a66795253b57ec541b454960a8c7aff544c7cf20105fc85dbdc23aa3f8640bd17615dabf606eb2f2cfe1b")
	validators := []common.Address{
		common.HexToAddress("0x0C0eaEC5552C93a22fB628De3bd18406E1e74989"),
		common.HexToAddress("0xB84B25a63BCEEB318FBC412203D6d70Fef8E8883"),
		common.HexToAddress("0x58e28D2cE00886dfd03Ec47543c15EA185922242"),
		common.HexToAddress("0x45c6971A31C15D8B4E11e22901014b2E5e37c1a8"),
		common.HexToAddress("0xce67056aD7C12bF52A1659FC9a474881ef17ab85"),
		common.HexToAddress("0xe5ac31f6890b4571F0acbE019d24F13E17Db428c"),
	}

	hash, _ := hex.DecodeString("9be6cc490c68327498647b5a846b34565b4358a806d8b7e25a64058cfec744a0")

	vs, rs, ss, _, _ := ParseSigs(sigs, hash, validators)

	concatSigs := ConcatSigs(vs, rs, ss)
	assert.Equal(t, concatSigs, sortedSigs, "signature mismatch")
}
