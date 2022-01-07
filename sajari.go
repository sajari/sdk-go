// Package sajari provides functionality for interacting with Search.io APIs.
package sajari // import "code.sajari.com/sdk-go"

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"code.sajari.com/sdk-go/internal"
	"code.sajari.com/sdk-go/internal/openapi"
)

const (
	endpoint         = "api-us-valkyrie.sajari.com:50051"
	userAgent        = "sdk-go-11032021"
	overrideHostname = "api.sajari.com"
)

// New creates a new Client which can be used to make requests to Search.io
// services.
func New(accountID, collectionID string, opts ...Opt) (*Client, error) {
	c := &Client{
		Project:    accountID,
		Collection: collectionID,
	}

	c.openAPI.config = openapi.NewConfiguration()
	c.openAPI.config.UserAgent = userAgent
	c.openAPI.config.DefaultHeader = map[string]string{
		"Account-Id": accountID,
	}

	defaultOpts := []Opt{
		WithEndpoint(endpoint),
		WithGRPCDialOption(grpc.WithUserAgent(userAgent)),
		WithGRPCDialOption(grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, overrideHostname))),
	}

	opts = append(defaultOpts, opts...)
	for _, opt := range opts {
		opt(c)
	}

	if c.ClientConn == nil {
		conn, err := grpc.Dial(c.endpoint, c.dialOpts...)
		if err != nil {
			return nil, err
		}
		c.ClientConn = conn
	}
	return c, nil
}

func (c *Client) newContext(ctx context.Context) context.Context {
	return internal.NewContext(ctx, c.Project, c.Collection)
}

// Client is a type which makes requests to Search.io.
type Client struct {
	Project    string
	Collection string
	endpoint   string

	ClientConn *grpc.ClientConn
	dialOpts   []grpc.DialOption

	v4      bool
	openAPI struct {
		config *openapi.Configuration
		client *openapi.APIClient
		auth   openapi.BasicAuth
	}
}

// Close releases all resources held by the Client.
func (c *Client) Close() error {
	return c.ClientConn.Close()
}
