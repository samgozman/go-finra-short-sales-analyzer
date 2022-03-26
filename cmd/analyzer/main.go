package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// TODO: Create tests for all methods
// TODO: Create e2e test for `/run`
// TODO: More error checks
// TODO: Logger with timestamps
// TODO: Provide more logging information (log on each stage)
// TODO: Use "generics" from go 1.18
// TODO: Add benchmark
// TODO: Optimize inner functions
// TODO: Connect as a service for tightshorts

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/run", runAnalyzerHandler)
	log.Fatal(http.ListenAndServe(":8000", r))
}

func runAnalyzerHandler(w http.ResponseWriter, r *http.Request) {
	Run()
	w.WriteHeader(http.StatusOK)
}
