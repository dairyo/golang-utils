package jwthandler

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
)

type Config struct {
	Issuer             string
	CacheTTL           time.Duration
	SignatureAlgorithm validator.SignatureAlgorithm
	Audiences          []string
	ValidatorOptions   []validator.Option
	MiddlewareOptions  []jwtmiddleware.Option
}

func Factory(c *Config) (func(next http.Handler) http.Handler, error) {
	issuerURL, err := url.Parse(fmt.Sprintf("https://%s/", c.Issuer))
	if err != nil {
		return nil, err
	}
	provider := jwks.NewCachingProvider(issuerURL, c.CacheTTL)
	jwtValidator, err := validator.New(
		provider.KeyFunc,
		c.SignatureAlgorithm,
		issuerURL.String(),
		c.Audiences,
		c.ValidatorOptions...,
	)
	if err != nil {
		return nil, err
	}
	middleware := jwtmiddleware.New(
		jwtValidator.ValidateToken,
		c.MiddlewareOptions...,
	)
	return func(next http.Handler) http.Handler {
		return middleware.CheckJWT(next)
	}, nil
}
