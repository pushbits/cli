package application

import (
	"fmt"
	"log"

	"github.com/pushbits/cli/internal/api"
	"github.com/pushbits/cli/internal/settings"
)

const (
	listEndpoint = "/application"
)

type listCommand struct {
}

func (c *listCommand) Execute(args []string) error {
	settings.Runner = c
	return nil
}

func (c *listCommand) Run(s settings.Settings, password string) {
	resp, err := api.Get(s.URL, listEndpoint, s.Username, password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
}
