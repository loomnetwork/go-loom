package plugin

import (
	"context"
	"errors"
	"time"

	extplugin "github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"

	loom "github.com/loomnetwork/go-loom"
	"github.com/loomnetwork/go-loom/plugin/types"
	"github.com/loomnetwork/go-loom/vm"
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
	resp, _ := c.client.Has(context.TODO(), &types.HasRequest{Key: key})
	return resp.Value
}

func (c *GRPCAPIClient) ValidatorPower(pubKey []byte) int64 {
	resp, _ := c.client.ValidatorPower(context.TODO(), &types.ValidatorPowerRequest{
		PubKey: pubKey,
	})

	return resp.Power
}

func (c *GRPCAPIClient) Set(key, value []byte) {
	c.client.Set(context.TODO(), &types.SetRequest{Key: key, Value: value})
}

func (c *GRPCAPIClient) Delete(key []byte) {
	c.client.Delete(context.TODO(), &types.DeleteRequest{Key: key})
}

func (c *GRPCAPIClient) StaticCall(addr loom.Address, input []byte) ([]byte, error) {
	resp, err := c.client.StaticCall(context.TODO(), &types.CallRequest{
		Address: addr.MarshalPB(),
		Input:   input,
	})
	if err != nil {
		return nil, err
	}

	return resp.Output, nil
}

func (c *GRPCAPIClient) Resolve(name string) (loom.Address, error) {
	resp, err := c.client.Resolve(context.TODO(), &types.ResolveRequest{
		Name: name,
	})
	if err != nil {
		return loom.Address{}, err
	}

	return loom.UnmarshalAddressPB(resp.Address), nil
}

func (c *GRPCAPIClient) call(addr loom.Address, input []byte, vmType vm.VMType) ([]byte, error) {
	resp, err := c.client.Call(context.TODO(), &types.CallRequest{
		Address: addr.MarshalPB(),
		Input:   input,
		VmType:  vmType,
	})
	if err != nil {
		return nil, err
	}

	return resp.Output, nil
}

func (c *GRPCAPIClient) Call(addr loom.Address, input []byte) ([]byte, error) {
	return c.call(addr, input, vm.VMType_PLUGIN)
}

func (c *GRPCAPIClient) CallEVM(addr loom.Address, input []byte) ([]byte, error) {
	return c.call(addr, input, vm.VMType_EVM)
}

func (c *GRPCAPIClient) SetValidatorPower(pubKey []byte, power int64) {
	c.client.SetValidatorPower(context.TODO(), &types.SetValidatorPowerRequest{
		PubKey: pubKey,
		Power:  power,
	})
}

type GRPCContext struct {
	*GRPCAPIClient
	message      *types.Message
	block        *loom.BlockHeader
	contractAddr loom.Address
}

var _ Context = &GRPCContext{}

func (c *GRPCContext) Block() loom.BlockHeader {
	return *c.block
}

func (c *GRPCContext) Now() time.Time {
	return time.Unix(c.block.Time, 0)
}

func (c *GRPCContext) ContractAddress() loom.Address {
	return c.contractAddr
}

func (c *GRPCContext) Message() Message {
	return Message{
		Sender: loom.UnmarshalAddressPB(c.message.Sender),
	}
}

func (c *GRPCContext) Emit(data []byte) {
}

func MakeGRPCContext(conn *grpc.ClientConn, req *types.ContractCallRequest) *GRPCContext {
	return &GRPCContext{
		GRPCAPIClient: &GRPCAPIClient{
			client: types.NewAPIClient(conn),
		},
		message:      req.Message,
		block:        req.Block,
		contractAddr: loom.UnmarshalAddressPB(req.ContractAddress),
	}
}

type GRPCContractServer struct {
	broker *extplugin.GRPCBroker
	Impl   Contract
}

var _ types.ContractServer = &GRPCContractServer{}

func (s *GRPCContractServer) Meta(ctx context.Context, req *types.MetaRequest) (*types.ContractMeta, error) {
	meta, err := s.Impl.Meta()
	return &meta, err
}

func (s *GRPCContractServer) Init(ctx context.Context, req *types.ContractCallRequest) (*types.InitResponse, error) {
	conn, err := s.broker.Dial(req.ApiServer)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	return &types.InitResponse{}, s.Impl.Init(MakeGRPCContext(conn, req), req.Request)
}

func (s *GRPCContractServer) Call(ctx context.Context, req *types.ContractCallRequest) (*types.Response, error) {
	conn, err := s.broker.Dial(req.ApiServer)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	return s.Impl.Call(MakeGRPCContext(conn, req), req.Request)
}

func (s *GRPCContractServer) StaticCall(ctx context.Context, req *types.ContractCallRequest) (*types.Response, error) {
	conn, err := s.broker.Dial(req.ApiServer)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	return s.Impl.StaticCall(MakeGRPCContext(conn, req), req.Request)
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
