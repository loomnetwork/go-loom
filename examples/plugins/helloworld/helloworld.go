package main

import (
	"github.com/loomnetwork/go-loom/examples/types"
	"github.com/loomnetwork/go-loom/plugin"
	contract "github.com/loomnetwork/go-loom/plugin/contractpb"
)

type HelloWorld struct {
}

func (c *HelloWorld) Meta() (plugin.Meta, error) {
	return plugin.Meta{
		Name:    "helloworld",
		Version: "1.0.0",
	}, nil
}

func (c *HelloWorld) Init(ctx contract.Context, req *types.HelloRequest) {
	data := &types.Dummy{
		Key:   "foo",
		Value: "bar",
	}

	ctx.Set([]byte("mydata"), data)
}

func (c *HelloWorld) Hello(ctx contract.StaticContext, req *types.HelloRequest) (*types.HelloResponse, error) {
	return &types.HelloResponse{
		Out: "Hello World!",
	}, nil
}

var Contract plugin.Contract = contract.MakePluginContract(&HelloWorld{})

func main() {
	plugin.Serve(Contract)
}
