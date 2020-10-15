package sajari

import (
	"google.golang.org/grpc"

	"code.sajari.com/sdk-go/internal/openapi"
)

// Opt is a type which defines Client options.
type Opt func(c *Client)

// WithV4Endpoint configures the client to use a v4 endpoint.
func WithV4Endpoint(endpoint string) Opt {
	return func(c *Client) {
		c.v4 = true

		c.openAPI.config.Servers = openapi.ServerConfigurations{
			{
				URL: endpoint,
			},
		}

		c.openAPI.client = openapi.NewAPIClient(c.openAPI.config)
	}
}

// WithEndpoint configures the client to use a custom endpoint.
func WithEndpoint(endpoint string) Opt {
	return func(c *Client) {
		c.endpoint = endpoint
	}
}

// WithCredentials sets the client credentials used in each request.
func WithCredentials(cr Credentials) Opt {
	return func(c *Client) {
		if kc, ok := cr.(keyCreds); ok {
			c.openAPI.auth = openapi.BasicAuth{
				UserName: kc.keyID,
				Password: kc.keySecret,
			}
		}
		opt := WithGRPCDialOption(grpc.WithPerRPCCredentials(creds{cr}))
		opt(c)
	}
}

// WithGRPCDialOption returns an Opt which appends a new grpc.DialOption
// to an underlying gRPC dial.
func WithGRPCDialOption(opt grpc.DialOption) Opt {
	return func(c *Client) {
		c.dialOpts = append(c.dialOpts, opt)
	}
}
