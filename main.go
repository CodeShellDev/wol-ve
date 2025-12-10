package main

import (
	"net/http"

	log "github.com/codeshelldev/gotl/pkg/logger"
	"github.com/codeshelldev/wol-ve/internals/config"
	"github.com/codeshelldev/wol-ve/internals/server"
)

func main() {
	config.Load()

	log.Init(config.ENV.LOG_LEVEL)

	log.Info("Initialized Logger with Level of ", log.Level())

	if log.Level() == "dev" {
		log.Dev("Welcome back Developer!")
	}

	config.Log()

	addr := config.ENV.ADDR + ":" + config.ENV.PORT

	srv := &http.Server{
		Addr:    addr,
		Handler: server.Handle(),
	}

	err := srv.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		log.Fatal("Server error: ", err.Error())
	}
}
