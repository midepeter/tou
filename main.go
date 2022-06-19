//tou is a work queue manager
package main

import (
	"log"
	"os"
	"github.com/midepeter/tou/cmd"
)

func main() {
	if err := cmd.Run().Run(os.Args); err != nil {
		log.Println("Error ", err)
		os.Exit(1)
	}
}