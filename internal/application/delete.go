package application

import (
	"fmt"
	"log"

	"github.com/pushbits/cli/internal/api"
	"github.com/pushbits/cli/internal/settings"
	"github.com/pushbits/cli/internal/ui"
)

const (
	deleteEndpoint = "/application/%d"
)

type deleteCommand struct {
	Arguments struct {
		ID uint `positional-arg-name:"id" description:"The ID of the application"`
	} `required:"true" positional-args:"true"`
}

func (c *deleteCommand) Execute(args []string) error {
	settings.Runner = c
	return nil
}

func (c *deleteCommand) Run(s *settings.Settings, password string) {
	populatedEndpoint := fmt.Sprintf(deleteEndpoint, c.Arguments.ID)

	resp, err := api.Delete(s.URL, populatedEndpoint, s.Proxy, s.Username, password)
	if err != nil {
		log.Fatal(err)
	}

	ui.PrintJSON(resp)
}
