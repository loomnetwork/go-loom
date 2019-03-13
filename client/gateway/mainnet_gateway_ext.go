// +build evm

package gateway

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (t *MainnetGatewayContractTransactor) DepositEthToGateway(opts *bind.TransactOpts) (*types.Transaction, error) {
	return t.contract.Transfer(opts)
}

// Low level call that directly returns the unsigend transaction without broadcasting it
func (t *MainnetGatewayContractTransactor) UnsignedWithdrawERC20(opts *bind.TransactOpts, amount *big.Int, sig []byte, tokenAddr common.Address) (*types.Transaction, error) {
	return t.contract.UnsignedTransact(opts, "withdrawERC20", amount, sig, tokenAddr)
}
