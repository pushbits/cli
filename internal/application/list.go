package application

import (
	"log"

	"github.com/pushbits/cli/internal/api"
	"github.com/pushbits/cli/internal/settings"
	"github.com/pushbits/cli/internal/ui"
)

const (
	listEndpoint = "/application"
)

type listCommand struct{}

func (c *listCommand) Execute(args []string) error {
	settings.Runner = c
	return nil
}

func (c *listCommand) Run(s *settings.Settings, password string) {
	resp, err := api.Get(s.URL, listEndpoint, s.Proxy, s.Username, password)
	if err != nil {
		log.Fatal(err)
	}

	ui.PrintJSON(resp)
}
