package main

import (
	"fmt"
	"net/http"
)

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    cfg.fileServerHits++
    next.ServeHTTP(w, r)
  })
}

func (cfg *apiConfig) metricsHandler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/plain; charset=utf-8")
  w.WriteHeader(http.StatusOK)
  fmt.Fprintf(w, "Hits: %v", cfg.fileServerHits)
}

func (cfg *apiConfig) metricsResetHandler(w http.ResponseWriter, r *http.Request) {
  cfg.fileServerHits = 0
  w.WriteHeader(http.StatusOK)
  w.Write([]byte("Hits reset to 0"))
}
