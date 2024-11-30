package db

import (
	"crypto/tls"
	"net/http"

	"github.com/norbix/demo1_fullstack_golang/backend/configs"
	"github.com/norbix/demo1_fullstack_golang/backend/internal/db/dbmodels"
)

// AccountRepository defines the interface for account-related database operations.
type AccountRepository interface {
	// CreateAccount persists a new account to the database.
	CreateAccount(dbmodels.Account) error

	// GetAccounts retrieves a list of accounts with pagination.
	GetAccounts(int, int) (map[string]interface{}, error)
}

type accountRepositoryImpl struct {
	config *configs.Config
	client *http.Client
}

// NewAccountRepo initializes the AccountRepo with the given config and HTTP client.
// If no client is provided, it defaults to http.DefaultClient.
func NewAccountRepo(config *configs.Config, client *http.Client) AccountRepository {
	if config.SkipTLS {
		client = &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		}
	} else {
		client = http.DefaultClient
	}

	return accountRepositoryImpl{
		config: config,
		client: client,
	}
}
