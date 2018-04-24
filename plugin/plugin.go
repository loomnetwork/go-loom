package plugin

import (
	"context"
	"errors"
	"time"

	extplugin "github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"

	loom "github.com/loomnetwork/loom-plugin"
	"github.com/loomnetwork/loom-plugin/types"
)

// Handshake is a common handshake that is shared by plugin and host.
var Handshake = extplugin.HandshakeConfig{
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

func (c *GRPCAPIClient) StaticCall(addr loom.Address, input []byte) ([]byte, error) {
	return nil, nil
}

func (c *GRPCAPIClient) Call(addr loom.Address, input []byte) ([]byte, error) {
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

func (c *GRPCContext) ContractAddress() loom.Address {
	return loom.Address{}
}

func (c *GRPCContext) Message() types.Message {
	return types.Message{}
}

func (c *GRPCContext) Emit(data []byte) {
}

type GRPCContractServer struct {
	broker *extplugin.GRPCBroker
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
	extplugin.NetRPCUnsupportedPlugin
	// Concrete implementation, written in Go. This is only used for plugins
	// that are written in Go.
	Impl Contract
}

var _ extplugin.GRPCPlugin = &ExternalPlugin{}

func (p *ExternalPlugin) GRPCServer(broker *extplugin.GRPCBroker, s *grpc.Server) error {
	types.RegisterContractServer(s, &GRPCContractServer{
		broker: broker,
		Impl:   p.Impl,
	})
	return nil
}

func (p *ExternalPlugin) GRPCClient(ctx context.Context, broker *extplugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return nil, errors.New("not implemented on plugin side")
}

func Serve(contract Contract) {
	extplugin.Serve(&extplugin.ServeConfig{
		HandshakeConfig: Handshake,
		Plugins: map[string]extplugin.Plugin{
			"contract": &ExternalPlugin{Impl: contract},
		},

		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: extplugin.DefaultGRPCServer,
	})
}
