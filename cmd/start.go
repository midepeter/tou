package cmd

import (
	"github.com/midepeter/tou/server"
	"github.com/urfave/cli/v2"
)

var StartCmd = &cli.Command{
	Name:  "start",
	Usage: "to start the workqueue manager on a server",
	Action: func(c *cli.Context) error {
		s := server.NewServer()
		return s.Serve("localhost:9000")
	},
}
