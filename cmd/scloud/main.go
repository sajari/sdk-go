package main

import (
	"fmt"
	"os"
	"strings"

	"code.sajari.com/sdk-go/cmd/scloud/config"
	"code.sajari.com/sdk-go/cmd/scloud/interaction"
	"code.sajari.com/sdk-go/cmd/scloud/pipeline"
	"code.sajari.com/sdk-go/cmd/scloud/record"
	"code.sajari.com/sdk-go/cmd/scloud/schema"
)

var (
	topLevelCommands = []string{"init", "auth", "config", "schema", "record", "pipeline", "interaction"}
)

// TODO
// auth login
// querying
// autocomplete -> not in SDK
// csv

func exit(str string, args ...interface{}) {
	fmt.Printf(str+"\n\n", args...)
	os.Exit(1)
}

func main() {
	if len(os.Args) == 1 {
		exit("usage: scloud <%v> [options...]", strings.Join(topLevelCommands, "|"))
	}
	cmd, params := os.Args[1], os.Args[2:]

	// Load current config
	c, err := config.Load()
	if err != nil {
		exit("error loading profile config: %v", err)
	}

	// Special case we are initializing
	if cmd == "init" {
		c.Init(params)
		return
	}

	p, ok := c.Get(c.Active)
	if !ok {
		// The active profile is gone. Should not happen
		if cmd == "config" {
			c.Settings(params)
		}
		exit("No default profile set. Run `scloud init` to get started or use `scloud config set <profile>` to use an existing saved profile")
	}
	client, fn, err := p.Client()
	if err != nil {
		exit("%v", err)
	}
	defer fn()

	// Main functions
	switch cmd {
	case "config":
		c.Settings(params)

	case "auth":
		fmt.Printf("to be implemented") // TODO

	case "schema":
		schema.Run(client, params)

	case "record":
		record.Run(client, params)

	case "pipeline":
		pipeline.Run(client, params)

	case "interaction":
		interaction.Run(client, params)

	default:
		exit("invalid command %q\nusage: engctl <%v> [options...]", cmd, strings.Join(topLevelCommands, "|"))
	}
}
