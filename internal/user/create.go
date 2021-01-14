package user

import (
	"fmt"
	"log"

	"github.com/pushbits/cli/internal/api"
	"github.com/pushbits/cli/internal/settings"
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

func (c *createCommand) Run(s settings.Settings, password string) {
	data := map[string]interface{}{
		"name":      c.Arguments.Name,
		"matrix_id": c.Arguments.MatrixID,
	}

	resp, err := api.Post(s.URL, createEndpoint, s.Username, password, data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
}
