package main

import (
	"os"

	"github.com/alecthomas/kong"
	log "github.com/sirupsen/logrus"

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

func init() {
	log.SetOutput(os.Stderr)
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{
		DisableTimestamp: true,
	})
}

func main() {
	ctx := kong.Parse(&cmd)

	if cmd.Verbose {
		log.SetLevel(log.DebugLevel)
	}

	err := ctx.Run(&cmd.Options)
	ctx.FatalIfErrorf(err)
}
