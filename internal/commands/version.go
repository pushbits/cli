package commands

import (
	"fmt"

	"github.com/pushbits/cli/internal/buildconfig"
	"github.com/pushbits/cli/internal/settings"
)

type VersionCommand struct{}

func (c *VersionCommand) Execute(args []string) error {
	settings.Runner = c
	return nil
}

func (c *VersionCommand) Run(s *settings.Settings, password string) {
	fmt.Printf("pbcli %s\n", buildconfig.Version)
}
