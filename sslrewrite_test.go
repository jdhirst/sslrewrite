package sslrewrite_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	sslrewrite "github.com/jdhirst/sslrewrite"
)

func TestSslRewrite(t *testing.T) {
	cfg := sslrewrite.CreateConfig()

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})
	handler, err := sslrewrite.New(ctx, next, cfg, "ssl-rewrite-plugin")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Set initial header to test rewriting
	req.Header.Set("X-Forwarded-Tls-Client-Cert", "dummy-cert-value")

	handler.ServeHTTP(recorder, req)

	// Verify header was rewritten
	assertHeader(t, req, "SSL_CLIENT_CERT", "dummy-cert-value")
}

func assertHeader(t *testing.T, req *http.Request, key, expected string) {
	t.Helper()

	if req.Header.Get(key) != expected {
		t.Errorf("invalid header value for %s: expected %s, got %s", key, expected, req.Header.Get(key))
	}
}
