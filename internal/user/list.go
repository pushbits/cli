package user

import (
	"log"

	"github.com/pushbits/cli/internal/settings"
)

type listCommand struct {
}

func (c *listCommand) Execute(args []string) error {
	settings.Runner = c
	return nil
}

func (c *listCommand) Run(s settings.Settings, password string) {
	log.Printf("listCommand")
}
