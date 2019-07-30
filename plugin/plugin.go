package plugin

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	extplugin "github.com/hashicorp/go-plugin"
	"github.com/loomnetwork/go-loom"
	cctypes "github.com/loomnetwork/go-loom/builtin/types/chainconfig"
	"github.com/loomnetwork/go-loom/plugin/types"
	ltypes "github.com/loomnetwork/go-loom/types"
	"github.com/loomnetwork/go-loom/vm"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
)

// Handshake is a common handshake that is shared by plugin and host.
var Handshake = extplugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "LOOM_CONTRACT",
	MagicCookieValue: "loomrocks",
}

// Create some standard server metrics.
var (
	grpcMetrics = grpc_prometheus.NewServerMetrics()
	reg         = prometheus.NewRegistry()
)

func init() {
	reg.MustRegister(grpcMetrics)
}

type GRPCAPIClient struct {
	client types.APIClient
}

func (c *GRPCAPIClient) Get(key []byte) []byte {
	resp, _ := c.client.Get(context.TODO(), &types.GetRequest{Key: key})
	return resp.Value
}

func (c *GRPCAPIClient) Range(prefix []byte) RangeData {
	ret := make(RangeData, 0)
	resp, _ := c.client.Range(context.TODO(), &types.RangeRequest{Prefix: prefix})
	for _, x := range resp.RangeEntries {
		r := &RangeEntry{
			Key:   x.Key,
			Value: x.Value,
		}
		ret = append(ret, r)
	}
	return ret
}

func (c *GRPCAPIClient) Has(key []byte) bool {
	resp, _ := c.client.Has(context.TODO(), &types.HasRequest{Key: key})
	return resp.Value
}

func (c *GRPCAPIClient) GetEvmTxReceipt(hash []byte) (types.EvmTxReceipt, error) {
	resp, err := c.client.GetEvmTxReceipt(context.TODO(), &types.EvmTxReceiptRequest{Value: hash})
	return *resp, err
}

func (c *GRPCAPIClient) Set(key, value []byte) {
	c.client.Set(context.TODO(), &types.SetRequest{Key: key, Value: value})
}

func (c *GRPCAPIClient) Delete(key []byte) {
	c.client.Delete(context.TODO(), &types.DeleteRequest{Key: key})
}

func (c *GRPCAPIClient) staticCall(addr loom.Address, input []byte, vmType vm.VMType) ([]byte, error) {
	resp, err := c.client.StaticCall(context.TODO(), &types.CallRequest{
		Address: addr.MarshalPB(),
		Input:   input,
		VmType:  vmType,
	})
	if err != nil {
		return nil, err
	}

	return resp.Output, nil
}

func (c *GRPCAPIClient) StaticCall(addr loom.Address, input []byte) ([]byte, error) {
	return c.staticCall(addr, input, vm.VMType_PLUGIN)
}

func (c *GRPCAPIClient) StaticCallEVM(addr loom.Address, input []byte) ([]byte, error) {
	return c.staticCall(addr, input, vm.VMType_EVM)
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

func (c *GRPCAPIClient) call(addr loom.Address, input []byte, vmType vm.VMType, value *loom.BigUInt) ([]byte, error) {
	if value == nil {
		value = loom.NewBigUIntFromInt(0)
	}
	resp, err := c.client.Call(context.TODO(), &types.CallRequest{
		Address: addr.MarshalPB(),
		Input:   input,
		VmType:  vmType,
		Value:   &ltypes.BigUInt{Value: *value},
	})
	if err != nil {
		return nil, err
	}

	return resp.Output, nil
}

func (c *GRPCAPIClient) Call(addr loom.Address, input []byte) ([]byte, error) {
	return c.call(addr, input, vm.VMType_PLUGIN, loom.NewBigUIntFromInt(0))
}

func (c *GRPCAPIClient) CallEVM(addr loom.Address, input []byte, value *loom.BigUInt) ([]byte, error) {
	return c.call(addr, input, vm.VMType_EVM, value)
}

func (c *GRPCAPIClient) ContractRecord(contractAddr loom.Address) (*ContractRecord, error) {
	resp, err := c.client.ContractRecord(context.TODO(), &types.ContractRecordRequest{
		Contract: contractAddr.MarshalPB(),
	})
	if err != nil {
		return nil, err
	}
	return &ContractRecord{
		ContractName:    resp.ContractName,
		ContractAddress: loom.UnmarshalAddressPB(resp.ContractAddress),
		CreatorAddress:  loom.UnmarshalAddressPB(resp.CreatorAddress),
	}, nil
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

func (c *GRPCContext) EmitTopics(data []byte, topics ...string) {
	if topics == nil {
		topics = []string{}
	}
	c.client.Emit(context.TODO(), &types.EmitRequest{Data: data, Topics: topics})
}

func (c *GRPCContext) Emit(data []byte) {
	c.EmitTopics(data)
}

func (c *GRPCContext) FeatureEnabled(name string, defaultVal bool) bool {
	return c.FeatureEnabled(name, defaultVal)
}

func (c *GRPCContext) Config() *cctypes.Config {
	return c.Config()
}

func (c *GRPCContext) EnabledFeatures() []string {
	return c.EnabledFeatures()
}

func (c *GRPCContext) Validators() []*ltypes.Validator {
	return c.Validators()
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
	meta, err := contract.Meta()
	if err != nil {
		fmt.Fprintf(os.Stderr, "contract meta error %v", err)
		os.Exit(1)
	}

	// default hostport for metrics
	var hostport = "127.0.0.1:9092"
	if meta.MetricAddr != "" {
		hostport = meta.MetricAddr
	}

	host, port, err := net.SplitHostPort(hostport)
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid metric address: %s", err)
		os.Exit(1)
	}
	// Serve promtheus http server
	httpServer := &http.Server{
		Handler: promhttp.HandlerFor(reg, promhttp.HandlerOpts{}),
		Addr:    net.JoinHostPort(host, port),
	}
	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			fmt.Fprintf(os.Stderr, "unable to start http server: %v", err)
			os.Exit(1)
		}
	}()

	// Serve the plugin
	extplugin.Serve(&extplugin.ServeConfig{
		HandshakeConfig: Handshake,
		Plugins: map[string]extplugin.Plugin{
			"contract": &ExternalPlugin{Impl: contract},
		},

		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: func(opts []grpc.ServerOption) *grpc.Server {
			// add prometheus plugin
			promOpts := []grpc.ServerOption{
				grpc.StreamInterceptor(grpcMetrics.StreamServerInterceptor()),
				grpc.UnaryInterceptor(grpcMetrics.UnaryServerInterceptor()),
			}
			opts = append(opts, promOpts...)
			s := grpc.NewServer(opts...)
			// initialize metrics
			grpcMetrics.InitializeMetrics(s)
			return s
		},
	})
}
