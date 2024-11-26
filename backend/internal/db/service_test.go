package db

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/norbix/demo1_fullstack_golang/backend/internal/db/dbmodels"
)

func TestDBService(t *testing.T) {
	repo := NewAccountRepo()
	service := NewDBService(repo)

	// Test CreateAccount
	account := dbmodels.Account{
		AccountNumber: "67890",
		AccountName:   "Jane Doe",
		IBAN:          "FR7630006000011234567890189",
		Address:       "5678 Oak St",
		Amount:        500.0,
		Type:          dbmodels.Receiving,
	}

	err := service.Repo.CreateAccount(account)
	assert.NoError(t, err, "CreateAccount should not return an error")

	// Test GetAccount
	retrievedAccount, err := service.Repo.GetAccount("67890")
	assert.NoError(t, err, "GetAccount should not return an error")
	assert.Equal(t, account, *retrievedAccount, "Retrieved account should match the created account")
}
