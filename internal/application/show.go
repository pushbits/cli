package application

import (
	"fmt"
	"log"

	"github.com/pushbits/cli/internal/api"
	"github.com/pushbits/cli/internal/options"
	"github.com/pushbits/cli/internal/ui"
)

const (
	showEndpoint = "/application/%d"
)

type showCommand struct {
	options.AuthOptions
	ID uint `arg:"" help:"The ID of the application"`
}

func (c *showCommand) Run(s *options.Options) error {
	password := ui.GetCurrentPassword(c.Username)

	populatedEndpoint := fmt.Sprintf(showEndpoint, c.ID)

	resp, err := api.Get(c.URL, populatedEndpoint, c.Proxy, c.Username, password)
	if err != nil {
		log.Fatal(err)
	}

	ui.PrintJSON(resp)

	return nil
}
