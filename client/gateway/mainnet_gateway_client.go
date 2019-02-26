// +build evm

package gateway

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/loomnetwork/go-loom/client"
)

type MainnetGatewayClient struct {
	ethClient *ethclient.Client
	contract  *MainnetGatewayContract
	// Mainnet Gateway contract address
	Address   common.Address
	TxTimeout time.Duration
}

func (c *MainnetGatewayClient) Contract() *MainnetGatewayContract {
	return c.contract
}

func (c *MainnetGatewayClient) DepositERC20(caller *client.Identity, amount *big.Int, tokenAddr common.Address) error {
	tx, err := c.contract.DepositERC20(client.DefaultTransactOptsForIdentity(caller), amount, tokenAddr)
	if err != nil {
		return err
	}
	return client.WaitForTxConfirmation(context.TODO(), c.ethClient, tx, c.TxTimeout)
}

func (c *MainnetGatewayClient) DepositETH(caller *client.Identity, amount *big.Int) (*big.Int, error) {
	opts := client.DefaultTransactOptsForIdentity(caller)
	opts.Value = amount
	tx, err := c.contract.DepositEthToGateway(opts)
	if err != nil {
		return nil, err
	}
	return client.WaitForTxConfirmationAndFee(context.TODO(), c.ethClient, tx, c.TxTimeout)
}

func (c *MainnetGatewayClient) ERC721Deposited(tokenID *big.Int, tokenAddr common.Address) (bool, error) {
	return c.contract.GetERC721(nil, tokenID, tokenAddr)
}

func (c *MainnetGatewayClient) ERC721XBalance(tokenID *big.Int, tokenAddr common.Address) (*big.Int, error) {
	return c.contract.GetERC721X(nil, tokenID, tokenAddr)
}

func (c *MainnetGatewayClient) ERC20Balance(tokenAddr common.Address) (*big.Int, error) {
	return c.contract.GetERC20(nil, tokenAddr)
}

func (c *MainnetGatewayClient) ETHBalance() (*big.Int, error) {
	return c.contract.GetETH(nil)
}

func (c *MainnetGatewayClient) WithdrawERC721(caller *client.Identity, tokenID *big.Int, tokenAddr common.Address, sig []byte) error {
	tx, err := c.contract.WithdrawERC721(client.DefaultTransactOptsForIdentity(caller), tokenID, sig, tokenAddr)
	if err != nil {
		return err
	}
	return client.WaitForTxConfirmation(context.TODO(), c.ethClient, tx, c.TxTimeout)
}

func (c *MainnetGatewayClient) WithdrawERC721X(caller *client.Identity, tokenID, amount *big.Int, tokenAddr common.Address, sig []byte) error {
	tx, err := c.contract.WithdrawERC721X(client.DefaultTransactOptsForIdentity(caller), tokenID, amount, sig, tokenAddr)
	if err != nil {
		return err
	}
	return client.WaitForTxConfirmation(context.TODO(), c.ethClient, tx, c.TxTimeout)
}

func (c *MainnetGatewayClient) WithdrawERC20(caller *client.Identity, amount *big.Int, tokenAddr common.Address, sig []byte) error {
	tx, err := c.contract.WithdrawERC20(client.DefaultTransactOptsForIdentity(caller), amount, sig, tokenAddr)
	if err != nil {
		return err
	}
	return client.WaitForTxConfirmation(context.TODO(), c.ethClient, tx, c.TxTimeout)
}

// WithdrawETH sends a tx to the Mainnet Gateway to withdraw the specified amount of ETH,
// and returns the tx fee.
func (c *MainnetGatewayClient) WithdrawETH(caller *client.Identity, amount *big.Int, sig []byte) (*big.Int, error) {
	tx, err := c.contract.WithdrawETH(client.DefaultTransactOptsForIdentity(caller), amount, sig)
	if err != nil {
		return nil, err
	}
	return client.WaitForTxConfirmationAndFee(context.TODO(), c.ethClient, tx, c.TxTimeout)
}


func ConnectToMainnetGateway(ethClient *ethclient.Client, gatewayAddr string) (*MainnetGatewayClient, error) {
	contractAddr := common.HexToAddress(gatewayAddr)
	contract, err := NewMainnetGatewayContract(contractAddr, ethClient)
	if err != nil {
		return nil, err
	}
	return &MainnetGatewayClient{
		ethClient: ethClient,
		contract:  contract,
		Address:   contractAddr,
	}, nil
}
