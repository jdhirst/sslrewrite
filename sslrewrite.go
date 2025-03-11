package sslrewrite

import (
	"context"
	"net/http"
)

// Config represents the plugin configuration.
type Config struct{}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{}
}

// SslRewrite is the Traefik plugin.
type SslRewrite struct {
	next http.Handler
	name string
}

// New creates a new instance of the SslRewrite plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &SslRewrite{
		next: next,
		name: name,
	}, nil
}

// ServeHTTP processes the request and modifies headers.
func (a *SslRewrite) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// Get the value of X-Forwarded-Tls-Client-Cert header
	if cert := req.Header.Get("X-Forwarded-Tls-Client-Cert"); cert != "" {
		req.Header.Set("SSL_CLIENT_CERT", cert)
	}

	// Call the next handler in the chain
	a.next.ServeHTTP(rw, req)
}
