// +build evm

package commands

import (
	"encoding/hex"
	"fmt"

	"github.com/loomnetwork/go-loom/cli"
	"github.com/loomnetwork/go-loom/client/gateway"

	ssha "github.com/miguelmota/go-solidity-sha3"

	"github.com/ethereum/go-ethereum/crypto"
	ethclient "github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
)

func SignRemoveValidator() *cobra.Command {
	return &cobra.Command{
		Use:   "remove [validator eth address] [gateway mainnet address] [signer eth key path] [eth endpoint]",
		Short: "Returns a signature for the addition of a new validator. Must be sorted together with the other signatures to match the multisig threshold",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			validator := args[0]
			privateKeyPath := args[1]
			gatewayAddr := cli.TxFlags.ContractAddr
			ethUri := cli.TxFlags.MainnetURI

			// Get the nonce
			ethClient, err := ethclient.Dial(ethUri)
			gatewayClient, err := gateway.ConnectToMainnetGateway(ethClient, gatewayAddr)
			nonce, err := gatewayClient.Contract().Nonce(nil)
			if err != nil {
				return err
			}

			operation := "remove"
			hash := ssha.SoliditySHA3(
				ssha.Address(gatewayAddr),
				ssha.Uint256(nonce),
				ssha.SoliditySHA3(
					ssha.String(operation),
					ssha.Address(validator),
				),
			)

			// TODO: Check if hsmconfigfile flag is given-> use that, otherwise privfile
			key, err := crypto.LoadECDSA(privateKeyPath)
			if err != nil {
				fmt.Println(err)
			}

			sig, err := crypto.Sign(hash, key)
			if err != nil {
				fmt.Println(err)

			}
			fmt.Printf("Operation: %s %s\n", operation, validator)
			fmt.Printf("Gateway Address: %s\n", gatewayAddr)
			fmt.Printf("Gateway Nonce: %d\n", nonce)
			fmt.Printf("Signer: %s\n", crypto.PubkeyToAddress(key.PublicKey).String())
			fmt.Printf("Msg hash: %s\n", hex.EncodeToString(hash))
			fmt.Printf("{signer: %s, signature: 0x%s}\n", crypto.PubkeyToAddress(key.PublicKey).String(), hex.EncodeToString(sig))
			return nil
		},
	}
}

func SignAddValidator() *cobra.Command {
	return &cobra.Command{
		Use:   "add [validator eth address] [gateway mainnet address] [signer eth key path] [eth endpoint",
		Short: "Returns a signature for the addition of a new validator. Must be sorted together with the other signatures to match the multisig threshold",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			validator := args[0]
			privateKeyPath := args[1]
			gatewayAddr := cli.TxFlags.ContractAddr
			ethUri := cli.TxFlags.MainnetURI

			// Get the nonce
			ethClient, err := ethclient.Dial(ethUri)
			gatewayClient, err := gateway.ConnectToMainnetGateway(ethClient, gatewayAddr)
			nonce, err := gatewayClient.Contract().Nonce(nil)
			if err != nil {
				return err
			}

			operation := "add"
			hash := ssha.SoliditySHA3(
				ssha.Address(gatewayAddr),
				ssha.Uint256(nonce),
				ssha.SoliditySHA3(
					ssha.String(operation),
					ssha.Address(validator),
				),
			)
			key, err := crypto.LoadECDSA(privateKeyPath)
			if err != nil {
				fmt.Println(err)
			}

			sig, err := crypto.Sign(hash, key)
			if err != nil {
				fmt.Println(err)

			}
			fmt.Printf("Operation: %s %s\n", operation, validator)
			fmt.Printf("Gateway Address: %s\n", gatewayAddr)
			fmt.Printf("Gateway Nonce: %d\n", nonce)
			fmt.Printf("Signer: %s\n", crypto.PubkeyToAddress(key.PublicKey).String())
			fmt.Printf("Msg hash: 0x%s\n", hex.EncodeToString(hash))
			fmt.Printf("{signer: %s, signature: 0x%s}\n", crypto.PubkeyToAddress(key.PublicKey).String(), hex.EncodeToString(sig))
			return nil
		},
	}
}

func AddValidatorCommands(root *cobra.Command) {
	root.AddCommand(
		SignAddValidator(),
		SignRemoveValidator(),
	)
}
