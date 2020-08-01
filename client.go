package siwa

import (
	"encoding/json"
	"net/http"
)

const (
	Endpoint = "https://appleid.apple.com/"
)

var (
	httpClient = &http.Client{}
)

// Client represents an API client for Apple Auth API.
type Client struct {
	httpClient *http.Client
	config     *Config
}

// Config is a setting for Apple API.
type Config struct {
	Endpoint string
}

// SetHTTPClient overrides the default HTTP client.
func SetHTTPClient(client *http.Client) {
	httpClient = client
}

func NewClient() *Client {
	if httpClient == nil {
		SetHTTPClient(&http.Client{})
	}
	return &Client{
		httpClient: httpClient,
		config: &Config{
			Endpoint: Endpoint,
		},
	}
}

func (c *Client) do(req *http.Request, res interface{}) error {
	response, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if res == nil {
		return nil
	}
	return json.NewDecoder(response.Body).Decode(&res)
}
