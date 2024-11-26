package db

import (
	"github.com/norbix/demo1_fullstack_golang/backend/internal/db/dbmodels"
)

// AccountService defines the interface for account-related operations.
type AccountService interface {
	// CreateAccount creates a new account.
	CreateAccount(account dbmodels.Account) error

	// GetAccount retrieves an account by its account number.
	GetAccount(accountNumber string) (*dbmodels.Account, error)

	// ListAccounts retrieves all accounts.
	ListAccounts() ([]dbmodels.Account, error)
}

// DBService provides an implementation of AccountService.
type DBService struct {
	// This can include any dependencies, e.g., a database client
	Repo AccountService
}

// NewDBService initializes a new DBService with an AccountService implementation.
func NewDBService(repo AccountService) *DBService {
	return &DBService{
		Repo: repo,
	}
}
