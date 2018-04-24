package loomplugin

import (
	"context"
	"errors"
	"time"

	plugin "github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"

	"github.com/loomnetwork/loom-plugin/types"
)

// Handshake is a common handshake that is shared by plugin and host.
var Handshake = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "LOOM_CONTRACT",
	MagicCookieValue: "loomrocks",
}

type GRPCAPIClient struct {
	client types.APIClient
}

func (c *GRPCAPIClient) Get(key []byte) []byte {
	resp, _ := c.client.Get(context.TODO(), &types.GetRequest{Key: key})
	return resp.Value
}

func (c *GRPCAPIClient) Has(key []byte) bool {
	return false
}

func (c *GRPCAPIClient) Set(key, value []byte) {

}

func (c *GRPCAPIClient) Delete(key []byte) {

}

func (c *GRPCAPIClient) StaticCall(addr Address, input []byte) ([]byte, error) {
	return nil, nil
}

func (c *GRPCAPIClient) Call(addr Address, input []byte) ([]byte, error) {
	return nil, nil
}

type GRPCContext struct {
	*GRPCAPIClient
	block *types.BlockHeader
}

var _ Context = &GRPCContext{}

func (c *GRPCContext) Block() types.BlockHeader {
	return *c.block
}

func (c *GRPCContext) Now() time.Time {
	return time.Unix(c.block.Time, 0)
}

func (c *GRPCContext) ContractAddress() Address {
	return Address{}
}

func (c *GRPCContext) Message() types.Message {
	return types.Message{}
}

func (c *GRPCContext) Emit(data []byte) {
}

type GRPCContractServer struct {
	broker *plugin.GRPCBroker
	Impl   Contract
}

var _ types.ContractServer = &GRPCContractServer{}

func (s *GRPCContractServer) Meta(ctx context.Context, req *types.MetaRequest) (*types.ContractMeta, error) {
	return nil, nil
}

func (s *GRPCContractServer) Init(ctx context.Context, req *types.ContractCallRequest) (*types.InitResponse, error) {
	conn, err := s.broker.Dial(req.ApiServer)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	pctx := &GRPCContext{
		GRPCAPIClient: &GRPCAPIClient{
			client: types.NewAPIClient(conn),
		},
		block: req.Block,
	}

	return &types.InitResponse{}, s.Impl.Init(pctx, req.Request)
}

func (s *GRPCContractServer) Call(ctx context.Context, req *types.ContractCallRequest) (*types.Response, error) {
	return nil, nil
}

func (s *GRPCContractServer) StaticCall(ctx context.Context, req *types.ContractStaticCallRequest) (*types.Response, error) {
	return nil, nil
}

type ExternalPlugin struct {
	plugin.NetRPCUnsupportedPlugin
	// Concrete implementation, written in Go. This is only used for plugins
	// that are written in Go.
	Impl Contract
}

var _ plugin.GRPCPlugin = &ExternalPlugin{}

func (p *ExternalPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	types.RegisterContractServer(s, &GRPCContractServer{
		broker: broker,
		Impl:   p.Impl,
	})
	return nil
}

func (p *ExternalPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return nil, errors.New("not implemented on plugin side")
}
