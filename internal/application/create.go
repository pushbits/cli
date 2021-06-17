package application

import (
	"log"

	"github.com/pushbits/cli/internal/api"
	"github.com/pushbits/cli/internal/settings"
	"github.com/pushbits/cli/internal/ui"
)

const (
	createEndpoint = "application"
)

type createCommand struct {
	Arguments struct {
		Name string `positional-arg-name:"name" description:"The name of the application"`
	} `required:"true" positional-args:"true"`
	StrictCompatibility bool `long:"compat" description:"Enforce strict compatibility with Gotify"`
}

func (c *createCommand) Execute(args []string) error {
	settings.Runner = c
	return nil
}

func (c *createCommand) Run(s settings.Settings, password string) {
	data := map[string]interface{}{
		"name":                 c.Arguments.Name,
		"strict_compatibility": c.StrictCompatibility,
	}

	resp, err := api.Post(s.URL, createEndpoint, s.Username, password, data)
	if err != nil {
		log.Fatal(err)
	}

	ui.PrintJSON(resp)
}
