package server

import (
	"api-gateway-template/server/internal/handler"
	"api-gateway-template/server/internal/middleware"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
)

const v1API string = "/api/v1"

// setupRoutes - the root route function.
func (s *Server) setupRoutes() {
	s.Router.Handle("/metrics", promhttp.Handler()).Name("Metrics")
	s.Router.HandleFunc("/_healthz", handler.Healthz).Methods(http.MethodGet).Name("Health")

	api := s.Router.PathPrefix(v1API).Subrouter()
	api.HandleFunc("/example", handler.Example(s.API)).Methods(http.MethodGet).Name("Example")

	addTracingAndMetrics(api)

}

// addTracingAndMetrics - Adds tracing and metrics to a router.
func addTracingAndMetrics(r *mux.Router) {
	r.Use(otelmux.Middleware("api-gateway"))

	tm := middleware.TraceMetrics{}
	r.Use(tm.MetricsMiddleware)

}
