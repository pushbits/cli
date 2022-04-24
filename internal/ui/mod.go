package ui

import (
	"encoding/json"
	"fmt"
	"syscall"

	log "github.com/sirupsen/logrus"

	"golang.org/x/term"
)

func getPassword(prompt string) string {
	fmt.Print(prompt)

	passwordBytes, err := term.ReadPassword(int(syscall.Stdin))
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

// PrintJSON pretty-prints JSON to the term.
func PrintJSON(obj interface{}) {
	pretty, err := json.MarshalIndent(obj, "", "\t")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(pretty))
}

// GetCurrentPassword reads the current password for `username` from the terminal.
func GetCurrentPassword(username string) string {
	return getPassword("Current password of user " + username + ": ")
}

// GetNewPassword reads the new password for `username` from the terminal.
func GetNewPassword(username string) string {
	return getPassword("New password of user " + username + ": ")
}
