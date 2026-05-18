package main

import (
	"log"
	"net/http"

	"ride-sharing/shared/env"
)

var (
	httpAddr       = env.GetString("GATEWAY_HTTP_ADDR", ":8081")
	tripServiceURL = env.GetString("TRIP_SERVICE_URL", "http://localhost:8083")
)

func main() {
	log.Println("Starting API Gateway")

	mux := http.NewServeMux()

	mux.HandleFunc("POST /trip/preview", handleTripReview)
	mux.HandleFunc("POST /trip/preview/", handleTripReview)

	server := &http.Server{
		Addr:    httpAddr,
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Printf("HTTP server error: %v", err)
	}
}
