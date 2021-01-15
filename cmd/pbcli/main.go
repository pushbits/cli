package main

import (
	"os"

	"github.com/jessevdk/go-flags"

	"github.com/pushbits/cli/internal/application"
	"github.com/pushbits/cli/internal/settings"
	"github.com/pushbits/cli/internal/ui"
	"github.com/pushbits/cli/internal/user"
)

type commands struct {
	settings.Settings
	Application application.Command `command:"application" alias:"a" description:"Configure applications"`
	User        user.Command        `command:"user" alias:"u" description:"Configure users"`
}

var (
	cmds   commands
	parser = flags.NewParser(&cmds, flags.Default)
)

func main() {
	_, err := parser.Parse()
	if err != nil {
		os.Exit(1)
	}

	s := settings.Settings{
		URL:      cmds.URL,
		Username: cmds.Username,
	}

	password := ui.GetPassword()

	settings.Runner.Run(s, password)
}
