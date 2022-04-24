package options

// Options represents the global options.
type Options struct {
	Verbose bool `short:"v" help:"Show debugging information"`
}

// AuthOptions represents the options for authentication against a server.
type AuthOptions struct {
	URL      string `name:"url" help:"The URL where the server listens for requests" required:""`
	Username string `name:"username" help:"The username for authenticating on the server" required:""`
	Proxy    string `name:"proxy" default:"" help:"The proxy to use for HTTP requests"`
}
