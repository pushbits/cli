package main

import (
	"os"

	"github.com/jessevdk/go-flags"

	"github.com/pushbits/cli/cmd/application"
	"github.com/pushbits/cli/cmd/settings"
	"github.com/pushbits/cli/cmd/user"
)

type Commands struct {
	settings.Settings
	Application application.Command `command:"application" alias:"a" description:"Configure applications"`
	User        user.Command        `command:"user" alias:"u" description:"Configure users"`
}

var (
	commands Commands
	parser   = flags.NewParser(&commands, flags.Default)
)

func main() {
	_, err := parser.Parse()
	if err != nil {
		os.Exit(1)
	}

	s := settings.Settings{
		URL:      commands.URL,
		Username: commands.Username,
	}

	settings.Runner.Run(s)
}
