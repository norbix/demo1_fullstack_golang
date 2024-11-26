package db

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/norbix/demo1_fullstack_golang/backend/configs"
	"github.com/norbix/demo1_fullstack_golang/backend/internal/db/dbmodels"
)

// MockRoundTripper mocks HTTP requests.
type MockRoundTripper struct {
	Response *http.Response
	Err      error
}

func (m *MockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return m.Response, m.Err
}

func TestAccountRepo_CreateAccount(t *testing.T) {
	// Mock HTTP client
	mockResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(bytes.NewBufferString(``)),
	}
	mockTransport := &MockRoundTripper{Response: mockResponse, Err: nil}
	client := &http.Client{Transport: mockTransport}

	// Initialize AccountRepo
	config := &configs.Config{
		BaseURL: "https://vault.immudb.io/ics/api/v1/ledger/default/collection/default",
		APIKey:  "test-write-api-key",
	}
	repo := NewAccountRepo(config, client)

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
}

func TestAccountRepo_GetAccount(t *testing.T) {
	// Mock HTTP client
	mockResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body: ioutil.NopCloser(bytes.NewBufferString(`{
			"documents": [{
				"account_number": "12345",
				"account_name": "John Doe",
				"iban": "DE89370400440532013000",
				"address": "1234 Elm St",
				"amount": 1000.0,
				"type": "sending"
			}]
		}`)),
	}
	mockTransport := &MockRoundTripper{Response: mockResponse, Err: nil}
	client := &http.Client{Transport: mockTransport}

	// Initialize AccountRepo
	config := &configs.Config{
		BaseURL: "https://vault.immudb.io/ics/api/v1/ledger/default/collection/default",
		APIKey:  "test-write-api-key",
	}
	repo := NewAccountRepo(config, client)

	// Test GetAccount
	account, err := repo.GetAccount("12345")
	assert.NoError(t, err, "GetAccount should not return an error")
	assert.Equal(t, "12345", account.AccountNumber, "AccountNumber should match")
	assert.Equal(t, "John Doe", account.AccountName, "AccountName should match")
}
