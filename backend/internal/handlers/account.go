package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/norbix/demo1_fullstack_golang/backend/internal/db/dbmodels"
)

// @Summary Create an account
// @Description Creates a new account with the provided details.
// @Tags accounts
// @Accept  json
// @Produce json
// @Param account body dbmodels.Account true "Account data"
// @Success 201 {string} string "Created"
// @Failure 400 {string} string "Invalid request body"
// @Failure 500 {string} string "Internal server error"
// @Router /accounts [put]
func (h accountHandlerImpl) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var account dbmodels.Account
	// Parse the request body into the account struct
	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate the input
	if account.AccountNumber == "" {
		http.Error(w, "Account number is required", http.StatusBadRequest)
		return
	}
	if account.Amount < 0 {
		http.Error(w, "Amount cannot be negative", http.StatusBadRequest)
		return
	}

	// Call the service layer to create the account
	response, err := h.accountService.CreateAccount(account)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

// @Summary Retrieve accounts
// @Description Retrieves accounts with pagination.
// @Tags accounts
// @Accept  json
// @Produce json
// @Param pagination body map[string]int true "Pagination details"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {string} string "Invalid request body"
// @Failure 500 {string} string "Internal server error"
// @Router /accounts/retrieve [post]
func (h accountHandlerImpl) GetAccounts(w http.ResponseWriter, r *http.Request) {
	var pagination struct {
		Page    int `json:"page"`
		PerPage int `json:"perPage"`
	}

	if err := json.NewDecoder(r.Body).Decode(&pagination); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	response, err := h.accountService.GetAccounts(pagination.Page, pagination.PerPage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
