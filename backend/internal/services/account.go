package services

import (
	"errors"

	"github.com/norbix/demo1_fullstack_golang/backend/internal/db/dbmodels"
)

func (s accountServiceImpl) CreateAccount(account dbmodels.Account) (map[string]interface{}, error) {
	// Business rule: Ensure account number is not empty
	if account.AccountNumber == "" {
		return nil, errors.New("account number is required")
	}

	// Business rule: Ensure amount is non-negative
	if account.Amount < 0 {
		return nil, errors.New("amount cannot be negative")
	}

	// Delegate persistence to the repository
	response, err := s.repo.CreateAccount(account)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetAccounts retrieves a list of accounts.
func (s accountServiceImpl) GetAccounts(page, perPage int) (map[string]interface{}, error) {
	response, err := s.repo.GetAccounts(page, perPage)
	if err != nil {
		return nil, err
	}

	return response, nil
}
