package plugin

import (
	"github.com/gogo/protobuf/proto"
	lp "github.com/loomnetwork/loom-plugin"
	pb "github.com/loomnetwork/loom-plugin/examples/types"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ed25519"
)

const (
	nodeUriFlag = "node"
)

// CreateTxCmdPlugin is a sample admin CLI cmd plugin that creates a new dummy tx & commits it to the DAppChain.
type CreateTxCmdPlugin struct {
	cmdPluginSystem lp.CmdPluginSystem
}

func (c *CreateTxCmdPlugin) Init(sys lp.CmdPluginSystem) error {
	c.cmdPluginSystem = sys
	return nil
}

func (c *CreateTxCmdPlugin) GetCmds() []*lp.Command {
	cmd := &lp.Command{
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

func (c *CreateTxCmdPlugin) runCmd(cmd *lp.Command, args []string) error {
	nodeUri, err := cmd.Flags().GetString(nodeUriFlag)
	if err != nil {
		return err
	}
	client, err := c.cmdPluginSystem.GetClient(nodeUri)
	if err != nil {
		return err
	}
	dummyValue := args[0]
	dummyTx := pb.DummyTx{
		Key: "hello",
		Val: dummyValue,
	}
	txBytes, err := proto.Marshal(&dummyTx)
	_, privKey, err := ed25519.GenerateKey(nil)
	if err != nil {
		return err
	}
	signer := lp.NewEd25519Signer(privKey)
	client.CommitTx(signer, txBytes)
	return nil
}
