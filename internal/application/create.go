package application

import (
	"log"

	"github.com/pushbits/cli/internal/api"
	"github.com/pushbits/cli/internal/options"
	"github.com/pushbits/cli/internal/ui"
)

const (
	createEndpoint = "/application"
)

type createCommand struct {
	options.AuthOptions
	Name                string `arg:"name" help:"The name of the application"`
	StrictCompatibility bool   `long:"compat" help:"Enforce strict compatibility with Gotify"`
}

func (c *createCommand) Run(s *options.Options) error {
	password := ui.GetCurrentPassword(c.Username)

	data := map[string]interface{}{
		"name":                 c.Name,
		"strict_compatibility": c.StrictCompatibility,
	}

	resp, err := api.Post(c.URL, createEndpoint, c.Proxy, c.Username, password, data)
	if err != nil {
		log.Fatal(err)
	}

	ui.PrintJSON(resp)

	return nil
}
