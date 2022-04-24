package user

import (
	"log"

	"github.com/pushbits/cli/internal/api"
	"github.com/pushbits/cli/internal/settings"
	"github.com/pushbits/cli/internal/ui"
)

const (
	createEndpoint = "/user"
)

type createCommand struct {
	Arguments struct {
		Name     string `positional-arg-name:"name" description:"The name of the user"`
		MatrixID string `positional-arg-name:"matrixid" description:"The Matrix ID of the user"`
	} `required:"true" positional-args:"true"`
}

func (c *createCommand) Execute(args []string) error {
	settings.Runner = c
	return nil
}

func (c *createCommand) Run(s *settings.Settings, password string) {
	newPassword := ui.GetPassword("New password of user " + c.Arguments.Name + ": ")

	data := map[string]interface{}{
		"name":      c.Arguments.Name,
		"password":  newPassword,
		"matrix_id": c.Arguments.MatrixID,
	}

	resp, err := api.Post(s.URL, createEndpoint, s.Proxy, s.Username, password, data)
	if err != nil {
		log.Fatal(err)
	}

	ui.PrintJSON(resp)
}
