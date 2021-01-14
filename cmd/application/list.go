package application

import (
	"log"

	"github.com/pushbits/cli/cmd/settings"
)

type listCommand struct {
}

func (c *listCommand) Execute(args []string) error {
	settings.Runner = c
	return nil
}

func (c *listCommand) Run(s settings.Settings) {
	log.Printf("listCommand")
}
