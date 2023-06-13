package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/gorilla/mux"
)

// TODO: Create tests for all methods
// TODO: Create e2e test for `/run`
// TODO: More error checks
// TODO: Logger with timestamps
// TODO: Provide more logging information (log on each stage)
// TODO: Add benchmark
// TODO: Optimize inner functions
// TODO: Connect as a service for tightshorts

func main() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              os.Getenv("ANALYZER_SENTRY_DSN"),
		TracesSampleRate: 0.2,
		SampleRate:       1.0,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	defer sentry.Flush(2 * time.Second)

	r := mux.NewRouter()
	r.HandleFunc("/run", runAnalyzerHandler).
		Headers("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("ANALYZER_TOKEN")))
	log.Fatal(http.ListenAndServe(":3030", r))
}

func runAnalyzerHandler(w http.ResponseWriter, r *http.Request) {
	Run()
	w.WriteHeader(http.StatusOK)
}
