//tou is a work queue manager
package main

import (
	"log"
	"os"

	"github.com/midepeter/tou/cmd/tou"
)

func main() {
	if err := tou.Run().Run(os.Args); err != nil {
		log.Println("Error ", err)
		os.Exit(1)
	}
}
