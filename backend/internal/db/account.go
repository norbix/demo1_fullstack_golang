package db

import (
	"errors"

	"github.com/norbix/demo1_fullstack_golang/backend/internal/db/dbmodels"
)

// AccountRepo is an in-memory implementation of AccountService.
type AccountRepo struct {
	accounts map[string]dbmodels.Account
}

// NewAccountRepo initializes a new AccountRepo.
func NewAccountRepo() *AccountRepo {
	return &AccountRepo{
		accounts: make(map[string]dbmodels.Account),
	}
}

// CreateAccount creates a new account and stores it in memory.
func (repo *AccountRepo) CreateAccount(account dbmodels.Account) error {
	// Ensure the account number is unique
	if _, exists := repo.accounts[account.AccountNumber]; exists {
		return errors.New("account already exists")
	}

	repo.accounts[account.AccountNumber] = account
	return nil
}

// GetAccount retrieves an account by its account number.
func (repo *AccountRepo) GetAccount(accountNumber string) (*dbmodels.Account, error) {
	account, exists := repo.accounts[accountNumber]
	if !exists {
		return nil, errors.New("account not found")
	}

	return &account, nil
}

// ListAccounts retrieves all accounts.
func (repo *AccountRepo) ListAccounts() ([]dbmodels.Account, error) {
	var result []dbmodels.Account
	for _, account := range repo.accounts {
		result = append(result, account)
	}
	return result, nil
}
