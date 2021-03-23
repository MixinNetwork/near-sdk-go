package rpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	MainnetRPCEndpoint = "https://rpc.mainnet.near.org"
	TestnetRPCEndpoint = "https://rpc.testnet.near.org"
)

type Client struct {
	endpoint   string
	httpclient *http.Client
}

func NewClient(endpoint string) *Client {
	return &Client{endpoint: endpoint, httpclient: new(http.Client)}
}

func (c *Client) request(method string, params interface{}, response interface{}) error {
	data, err := json.Marshal(map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      time.Now().UnixNano(),
		"method":  method,
		"params":  params,
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", c.endpoint, bytes.NewReader(data))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Close = true
	res, err := c.httpclient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if len(body) != 0 {
		if err := json.Unmarshal(body, &response); err != nil {
			return err
		}
	}
	if res.StatusCode < 200 || res.StatusCode > 300 {
		return fmt.Errorf("status code: %d", res.StatusCode)
	}
	return nil
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (err *ErrorResponse) Error() string {
	return fmt.Sprintf("RPC ERROR %d %s", err.Code, err.Message)
}

type GeneralResponse struct {
	JsonRPC string         `json:"jsonrpc"`
	ID      uint64         `json:"id"`
	Error   *ErrorResponse `json:"error"`
}
