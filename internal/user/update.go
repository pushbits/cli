package user

import (
	"fmt"
	"log"

	"github.com/pushbits/cli/internal/api"
	"github.com/pushbits/cli/internal/settings"
	"github.com/pushbits/cli/internal/ui"
)

const (
	updateEndpoint = "/user/%d"
)

type updateCommand struct {
	Arguments struct {
		ID uint `positional-arg-name:"id" description:"The ID of the user"`
	} `required:"true" positional-args:"true"`
	NewName     *string `long:"new-name" description:"The new name of the user"`
	NewMatrixID *string `long:"new-matrixid" description:"The new Matrix ID of the user"`
}

func (c *updateCommand) Execute(args []string) error {
	settings.Runner = c
	return nil
}

func (c *updateCommand) Run(s *settings.Settings, password string) {
	data := map[string]interface{}{}

	if c.NewName != nil {
		data["name"] = c.NewName
	}

	if c.NewMatrixID != nil {
		data["matrix_id"] = c.NewMatrixID
	}

	populatedEndpoint := fmt.Sprintf(updateEndpoint, c.Arguments.ID)

	resp, err := api.Put(s.URL, populatedEndpoint, s.Username, password, data)
	if err != nil {
		log.Fatal(err)
	}

	ui.PrintJSON(resp)
}
