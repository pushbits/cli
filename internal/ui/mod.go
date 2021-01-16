package ui

import (
	"encoding/json"
	"fmt"
	"log"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

// GetPassword reads a password from the terminal.
func GetPassword(prompt string) string {
	fmt.Print(prompt)
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

// PrintJSON pretty-prints JSON to the terminal.
func PrintJSON(obj interface{}) {
	pretty, err := json.MarshalIndent(obj, "", "\t")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(pretty))
}
