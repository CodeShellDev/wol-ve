package config

import (
	"os"
	"strconv"

	"github.com/codeshelldev/gotl/pkg/logger"
	"github.com/codeshelldev/wol-ve/internals/config/structure"
)

var ENV = &structure.ENV{
	LOG_LEVEL: "info",
	PORT: "9000",
	ADDR: "0.0.0.0",
	PING_INTERVAL: 5,
	PING_RETRIES: 3,
}

func Load() {
	logLevel := os.Getenv("LOG_LEVEL")

	if logLevel != "" {
		ENV.LOG_LEVEL = logLevel
	}

	port := os.Getenv("PORT")

	if port != "" {
		ENV.PORT = port
	}

	addr := os.Getenv("ADDR")

	if addr != "" {
		ENV.ADDR = addr
	}

	pingInterval := os.Getenv("PING_INTERVAL")

	if pingInterval != "" {
		interval, err := strconv.Atoi(pingInterval)

		if err != nil {
			logger.Error("Invalid ping interval: ", err.Error())
		} else {
			ENV.PING_INTERVAL = interval
		}
	}

	pingRetries := os.Getenv("PING_RETRIES")

	if pingRetries != "" {
		retries, err := strconv.Atoi(pingRetries)

		if err != nil {
			logger.Error("Invalid ping retries: ", err.Error())
		} else {
			ENV.PING_RETRIES = retries
		}
	}
}

func Log() {
	logger.Dev("Loaded Environment:", ENV)
}