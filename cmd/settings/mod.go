package settings

type Settings struct {
	URL      string `long:"url" description:"The URL where the server listens for requests" required:"true"`
	Username string `long:"username" description:"The username for authenticating on the server" required:"true"`
}

type Runnable interface {
	Run(Settings)
}

var Runner Runnable
