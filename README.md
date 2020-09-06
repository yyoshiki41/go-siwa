# go-siwa

[![godoc](https://godoc.org/github.com/yyoshiki41/go-siwa?status.svg)](https://pkg.go.dev/github.com/yyoshiki41/go-siwa?tab=doc)

Go library for SIWA (Sign In With Apple)

## [Sign in with Apple REST API](https://developer.apple.com/documentation/sign_in_with_apple/sign_in_with_apple_rest_api)

This client supports Sign in with Apple REST API to generate and validate the identity tokens.

## Quick Start

1. Generate sigend JWT from your private key
2. Generate and Validate Tokens
   - For authorization code validation, use `TokenGrantTypeAuthorizationCode`.
   - For refresh token validation requests, use `TokenGrantTypeRefreshToken`.

```go
package main

import (
	"context"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/yyoshiki41/go-siwa"
)

func main() {
	kid := "kid"
	teamID := "teamID"
	bundleID := "bundleID"
	privateKey := `-----BEGIN PRIVATE KEY-----
		YourPrivateKey
		-----END PRIVATE KEY-----`

	// Decode your private key
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		log.Fatal(errors.New("block is nil"))
	}
	pKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		log.Fatal(err)
	}

	// 1. Generate sigend JWT from your private key
	now := time.Now()
	secret, err := siwa.NewJWTString(
		siwa.NewJWTHeader(kid),
		siwa.NewJWTPayload(teamID, bundleID,
			now.Unix(), now.Add(90*24*time.Hour).Unix()),
		pKey,
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", secret)

	// 2. Generate and Validate Tokens
	authorizationCode := "authorization.code"
	redirectURI := ""
	client := siwa.NewClient()
	token, err := client.TokenGrantTypeAuthorizationCode(
		context.Background(), bundleID, secret, authorizationCode, redirectURI)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", token)
}
```

## Author

Yoshiki Nakagawa
