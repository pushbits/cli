package user

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/pushbits/cli/internal/api"
	"github.com/pushbits/cli/internal/options"
	"github.com/pushbits/cli/internal/ui"
)

const (
	deleteEndpoint = "/user/%d"
)

type deleteCommand struct {
	options.AuthOptions
	ID uint `arg:"" help:"The ID of the user"`
}

func (c *deleteCommand) Run(_ *options.Options) error {
	password := ui.GetCurrentPassword(c.Username)

	populatedEndpoint := fmt.Sprintf(deleteEndpoint, c.ID)

	resp, err := api.Delete(c.URL, populatedEndpoint, c.Proxy, c.Username, password)
	if err != nil {
		log.Fatal(err)
	}

	ui.PrintJSON(resp)

	return nil
}
