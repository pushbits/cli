package user

import (
	"log"

	"github.com/pushbits/cli/internal/settings"
)

type showCommand struct {
	Arguments struct {
		ID uint `positional-arg-name:"id" description:"The ID of the user"`
	} `required:"true" positional-args:"true"`
}

func (c *showCommand) Execute(args []string) error {
	settings.Runner = c
	return nil
}

func (c *showCommand) Run(s settings.Settings, password string) {
	log.Printf("showCommand")
}
