package services

import (
	"errors"

	"github.com/norbix/demo1_fullstack_golang/backend/internal/db"
	"github.com/norbix/demo1_fullstack_golang/backend/internal/db/dbmodels"
)

// AccountService provides account-related business logic.
type AccountService struct {
	repo db.AccountService
}

// NewAccountService initializes an AccountService.
func NewAccountService(repo db.AccountService) *AccountService {
	return &AccountService{repo: repo}
}

// CreateAccount validates and creates a new account.
func (s *AccountService) CreateAccount(account dbmodels.Account) error {
	// Business rule: Ensure account number is not empty
	if account.AccountNumber == "" {
		return errors.New("account number is required")
	}

	// Business rule: Ensure amount is non-negative
	if account.Amount < 0 {
		return errors.New("amount cannot be negative")
	}

	// Delegate persistence to the repository
	return s.repo.CreateAccount(account)
}

// GetAccounts retrieves a list of accounts.
func (s *AccountService) GetAccounts(page, perPage int) (map[string]interface{}, error) {
	response, err := s.repo.GetAccounts(page, perPage)
	if err != nil {
		return nil, err
	}

	return response, nil
}
