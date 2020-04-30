package sajari_test

import (
	"context"
	"log"
	"time"

	sajari "code.sajari.com/sdk-go"
)

func ExampleNew() {
	creds := sajari.KeyCredentials("key-id", "key-secret")
	client, err := sajari.New("project", "collection", sajari.WithCredentials(creds))
	if err != nil {
		// handle
	}
	defer client.Close()
}

func ExampleClient_GetRecord() {
	creds := sajari.KeyCredentials("key-id", "key-secret")
	client, err := sajari.New("project", "collection", sajari.WithCredentials(creds))
	if err != nil {
		// handle
	}
	defer client.Close()

	key := sajari.NewKey("id", "12345") // or using your Key returned from another call

	record, err := client.GetRecord(context.Background(), key)
	if err != nil {
		// handle
	}
	_ = record // use record
}

func ExampleClient_MutateRecord() {
	creds := sajari.KeyCredentials("key-id", "key-secret")
	client, err := sajari.New("project", "collection", sajari.WithCredentials(creds))
	if err != nil {
		// handle
	}
	defer client.Close()

	key := sajari.NewKey("id", "12345") // or using your Key returned from another call

	// update a single field
	err = client.MutateRecord(context.Background(), key, sajari.SetFieldValue("updated_at", time.Now().String()))
	if err != nil {
		// handle
	}

	// unset a single field
	err = client.MutateRecord(context.Background(), key, sajari.SetFieldValue("available", nil))
	if err != nil {
		// handle
	}

	// set multiple fields at once
	err = client.MutateRecord(context.Background(), key, sajari.SetFields(map[string]interface{}{
		"updated_at": time.Now().String(),
		"available":  nil,
	})...)
	if err != nil {
		// handle
	}
}

func ExampleClient_DeleteRecord() {
	creds := sajari.KeyCredentials("key-id", "key-secret")
	client, err := sajari.New("project", "collection", sajari.WithCredentials(creds))
	if err != nil {
		// handle
	}
	defer client.Close()

	key := sajari.NewKey("id", "12345") // or using your Key returned from another call

	err = client.DeleteRecord(context.Background(), key)
	if err != nil {
		// handle
	}
}

func ExamplePipeline_CreateRecord() {
	creds := sajari.KeyCredentials("key-id", "key-secret")
	client, err := sajari.New("project", "collection", sajari.WithCredentials(creds))
	if err != nil {
		// handle
	}
	defer client.Close()

	pipeline := client.Pipeline("record", "v5")

	values := map[string]string{
		// ...
	}

	record := sajari.Record{
		"id":    12345,
		"name":  "Smart TV",
		"brand": "Sunny",
		"price": 999,
	}

	key, _, err := pipeline.CreateRecord(context.Background(), values, record)
	if err != nil {
		// handle
	}
	_ = key // use key
}

func ExamplePipeline_ReplaceRecord() {
	creds := sajari.KeyCredentials("key-id", "key-secret")
	client, err := sajari.New("project", "collection", sajari.WithCredentials(creds))
	if err != nil {
		// handle
	}
	defer client.Close()

	pipeline := client.Pipeline("record", "v5")

	values := map[string]string{
		// ...
	}

	key := sajari.NewKey("id", "12345") // or using your Key returned from another call

	record := sajari.Record{
		"id":    12345,
		"name":  "Smart TV",
		"brand": "Sunny",
		"price": 899,
	}

	key, _, err = pipeline.ReplaceRecord(context.Background(), values, key, record)
	if err != nil {
		// handle
	}
	_ = key // use key
}

func ExamplePipeline_Search() {
	creds := sajari.KeyCredentials("key-id", "key-secret")
	client, err := sajari.New("project", "collection", sajari.WithCredentials(creds))
	if err != nil {
		// handle
	}
	defer client.Close()

	pipeline := client.Pipeline("record", "v5")

	values := map[string]string{
		"q":              "your search terms",
		"resultsPerPage": "10",
		"page":           "1",
	}

	res, _, err := pipeline.Search(context.Background(), values, sajari.NonTrackedSession())
	if err != nil {
		// handle
	}

	for _, r := range res.Results {
		log.Printf("Values: %v", r.Values)
		log.Printf("Tokens: %v", r.Tokens)
	}
}
