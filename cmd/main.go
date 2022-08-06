//tou is a work queue manager
package main

import (
	"fmt"
	"os"

	"github.com/midepeter/tou/cmd/tou"
	"github.com/midepeter/tou/logger"
)

func main() {
	log, err := logger.SetUpLogger()
	if err != nil {
		panic(fmt.Sprintf("Error trying to set up logger %v", err))
	}

	log.Info("Hello we are here it is working!!!")
	if err := tou.Run().Run(os.Args); err != nil {
		os.Exit(1)
	}
}
