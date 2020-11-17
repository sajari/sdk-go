package sajari

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"

	"code.sajari.com/sdk-go/internal/openapi"
)

// ErrNoSuchCollection is returned when a collection was requested but there is
// no such collection.
var ErrNoSuchCollection = errors.New("no such collection")

// UpdateCollectionOpt is a type which defines options to update a collection.
type UpdateCollectionOpt func(c *openapi.V4beta1Collection, updateMask map[string]struct{})

// SetCollectionDisplayName is a collection mutation that set a collection's
// display name.
func SetCollectionDisplayName(displayName string) UpdateCollectionOpt {
	return func(c *openapi.V4beta1Collection, updateMask map[string]struct{}) {
		c.DisplayName = displayName
		updateMask["display_name"] = struct{}{}
	}
}

// UpdateCollection updates a collection identified by the provided ID.
//
// If there is no such collection matching the given ID this method returns an
// error wrapping ErrNoSuchCollection.
func (c *Client) UpdateCollection(ctx context.Context, id string, opts ...UpdateCollectionOpt) error {
	if !c.v4 {
		return errors.New("not supported on non-v4 endpoints")
	}

	col := &openapi.V4beta1Collection{}
	updateMask := map[string]struct{}{}

	for _, opt := range opts {
		opt(col, updateMask)
	}

	um := make([]string, 0, len(updateMask))
	for f := range updateMask {
		um = append(um, f)
	}

	ctx = context.WithValue(ctx, openapi.ContextBasicAuth, c.openAPI.auth)

	req := c.openAPI.client.CollectionsApi.UpdateCollection(ctx, id)
	req.V4beta1Collection(*col)
	req.UpdateMask(strings.Join(um, ","))

	_, _, err := req.Execute()
	if err != nil {
		ok, err := handleGenericOpenAPIError(err)
		if ok {
			return err
		}
		return fmt.Errorf("could not update collection: %w", err)
	}

	return nil
}

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
		ok, err := handleGenericOpenAPIError(err)
		if ok {
			return err
		}
		return fmt.Errorf("could not delete collection: %w", err)
	}

	return nil
}

func handleGenericOpenAPIError(err error) (handled bool, rerr error) {
	switch x := err.(type) {
	case openapi.GenericOpenAPIError:
		m := x.Model()

		if m, ok := m.(openapi.GatewayruntimeError); ok {
			switch codes.Code(m.GetCode()) {
			case codes.NotFound:
				return true, fmt.Errorf("%w", ErrNoSuchCollection)
			}
		}
	}
	return false, nil
}
