package lottery

import (
	loom "github.com/loomnetwork/go-loom"
	"github.com/loomnetwork/go-loom/builtin/types/coin"
	"github.com/loomnetwork/go-loom/plugin"
	contract "github.com/loomnetwork/go-loom/plugin/contractpb"
	"github.com/loomnetwork/go-loom/types"
)

type Lottery struct {
}

var coinContractKey = []byte("coincontract")

func transfer(ctx contract.Context, to loom.Address, amount *loom.BigUInt) error {
	req := &coin.TransferRequest{
		To:     to.MarshalPB(),
		Amount: &types.BigUInt{*amount},
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
	transfer(ctx, winnerAddr, loom.NewBigUIntFromInt(1000))
}

var Contract plugin.Contract = contract.MakePluginContract(&Lottery{})

func main() {
	plugin.Serve(Contract)
}
