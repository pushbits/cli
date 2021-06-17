package application

import (
	"fmt"
	"log"

	"github.com/pushbits/cli/internal/api"
	"github.com/pushbits/cli/internal/settings"
	"github.com/pushbits/cli/internal/ui"
)

const (
	updateEndpoint = "application/%d"
)

type updateCommand struct {
	Arguments struct {
		ID uint `required:"true" positional-arg-name:"id" description:"The ID of the application"`
	} `positional-args:"true"`
	NewName             *string `long:"new-name" description:"The new name of the application"`
	RefreshToken        bool    `long:"refresh" description:"Refresh the token of the application"`
	StrictCompatibility bool    `long:"compat" description:"Enforce strict compatibility with Gotify"`
}

func (c *updateCommand) Execute(args []string) error {
	settings.Runner = c
	return nil
}

func (c *updateCommand) Run(s settings.Settings, password string) {
	if !c.RefreshToken && c.StrictCompatibility {
		log.Fatal("Can only enforce compatibility when refreshing the token of the application")
	}

	data := map[string]interface{}{
		"refresh_token":        c.RefreshToken,
		"strict_compatibility": c.StrictCompatibility,
	}

	if c.NewName != nil {
		data["new_name"] = c.NewName
	}

	populatedEndpoint := fmt.Sprintf(updateEndpoint, c.Arguments.ID)

	resp, err := api.Put(s.URL, populatedEndpoint, s.Username, password, data)
	if err != nil {
		log.Fatal(err)
	}

	ui.PrintJSON(resp)
}
