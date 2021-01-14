package application

import (
	"fmt"
	"log"

	"github.com/pushbits/cli/internal/api"
	"github.com/pushbits/cli/internal/settings"
)

const (
	createEndpoint = "/application"
)

type createCommand struct {
	Arguments struct {
		Name string `positional-arg-name:"name" description:"The name of the application"`
	} `required:"true" positional-args:"true"`
}

func (c *createCommand) Execute(args []string) error {
	settings.Runner = c
	return nil
}

func (c *createCommand) Run(s settings.Settings, password string) {
	data := map[string]interface{}{
		"name": c.Arguments.Name,
	}

	resp, err := api.Post(s.URL, createEndpoint, s.Username, password, data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
}
