package main

import (
	"fmt"
	"log"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/norbix/demo1_fullstack_golang/backend/configs"
	"github.com/norbix/demo1_fullstack_golang/backend/internal/db"
	"github.com/norbix/demo1_fullstack_golang/backend/internal/handlers"
	"github.com/norbix/demo1_fullstack_golang/backend/internal/services"

	_ "github.com/norbix/demo1_fullstack_golang/backend/docs"

	"github.com/gorilla/mux"
)

// @title Backend Component API
// @version 1.0
// @description This is a sample server for managing accounts.
// @host localhost:8080
// @BasePath /
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
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	router.HandleFunc("/accounts", accountHandler.CreateAccount).Methods("PUT")
	router.HandleFunc("/accounts/retrieve", accountHandler.GetAccounts).Methods("POST")

	// Start the server
	fmt.Println("Starting Backend Component on port 8080...")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}

// @Summary Health check
// @Description Responds with "Backend is healthy!" if the service is up.
// @Tags health
// @Accept  json
// @Produce json
// @Success 200 {string} string "Backend is healthy!"
// @Router /healthz [get]
func healthHandler(w http.ResponseWriter, r *http.Request) {
	// Respond with HTTP 200 OK
	w.WriteHeader(http.StatusOK)

	// Write a simple health status message
	_, _ = w.Write([]byte("Backend is healthy!"))
}
