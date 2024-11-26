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

// GetAccount retrieves an account and performs additional checks.
func (s *AccountService) GetAccount(accountNumber string) (*dbmodels.Account, error) {
	// Validate input
	if accountNumber == "" {
		return nil, errors.New("account number is required")
	}

	// Retrieve the account from the repository
	account, err := s.repo.GetAccount(accountNumber)
	if err != nil {
		return nil, err
	}

	return account, nil
}
