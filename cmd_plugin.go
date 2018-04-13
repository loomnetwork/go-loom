package loomplugin

import "github.com/spf13/cobra"

// Command is an alias for cobra.Command that must be used in cmd plugins in order to avoid type
// collisions with the Loom SDK.
type Command = cobra.Command

// CmdPluginSystem interface is used by command plugins to hook into the Loom admin CLI.
type CmdPluginSystem interface {
	// GetClient returns a DAppChainClient that can be used to commit txs to a Loom DAppChain.
	GetClient(nodeURI string) (DAppChainClient, error)
}

type CmdPlugin interface {
	Init(sys CmdPluginSystem) error
	GetCmds() []*Command
}
