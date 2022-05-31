package cmd

import (
	"log"

	"github.com/urfave/cli"
)

func Run()  *cli.App {
	app := cli.NewApp()
	app.Name = "tou"
	app.Description = "tou is a message queuing system"
	app.Action = func(c *cli.Context) error {
		log.Println("tou is running")
		return nil
	}

	return app
}
