package pipeline

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"sort"
	"text/tabwriter"

	"google.golang.org/grpc"

	sajari "code.sajari.com/sdk-go"
)

// Query executes a query pipeline
func Query(client *sajari.Client, args []string) error {

	iflags := flag.NewFlagSet("query", flag.ExitOnError)
	name := iflags.String("name", "", "pipeline `name` to run")
	version := iflags.String("version", "", "pipeline `version` to run (optional, blank will use the default version)")
	inputs := iflags.String("inputs", "", "pipeline `inputs` object as a JSON string")
	format := iflags.String("format", "table", "output format. Supports `json`, `table` (default)") // csv

	if len(args) == 0 {
		defer iflags.Usage()
		return fmt.Errorf("\nusage: scloud pipeline query [options...]\n\n")
	}
	iflags.Parse(args)

	in := map[string]string{}
	if err := json.Unmarshal([]byte(*inputs), &in); err != nil {
		return fmt.Errorf("error unmarshalling JSON inputs: %v", err)
	}

	if *name == "" {
		return fmt.Errorf("-name is blank, you must specify a pipeline `name`")
	}
	if len(in) == 0 {
		return fmt.Errorf("-inputs is blank, expecting a JSON object. e.g. {'q':'yogi'}")
	}

	ctx := context.Background()
	ctx = newContext(ctx, client)

	tracking := sajari.NewSession(sajari.TrackingNone, "", nil) // TODO: this is dumb. Fix the SDK

	resp, outputs, err := client.Pipeline(*name, *version).Search(ctx, in, tracking)
	if err != nil {
		return fmt.Errorf("Code: %v Message: %v", grpc.Code(err), grpc.ErrorDesc(err))
	}

	switch *format {

	case "json":
		for _, result := range resp.Results {
			b, err := json.MarshalIndent(result, "", "  ")
			if err != nil {
				return fmt.Errorf("could not write out result (%v): %v", result, err)
			}
			fmt.Println(string(b))
		}
	case "table":
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.Debug)
		fmt.Fprintf(w, "\n")
		var headers []string
		for i, result := range resp.Results {
			if i == 0 {
				for k := range result.Values {
					headers = append(headers, k)
				}
				sort.Strings(headers)
				fmt.Fprintf(w, "index score\tscore\t")
				for _, k := range headers {
					fmt.Fprintf(w, " %v\t", k)
				}
				fmt.Fprintf(w, "\n")
			}
			fmt.Fprintf(w, "%.4f\t %.4f\t", result.IndexScore, result.Score)
			for _, k := range headers {
				fmt.Fprintf(w, " %v\t", result.Values[k])
			}
			fmt.Fprintf(w, "\n")
		}
		fmt.Fprintf(w, "\n")
		w.Flush()
	}

	if len(outputs) > 0 {
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.Debug)
		fmt.Fprintf(w, "Outputs:\nkey\t value\n")
		for k, v := range outputs {
			fmt.Fprintf(w, "%v\t %v\n", k, v)
		}
		w.Flush()
	}

	fmt.Printf("\nTotal Results: %v", len(resp.Results))
	fmt.Printf("\nReads: %v", resp.Reads)
	fmt.Printf("\nTime: %v", resp.Latency)

	return nil
}
