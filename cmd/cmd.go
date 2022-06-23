package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func Run() *cli.App {
	app := cli.NewApp()
	app.Name = "tou"
	app.Description = "tou is a message queuing system"
	app.Action = func(c *cli.Context) error {
		fmt.Println("Tou is a work queue manager library that can be used as a middelware to communicated between two application processes")
		return nil
	}

	app.Commands = []*cli.Command{
		VersionCmd,
		StartCmd,
		ConnectCmd,
	}

	return app
}
