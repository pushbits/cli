package main

import (
	"flag"
	"fmt"
	"log"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

type Credentials struct {
	Url string
	Username string
	Password string
}

func getCredentials(args []string) *Credentials {
	credFlags := flag.NewFlagSet("cred", flag.ExitOnError)
	url := credFlags.String("url", "", "The URL on which the server is listening")
	username := credFlags.String("username", "", "The username to authenticate on the server")

	credFlags.Parse(args)

	if *url == "" {
		log.Fatal("No URL was supplied.")
	}

	if *username == "" {
		log.Fatal("No username was supplied.")
	}

	fmt.Print("Password: ")
	password, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Fatal("No password was supplied.")
	}

	return &Credentials{
		Url: *url,
		Username: *username,
		Password: string(password),
	}
}
