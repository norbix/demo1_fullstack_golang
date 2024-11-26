package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/norbix/demo1_fullstack_golang/backend/internal/db/dbmodels"
	"github.com/norbix/demo1_fullstack_golang/backend/internal/services"
)

// AccountHandler provides HTTP handlers for account-related endpoints.
type AccountHandler struct {
	accountService *services.AccountService
}

// NewAccountHandler initializes an AccountHandler with the given service.
func NewAccountHandler(accountService *services.AccountService) *AccountHandler {
	return &AccountHandler{accountService: accountService}
}

// CreateAccount handles the "POST /accounts" endpoint.
func (h *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var account dbmodels.Account
	// Parse the request body into the account struct
	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Call the service layer to create the account
	if err := h.accountService.CreateAccount(account); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with 201 Created
	w.WriteHeader(http.StatusCreated)
}

// GetAccount handles the "GET /accounts/{accountNumber}" endpoint.
func (h *AccountHandler) GetAccount(w http.ResponseWriter, r *http.Request) {
	// Extract the account number from the URL path
	accountNumber := r.URL.Query().Get("account_number")
	if accountNumber == "" {
		http.Error(w, "Account number is required", http.StatusBadRequest)
		return
	}

	// Call the service layer to retrieve the account
	account, err := h.accountService.GetAccount(accountNumber)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Respond with the account data as JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(account); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
