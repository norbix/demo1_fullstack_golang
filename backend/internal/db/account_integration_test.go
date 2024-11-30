package db

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/norbix/demo1_fullstack_golang/backend/configs"
	"github.com/norbix/demo1_fullstack_golang/backend/internal/db/dbmodels"
)

func TestAccountRepository_CreateAccount_Integration(t *testing.T) {
	config := &configs.Config{
		BaseURL: "https://vault.immudb.io/ics/api/v1/ledger/default/collection/default",
		APIKey:  "default.jKXkTQquKyXAEfz1qHei1A.gTmSG38ipa8QNz4jPVLUJuw6etoejMTkqZ9fxwvovQ9xNBV_",
		SkipTLS: true,
	}
	client := &http.Client{}
	repo := NewAccountRepo(config, client)

	t.Run("Valid Account", func(t *testing.T) {
		account := dbmodels.Account{
			AccountNumber: "12345",
			Amount:        100.0,
		}

		// Call the actual implementation
		response, err := repo.CreateAccount(account)

		// Assert the results
		assert.NotNil(t, response)
		assert.NoError(t, err)
	})
}
