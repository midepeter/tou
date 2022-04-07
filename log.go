package main

import (
	"encoding/json"
	"errors"

	"go.uber.org/zap"
)

func initLogger() (*zap.Logger, error) {
	//config := zap.NewDevelopmentConfig()

	rawJSON := []byte(`{
		"level": "debug",
		"encoding": "json",
		"outputPaths": ["stdout", "/tmp/logs"],
		"errorOutputPaths": ["stderr"],
		"encoderConfig": {
			"messageKey": "message",
			"levelKey": "level",
			"levelEncoder": "lowercase"
		}
	}`)

	var cfg zap.Config

	err := json.Unmarshal(rawJSON, &cfg)
	if err != nil {
		return nil, errors.New("Error config failed")
	}

	logger, err := cfg.Build()
	if err != nil {
		return nil, errors.New("Error config failed")
	}
	defer logger.Sync()
	logger.Info("Logger successfully constructed")

	return logger, nil
}
