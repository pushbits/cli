package application

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/pushbits/cli/internal/api"
	"github.com/pushbits/cli/internal/options"
	"github.com/pushbits/cli/internal/ui"
)

const (
	updateEndpoint = "/application/%d"
)

type updateCommand struct {
	options.AuthOptions
	ID                  uint   `arg:"" help:"The ID of the application"`
	NewName             string `long:"new-name" help:"The new name of the application" optional:""`
	RefreshToken        bool   `long:"refresh" help:"Refresh the token of the application"`
	StrictCompatibility bool   `long:"compat" help:"Enforce strict compatibility with Gotify"`
}

func (c *updateCommand) Run(s *options.Options) error {
	password := ui.GetCurrentPassword(c.Username)

	if !c.RefreshToken && c.StrictCompatibility {
		log.Fatal("Can only enforce compatibility when refreshing the token of the application")
	}

	data := map[string]interface{}{
		"refresh_token":        c.RefreshToken,
		"strict_compatibility": c.StrictCompatibility,
	}

	if len(c.NewName) > 0 {
		data["new_name"] = c.NewName
	}

	populatedEndpoint := fmt.Sprintf(updateEndpoint, c.ID)

	resp, err := api.Put(c.URL, populatedEndpoint, c.Proxy, c.Username, password, data)
	if err != nil {
		log.Fatal(err)
	}

	ui.PrintJSON(resp)

	return nil
}
