package siwa

import (
	"context"
	"net/http"
	"net/url"
	"path"
	"strings"
)

// https://developer.apple.com/documentation/sign_in_with_apple/generate_and_validate_tokens

const (
	pathAuthToken  = "/auth/token"
	pathAuthRevoke = "/auth/revoke"
)

type GrantType string

const (
	GrantTypeAuthorizationCode GrantType = "authorization_code"
	GrantTypeRefreshToken      GrantType = "refresh_token"
)

type TokenTypeHint string

const (
	TokenTypeHintAccessToken  TokenTypeHint = "access_token"
	TokenTypeHintRefreshToken TokenTypeHint = "refresh_token"
)

type Token struct {
	// (Reserved for future use) A token used to access allowed data. Currently, no data set has been defined for access.
	AccessToken string `json:"access_token"`
	// The amount of time, in seconds, before the access token expires.
	ExpiresIn int `json:"expires_in"`
	// A JSON Web Token that contains the user’s identity information.
	IDToken string `json:"id_token"`
	// The refresh token used to regenerate new access tokens. Store this token securely on your server.
	RefreshToken string `json:"refresh_token"`
	// The type of access token. It will always be bearer
	TokenType string `json:"token_type"`
}

func (c *Client) TokenGrantTypeAuthorizationCode(
	ctx context.Context, clientID, clientSecret, code, redirectURI string,
) (*Token, error) {
	u, err := url.Parse(c.config.Endpoint)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, pathAuthToken)
	v := formValues(GrantTypeAuthorizationCode, clientID, clientSecret, code, redirectURI, "")

	req, err := http.NewRequestWithContext(ctx, http.MethodPost,
		u.String(), strings.NewReader(v.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	var result Token
	err = c.do(req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) TokenGrantTypeRefreshToken(
	ctx context.Context, clientID, clientSecret, refreshToken string,
) (*Token, error) {
	u, err := url.Parse(c.config.Endpoint)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, pathAuthToken)
	v := formValues(GrantTypeRefreshToken, clientID, clientSecret, "", "", refreshToken)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost,
		u.String(), strings.NewReader(v.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	var result Token
	err = c.do(req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func formValues(grantType GrantType, clientID, clientSecret, code, redirectURI, refreshToken string) url.Values {
	v := url.Values{}
	v.Set("client_id", clientID)
	v.Set("client_secret", clientSecret)
	v.Set("grant_type", string(grantType))

	switch grantType {
	case GrantTypeAuthorizationCode:
		v.Set("code", code)
		v.Set("redirect_uri", redirectURI)
	case GrantTypeRefreshToken:
		v.Set("refresh_token", refreshToken)
	}
	return v
}

func (c *Client) RevokeToken(
	ctx context.Context, clientID, clientSecret, token string, tokenTypeHint TokenTypeHint,
) error {
	u, err := url.Parse(c.config.Endpoint)
	if err != nil {
		return err
	}
	u.Path = path.Join(u.Path, pathAuthRevoke)
	v := url.Values{}
	v.Set("client_id", clientID)
	v.Set("client_secret", clientSecret)
	v.Set("token", token)
	v.Set("token_type_hint", string(tokenTypeHint))

	req, err := http.NewRequestWithContext(ctx, http.MethodPost,
		u.String(), strings.NewReader(v.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if err := c.do(req, nil); err != nil {
		return err
	}
	return nil
}
