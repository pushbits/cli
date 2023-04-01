// Package user provides commands related to managing users.
package user

// Command contains all subcommands provided by this package.
type Command struct {
	Create createCommand `cmd:"create" aliases:"c" help:"Create a new user"`
	Delete deleteCommand `cmd:"delete" aliases:"d" help:"Delete an existing user"`
	Update updateCommand `cmd:"update" aliases:"u" help:"Update an existing user"`
	List   listCommand   `cmd:"list" aliases:"l" help:"List all existing users"`
	Show   showCommand   `cmd:"show" aliases:"s" help:"Show details of an existing user"`
}
