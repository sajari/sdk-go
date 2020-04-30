# Sajari APIs Go client

[![GoDoc](https://godoc.org/code.sajari.com/sdk-go?status.svg)](http://godoc.org/code.sajari.com/sdk-go)
[![Build status](https://travis-ci.org/sajari/sdk-go.svg?branch=master)](https://travis-ci.org/sajari/sdk-go)
[![Report card](https://goreportcard.com/badge/code.sajari.com/sdk-go)](https://goreportcard.com/report/code.sajari.com/sdk-go)
[![Sourcegraph](https://sourcegraph.com/code.sajari.com/sdk-go/-/badge.svg)](https://sourcegraph.com/code.sajari.com/sdk-go?badge)

The official [Sajari](https://www.sajari.com) Go client library.

## Table of contents

- [Requirements](#requirements)
- [Installation](#installation)
- [Documentation](#documentation)
  - [Creating a client](#creating-a-client)
  - [Adding a record](#adding-a-record)
  - [Getting a record](#getting-a-record)
  - [Replacing a record](#replacing-a-record)
  - [Mutating a record](#mutating-a-record)
  - [Deleting a record](#deleting-a-record)
  - [Searching for records](#searching-for-records)
- [Development](#development)
- [Test](#test)
- [License](#license)

## Requirements

Requires [Go](https://golang.org/dl/) version 1.13 or higher.

## Installation

Install `sdk-go` with:

```shell
go get -u code.sajari.com/sdk-go
```

Then, import it using:

```go
import "code.sajari.com/sdk-go"
```

## Documentation

Below are a few simple examples that will help get you up and running.

### Creating a client

To start you need to create a client to make calls to the API.

You can get your project, collection, key ID and key secret from the [Sajari console](https://www.sajari.com/console).

```go
creds := sajari.KeyCredentials("key-id", "key-secret")
client, err := sajari.New("project", "collection", sajari.WithCredentials(creds))
if err != nil {
	// handle
}
defer client.Close()
```

> Note: do not forget to close the client when you are finished with it.

### Adding a record

A record can be added to a collection using the `CreateRecord` method on a record `Pipeline`.

First, you should initialise the record pipeline by passing in its name and the version you want to use.

```go
pipeline := client.Pipeline("record", "v5")
```

Next, set up any values that you need to pass to the record pipeline, define your record and call `CreateRecord`.

```go
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
```

You can use the returned key to uniquely identify the newly inserted record. This can be used in various calls such as `GetRecord`, `MutateRecord`, `DeleteRecord` and `ReplaceRecord`.

### Getting a record

An existing record in your collection can be retrieved using the `GetRecord` method on your `Client`.

```go
key := sajari.NewKey("id", "12345") // or using your Key returned from another call
record, err := client.GetRecord(context.Background(), key)
if err != nil {
	// handle
}
```

### Replacing a record

An existing record in your collection can be entirely replaced using the `ReplaceRecord` method on a record `Pipeline`.

When calling `ReplaceRecord` Sajari actually performs an upsert. If the record is an existing record, Sajari will perform a diff between the old and new records and applies your changes—this is extremely efficient. Because `ReplaceRecord` can both insert and update it is typically preferred over `CreateRecord` when [adding a record](#adding-a-record).

> Note: if you want to make granular changes to the record it is best to use [`MutateRecord`](#mutating-a-record).

> Note: if you want to change an indexed field you will need to use `ReplaceRecord`.

First, you should initialise the record pipeline by passing in its name and the version you want to use.

```go
pipeline := client.Pipeline("record", "v5")
```

Next, set up any values that you need to pass to the record pipeline, define your record and call `ReplaceRecord`.

```go
values := map[string]string{
	// ...
}

key := sajari.NewKey("id", "12345") // or using your Key returned from another call

record := sajari.Record{
	"id":    12345,
	"name":  "Large Smart TV",
	"brand": "Sunny",
	"price": 999,
}

key, _, err = pipeline.ReplaceRecord(context.Background(), values, key, record)
if err != nil {
	// handle
}
```

### Mutating a record

An existing record in your collection can be mutated using the `MutateRecord` method on your `Client`. You might need this method if you need to update a single field or unset a single field.

As an example, if you were storing products in your collection and you needed to update a product's price or stock levels this method would be useful.

> Note: if you want to replace the entire record it is best to use [`ReplaceRecord`](#replacing-a-record).

> Note: if you want to change an indexed field you will need to use `ReplaceRecord`.

You will need to pass one or more mutation operations to `MutateRecord` that will be appled to your record. For example, you can pass an operation to set a field, unset a field or set multiple fields at once.

```go
key := sajari.NewKey("id", "12345") // or using your Key returned from another call

// update a single field
err := client.MutateRecord(context.Background(), key, sajari.SetFieldValue("updated_at", time.Now().String()))
if err != nil {
	// handle
}

// unset a single field
err := client.MutateRecord(context.Background(), key, sajari.SetFieldValue("available", nil))
if err != nil {
	// handle
}

// set multiple fields at once
err := client.MutateRecord(context.Background(), key, sajari.SetFields(map[string]interface{}{
	"updated_at": time.Now().String(),
	"available":  nil,
})...)
if err != nil {
	// handle
}
```

### Deleting a record

An existing record in your collection can be deleted using the `DeleteRecord` method on your `Client`.

```go
key := sajari.NewKey("id", "12345") // or using your Key returned from another call
err := client.DeleteRecord(context.Background(), key)
if err != nil {
	// handle
}
```

### Searching for records

You can search for records in your collection using the `Search` method with a query `Pipeline`.

First, you should initialise the query pipeline by passing in its name and the version you want to use.

```go
pipeline := client.Pipeline("search", "v5")
```

Next, set up any values that you need to pass to the query pipeline, create a session and run your search.

```go
values := map[string]string{
	"q":              "your search terms",
	"resultsPerPage": "10",
}

res, _, err := pipeline.Search(context.Background(), values, sajari.NonTrackedSession())
if err != nil {
	// handle
}

for _, r := range res.Results {
	log.Printf("Values: %v", r.Values)
	log.Printf("Tokens: %v", r.Tokens)
}
```

## Development

Pull requests from the community are welcome. If you submit one, please keep the following guidelines in mind:

1. Code must be `go fmt` compliant.
2. All types, structs and funcs should be documented.
3. Ensure that `go test ./...` succeeds.

## Test

Run all tests:

```shell
go test ./...
```

## License

We use the [MIT License](./LICENSE.md).
