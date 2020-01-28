package pipeline

import (
	"context"
	"encoding/json"
	"fmt"

	"google.golang.org/grpc"

	sajari "code.sajari.com/sdk-go"
)

// Query executes a query pipeline
func Query(client *sajari.Client, args []string, name, version string, inputs map[string]string) {

	ctx := context.Background()
	ctx = newContext(ctx, client)

	tracking := sajari.NewSession(sajari.TrackingNone, "", nil) // TODO: this is dumb. Fix the SDK

	resp, _, err := client.Pipeline(name, version).Search(ctx, inputs, tracking)
	if err != nil {
		fmt.Printf("Code: %v Message: %v", grpc.Code(err), grpc.ErrorDesc(err))
		return
	}

	for _, result := range resp.Results {
		b, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			fmt.Printf("could not write out result (%v): %v", result, err)
		}
		fmt.Println(string(b))
	}

	fmt.Println("Total Results", len(resp.Results))
	fmt.Println("Reads", resp.Reads)
	fmt.Println("Time", resp.Latency)

}
