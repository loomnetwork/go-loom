package contractpb

import (
	"errors"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/gogo/protobuf/proto"

	"github.com/loomnetwork/go-loom"
	"github.com/loomnetwork/go-loom/plugin"
	"github.com/loomnetwork/go-loom/plugin/types"
)

var (
	ErrNotFound = errors.New("not found")
)

// StaticContext is the high-level context provided to Go contract methods that don't mutate state.
type StaticContext interface {
	plugin.StaticAPI
	Get(key []byte, pb proto.Message) error
	Range(prefix []byte) (plugin.RangeData, error)
	Has(key []byte) bool
	Block() loom.BlockHeader
	Now() time.Time
	Message() plugin.Message
	ContractAddress() loom.Address
	Logger() *loom.Logger
	GetEvmTxReceipt([]byte) (types.EvmTxReceipt, error)
	HasPermissionFor(addr loom.Address, token []byte, roles []string) (bool, []string, error)
	FeatureEnabled(name string, defaultVal bool) bool

	// ContractRecord retrieves the contract meta data stored in the Registry.
	// NOTE: This method requires Registry v2.
	ContractRecord(contractAddr loom.Address) (*plugin.ContractRecord, error)
}

// Context is the high-level context provided to Go contract methods that mutate state.
type Context interface {
	plugin.VolatileAPI
	StaticContext
	Set(key []byte, pb proto.Message) error
	Delete(key []byte) error
	HasPermission(token []byte, roles []string) (bool, []string, error)
	GrantPermissionTo(addr loom.Address, token []byte, role string)
	RevokePermissionFrom(addr loom.Address, token []byte, role string)
	GrantPermission(token []byte, roles []string)
}

type Contract interface {
	Meta() (plugin.Meta, error)
}

// Implements the StaticContext interface for Go contract methods.
type wrappedPluginStaticContext struct {
	plugin.StaticContext
	logger *loom.Logger
}

var _ StaticContext = &wrappedPluginStaticContext{}

func (c *wrappedPluginStaticContext) Logger() *loom.Logger {
	return c.logger
}

func (c *wrappedPluginStaticContext) Range(prefix []byte) ( plugin.RangeData,error) {
	ret := make( plugin.RangeData, 0)
	return ret,nil
}

func (c *wrappedPluginStaticContext) Has (key []byte) bool {
	return false
}

func (c *wrappedPluginStaticContext) Delete (key []byte) error {
	return nil
}

func (c *wrappedPluginStaticContext) Get(key []byte, pb proto.Message) error {
	data,err := c.StaticContext.Get(key)
	if err != nil {
		return err
	}
	if len(data) == 0 {
		return ErrNotFound
	}

	return proto.Unmarshal(data, pb)
}

// HasPermissionFor checks whether the given `addr` has any of the permission given in `roles` on `token`
func (c *wrappedPluginStaticContext) HasPermissionFor(addr loom.Address, token []byte, roles []string) (bool, []string, error) {
	found := false
	foundRoles := []string{}
	for _, role := range roles {
		v,err := c.StaticContext.Get(rolePermKey(addr, token, role))
		if err != nil {
			return false,nil,err
		}
		if v != nil && string(v) == "true" {
			found = true
			foundRoles = append(foundRoles, role)
		}
	}
	return found, foundRoles, nil
}

// FeatureEnabled checks whether the feature is enabled on chain
func (c *wrappedPluginStaticContext) FeatureEnabled(name string, defaultVal bool) bool {
	return c.StaticContext.FeatureEnabled(name, defaultVal)
}

// Implements the Context interface for Go contract methods.
type wrappedPluginContext struct {
	plugin.Context
	wrappedPluginStaticContext
}

var _ Context = &wrappedPluginContext{}

func (c *wrappedPluginContext) Get(key []byte, pb proto.Message) error {
	return c.wrappedPluginStaticContext.Get(key, pb)
}

func (c *wrappedPluginContext) Has (key []byte) bool {
	return false
}

func (c *wrappedPluginContext) Delete (key []byte) error {
	return nil
}


func (c *wrappedPluginContext) Set(key []byte, pb proto.Message) error {
	enc, err := proto.Marshal(pb)
	if err != nil {
		return err
	}
	c.Context.Set(key, enc)
	return nil
}

// HasPermission checks whether the sender of the tx has any of the permission given in `roles` on `token`
func (c *wrappedPluginContext) HasPermission(token []byte, roles []string) (bool, []string, error) {
	addr := c.Message().Sender
	found, foundRoles, err :=  c.HasPermissionFor(addr, token, roles)
	return found, foundRoles, err
}


func (c *wrappedPluginContext) Range(prefix []byte) ( plugin.RangeData,error) {
	ret := make( plugin.RangeData, 0)
	return ret,nil
}

// GrantPermissionTo sets a given `role` permission on `token` for the given `addr`
func (c *wrappedPluginContext) GrantPermissionTo(addr loom.Address, token []byte, role string) {
	c.Context.Set(rolePermKey(addr, token, role), []byte("true"))
}

// RevokePermissionFrom removes a permission previously granted by GrantPermissionTo
func (c *wrappedPluginContext) RevokePermissionFrom(addr loom.Address, token []byte, role string) {
	c.Context.Delete(rolePermKey(addr, token, role))
}

// Check if feature is enabled on chain
func (c *wrappedPluginContext) FeatureEnabled(name string, defaultVal bool) bool {
	return c.Context.FeatureEnabled(name, defaultVal)
}

func rolePermKey(addr loom.Address, token []byte, role string) []byte {
	// TODO: This generates an overly long key, the key generated here is prefixed by the contract
	//       address, but the wrappedPluginContext only has access to the state prefixed by the
	//       contract address, so all the permission keys are effectively prefixed by the contract
	//       address twice!
	return []byte(fmt.Sprintf("%stoken:%s:role:%s", loom.PermPrefix(addr), token, []byte(role)))
}

// GrantPermission sets a given `role` permission on `token` for the sender of the tx
func (c *wrappedPluginContext) GrantPermission(token []byte, roles []string) {
	for _, r := range roles {
		c.GrantPermissionTo(c.Message().Sender, token, r)
	}
}

func MakePluginContract(c Contract) plugin.Contract {
	r, err := NewRequestDispatcher(c)
	if err != nil {
		panic(err)
	}
	setupLogger()

	return r
}

func Call(ctx Context, addr loom.Address, inpb proto.Message, outpb proto.Message) error {
	input, err := makeEnvelope(inpb)
	if err != nil {
		return err
	}

	output, err := ctx.Call(addr, input)
	if err != nil {
		return err
	}

	var resp plugin.Response
	err = proto.Unmarshal(output, &resp)
	if err != nil {
		return err
	}

	if outpb != nil {
		err = proto.Unmarshal(resp.Body, outpb)
		if err != nil {
			return err
		}
	}

	return nil
}

func CallMethod(ctx Context, addr loom.Address, method string, inpb proto.Message, outpb proto.Message) error {
	args, err := proto.Marshal(inpb)
	if err != nil {
		return err
	}

	query := &types.ContractMethodCall{
		Method: method,
		Args:   args,
	}

	return Call(ctx, addr, query, outpb)
}

var logger *loom.Logger
var onceSetup sync.Once

func setupLogger() {
	onceSetup.Do(func() {
		level := "info"
		envLevel := os.Getenv("CONTRACT_LOG_LEVEL")
		if envLevel != "" {
			level = envLevel
		}
		dest := "file://-"
		envDest := os.Getenv("CONTRACT_LOG_DESTINATION")
		if envDest != "" {
			dest = envDest
		}
		logger = loom.NewLoomLogger(level, dest)
	})
}

func StaticCall(ctx StaticContext, addr loom.Address, inpb proto.Message, outpb proto.Message) error {
	input, err := makeEnvelope(inpb)
	if err != nil {
		return err
	}

	output, err := ctx.StaticCall(addr, input)
	if err != nil {
		return err
	}

	var resp plugin.Response
	err = proto.Unmarshal(output, &resp)
	if err != nil {
		return err
	}

	if outpb != nil {
		err = proto.Unmarshal(resp.Body, outpb)
		if err != nil {
			return err
		}
	}

	return nil
}

func StaticCallMethod(ctx StaticContext, addr loom.Address, method string, inpb proto.Message, outpb proto.Message) error {
	args, err := proto.Marshal(inpb)
	if err != nil {
		return err
	}

	query := &types.ContractMethodCall{
		Method: method,
		Args:   args,
	}

	return StaticCall(ctx, addr, query, outpb)
}

func CallEVM(ctx Context, addr loom.Address, input []byte, output *[]byte) error {
	resp, err := ctx.CallEVM(addr, input, loom.NewBigUIntFromInt(0))
	*output = resp
	return err
}

func StaticCallEVM(ctx StaticContext, addr loom.Address, input []byte, output *[]byte) error {
	resp, err := ctx.StaticCallEVM(addr, input)
	*output = resp
	return err
}

func WrapPluginContext(ctx plugin.Context) Context {
	return &wrappedPluginContext{ctx, wrappedPluginStaticContext{ctx, logger}}
}

func WrapPluginStaticContext(ctx plugin.StaticContext) StaticContext {
	return &wrappedPluginStaticContext{ctx, logger}
}

func makeEnvelope(inpb proto.Message) ([]byte, error) {
	body, err := proto.Marshal(inpb)
	if err != nil {
		return nil, err
	}

	req := &plugin.Request{
		ContentType: plugin.EncodingType_PROTOBUF3,
		Accept:      plugin.EncodingType_PROTOBUF3,
		Body:        body,
	}

	return proto.Marshal(req)
}
