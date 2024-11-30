package db

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/norbix/demo1_fullstack_golang/backend/internal/db/dbmodels"
	"github.com/norbix/demo1_fullstack_golang/backend/internal/mocks"
)

func TestAccountRepository_CreateAccount(t *testing.T) {
	mockRepo := new(mocks.AccountRepository)

	t.Run("Valid Account", func(t *testing.T) {
		account := dbmodels.Account{
			AccountNumber: "12345",
			Amount:        100.0,
		}

		expectedResponse := map[string]interface{}{
			"documentId":    "abc123",
			"transactionId": "txn456",
		}

		mockRepo.EXPECT().
			CreateAccount(account).
			Return(expectedResponse, nil)

		response, err := mockRepo.CreateAccount(account)

		assert.NoError(t, err)
		assert.Equal(t, expectedResponse, response)

		mockRepo.AssertCalled(t, "CreateAccount", account)
	})

	t.Run("Serialization Error", func(t *testing.T) {
		account := dbmodels.Account{
			AccountNumber: "invalid-\xe9", // Invalid UTF-8 sequence
		}

		mockRepo.EXPECT().
			CreateAccount(account).
			Return(nil, errors.New("serialization error"))

		response, err := mockRepo.CreateAccount(account)

		assert.Nil(t, response)
		assert.EqualError(t, err, "serialization error")
	})
}

func TestAccountRepository_GetAccounts(t *testing.T) {
	mockRepo := new(mocks.AccountRepository)

	t.Run("Valid Pagination", func(t *testing.T) {
		page := 1
		perPage := 10

		expectedResponse := map[string]interface{}{
			"accounts": []map[string]interface{}{
				{"AccountNumber": "12345", "Amount": 100.0},
				{"AccountNumber": "67890", "Amount": 200.0},
			},
			"total": 2,
		}

		// Configure mock for Valid Pagination
		mockRepo.EXPECT().
			GetAccounts(page, perPage).
			Return(expectedResponse, nil)

		response, err := mockRepo.GetAccounts(page, perPage)

		assert.NoError(t, err)
		assert.Equal(t, expectedResponse, response)

		mockRepo.AssertCalled(t, "GetAccounts", page, perPage)
	})
}
