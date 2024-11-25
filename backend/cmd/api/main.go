package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Create a simple HTTP server with a /healthz endpoint
	http.HandleFunc("/healthz", healthHandler)

	// Print a startup message
	fmt.Println("Starting Backend Component on port 8080...")

	// Start the HTTP server on port 8080
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

// healthHandler responds to /healthz with a health status message
func healthHandler(w http.ResponseWriter, r *http.Request) {
	// Respond with HTTP 200 OK
	w.WriteHeader(http.StatusOK)

	// Write a simple health status message
	_, _ = w.Write([]byte("Backend is healthy!"))
}
