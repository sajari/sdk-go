package main

import (
	"fmt"
	"os"
	"strings"

	"code.sajari.com/sdk-go/cmd/scloud/config"
	csvcmd "code.sajari.com/sdk-go/cmd/scloud/csv"
	"code.sajari.com/sdk-go/cmd/scloud/interaction"
	"code.sajari.com/sdk-go/cmd/scloud/pipeline"
	"code.sajari.com/sdk-go/cmd/scloud/record"
	"code.sajari.com/sdk-go/cmd/scloud/schema"
)

var (
	version = "0.1"

	topLevelCommands = []string{"version", "init", "auth", "config", "schema", "record", "pipeline", "interaction"}
)

// TODO
// auth login
// autocomplete -> not in SDK
// csv

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("usage: scloud <%v> [options...]", strings.Join(topLevelCommands, "|"))
		return
	}
	cmd, params := os.Args[1], os.Args[2:]

	c, err := config.Load()
	if err != nil {
		fmt.Printf("error loading profile config: %v", err)
		return
	}

	if cmd == "init" {
		if err := c.Init(params); err != nil {
			fmt.Printf("%v", err)
		}
		return
	}

	p, ok := c.Get(c.Active)
	if !ok {
		if cmd == "config" {
			if err := c.Settings(params); err != nil {
				fmt.Printf("%v", err)
				return
			}
			return
		}
		fmt.Printf("No default profile set. Run `scloud init` to get started or use `scloud config set <profile>` to use an existing saved profile")
		return
	}
	client, fn, err := p.Client()
	if err != nil {
		fmt.Printf("%v", err)
	}
	defer fn()

	switch cmd {
	case "version":
		fmt.Printf("version: %v\n", version)

	case "config":
		if err := c.Settings(params); err != nil {
			fmt.Printf("%v", err)
		}

	case "auth":
		fmt.Printf("to be implemented") // TODO

	case "schema":
		if err := schema.Run(client, params); err != nil {
			fmt.Printf("%v", err)
		}

	case "record":
		if err := record.Run(client, params); err != nil {
			fmt.Printf("%v", err)
		}

	case "pipeline":
		if err := pipeline.Run(client, params); err != nil {
			fmt.Printf("%v", err)
		}

	case "interaction":
		if err := interaction.Run(client, params); err != nil {
			fmt.Printf("%v", err)
		}

	case "csv":
		if err := csvcmd.Run(client, params); err != nil {
			fmt.Printf("%v", err)
		}

	default:
		fmt.Printf("usage: scloud <%v> [options...]", strings.Join(topLevelCommands, "|"))
	}
}
