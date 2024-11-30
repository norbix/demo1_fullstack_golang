package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/norbix/demo1_fullstack_golang/backend/internal/db/dbmodels"
	"github.com/norbix/demo1_fullstack_golang/backend/internal/mocks"
)

func TestAccountHandler_CreateAccount(t *testing.T) {
	mockService := new(mocks.AccountService)
	handler := NewAccountHandler(mockService)

	t.Run("Valid Account", func(t *testing.T) {
		account := dbmodels.Account{
			AccountNumber: "12345",
			Amount:        100.0,
		}

		expectedResponse := map[string]interface{}{
			"documentId":    "674b8ac7000000000000000a53d8e43d",
			"transactionId": "11",
		}

		mockService.EXPECT().CreateAccount(account).Return(expectedResponse, nil)

		body, _ := json.Marshal(account)
		req := httptest.NewRequest(http.MethodPut, "/accounts", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()

		handler.CreateAccount(rr, req)

		assert.Equal(t, http.StatusCreated, rr.Code) // Expecting 201 Created
		var response map[string]interface{}
		json.NewDecoder(rr.Body).Decode(&response)
		assert.Equal(t, expectedResponse, response)
	})

	t.Run("Invalid Account Number", func(t *testing.T) {
		account := dbmodels.Account{
			AccountNumber: "",
			Amount:        100.0,
		}

		body, _ := json.Marshal(account)
		req := httptest.NewRequest(http.MethodPut, "/accounts", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()

		handler.CreateAccount(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		mockService.AssertNotCalled(t, "CreateAccount")
	})

	t.Run("Negative Amount", func(t *testing.T) {
		account := dbmodels.Account{
			AccountNumber: "12345",
			Amount:        -10.0,
		}

		body, _ := json.Marshal(account)
		req := httptest.NewRequest(http.MethodPut, "/accounts", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()

		handler.CreateAccount(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		mockService.AssertNotCalled(t, "CreateAccount")
	})
}
func TestAccountHandler_GetAccounts(t *testing.T) {
	mockService := new(mocks.AccountService)
	handler := NewAccountHandler(mockService)

	t.Run("Valid Pagination", func(t *testing.T) {
		pagination := map[string]int{
			"page":    1,
			"perPage": 10,
		}

		expectedResponse := map[string]interface{}{
			"accounts": []interface{}{
				map[string]interface{}{"AccountNumber": "12345", "Amount": 100.0},
				map[string]interface{}{"AccountNumber": "67890", "Amount": 200.0},
			},
			"total": float64(2),
		}

		mockService.EXPECT().GetAccounts(1, 10).Return(expectedResponse, nil)

		body, _ := json.Marshal(pagination)
		req := httptest.NewRequest(http.MethodPost, "/accounts/retrieve", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()

		handler.GetAccounts(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		var response map[string]interface{}
		json.NewDecoder(rr.Body).Decode(&response)
		assert.Equal(t, expectedResponse, response)
	})

	t.Run("Invalid Request Body", func(t *testing.T) {
		body := []byte("invalid-json")

		req := httptest.NewRequest(http.MethodPost, "/accounts/retrieve", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()

		handler.GetAccounts(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, "Invalid request body\n", rr.Body.String())
	})
}
