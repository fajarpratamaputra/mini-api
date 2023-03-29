package main

import (
	"context"
	"fmt"
	"github.com/labstack/gommon/log"
	"interaction-api/api"
	"interaction-api/config"
	"os"
	"os/signal"
	"time"
)

func main() {
	e := api.RegisterPath()
	// run server
	go func() {
		address := fmt.Sprintf("0.0.0.0:%s", config.AppConfig.Port)

		if err := e.Start(address); err != nil {
			log.Info("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// a timeout of 10 seconds to shutdown the server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
