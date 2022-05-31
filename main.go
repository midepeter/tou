package main

import (
	"os"
	"workqueue/cmd"

	"go.uber.org/zap"
)

var log *zap.Logger

func main() {
	err := cmd.Run()
	if err != nil {
		log.Error("App failed")
		os.Exit(1)
	}
}
