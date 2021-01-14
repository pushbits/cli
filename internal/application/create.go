package application

import (
	"log"

	"github.com/pushbits/cli/internal/settings"
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
	log.Printf("createCommand")
}
