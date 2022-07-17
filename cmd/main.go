//tou is a work queue manager
package main

import (
	"os"

	"github.com/midepeter/tou/cmd/tou"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.FatalLevel)

	if err := tou.Run().Run(os.Args); err != nil {
		log.Err(err).Msg("Exiting tou.....")
		os.Exit(1)
	}
}
