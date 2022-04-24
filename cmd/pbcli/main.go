package main

import (
	"github.com/alecthomas/kong"

	"github.com/pushbits/cli/internal/application"
	"github.com/pushbits/cli/internal/commands"
	"github.com/pushbits/cli/internal/options"
	"github.com/pushbits/cli/internal/user"
)

var cmd struct {
	options.Options
	Application application.Command     `cmd:"application" aliases:"a" help:"Configure applications"`
	User        user.Command            `cmd:"user" aliases:"u" help:"Configure users"`
	Version     commands.VersionCommand `cmd:"version" aliases:"v" help:"Print the program version"`
}

func main() {
	ctx := kong.Parse(&cmd)
	err := ctx.Run(&cmd.Options)
	ctx.FatalIfErrorf(err)
}
