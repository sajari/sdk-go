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

// Run executes several different record request options
func Run(client *sajari.Client, args []string) error {
	iflags := flag.NewFlagSet("record", flag.ExitOnError)
	field := iflags.String("field", "", "`field` to count unique keys")
	value := iflags.String("value", "", "`value` of a record field")
	data := iflags.String("data", "", "`json` encoded map of keys to values")

	if len(args) == 0 {
		defer iflags.Usage()
		return fmt.Errorf("\nusage: scloud record <%v> [options...]\n\n", strings.Join(topLevelCommands, "|"))
	}
	iflags.Parse(args[1:])

	switch args[0] {
	case "get":
		if *field == "" || *value == "" {
			return fmt.Errorf("mutate must specify both `field` and `value` flags")
		}

		k := sajari.NewKey(*field, *value)
		d, err := client.GetRecord(context.Background(), k)
		if err != nil {
			return fmt.Errorf("error from Get(%v): %v\n", k, err)
		}

		b, err := json.MarshalIndent(d, "", "  ")
		if err != nil {
			return fmt.Errorf("error marshaling JSON output: %v\n", err)
		}

		fmt.Println(string(b))
		return nil

	case "mutate":
		if *data == "" {
			return fmt.Errorf("no data found, supply json input using the `data` flag")
		}
		if *field == "" || *value == "" {
			return fmt.Errorf("mutate must specify both `field` and `value` flags")
		}
		d := map[string]interface{}{}
		if err := json.Unmarshal([]byte(*data), &d); err != nil {
			return fmt.Errorf("got error unmarshalling json from `data`: %v\n", err)
		}

		ctx := context.Background()
		k := sajari.NewKey(*field, *value)
		if err := client.MutateRecord(ctx, k, sajari.SetFields(d)...); err != nil {
			return fmt.Errorf("error mutating record: %v\n", err)
		}

	case "delete":
		if *field == "" || *value == "" {
			return fmt.Errorf("delete must specify both `field` and `value` flags")
		}

		k := sajari.NewKey(*field, *value)
		if err := client.DeleteRecord(context.Background(), k); err != nil {
			return fmt.Errorf("error from Delete(%v): %v\n", k, err)
		}

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
				return fmt.Errorf("Could not get key: %v", err)
			}
			total++
		}
		return fmt.Errorf("Total: %d keys", total)

	default:
		return fmt.Errorf("usage: scloud schema <%v> [options...]\n", strings.Join(topLevelCommands, "|"))
	}
	return nil
}
