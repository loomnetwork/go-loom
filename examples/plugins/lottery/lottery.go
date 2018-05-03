package lottery

import (
	"math/big"

	loom "github.com/loomnetwork/go-loom"
	"github.com/loomnetwork/go-loom/builtin/types/coin"
	"github.com/loomnetwork/go-loom/plugin"
	contract "github.com/loomnetwork/go-loom/plugin/contractpb"
)

func UnmarshalBigUIntPB(b *loom.BigUInt) *big.Int {
	return new(big.Int).SetBytes(b.Value)
}

func MarshalBigIntPB(b *big.Int) *loom.BigUInt {
	return &loom.BigUInt{
		Value: b.Bytes(),
	}
}

type Lottery struct {
}

var coinContractKey = []byte("coincontract")

func transfer(ctx contract.Context, to loom.Address, amount *big.Int) error {
	req := &coin.TransferRequest{
		To:     to.MarshalPB(),
		Amount: MarshalBigIntPB(amount),
	}

	coinAddr, err := ctx.Resolve("coin")
	if err != nil {
		return err
	}
	return contract.Call(ctx, coinAddr, req, nil)
}

func (c *Lottery) Meta() (plugin.Meta, error) {
	return plugin.Meta{
		Name:    "lottery",
		Version: "1.0.0",
	}, nil
}

func (c *Lottery) Init(ctx contract.Context, req *LotteryInit) {
	winnerAddr := loom.UnmarshalAddressPB(req.Winner)
	transfer(ctx, winnerAddr, big.NewInt(1000))
}

var Contract plugin.Contract = contract.MakePluginContract(&Lottery{})

func main() {
	plugin.Serve(Contract)
}
