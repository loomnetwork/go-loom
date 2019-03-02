// +build evm

package gateway_v2

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
)

func (t *MainnetGatewayContractTransactor) DepositEthToGateway(opts *bind.TransactOpts) (*types.Transaction, error) {
	return t.contract.Transfer(opts)
}
