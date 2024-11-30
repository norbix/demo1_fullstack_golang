package db

import (
	"crypto/tls"
	"net/http"

	"github.com/norbix/demo1_fullstack_golang/backend/configs"
	"github.com/norbix/demo1_fullstack_golang/backend/internal/db/dbmodels"
)

type AccountRepository interface {
	CreateAccount(dbmodels.Account) (map[string]interface{}, error)
	GetAccounts(int, int) (map[string]interface{}, error)
}

type accountRepositoryImpl struct {
	config *configs.Config
	client *http.Client
}

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
