package record

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"strings"

	sajari "code.sajari.com/sdk-go"
)

var (
	topLevelCommands = []string{"get", "mutate", "delete", "count"}
)

func Run(client *sajari.Client, args []string) {
	if len(args) == 0 {
		fmt.Printf("usage: scloud record <%v> [options...]\n", strings.Join(topLevelCommands, "|"))
		return
	}

	iflags := flag.NewFlagSet("record", flag.ExitOnError)
	field := iflags.String("field", "", "`field` to count unique keys")
	value := iflags.String("value", "", "`value` of a record field")
	data := iflags.String("data", "", "`json` encoded map of keys to values")
	iflags.Parse(args[1:])

	switch args[0] {
	case "get":
		if *field == "" || *value == "" {
			fmt.Printf("mutate must specify both `field` and `value` flags")
			return
		}

		k := sajari.NewKey(*field, *value)
		d, err := client.GetRecord(context.Background(), k)
		if err != nil {
			fmt.Printf("error from Get(%v): %v\n", k, err)
			return
		}

		b, err := json.MarshalIndent(d, "", "  ")
		if err != nil {
			fmt.Printf("error marshaling JSON output: %v\n", err)
			return
		}

		fmt.Println(string(b))
		return

	case "mutate":
		if *data == "" {
			fmt.Println("no data found, supply json input using the `data` flag")
			return
		}
		if *field == "" || *value == "" {
			fmt.Printf("mutate must specify both `field` and `value` flags")
			return
		}
		d := map[string]interface{}{}
		if err := json.Unmarshal([]byte(*data), &d); err != nil {
			fmt.Printf("got error unmarshalling json from `data`: %v\n", err)
			return
		}

		ctx := context.Background()
		k := sajari.NewKey(*field, *value)
		if err := client.MutateRecord(ctx, k, sajari.SetFields(d)...); err != nil {
			fmt.Printf("error mutating record: %v\n", err)
			return
		}
		return

	case "delete":
		if *field == "" || *value == "" {
			fmt.Printf("delete must specify both `field` and `value` flags")
			return
		}

		k := sajari.NewKey(*field, *value)
		if err := client.DeleteRecord(context.Background(), k); err != nil {
			fmt.Printf("error from Delete(%v): %v\n", k, err)
		}
		return

	case "count":
		limit := 1000
		it := client.Keys(context.Background(), *field, limit)
		var total int
		for {
			_, err := it.Next()
			if err == sajari.ErrDone {
				break
			}
			if err != nil {
				fmt.Printf("Could not get key: %v", err)
				return
			}
			total++
		}
		fmt.Printf("Total: %d keys", total)
		return

	default:
		fmt.Printf("usage: scloud schema <%v> [options...]\n", strings.Join(topLevelCommands, "|"))
	}
}
