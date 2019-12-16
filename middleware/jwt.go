package middleware

import (
	"crypto/rsa"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/totoro081295/daily-report-api/token"
)

type jwtMiddleware struct {
	key *rsa.PublicKey
}

// NewJWTMiddleware mount jwt middleware
func NewJWTMiddleware(tokenHandler token.Handler) (JWTMiddleware, error) {
	key, err := tokenHandler.LoadJWTPublicKeys()
	if err != nil {
		return nil, err
	}
	return &jwtMiddleware{key: key}, nil
}

// JWTMiddleware jwt middleware interface
type JWTMiddleware interface {
	JWT() echo.MiddlewareFunc
}

func (j *jwtMiddleware) JWT() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:    j.key,
		SigningMethod: "RS256",
	})
}
