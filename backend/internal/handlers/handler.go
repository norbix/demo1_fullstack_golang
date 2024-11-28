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
func (h *AccountHandler) GetAccounts(w http.ResponseWriter, r *http.Request) {
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
	json.NewEncoder(w).Encode(response)
}
