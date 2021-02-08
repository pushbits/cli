package user

// Command contains all subcommands provided by this package.
type Command struct {
	Create createCommand `command:"create" alias:"c" description:"Create a new user"`
	Delete deleteCommand `command:"delete" alias:"d" description:"Delete an existing user"`
	Update updateCommand `command:"update" alias:"u" description:"Update an existing user"`
	List   listCommand   `command:"list" alias:"l" description:"List all existing users"`
	Show   showCommand   `command:"show" alias:"s" description:"Show details of an existing user"`
}
