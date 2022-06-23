package cmd

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

var port string

var ConnectCmd = &cli.Command{
	Name:  "connect",
	Usage: "The connect commad is used for connecting to the running workqueue",
	Action: func(c *cli.Context) error {
		fmt.Fprintf(os.Stdout, "The port specified is %s\n", port)
		return nil
	},
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "port",
			Usage:       "to specify port on which work queue server instance is to run",
			Aliases:     []string{"p"},
			Required:    false,
			Destination: &port,
		},
	},
}
