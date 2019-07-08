package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
)

type RPCRequest struct {
	Version string          `json:"jsonrpc"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params"` // map[string]interface{} or []interface{}
	ID      string          `json:"id"`
}

type RPCResponse struct {
	Version string          `json:"jsonrpc"`
	ID      string          `json:"id"`
	Result  json.RawMessage `json:"result,omitempty"`
	Error   *RPCError       `json:"error,omitempty"`
}

type RPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data,omitempty"`
}

func (err RPCError) Error() string {
	if err.Data != "" {
		return fmt.Sprintf("RPC error %v - %s: %s", err.Code, err.Message, err.Data)
	}
	return fmt.Sprintf("RPC error %v - %s", err.Code, err.Message)
}

func NewRPCRequest(method string, params json.RawMessage, id string) RPCRequest {
	return RPCRequest{
		Version: "2.0",
		Method:  method,
		Params:  params,
		ID:      id,
	}
}

type JSONRPCClient struct {
	host   string
	client *http.Client
}

func newHTTPDialer(host string) func(string, string) (net.Conn, error) {
	u, err := url.Parse(host)
	// default to tcp if nothing specified
	if err != nil || u == nil {
		return func(_ string, _ string) (net.Conn, error) {
			return nil, fmt.Errorf("Invalid host: %s", host)
		}
	}
	protocol := u.Scheme
	if protocol == "http" {
		protocol = "tcp"
	}
	return func(p, a string) (net.Conn, error) {
		return net.Dial(protocol, u.Host)
	}
}

func NewJSONRPCClient(host string) *JSONRPCClient {
	return NewJSONRPCClientShareTransport(host, &http.Transport{})
}

func NewJSONRPCClientShareTransport(host string, transport *http.Transport) *JSONRPCClient {
	transport.Dial = newHTTPDialer(host)
	rpcClient := &JSONRPCClient{
		host: host,
		client: &http.Client{
			Transport: transport,
		},
	}
	// Replace tcp:// with http:// so http.Client doesn't get confused
	parts := strings.SplitN(host, "://", 2)
	if len(parts) == 2 && parts[0] == "tcp" {
		rpcClient.host = "http://" + parts[1]
	}
	return rpcClient
}

func (c *JSONRPCClient) Call(method string, params map[string]interface{}, id string, result interface{}) error {
	paramsBytes, err := json.Marshal(params)
	if err != nil {
		return err
	}
	req := NewRPCRequest(method, paramsBytes, id)
	reqBytes, err := json.Marshal(req)
	if err != nil {
		return err
	}
	resp, err := c.client.Post(c.host, "text/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return fmt.Errorf("post status not OK: %s, response body: %s", resp.Status, string(respBytes))
	}
	if err != nil {
		return err
	}
	var rpcResp RPCResponse
	if err := json.Unmarshal(respBytes, &rpcResp); err != nil {
		return fmt.Errorf("error unmarshalling rpc response: %v", err)
	}
	if rpcResp.Error != nil {
		return fmt.Errorf("Response error: %v", rpcResp.Error)
	}
	if err := json.Unmarshal(rpcResp.Result, result); err != nil {
		return fmt.Errorf("error unmarshalling rpc response result: %v", err)
	}
	return nil
}
