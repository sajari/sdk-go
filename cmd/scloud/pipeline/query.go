package pipeline

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"

	"google.golang.org/grpc"

	sajari "code.sajari.com/sdk-go"
)

// Query executes a query pipeline
func Query(client *sajari.Client, args []string) error {

	iflags := flag.NewFlagSet("query", flag.ExitOnError)
	name := iflags.String("name", "", "pipeline `name` to run")
	version := iflags.String("version", "", "pipeline `version` to run (optional, blank will use the default version)")
	inputs := iflags.String("inputs", "", "pipeline `inputs` object as a JSON string")

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

	resp, _, err := client.Pipeline(*name, *version).Search(ctx, in, tracking)
	if err != nil {
		return fmt.Errorf("Code: %v Message: %v", grpc.Code(err), grpc.ErrorDesc(err))
	}

	for _, result := range resp.Results {
		b, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return fmt.Errorf("could not write out result (%v): %v", result, err)
		}
		fmt.Println(string(b))
	}

	fmt.Println("Total Results", len(resp.Results))
	fmt.Println("Reads", resp.Reads)
	fmt.Println("Time", resp.Latency)

	return nil
}
