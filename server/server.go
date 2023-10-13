// Package server provides functionality to easily set up an HTTTP server.
//
// The server holds all the clients it needs and they should be set up in the Create method.
//
// The HTTP routes and middleware are set up in the setupRouter method.
package server

import (
	"api-gateway-template/client/fakeapi"
	"api-gateway-template/config"
	"api-gateway-template/monitoring/metrics"
	"api-gateway-template/monitoring/trace"
	"context"
	"fmt"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"

	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// Server holds the HTTP server, router, config and all clients.
type Server struct {
	Config         *config.Config
	API            *fakeapi.Client
	HTTP           *http.Server
	Router         *mux.Router
	TracerProvider *tracesdk.TracerProvider
}

// Create sets up the HTTP server, router and all clients.
// Returns an error if an error occurs.
func (s *Server) Create(ctx context.Context, config *config.Config) error {
	metrics.RegisterPrometheusCollectors()

	var apiClient fakeapi.Client
	if err := apiClient.Init(config); err != nil {
		return fmt.Errorf("api client: %w", err)
	}

	s.API = &apiClient
	s.Config = config
	s.Router = mux.NewRouter()
	s.HTTP = &http.Server{
		Addr:    fmt.Sprintf(":%s", s.Config.Port),
		Handler: s.Router,
	}

	s.setupRoutes()

	return nil
}

// Serve tells the server to start listening and serve HTTP requests.
// It also makes sure that the server gracefully shuts down on exit.
// Returns an error if an error occurs.
func (s *Server) Serve(ctx context.Context) error {
	var err error
	s.TracerProvider, err = trace.TracerProvider(s.Config)
	if err != nil {
		return fmt.Errorf("init global tracer: %w", err)
	}

	idleConnsClosed := make(chan struct{}) // this is used to signal that we can not exit
	go func(ctx context.Context) {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

		<-stop

		log.Info("Shutdown signal received")
		s.shutdown(ctx)

		close(idleConnsClosed) // call close to say we can now exit the function
	}(ctx)

	log.Infof("Ready at: %s", s.Config.Port)

	if err := s.HTTP.ListenAndServe(); err != http.ErrServerClosed {
		return fmt.Errorf("unexpected server error: %w", err)
	}
	<-idleConnsClosed // this will block until close is called

	return nil
}

func (s *Server) shutdown(ctx context.Context) {
	if err := s.TracerProvider.Shutdown(ctx); err != nil {
		log.Error(err.Error())
	}

	if err := s.API.Close(); err != nil {
		log.Error(err.Error())
	}

	if err := s.HTTP.Shutdown(ctx); err != nil {
		log.Error(err.Error())
	}

}
