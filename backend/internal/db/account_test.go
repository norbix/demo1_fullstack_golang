package db

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/norbix/demo1_fullstack_golang/backend/configs"
	"github.com/norbix/demo1_fullstack_golang/backend/internal/db/dbmodels"

	"github.com/stretchr/testify/assert"
)

// MockRoundTripper mocks HTTP requests.
type MockRoundTripper struct {
	Response *http.Response
	Err      error
	Requests []*http.Request
}

func (m *MockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	m.Requests = append(m.Requests, req)
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

	// Verify results
	assert.NoError(t, err, "CreateAccount should not return an error")
	assert.Len(t, mockTransport.Requests, 1, "One HTTP request should be made")

	// Verify the request method and URL
	req := mockTransport.Requests[0]
	assert.Equal(t, "POST", req.Method, "HTTP method should be POST")
	assert.Equal(t, config.BaseURL+"/document", req.URL.String(), "URL should match the expected endpoint")
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

	// Verify results
	assert.NoError(t, err, "GetAccount should not return an error")
	assert.NotNil(t, account, "Account should not be nil")
	assert.Equal(t, "12345", account.AccountNumber, "AccountNumber should match")
	assert.Equal(t, "John Doe", account.AccountName, "AccountName should match")
	assert.Len(t, mockTransport.Requests, 1, "One HTTP request should be made")

	// Verify the request method, URL, and payload
	req := mockTransport.Requests[0]
	assert.Equal(t, "POST", req.Method, "HTTP method should be POST")
	assert.Equal(t, config.BaseURL+"/documents/search", req.URL.String(), "URL should match the expected endpoint")

	// Verify the request body
	expectedBody := `{"query":{"account_number":"12345"}}`
	body, _ := ioutil.ReadAll(req.Body)
	assert.JSONEq(t, expectedBody, string(body), "Request body should match expected JSON")
}
