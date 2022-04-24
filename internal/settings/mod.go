package settings

// Settings represents the global configuration for the program.
type Settings struct {
	URL      string `long:"url" description:"The URL where the server listens for requests" required:"true"`
	Username string `long:"username" description:"The username for authenticating on the server" required:"true"`
	Proxy    string `short:"x" long:"proxy" default:"" description:"The proxy to use for HTTP requests"`
}

// Runnable defines an interface for subcommands that take the global settings and a password.
type Runnable interface {
	Run(*Settings, string)
}

// Runner is the subcommand to run after all arguments were parsed.
var Runner Runnable
