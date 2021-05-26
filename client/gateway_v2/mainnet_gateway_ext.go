// +build evm

package gateway_v2

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (t *MainnetGatewayContractTransactor) DepositEthToGateway(opts *bind.TransactOpts) (*types.Transaction, error) {
	return t.contract.Transfer(opts)
}

// UnsignedWithdrawERC20 returns an unsigend transaction that calls WithdrawERC20, the transaction
// is not broadcasting to Ethereum.
func (t *MainnetGatewayContractTransactor) UnsignedWithdrawERC20(
	opts *bind.TransactOpts, amount *big.Int, contractAddress common.Address,
	signersIndexes []*big.Int, vs []uint8, rs [][32]byte, ss [][32]byte,
) (*types.Transaction, error) {
	return t.contract.UnsignedTransact(opts, "withdrawERC20", amount, contractAddress, signersIndexes, vs, rs, ss)
}
