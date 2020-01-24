package main

import (
	"fmt"
	"os"
	"strings"

	"code.sajari.com/sdk-go/cmd/scloud/config"
	"code.sajari.com/sdk-go/cmd/scloud/pipeline"
	"code.sajari.com/sdk-go/cmd/scloud/record"
	"code.sajari.com/sdk-go/cmd/scloud/schema"
)

var (
	topLevelCommands = []string{"init", "config", "schema", "record", "pipeline"}
)

// TODO
// init cmd
// config set
// auth login
// run actual cmd line tools

func exit(str string, args ...interface{}) {
	fmt.Printf(str+"\n\n", args...)
	os.Exit(1)
}

func main() {
	// Try to open config and use the current active profile
	// if it doesn't exist, bail and suggest init or auth

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

	case "schema":
		schema.Run(client, params)

	case "record":
		record.Run(client, params)

	case "pipeline":
		pipeline.Run(client, params)

	default:
		exit("invalid command %q\nusage: engctl <%v> [options...]", cmd, strings.Join(topLevelCommands, "|"))
	}
}
