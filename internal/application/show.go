package application

import (
	"fmt"
	"log"

	"github.com/pushbits/cli/internal/api"
	"github.com/pushbits/cli/internal/settings"
)

const (
	showEndpoint = "/application/%d"
)

type showCommand struct {
	Arguments struct {
		ID uint `positional-arg-name:"id" description:"The ID of the application"`
	} `required:"true" positional-args:"true"`
}

func (c *showCommand) Execute(args []string) error {
	settings.Runner = c
	return nil
}

func (c *showCommand) Run(s settings.Settings, password string) {
	populatedEndpoint := fmt.Sprintf(showEndpoint, c.Arguments.ID)

	resp, err := api.Get(s.URL, populatedEndpoint, s.Username, password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
}
