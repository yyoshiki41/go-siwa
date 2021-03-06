package siwa

import (
	"github.com/dgrijalva/jwt-go"
)

const (
	AudienceApple = "https://appleid.apple.com"
)

func NewJWTHeader(kid string) map[string]interface{} {
	return map[string]interface{}{
		"kid": kid,
		"alg": "ES256",
	}
}

func NewJWTPayload(iss, sub string, iat, exp int64) jwt.Claims {
	return jwt.StandardClaims{
		Audience:  AudienceApple,
		Issuer:    iss,
		Subject:   sub,
		IssuedAt:  iat,
		ExpiresAt: exp,
	}
}

func NewJWTString(
	header map[string]interface{}, claims jwt.Claims, key interface{},
) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	token.Header = header
	return token.SignedString(key)
}
