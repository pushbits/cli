package credentials

import (
	"fmt"
	"log"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

// GetPassword reads a password from the terminal.
func GetPassword() string {
	fmt.Print("Password: ")
	passwordBytes, err := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Print("\n")

	if err != nil {
		log.Fatal("No password was supplied.")
	}

	password := string(passwordBytes)

	if len(password) == 0 {
		log.Fatal("Password may not be empty.")
	}

	return password
}
