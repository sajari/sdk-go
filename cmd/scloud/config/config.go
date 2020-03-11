package config

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	configPath = ".config/sajari"
	configName = "config"

	topLevelCommands = []string{"list", "get", "set", "delete"}
)

// config contains the credentials and onfig to use the scloud tool
type config struct {
	Active   string              `json:"active"`
	Profiles map[string]*Profile `json:"profiles"`
}

// New creates a new config profile
func New() *config {
	return &config{
		Profiles: make(map[string]*Profile),
	}
}

// Get returns the profile matching the name
func (c *config) Get(name string) (*Profile, bool) {
	p, ok := c.Profiles[name]
	return p, ok
}

// setDefault sets the default config profile
func (c *config) setDefault(name string) error {
	if _, ok := c.Profiles[name]; ok {
		c.Active = name
		return c.Save()
	}
	return fmt.Errorf("no profile named [%v]", name)
}

// delete removes a config profile by name
func (c *config) delete(name string) error {
	if _, ok := c.Profiles[name]; ok {
		scanner := bufio.NewScanner(os.Stdin)
		if confirmInput(scanner, "Delete profile. Are you sure? y/n") {
			delete(c.Profiles, name)
			if c.Active == name {
				c.Active = ""
			}
			return c.Save()
		}
		fmt.Println("config deletion aborted")
		return nil
	}
	return fmt.Errorf("no profile named [%v]", name)
}

// Load opens the existing config from local disk
func Load() (*config, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	path := filepath.Join(home, configPath, configName)
	f, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return New(), nil
		}
		return nil, err
	}
	defer f.Close()
	dec := json.NewDecoder(f)
	c := &config{}
	err = dec.Decode(c)
	return c, err
}

// Save writes the config to disk
func (c *config) Save() error {
	out, err := json.Marshal(c)
	if err != nil {
		return err
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	dir := filepath.Join(home, configPath)
	os.MkdirAll(dir, 0700)
	path := filepath.Join(dir, configName)
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(out)
	if err != nil {
		return err
	}
	return nil
}

func cmdInput(s *bufio.Scanner, q string) string {
	fmt.Println(q)
	s.Scan()
	return s.Text()
}

func confirmInput(s *bufio.Scanner, q string) bool {
	res := cmdInput(s, q)
	if res == "y" || res == "yes" {
		return true
	}
	return false
}

// Init creates a new config profile from cmd line input
func (c *config) Init(args []string) error {
	iflags := flag.NewFlagSet("init", flag.ExitOnError)
	endpoint := iflags.String("endpoint", "", "endpoint `address`")
	project := iflags.String("project", "", "project `id`")
	collection := iflags.String("collection", "", "collection `name`")
	creds := iflags.String("creds", "", "calling credentials `key-id,key-secret`")
	name := iflags.String("profile", "", "profile `name`")
	iflags.Parse(args)

	if len(args) == 1 {
		*name = args[0]
	}

	fmt.Println("scloud init")
	scanner := bufio.NewScanner(os.Stdin)

	p := &Profile{}

	if *project != "" {
		p.Project = *project
	} else {
		p.Project = cmdInput(scanner, "enter a project id:")
	}
	if *collection != "" {
		p.Collection = *collection
	} else {
		p.Collection = cmdInput(scanner, "enter a collection id (optional):")
	}
	if *endpoint != "" {
		p.Endpoint = *endpoint
	} else {
		p.Endpoint = cmdInput(scanner, "specify endpoint (optional):")
	}
	if *creds != "" {
		credsSplit := strings.Split(*creds, ",")
		if len(credsSplit) != 2 {
			return fmt.Errorf("creds: expected 'id,secret', got '%v'", *creds)
		}
		p.Key = credsSplit[0]
		p.Secret = credsSplit[1]
	} else {
		p.Key = cmdInput(scanner, "enter your key:")
		p.Secret = cmdInput(scanner, "enter your secret:")
	}
	if *name == "" {
		*name = cmdInput(scanner, "name this profile (optional):")
	}

	if _, ok := c.Profiles[*name]; ok {
		if confirmInput(scanner, fmt.Sprintf("overwrite existing %q profile? y/n:", *name)) {
			c.Profiles[*name] = p
		} else {
			*name = cmdInput(scanner, "name this profile (optional):")
			c.Profiles[*name] = p
		}
	} else {
		c.Profiles[*name] = p
	}

	// Add a flag to set this to the active profile
	c.Active = *name
	fmt.Println("profile has been set to default")

	if err := c.Save(); err != nil {
		return fmt.Errorf("could not save config: %v", err)
	}

	fmt.Printf("profile [%v] saved (and set to default)", *name)
	return nil
}

// Settings changes the active profile
func (c *config) Settings(args []string) error {
	var name string
	if len(args) < 1 {
		return fmt.Errorf("usage: scloud config <%v> [options...]\n", strings.Join(topLevelCommands, "|"))
	}
	if len(args) == 2 {
		name = args[1]
	}
	switch args[0] {
	case "list":
		for name, p := range c.Profiles {
			fmt.Printf("[%v]\n%v\n", name, p)
		}
		return nil

	case "get":
		if p, ok := c.Profiles[c.Active]; ok {
			fmt.Printf("[%v]\n%v\n", c.Active, p)
		}
		return nil

	case "set":
		err := c.setDefault(name)
		if err != nil {
			return err
		}
		fmt.Printf("default profile now [%v]\n", name)
		return nil

	case "delete":
		err := c.delete(name)
		if err != nil {
			return err
		}
		fmt.Printf("profile [%v] successfully deleted\n", name)
		return nil

	default:
		return fmt.Errorf("usage: scloud config <%v> [options...]\n", strings.Join(topLevelCommands, "|"))
	}

	return nil
}
