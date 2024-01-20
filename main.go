package main

import (
	"log"
	"net/http"
)

type apiConfig struct {
  fileServerHits int
}

func main() {
  const filepathRoot = "./app"
  const port = "8080"

  apiCfg := &apiConfig{
    fileServerHits: 0,
  }

  mux := http.NewServeMux()
  fileHandler := http.StripPrefix("/app/", http.FileServer(http.Dir(filepathRoot)))
  mux.Handle("/app/", apiCfg.middlewareMetricsInc(fileHandler))
  mux.HandleFunc("/healthz", readinessHandler)
  mux.HandleFunc("/metrics", apiCfg.metricsHandler)
  mux.HandleFunc("/reset", apiCfg.metricsResetHandler)
  corsMux := middlewareCors(mux)

  srv := &http.Server{
    Addr: ":" + port,
    Handler: corsMux,
  }

  log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
  log.Fatal(srv.ListenAndServe())
}
