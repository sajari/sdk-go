package main

import (
	"context"
	"flag"
	"log"
	"strings"

	sajari "code.sajari.com/sdk-go"
)

var (
	endpoint   = flag.String("endpoint", "", "endpoint `address`, uses default if not set")
	project    = flag.String("project", "", "project `name` to query")
	collection = flag.String("collection", "", "collection `name` to query")
	creds      = flag.String("creds", "", "calling credentials `key-id,key-secret`")

	token = flag.String("token", "", "`token` to submit")
)

func main() {
	flag.Parse()

	if *token == "" {
		log.Fatalf("-token is not set")
	}

	var opts []sajari.Opt
	if *endpoint != "" {
		opts = append(opts, sajari.WithEndpoint(*endpoint))
	}

	if *creds != "" {
		credsSplit := strings.Split(*creds, ",")
		if len(credsSplit) != 2 {
			log.Printf("creds: expected 'id,secret', got '%v'", *creds)
			return
		}
		kc := sajari.KeyCredentials(credsSplit[0], credsSplit[1])
		opts = append(opts, sajari.WithCredentials(kc))
	}

	if *project == "" {
		log.Fatal("-project not set")
	}

	if *collection == "" {
		log.Fatal("-collection not set")
	}

	client, err := sajari.New(*project, *collection, opts...)
	if err != nil {
		log.Printf("error from sajari.New(): %v", err)
		return
	}
	defer func() {
		if err := client.Close(); err != nil {
			log.Printf("error closing Client: %v", err)
		}
	}()

	if err := client.Interaction().ConsumeToken(context.Background(), *token, sajari.InteractionOptions{}); err != nil {
		log.Println(err)
	}
}
