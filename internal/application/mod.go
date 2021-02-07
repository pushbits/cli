package application

// Command contains all subcommands provided by this package.
type Command struct {
	Create createCommand `command:"create" alias:"c" description:"Create a new application for a user"`
	Delete deleteCommand `command:"delete" alias:"d" description:"Delete an existing application for a user"`
	Update updateCommand `command:"update" alias:"u" description:"Update an existing application for a user"`
	List   listCommand   `command:"list" alias:"l" description:"List all existing applications of the user"`
	Show   showCommand   `command:"show" alias:"s" description:"Show details of an existing application of a user"`
}
