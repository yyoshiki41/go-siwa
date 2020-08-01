package siwa

import (
	"context"
	"net/http"
	"net/url"
	"path"
)

const (
	pathKeys = "/auth/keys"
)

type Key struct {
	Kty string `json:"kty"`
	Kid string `json:"kid"`
	Use string `json:"use"`
	Alg string `json:"alg"`
	N   string `json:"n"`
	E   string `json:"e"`
}

func (c *Client) Keys(ctx context.Context) ([]Key, error) {
	u, err := url.Parse(c.config.Endpoint)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, pathKeys)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet,
		u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	type KeysResponse struct {
		Keys []Key `json:"keys"`
	}

	var res KeysResponse
	err = c.do(req, &res)
	if err != nil {
		return nil, err
	}
	return res.Keys, nil
}
