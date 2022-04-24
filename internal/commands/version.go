package commands

import (
	"fmt"

	"github.com/pushbits/cli/internal/buildconfig"
	"github.com/pushbits/cli/internal/options"
)

type VersionCommand struct{}

func (c *VersionCommand) Run(s *options.Options) error {
	fmt.Printf("pbcli %s\n", buildconfig.Version)

	return nil
}
