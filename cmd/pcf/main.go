package main

import (
	"log"
	"net/http"
)

// Phase C0 placeholder. No Chi router, no SBI handlers yet (that starts in
// C1, once contracts/openapi/ exists). The only purpose of this file is to
// give `go build ./...` something real to compile, and to have an endpoint
// that can be checked manually.
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("PulseGrid PCF — alive"))
	})

	addr := ":8080"
	log.Printf("pcf-api listening on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
