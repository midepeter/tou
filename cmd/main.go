package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "tou"
	app.Description = "tou is a message queuing system"
	app.Action = func(c *cli.Context) error {
		log.Println("tou is running")
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal("The error ", err)
		os.Exit(1)
	}
}
