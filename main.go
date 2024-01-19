package main

import (
	"log"
	"net/http"
)

func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func readinessHandler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/plain; charset=utf-8")
  w.WriteHeader(200)
  w.Write([]byte("OK"))
}

func main() {
  const filepathRoot = "./app"
  const port = "8080"

  mux := http.NewServeMux()
  mux.Handle("/app/", http.StripPrefix("/app/", http.FileServer(http.Dir(filepathRoot))))
  mux.HandleFunc("/healthz", readinessHandler)
  corsMux := middlewareCors(mux)

  srv := &http.Server{
    Addr: ":" + port,
    Handler: corsMux,
  }

  log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
  log.Fatal(srv.ListenAndServe())
}
