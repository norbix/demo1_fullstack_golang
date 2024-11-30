package services

import (
	"github.com/norbix/demo1_fullstack_golang/backend/internal/db"
	"github.com/norbix/demo1_fullstack_golang/backend/internal/db/dbmodels"
)

// AccountService defines the interface for account-related operations.
type AccountService interface {
	CreateAccount(dbmodels.Account) error
	GetAccounts(int, int) (map[string]interface{}, error)
}

type accountServiceImpl struct {
	repo db.AccountRepository
}

func NewAccountService(repo db.AccountRepository) AccountService {
	return accountServiceImpl{repo: repo}
}
