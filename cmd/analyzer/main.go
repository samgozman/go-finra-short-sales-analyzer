package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/run", runAnalyzerHandler)
	log.Fatal(http.ListenAndServe(":8000", r))
}

func runAnalyzerHandler(w http.ResponseWriter, r *http.Request) {
	Run()
	w.WriteHeader(http.StatusOK)
}
