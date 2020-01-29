package interaction

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"strconv"
	"strings"

	sajari "code.sajari.com/sdk-go"
)

var (
	topLevelCommands = []string{"consume"}
)

// Run sends an interaction token to the API
func Run(client *sajari.Client, args []string) error {
	iflags := flag.NewFlagSet("interaction", flag.ExitOnError)
	token := iflags.String("token", "", "`token` to send back")
	weight := iflags.String("weight", "", "`weight` of the interaction")
	identifier := iflags.String("identifier", "", "`identifier` of the interaction, e.g. click, purchase")
	data := iflags.String("data", "", "`json` encoded map of keys to values")

	if len(args) == 0 {
		defer iflags.Usage()
		return fmt.Errorf("\nusage: scloud interaction <%v> [options...]\n\n", strings.Join(topLevelCommands, "|"))
	}
	iflags.Parse(args[1:])

	opts := sajari.InteractionOptions{
		Identifier: *identifier,
	}

	if *data != "" {
		d := map[string]string{}
		if err := json.Unmarshal([]byte(*data), &d); err != nil {
			return fmt.Errorf("got error unmarshalling json from `data`: %v\n", err)
		}
		opts.Data = d
	}

	if *weight != "" {
		w, err := strconv.Atoi(*weight)
		if err != nil {
			return fmt.Errorf("`weight must be an integer`")
		}
		opts.Weight = int32(w)
	}

	if err := client.Interaction().ConsumeToken(context.Background(), *token, opts); err != nil {
		return err
	}

	fmt.Printf("interaction token successfully sent")
	return nil
}
