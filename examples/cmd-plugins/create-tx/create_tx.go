package main

import (
	"github.com/gogo/protobuf/proto"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ed25519"

	loom "github.com/loomnetwork/go-loom"
	"github.com/loomnetwork/go-loom/examples/types"
)

const (
	nodeUriFlag = "node"
)

// CreateTxCmdPlugin is a sample admin CLI cmd plugin that creates a new dummy tx & commits it to the DAppChain.
type CreateTxCmdPlugin struct {
	cmdPluginSystem loom.CmdPluginSystem
}

func (c *CreateTxCmdPlugin) Init(sys loom.CmdPluginSystem) error {
	c.cmdPluginSystem = sys
	return nil
}

func (c *CreateTxCmdPlugin) GetCmds() []*loom.Command {
	cmd := &loom.Command{
		Use:   "create-tx <value>",
		Short: "Create & commit a dummy tx to the DAppChain",
		Args:  cobra.ExactArgs(1),
		RunE:  c.runCmd,
	}
	cmd.Flags().StringP(
		nodeUriFlag, "n", "tcp://0.0.0.0:46657",
		"URI of node to administer, in the form tcp://<host>:<port>")
	return []*cobra.Command{cmd}
}

func (c *CreateTxCmdPlugin) runCmd(cmd *loom.Command, args []string) error {
	nodeUri, err := cmd.Flags().GetString(nodeUriFlag)
	if err != nil {
		return err
	}
	client, err := c.cmdPluginSystem.GetClient(nodeUri)
	if err != nil {
		return err
	}
	dummyValue := args[0]
	dummyTx := types.Dummy{
		Key:   "hello",
		Value: dummyValue,
	}
	txBytes, err := proto.Marshal(&dummyTx)
	_, privKey, err := ed25519.GenerateKey(nil)
	if err != nil {
		return err
	}
	signer := loom.NewEd25519Signer(privKey)
	client.CommitTx(signer, txBytes)
	return nil
}

// Create an instance of the plugin that will be loaded by the plugin manager.
var CmdPlugin CreateTxCmdPlugin

// go-code-check throws up errors if this is missing...
func main() {}
