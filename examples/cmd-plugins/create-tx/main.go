package main

import (
	"github.com/loomnetwork/go-loom/examples/cmd-plugins/create-tx/plugin"
)

// Create an instance of the plugin that will be loaded by the plugin manager.
var CmdPlugin plugin.CreateTxCmdPlugin

// go-code-check throws up errors if this is missing...
func main() {}
