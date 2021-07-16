package httphandlers

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func InitDefaultHandlers(mux http.Handler) {
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/health-check", HealthCheck)
	http.Handle("/api/v1/", AccessControl(mux))
}
