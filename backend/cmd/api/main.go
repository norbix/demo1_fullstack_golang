package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/norbix/demo1_fullstack_golang/backend/configs"
	"github.com/norbix/demo1_fullstack_golang/backend/internal/db"
	"github.com/norbix/demo1_fullstack_golang/backend/internal/handlers"
	"github.com/norbix/demo1_fullstack_golang/backend/internal/services"

	"github.com/gorilla/mux" // Router library
)

func main() {
	// Load configuration
	config, err := configs.LoadConfig() // Load configuration
	if err != nil {
		log.Fatalf("Error loading configuration: %s", err)
	}

	// Initialize dependencies
	accountRepo := db.NewAccountRepo(config, nil)                // Repository layer
	accountService := services.NewAccountService(accountRepo)    // Service layer
	accountHandler := handlers.NewAccountHandler(accountService) // HTTP handlers

	// Create a router
	router := mux.NewRouter()

	// Register endpoints
	router.HandleFunc("/healthz", healthHandler).Methods("GET")
	router.HandleFunc("/accounts", accountHandler.CreateAccount).Methods("POST")
	router.HandleFunc("/accounts/retrieve", accountHandler.GetAccount).Methods("POST")

	// Start the server
	fmt.Println("Starting Backend Component on port 8080...")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}

// healthHandler responds to /healthz with a health status message
func healthHandler(w http.ResponseWriter, r *http.Request) {
	// Respond with HTTP 200 OK
	w.WriteHeader(http.StatusOK)

	// Write a simple health status message
	_, _ = w.Write([]byte("Backend is healthy!"))
}
