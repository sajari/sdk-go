package internal

import (
	"context"

	"google.golang.org/grpc/metadata"
)

const (
	projectKey    = "project"
	collectionKey = "collection"
)

// NewContext creates a new context with account ID and collection ID attached.
func NewContext(ctx context.Context, accountID, collectionID string) context.Context {
	m := map[string]string{
		projectKey:    accountID,
		collectionKey: collectionID,
	}
	return metadata.NewOutgoingContext(ctx, metadata.New(m))
}
