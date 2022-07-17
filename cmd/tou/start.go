package tou

import (
	"github.com/urfave/cli/v2"

	"github.com/midepeter/tou/internal/handlers"
	"github.com/midepeter/tou/internal/queue"
	"github.com/midepeter/tou/server"
)

var StartCmd = &cli.Command{
	Name:  "start",
	Usage: "to start the workqueue manager on a server",
	Action: func(c *cli.Context) error {
		newQueue := queue.NewQueue("first-queue")

		h := handlers.Handler{
			Queue: newQueue,
		}

		s := server.NewServer(h)
		return s.Serve("localhost:9000")
	},
}
