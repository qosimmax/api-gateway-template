// Package middleware provides HTTP middleware.
package middleware

import (
	"api-gateway-template/monitoring/metrics"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// TraceMetrics is the configuration for trace and metrics middleware.
type TraceMetrics struct{}

// MetricsMiddleware collects HTTP request metrics for Prometheus.
// Collects request duration and response code.
func (tm *TraceMetrics) MetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		routeName := mux.CurrentRoute(r).GetName()

		crw := customResponseWriter{ResponseWriter: w}
		start := time.Now()

		next.ServeHTTP(&crw, r)

		duration := time.Since(start)

		ctx := r.Context()
		metrics.ObserveTimeToProcess(ctx, routeName, duration.Seconds())
		metrics.ReceivedRequest(ctx, crw.status, routeName)
	})
}

type customResponseWriter struct {
	http.ResponseWriter
	status int
}

func (crw *customResponseWriter) WriteHeader(status int) {
	crw.status = status
	crw.ResponseWriter.WriteHeader(status)
}
