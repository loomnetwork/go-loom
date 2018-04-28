package lottery

import (
	"math/big"

	loom "github.com/loomnetwork/go-loom"
	"github.com/loomnetwork/go-loom/builtin/plugins/coin"
	"github.com/loomnetwork/go-loom/plugin"
	contract "github.com/loomnetwork/go-loom/plugin/contractpb"
)

type Lottery struct {
}

var coinContractKey = []byte("coincontract")

func transfer(ctx contract.Context, coinAddr loom.Address, to loom.Address, amount *big.Int) error {
	req := &coin.TransferRequest{
		To:     to.MarshalPB(),
		Amount: coin.MarshalBigIntPB(amount),
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
	coinAddr := loom.UnmarshalAddressPB(req.CoinContract)
	winnerAddr := loom.UnmarshalAddressPB(req.Winner)
	transfer(ctx, coinAddr, winnerAddr, big.NewInt(1000))
}

var Contract plugin.Contract = contract.MakePluginContract(&Lottery{})

func main() {
	plugin.Serve(Contract)
}
