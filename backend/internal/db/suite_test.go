package db

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/norbix/demo1_fullstack_golang/backend/configs"
	"github.com/norbix/demo1_fullstack_golang/backend/internal/db/dbmodels"
)

// AccountRepoSuite defines the test suite for AccountRepo.
type AccountRepoSuite struct {
	suite.Suite
	config *configs.Config
	client *http.Client
	repo   *AccountRepo
}

// SetupSuite runs once before all tests in the suite.
func (suite *AccountRepoSuite) SetupSuite() {
	suite.config = &configs.Config{
		BaseURL: "https://vault.immudb.io/ics/api/v1/ledger/default/collection/default",
		APIKey:  "test-api-key",
	}
}

// SetupTest runs before each test in the suite.
func (suite *AccountRepoSuite) SetupTest() {
	// Reset the mock HTTP client before each test
	suite.client = &http.Client{
		Transport: &MockRoundTripper{},
	}
	suite.repo = NewAccountRepo(suite.config, suite.client)
}

func (suite *AccountRepoSuite) TestCreateAccount_Success() {
	// Given: A valid account and a successful HTTP response
	account := dbmodels.Account{
		AccountNumber: "12345",
		AccountName:   "John Doe",
		IBAN:          "DE89370400440532013000",
		Address:       "123 Elm St",
		Amount:        1000.0,
		Type:          dbmodels.Sending,
	}
	mockResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(bytes.NewBufferString(``)),
	}
	suite.client.Transport.(*MockRoundTripper).Response = mockResponse

	// When: CreateAccount is called
	err := suite.repo.CreateAccount(account)

	// Then: No error should occur
	assert.NoError(suite.T(), err)
}

func (suite *AccountRepoSuite) TestCreateAccount_Failure() {
	// Given: A valid account and a failing HTTP response
	account := dbmodels.Account{
		AccountNumber: "12345",
		AccountName:   "John Doe",
		IBAN:          "DE89370400440532013000",
		Address:       "123 Elm St",
		Amount:        1000.0,
		Type:          dbmodels.Sending,
	}
	mockResponse := &http.Response{
		StatusCode: http.StatusInternalServerError,
		Body:       ioutil.NopCloser(bytes.NewBufferString(``)),
	}
	suite.client.Transport.(*MockRoundTripper).Response = mockResponse

	// When: CreateAccount is called
	err := suite.repo.CreateAccount(account)

	// Then: An error should occur
	assert.Error(suite.T(), err)
	assert.Contains(suite.T(), err.Error(), "unexpected status code")
}

func (suite *AccountRepoSuite) TestGetAccount_Success() {
	// Given: A valid account and a successful HTTP response
	mockResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body: ioutil.NopCloser(bytes.NewBufferString(`{
			"documents": [{
				"account_number": "12345",
				"account_name": "John Doe",
				"iban": "DE89370400440532013000",
				"address": "123 Elm St",
				"amount": 1000.0,
				"type": "sending"
			}]
		}`)),
	}
	suite.client.Transport.(*MockRoundTripper).Response = mockResponse

	// When: GetAccount is called with an existing account number
	account, err := suite.repo.GetAccount("12345")

	// Then: The account should be returned with no error
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), account)
	assert.Equal(suite.T(), "12345", account.AccountNumber)
	assert.Equal(suite.T(), "John Doe", account.AccountName)
}

func (suite *AccountRepoSuite) TestGetAccount_NotFound() {
	// Given: A failing HTTP response indicating account not found
	mockResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body: ioutil.NopCloser(bytes.NewBufferString(`{
			"documents": []
		}`)),
	}
	suite.client.Transport.(*MockRoundTripper).Response = mockResponse

	// When: GetAccount is called with a non-existent account number
	account, err := suite.repo.GetAccount("99999")

	// Then: An error should occur and no account should be returned
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), account)
	assert.Contains(suite.T(), err.Error(), "account not found")
}

func (suite *AccountRepoSuite) TestGetAccount_HTTPFailure() {
	// Given: An HTTP error occurs
	suite.client.Transport.(*MockRoundTripper).Err = errors.New("http request failed")

	// When: GetAccount is called
	account, err := suite.repo.GetAccount("12345")

	// Then: An error should occur and no account should be returned
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), account)
	assert.Contains(suite.T(), err.Error(), "http request failed")
}

// TestAccountRepoSuite runs the test suite.
func TestAccountRepoSuite(t *testing.T) {
	suite.Run(t, new(AccountRepoSuite))
}
