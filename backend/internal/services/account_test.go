package services

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/norbix/demo1_fullstack_golang/backend/internal/db/dbmodels"
	"github.com/norbix/demo1_fullstack_golang/backend/internal/mocks"
)

func TestAccountService_CreateAccount(t *testing.T) {
	mockRepo := new(mocks.AccountRepository)
	service := NewAccountService(mockRepo)

	t.Run("Valid Account", func(t *testing.T) {
		account := dbmodels.Account{
			AccountNumber: "12345",
			Amount:        100.0,
		}

		expectedResponse := map[string]interface{}{
			"documentId":    "abc123",
			"transactionId": "txn456",
		}

		// Mocking repository behavior for valid input
		mockRepo.On("CreateAccount", account).Return(expectedResponse, nil)

		response, err := service.CreateAccount(account)

		assert.NoError(t, err)
		assert.Equal(t, expectedResponse, response)

		mockRepo.AssertCalled(t, "CreateAccount", account)
	})
}

func TestAccountService_GetAccounts(t *testing.T) {
	mockRepo := new(mocks.AccountRepository)
	service := NewAccountService(mockRepo)

	t.Run("Valid Request", func(t *testing.T) {
		page := 1
		perPage := 10

		expectedResponse := map[string]interface{}{
			"accounts": []interface{}{
				map[string]interface{}{"AccountNumber": "12345", "Amount": 100.0},
				map[string]interface{}{"AccountNumber": "67890", "Amount": 200.0},
			},
			"total": float64(2),
		}

		mockRepo.On("GetAccounts", page, perPage).Return(expectedResponse, nil)

		response, err := service.GetAccounts(page, perPage)

		assert.NoError(t, err)
		assert.Equal(t, expectedResponse, response)

		mockRepo.AssertCalled(t, "GetAccounts", page, perPage)
	})
}
