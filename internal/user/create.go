package user

import (
	log "github.com/sirupsen/logrus"

	"github.com/pushbits/cli/internal/api"
	"github.com/pushbits/cli/internal/options"
	"github.com/pushbits/cli/internal/ui"
)

const (
	createEndpoint = "/user"
)

type createCommand struct {
	options.AuthOptions
	Name     string `arg:"" help:"The name of the user"`
	MatrixID string `arg:"" help:"The Matrix ID of the user"`
}

func (c *createCommand) Run(s *options.Options) error {
	password := ui.GetCurrentPassword(c.Username)

	newPassword := ui.GetNewPassword(c.Name)

	data := map[string]interface{}{
		"name":      c.Name,
		"password":  newPassword,
		"matrix_id": c.MatrixID,
	}

	resp, err := api.Post(c.URL, createEndpoint, c.Proxy, c.Username, password, data)
	if err != nil {
		log.Fatal(err)
	}

	ui.PrintJSON(resp)

	return nil
}
