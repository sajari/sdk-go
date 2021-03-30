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

// A Collection stores the records that can be searched.
type Collection = openapi.Collection

// GetCollection gets a collection identified by the provided ID.
//
// If there is no such collection matching the given ID this method returns an
// error wrapping ErrNoSuchCollection.
func (c *Client) GetCollection(ctx context.Context, id string) (*Collection, error) {
	if !c.v4 {
		return nil, errors.New("not supported on non-v4 endpoints")
	}

	if id == "" {
		return nil, errors.New("collection id cannot be empty")
	}

	ctx = context.WithValue(ctx, openapi.ContextBasicAuth, c.openAPI.auth)

	req := c.openAPI.client.CollectionsApi.GetCollection(ctx, id)

	collection, _, err := req.Execute()
	if err != nil {
		if ok, err := collectionsHandleGenericOpenAPIError(err); ok {
			return nil, err
		}
		return nil, fmt.Errorf("could not get collection: %w", err)
	}

	return &collection, nil
}

// UpdateCollectionOpt is a type which defines options to update a collection.
type UpdateCollectionOpt func(c *openapi.Collection, updateMask map[string]struct{})

// SetCollectionDisplayName is a collection mutation that set a collection's
// display name.
func SetCollectionDisplayName(displayName string) UpdateCollectionOpt {
	return func(c *openapi.Collection, updateMask map[string]struct{}) {
		c.DisplayName = displayName
		updateMask["display_name"] = struct{}{}
	}
}

// SetAuthorizedQueryDomains is a collection mutation that set a collection's
// authorized query domains.
func SetAuthorizedQueryDomains(domains []string) UpdateCollectionOpt {
	return func(c *openapi.Collection, updateMask map[string]struct{}) {
		c.AuthorizedQueryDomains = &domains
		updateMask["authorized_query_domains"] = struct{}{}
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

	if id == "" {
		return errors.New("collection id cannot be empty")
	}

	col := &openapi.Collection{}
	updateMask := map[string]struct{}{}

	for _, opt := range opts {
		opt(col, updateMask)
	}

	um := make([]string, 0, len(updateMask))
	for f := range updateMask {
		um = append(um, f)
	}

	ctx = context.WithValue(ctx, openapi.ContextBasicAuth, c.openAPI.auth)

	req := c.openAPI.client.CollectionsApi.
		UpdateCollection(ctx, id).
		Collection(*col).
		UpdateMask(strings.Join(um, ","))

	_, _, err := req.Execute()
	if err != nil {
		if ok, err := collectionsHandleGenericOpenAPIError(err); ok {
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
		if ok, err := collectionsHandleGenericOpenAPIError(err); ok {
			return err
		}
		return fmt.Errorf("could not delete collection: %w", err)
	}

	return nil
}

// GetDefaultPipeline gets the default pipeline for a collection.
func (c *Client) GetDefaultPipeline(ctx context.Context, id string, typ PipelineType) (string, error) {
	if !c.v4 {
		return "", errors.New("not supported on non-v4 endpoints")
	}

	if id == "" {
		return "", errors.New("collection id cannot be empty")
	}
	if typ == "" {
		return "", errors.New("type cannot be empty")
	}

	ctx = context.WithValue(ctx, openapi.ContextBasicAuth, c.openAPI.auth)

	req := c.openAPI.client.PipelinesApi.
		GetDefaultPipeline(ctx, id).
		Type_(string(typ))

	resp, _, err := req.Execute()
	if err != nil {
		if ok, err := collectionsHandleGenericOpenAPIError(err); ok {
			return "", err
		}
		return "", fmt.Errorf("could not get default pipeline: %w", err)
	}

	return resp.GetPipeline(), nil
}

// collectionsHandleGenericOpenAPIError handles generic OpenAPI errors in the
// context of collections. E.g. 404 is converted to ErrNoSuchCollection.
func collectionsHandleGenericOpenAPIError(err error) (handled bool, rerr error) {
	switch x := err.(type) {
	case openapi.GenericOpenAPIError:
		m := x.Model()

		if m, ok := m.(openapi.Error); ok {
			switch codes.Code(m.GetCode()) {
			case codes.NotFound:
				return true, fmt.Errorf("%w", ErrNoSuchCollection)
			default:
				return true, fmt.Errorf("%s: %w", m.GetMessage(), err)
			}
		}
	}
	return false, nil
}
