package siwa

import (
	"encoding/json"
	"io"
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

	var r io.Reader = response.Body
	// r = io.TeeReader(r, os.Stderr)

	code := response.StatusCode
	if code >= http.StatusBadRequest {
		var e ErrorResponse
		json.NewDecoder(r).Decode(&e)
		return &e
	}

	if res == nil {
		return nil
	}
	return json.NewDecoder(r).Decode(&res)
}

type ErrorResponse struct {
	Err string `json:"error"`
}

func (e *ErrorResponse) Error() string {
	return e.Err
}
