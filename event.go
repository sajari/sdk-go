package sajari

import (
	"context"
	"errors"
	"fmt"

	"google.golang.org/grpc/codes"

	"code.sajari.com/sdk-go/internal/openapi"
)

type SendEventRequest = openapi.SendEventRequest

var NewSendEventRequest = openapi.NewSendEventRequest

// SendEvent sends an event to the ranking system after a user interacts with a
// search result.
func (c *Client) SendEvent(ctx context.Context, r *SendEventRequest) error {
	if !c.v4 {
		return errors.New("not supported on non-v4 endpoints")
	}

	if r.GetName() == "" {
		return errors.New("name cannot be empty")
	}
	if r.GetToken() == "" {
		return errors.New("token cannot be empty")
	}

	ctx = context.WithValue(ctx, openapi.ContextBasicAuth, c.openAPI.auth)

	req := c.openAPI.client.EventsApi.SendEvent(ctx).SendEventRequest(*r)

	_, _, err := req.Execute()
	if err != nil {
		if ok, err := eventsHandleGenericOpenAPIError(err); ok {
			return err
		}
		return fmt.Errorf("could not execute request: %w", err)
	}

	return nil
}

func eventsHandleGenericOpenAPIError(err error) (handled bool, rerr error) {
	switch x := err.(type) {
	case openapi.GenericOpenAPIError:
		m := x.Model()

		if m, ok := m.(openapi.Error); ok {
			switch codes.Code(m.GetCode()) {
			default:
				return true, fmt.Errorf("%s: %w", m.GetMessage(), err)
			}
		}
	}
	return false, nil
}
