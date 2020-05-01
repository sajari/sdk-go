package internal

import (
	"context"

	"google.golang.org/grpc/metadata"
)

const (
	projectKey    = "project"
	collectionKey = "collection"
)

// NewContext creates a new context with project and collection attached.
func NewContext(ctx context.Context, project, collection string) context.Context {
	m := map[string]string{
		projectKey:    project,
		collectionKey: collection,
	}
	return metadata.NewOutgoingContext(ctx, metadata.New(m))
}
