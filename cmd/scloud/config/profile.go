package config

import (
	"fmt"
	"log"

	sajari "code.sajari.com/sdk-go"
)

type Profile struct {
	Endpoint   string `json:"endpoint"`
	Project    string `json:"project"`
	Collection string `json:"collection"`
	Key        string `json:"key"`
	Secret     string `json:"secret"`
	Token      string `json:"token"`
}

func (p *Profile) String() string {
	return fmt.Sprintf("project = %v\ncollection = %v\nendpoint = %v\n", p.Project, p.Collection, p.Endpoint)
}

func (p *Profile) Client() (*sajari.Client, func(), error) {
	var opts []sajari.Opt
	if p.Endpoint != "" {
		opts = append(opts, sajari.WithEndpoint(p.Endpoint))
	}

	if p.Token != "" {
		// TODO
	} else if p.Key != "" && p.Secret != "" {
		kc := sajari.KeyCredentials(p.Key, p.Secret)
		opts = append(opts, sajari.WithCredentials(kc))
	}

	if p.Project == "" {
		return nil, func() {}, fmt.Errorf("project not set")
	}

	if p.Collection == "" {
		return nil, func() {}, fmt.Errorf("collection not set")
	}

	client, err := sajari.New(p.Project, p.Collection, opts...)
	if err != nil {
		return nil, func() {}, err
	}
	fn := func() {
		if err := client.Close(); err != nil {
			log.Printf("error closing Client: %v", err)
		}
	}
	return client, fn, nil
}
