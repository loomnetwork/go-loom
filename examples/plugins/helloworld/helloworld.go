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
	data := &types.MapEntry{
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

func (c *HelloWorld) SetMsg(ctx contract.Context, req *types.MapEntry) error {
	return ctx.Set([]byte(req.Key), req)
}

func (c *HelloWorld) GetMsg(ctx contract.StaticContext, req *types.MapEntry) (*types.MapEntry, error) {
	var result types.MapEntry
	if err := ctx.Get([]byte(req.Key), &result); err != nil {
		return nil, err
	}
	return &result, nil
}

var Contract plugin.Contract = contract.MakePluginContract(&HelloWorld{})

func main() {
	plugin.Serve(Contract)
}
