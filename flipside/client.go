package flipside

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

const (
	apiURL = "https://api-v2.flipsidecrypto.xyz/json-rpc"
)

type Client struct {
	apiKey string
	http   *http.Client
	log    *logrus.Entry
}

func NewClient(apiKey string, log *logrus.Logger) *Client {
	return &Client{
		apiKey: apiKey,
		http:   &http.Client{},
		log:    log.WithField("module", "flipside"),
	}
}

func (c *Client) sendRequest(method string, params interface{}) (*http.Response, error) {
	payload := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  method,
		"params":  []interface{}{params},
		"id":      1,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", c.apiKey)

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	responseBody, _ := io.ReadAll(resp.Body)
	resp.Body = io.NopCloser(bytes.NewBuffer(responseBody))

	if resp.StatusCode != http.StatusOK {
		fmt.Println(string(responseBody))
		return nil, fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	return resp, nil
}
