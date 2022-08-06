package tou

import (
	"github.com/urfave/cli/v2"

	"github.com/midepeter/tou/internal/handlers"
	"github.com/midepeter/tou/internal/queue"
	"github.com/midepeter/tou/logger"
	"github.com/midepeter/tou/server"
)

var StartCmd = &cli.Command{
	Name:  "start",
	Usage: "to start the workqueue manager on a server",
	Action: func(c *cli.Context) error {
		newQueue := queue.NewQueue("first-queue")

		logger, err := logger.SetUpLogger()
		if err != nil {
			panic(err)
		}

		h := handlers.Handler{
			Queue: newQueue,
			Log:   logger,
		}

		s := server.NewServer(h, logger)
		return s.Serve("localhost:9000")
	},
}
