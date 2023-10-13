package server

import (
	"api-gateway-template/server/internal/handler"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

const v1API string = "/api/v1"

// setupRoutes - the root route function.
func (s *Server) setupRoutes() {
	s.Router.Handle("/metrics", promhttp.Handler()).Name("Metrics")
	s.Router.HandleFunc("/_healthz", handler.Healthz).Methods(http.MethodGet).Name("Health")

	api := s.Router.PathPrefix(v1API).Subrouter()
	api.HandleFunc("/example", handler.Example(s.API)).Methods(http.MethodGet).Name("Example")
}
