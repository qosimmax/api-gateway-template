package main

import (
	"api-gateway-template/config"
	"api-gateway-template/server"
	"context"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	log.Info("Starting ...")

	ctx := context.Background()
	config, err := config.LoadConfig()

	if err != nil {
		log.WithField("err", err.Error()).Fatal("Failed to load config")
	}

	var s server.Server

	if err := s.Create(ctx, config); err != nil {
		log.WithField("err", err.Error()).Fatal("Server error from s.Create()")
	}

	if err := s.Serve(ctx); err != nil {
		log.WithField("err", err.Error()).Fatal("Server error from s.Serve()")
	}
}
