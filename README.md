# Traefik SSL Rewrite Plugin

## Overview
This is a Traefik middleware plugin that copies the value of the `X-Forwarded-Tls-Client-Cert` header to `SSL_CLIENT_CERT`. This allows backend services to access client certificate information forwarded by a proxy.

## Installation
To use this plugin, add it to your Traefik configuration as a middleware.

### Example Static Configuration (YAML)
```yaml
experimental:
  plugins:
    sslrewrite:
      moduleName: "github.com/jdhirst/sslrewrite"
      version: "v1.0.0"
```

### Example Dynamic Configuration (YAML)
```yaml
http:
  middlewares:
    ssl-rewrite:
      plugin:
        sslrewrite: {}
  routers:
    my-router:
      rule: "Host(`example.com`)"
      middlewares:
        - ssl-rewrite
      service: my-service
```

## Usage
When a request passes through this middleware, if the `X-Forwarded-Tls-Client-Cert` header is present, its value is copied to the `SSL_CLIENT_CERT` header.

## Testing
A test is provided to verify that the middleware correctly rewrites the header.

