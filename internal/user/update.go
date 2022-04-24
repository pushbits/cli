package user

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/pushbits/cli/internal/api"
	"github.com/pushbits/cli/internal/options"
	"github.com/pushbits/cli/internal/ui"
)

const (
	updateEndpoint = "/user/%d"
)

type updateCommand struct {
	options.AuthOptions
	ID          uint   `arg:"" help:"The ID of the user"`
	NewName     string `long:"new-name" help:"The new name of the user"`
	NewMatrixID string `long:"new-matrixid" help:"The new Matrix ID of the user"`
}

func (c *updateCommand) Run(s *options.Options) error {
	password := ui.GetCurrentPassword(c.Username)

	data := map[string]interface{}{}

	if len(c.NewName) > 0 {
		data["name"] = c.NewName
	}

	if len(c.NewMatrixID) > 0 {
		data["matrix_id"] = c.NewMatrixID
	}

	populatedEndpoint := fmt.Sprintf(updateEndpoint, c.ID)

	resp, err := api.Put(c.URL, populatedEndpoint, c.Proxy, c.Username, password, data)
	if err != nil {
		log.Fatal(err)
	}

	ui.PrintJSON(resp)

	return nil
}
