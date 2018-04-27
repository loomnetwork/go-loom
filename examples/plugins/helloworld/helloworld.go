package main

import (
	"encoding/json"
	"errors"

	"github.com/loomnetwork/go-loom/plugin"
)

type rpcRequest struct {
	Body string `json:"body"`
}
type rpcResponse struct {
	Body string `json:"body"`
}

type HelloWorld struct {
}

func (c *HelloWorld) Meta() (plugin.Meta, error) {
	return plugin.Meta{
		Name:    "helloworld",
		Version: "1.0.0",
	}, nil
}

func (c *HelloWorld) Init(ctx plugin.Context, req *plugin.Request) error {
	println("init contract")
	ctx.Set([]byte("foo"), []byte("bar"))
	return nil
}

func (c *HelloWorld) Call(ctx plugin.Context, req *plugin.Request) (*plugin.Response, error) {
	return &plugin.Response{}, nil
}

func (c *HelloWorld) StaticCall(ctx plugin.StaticContext, req *plugin.Request) (*plugin.Response, error) {
	rr := &rpcRequest{}
	if req.ContentType == plugin.EncodingType_JSON {
		if err := json.Unmarshal(req.Body, rr); err != nil {
			return nil, err
		}
	} else {
		// content type could also be protobuf
		return nil, errors.New("unsupported content type")
	}
	if "hello" == rr.Body {
		var body []byte
		var err error
		if req.Accept == plugin.EncodingType_JSON {
			body, err = json.Marshal(&rpcResponse{Body: "world"})
			if err != nil {
				return nil, err
			}
			return &plugin.Response{
				ContentType: plugin.EncodingType_JSON,
				Body:        body,
			}, nil
		}
		// accepted content type could also be protobuf
		return nil, errors.New("unsupported content type")
	}
	return nil, errors.New("invalid query")
}

var Contract plugin.Contract = &HelloWorld{}

func main() {
	plugin.Serve(Contract)
}
