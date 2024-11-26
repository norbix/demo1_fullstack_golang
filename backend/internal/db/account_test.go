package db

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/norbix/demo1_fullstack_golang/backend/internal/db/dbmodels"
)

func TestAccountRepo(t *testing.T) {
	repo := NewAccountRepo()

	// Test CreateAccount
	account := dbmodels.Account{
		AccountNumber: "12345",
		AccountName:   "John Doe",
		IBAN:          "DE89370400440532013000",
		Address:       "1234 Elm St",
		Amount:        1000.0,
		Type:          dbmodels.Sending,
	}

	err := repo.CreateAccount(account)
	assert.NoError(t, err, "CreateAccount should not return an error")

	// Test GetAccount
	retrievedAccount, err := repo.GetAccount("12345")
	assert.NoError(t, err, "GetAccount should not return an error")
	assert.Equal(t, account, *retrievedAccount, "Retrieved account should match the created account")

	// Test ListAccounts
	accounts, err := repo.ListAccounts()
	assert.NoError(t, err, "ListAccounts should not return an error")
	assert.Equal(t, 1, len(accounts), "ListAccounts should return one account")

	// Test CreateAccount with duplicate account number
	err = repo.CreateAccount(account)
	assert.Error(t, err, "CreateAccount should return an error for duplicate account number")
}
