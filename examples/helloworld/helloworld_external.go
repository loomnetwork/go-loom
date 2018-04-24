// +build external

// go build -tags external -o contracts/helloworld.1.0.0 github.com/loomnetwork/loom-plugin/examples/helloworld
package main

import (
	plugin "github.com/hashicorp/go-plugin"

	lp "github.com/loomnetwork/loom-plugin"
)

func main() {
	var contract lp.Contract = &HelloWorld{}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: lp.Handshake,
		Plugins: map[string]plugin.Plugin{
			"contract": &lp.ExternalPlugin{Impl: contract},
		},

		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
