package sajari

import (
	"context"

	interactionv2pb "code.sajari.com/protogen-go/sajari/interaction/v2"
)

// Interaction creates a new Interaction which can be used to register
// interactions.
func (c *Client) Interaction() *Interaction {
	return &Interaction{
		c: c,
	}
}

// Interaction is used to register interactions.
type Interaction struct {
	c *Client
}

// InteractionOptions are passed with the token.
type InteractionOptions struct {
	Identifier string
	Weight     int32
	Data       map[string]string
}

// ConsumeToken registers an interaction corresponding to a token.
func (i *Interaction) ConsumeToken(ctx context.Context, token string, options InteractionOptions) error {
	_, err := interactionv2pb.NewInteractionClient(i.c.ClientConn).ConsumeToken(ctx, &interactionv2pb.ConsumeTokenRequest{
		Token:      token,
		Identifier: options.Identifier,
		Weight:     options.Weight,
		Data:       options.Data,
	})
	return err
}
