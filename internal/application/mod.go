package application

// Command contains all subcommands provided by this package.
type Command struct {
	Create createCommand `cmd:"create" aliases:"c" help:"Create a new application for a user"`
	Delete deleteCommand `cmd:"delete" aliases:"d" help:"Delete an existing application for a user"`
	Update updateCommand `cmd:"update" aliases:"u" help:"Update an existing application for a user"`
	List   listCommand   `cmd:"list" aliases:"l" help:"List all existing applications of the user"`
	Show   showCommand   `cmd:"show" aliases:"s" help:"Show details of an existing application of a user"`
}
