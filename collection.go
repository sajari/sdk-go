package sajari

import (
	"context"
	"errors"
	"fmt"

	"google.golang.org/grpc/codes"

	"code.sajari.com/sdk-go/internal/openapi"
)

// ErrNoSuchCollection is returned when a collection was requested but there is
// no such collection.
var ErrNoSuchCollection = errors.New("no such collection")

// DeleteCollection removes a collection identified by the provided ID.
//
// If there is no such collection matching the given ID this method returns an
// error wrapping ErrNoSuchCollection.
func (c *Client) DeleteCollection(ctx context.Context, id string) error {
	if !c.v4 {
		return errors.New("not supported on non-v4 endpoints")
	}

	ctx = context.WithValue(ctx, openapi.ContextBasicAuth, c.openAPI.auth)

	_, _, err := c.openAPI.client.CollectionsApi.DeleteCollection(ctx, id).Execute()
	if err != nil {
		switch x := err.(type) {
		case openapi.GenericOpenAPIError:
			m := x.Model()

			if m, ok := m.(openapi.GatewayruntimeError1); ok {
				switch codes.Code(m.GetCode()) {
				case codes.NotFound:
					return fmt.Errorf("%v: %w", id, ErrNoSuchCollection)
				}
			}
		}
		return fmt.Errorf("could not delete collection: %w", err)
	}

	return nil
}
