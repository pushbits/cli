package main

import (
	"os"

	"github.com/jessevdk/go-flags"

	"github.com/pushbits/cli/internal/application"
	"github.com/pushbits/cli/internal/commands"
	"github.com/pushbits/cli/internal/settings"
	"github.com/pushbits/cli/internal/ui"
	"github.com/pushbits/cli/internal/user"
)

type options struct {
	settings.Settings
	Application application.Command     `command:"application" alias:"a" description:"Configure applications"`
	User        user.Command            `command:"user" alias:"u" description:"Configure users"`
	Version     commands.VersionCommand `command:"version" alias:"v" description:"Print the program version"`
}

var (
	cmds   options
	parser = flags.NewParser(&cmds, flags.Default)
)

func main() {
	_, err := parser.Parse()
	if err != nil {
		os.Exit(1)
	}

	s := &settings.Settings{
		URL:      cmds.URL,
		Username: cmds.Username,
	}

	password := ui.GetPassword("Current password of user " + s.Username + ": ")

	settings.Runner.Run(s, password)
}
