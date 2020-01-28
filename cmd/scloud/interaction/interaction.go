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

func Run(client *sajari.Client, args []string) {
	iflags := flag.NewFlagSet("interaction", flag.ExitOnError)
	token := iflags.String("token", "", "`token` to send back")
	weight := iflags.String("weight", "", "`weight` of the interaction")
	identifier := iflags.String("identifier", "", "`identifier` of the interaction, e.g. click, purchase")
	data := iflags.String("data", "", "`json` encoded map of keys to values")

	if len(args) == 0 {
		fmt.Printf("\nusage: scloud interaction <%v> [options...]\n\n", strings.Join(topLevelCommands, "|"))
		iflags.Usage()
		return
	}
	iflags.Parse(args[1:])

	opts := sajari.InteractionOptions{
		Identifier: *identifier,
	}

	if *data != "" {
		d := map[string]string{}
		if err := json.Unmarshal([]byte(*data), &d); err != nil {
			fmt.Printf("got error unmarshalling json from `data`: %v\n", err)
			return
		}
		opts.Data = d
	}

	if *weight != "" {
		w, err := strconv.Atoi(*weight)
		if err != nil {
			fmt.Printf("`weight must be an integer`")
			return
		}
		opts.Weight = int32(w)
	}

	if err := client.Interaction().ConsumeToken(context.Background(), *token, opts); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("interaction token successfully sent")
}
