package pipeline

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"strings"

	"google.golang.org/grpc"

	sajari "code.sajari.com/sdk-go"
)

// Replace executes a record replace pipeline
func Replace(client *sajari.Client, args []string) error {
	iflags := flag.NewFlagSet("replace", flag.ExitOnError)
	name := iflags.String("name", "", "pipeline `name` to run")
	version := iflags.String("version", "", "pipeline `version` to run (optional, if blank the default version will run)")
	inputs := iflags.String("inputs", "", "pipeline `inputs` object as a JSON string")
	record := iflags.String("record", "", "`record` object as a JSON string")
	field := iflags.String("field", "", "`field` to lookup a unique key")
	value := iflags.String("value", "", "`value` of the unique key `field`")

	if len(args) == 0 {
		defer iflags.Usage()
		return fmt.Errorf("\nusage: scloud pipeline replace [options...]\n\n", strings.Join(topLevelCommands, "|"))
	}
	iflags.Parse(args)

	in := map[string]string{}
	if err := json.Unmarshal([]byte(*inputs), &in); err != nil {
		return fmt.Errorf("error unmarshalling JSON inputs: %v", err)
	}

	var rec sajari.Record
	if err := json.Unmarshal([]byte(*record), &rec); err != nil {
		return fmt.Errorf("error unmarshalling JSON record: %v", err)
	}

	if *name == "" {
		return fmt.Errorf("-name is blank, you must specify a pipeline `name`")
	}
	if len(in) == 0 {
		return fmt.Errorf("-inputs is blank, expecting a JSON object. e.g. {'q':'yogi'}")
	}
	if len(rec) == 0 {
		return fmt.Errorf("-record is blank, expecting a record as a JSON object. e.g. {'name':'yogi', 'location': 'Jellystone'}")
	}
	if *field == "" || *value == "" {
		return fmt.Errorf("`field` and `value` must both be specified to create a valid record key")
	}

	key := sajari.NewKey(*field, *value)

	ctx := context.Background()
	ctx = newContext(ctx, client)

	_, _, err := client.Pipeline(*name, *version).ReplaceRecord(ctx, in, key, rec)
	if err != nil {
		return fmt.Errorf("Code: %v Message: %v", grpc.Code(err), grpc.ErrorDesc(err))
	}

	fmt.Print("%v successfullly replaced", key)
	return nil
}
