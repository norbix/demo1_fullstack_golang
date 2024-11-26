package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/norbix/demo1_fullstack_golang/backend/internal/db/dbmodels"
)

// DBTestSuite defines the test suite for the db package.
type DBTestSuite struct {
	suite.Suite
	Repo    *AccountRepo // Shared repository for all tests
	Service *DBService   // Shared service for all tests
}

// SetupSuite runs once before any tests in the suite.
func (suite *DBTestSuite) SetupSuite() {
	// Initialize the repository and service
	suite.Repo = NewAccountRepo()
	suite.Service = NewDBService(suite.Repo)
}

// TearDownSuite runs once after all tests in the suite.
func (suite *DBTestSuite) TearDownSuite() {
	// Perform any necessary cleanup (not needed for in-memory data)
}

// SetupTest runs before each test in the suite.
func (suite *DBTestSuite) SetupTest() {
	// Reset the in-memory repository before each test
	suite.Repo.accounts = make(map[string]dbmodels.Account)
}

// TestCreateAccount tests the CreateAccount functionality.
func (suite *DBTestSuite) TestCreateAccount() {
	// Given: A valid account
	account := dbmodels.Account{
		AccountNumber: "12345",
		AccountName:   "John Doe",
		IBAN:          "DE89370400440532013000",
		Address:       "123 Elm St",
		Amount:        1000.0,
		Type:          dbmodels.Sending,
	}

	// When: The account is created
	err := suite.Service.Repo.CreateAccount(account)

	// Then: No error should occur, and the account should be retrievable
	assert.NoError(suite.T(), err, "CreateAccount should not return an error")

	retrievedAccount, err := suite.Service.Repo.GetAccount("12345")
	assert.NoError(suite.T(), err, "GetAccount should not return an error")
	assert.Equal(suite.T(), account, *retrievedAccount, "Retrieved account should match the created account")
}

// TestGetAccount_NotFound tests retrieving an account that doesn't exist.
func (suite *DBTestSuite) TestGetAccount_NotFound() {
	// When: Attempting to retrieve a non-existent account
	account, err := suite.Service.Repo.GetAccount("99999")

	// Then: An error should occur, and the account should be nil
	assert.Error(suite.T(), err, "GetAccount should return an error for non-existent account")
	assert.Nil(suite.T(), account, "Account should be nil for non-existent account")
}

// TestListAccounts tests retrieving all accounts.
func (suite *DBTestSuite) TestListAccounts() {
	// Given: Two accounts are created
	account1 := dbmodels.Account{
		AccountNumber: "12345",
		AccountName:   "John Doe",
		IBAN:          "DE89370400440532013000",
		Address:       "123 Elm St",
		Amount:        1000.0,
		Type:          dbmodels.Sending,
	}
	account2 := dbmodels.Account{
		AccountNumber: "67890",
		AccountName:   "Jane Doe",
		IBAN:          "FR7630006000011234567890189",
		Address:       "567 Oak St",
		Amount:        2000.0,
		Type:          dbmodels.Receiving,
	}

	_ = suite.Service.Repo.CreateAccount(account1)
	_ = suite.Service.Repo.CreateAccount(account2)

	// When: Listing all accounts
	accounts, err := suite.Service.Repo.ListAccounts()

	// Then: The correct number of accounts should be returned
	assert.NoError(suite.T(), err, "ListAccounts should not return an error")
	assert.Equal(suite.T(), 2, len(accounts), "ListAccounts should return the correct number of accounts")
	assert.Contains(suite.T(), accounts, account1, "ListAccounts should include account1")
	assert.Contains(suite.T(), accounts, account2, "ListAccounts should include account2")
}

// Run the test suite
func TestDBTestSuite(t *testing.T) {
	suite.Run(t, new(DBTestSuite))
}
