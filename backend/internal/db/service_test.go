package db

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/norbix/demo1_fullstack_golang/backend/internal/db/dbmodels"
)

// MockAccountRepo is a mock implementation of AccountRepo.
type MockAccountRepo struct {
	mock.Mock
}

// CreateAccount mocks the CreateAccount method.
func (m *MockAccountRepo) CreateAccount(account dbmodels.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

// GetAccount mocks the GetAccount method.
func (m *MockAccountRepo) GetAccount(accountNumber string) (*dbmodels.Account, error) {
	args := m.Called(accountNumber)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dbmodels.Account), args.Error(1)
}

func TestDBService_CreateAccount(t *testing.T) {
	// Initialize the mock repository
	mockRepo := new(MockAccountRepo)
	service := NewDBService(mockRepo)

	// Given: A valid account
	account := dbmodels.Account{
		AccountNumber: "12345",
		AccountName:   "John Doe",
		IBAN:          "DE89370400440532013000",
		Address:       "1234 Elm St",
		Amount:        1000.0,
		Type:          dbmodels.Sending,
	}

	// Mock the behavior of CreateAccount
	mockRepo.On("CreateAccount", account).Return(nil)

	// When: CreateAccount is called
	err := service.Repo.CreateAccount(account)

	// Then: Ensure no error is returned
	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "CreateAccount", account)
}

func TestDBService_GetAccount(t *testing.T) {
	// Initialize the mock repository
	mockRepo := new(MockAccountRepo)
	service := NewDBService(mockRepo)

	// Given: An account exists with a specific account number
	account := &dbmodels.Account{
		AccountNumber: "12345",
		AccountName:   "John Doe",
		IBAN:          "DE89370400440532013000",
		Address:       "1234 Elm St",
		Amount:        1000.0,
		Type:          dbmodels.Sending,
	}

	// Mock the behavior of GetAccount
	mockRepo.On("GetAccount", "12345").Return(account, nil)

	// When: GetAccount is called with the existing account number
	result, err := service.Repo.GetAccount("12345")

	// Then: Ensure the returned account matches the mock and no error occurs
	assert.NoError(t, err)
	assert.Equal(t, account, result)
	mockRepo.AssertCalled(t, "GetAccount", "12345")
}

func TestDBService_GetAccount_NotFound(t *testing.T) {
	// Initialize the mock repository
	mockRepo := new(MockAccountRepo)
	service := NewDBService(mockRepo)

	// Given: No account exists for the provided account number
	mockRepo.On("GetAccount", "99999").Return(nil, errors.New("account not found"))

	// When: GetAccount is called with a non-existent account number
	result, err := service.Repo.GetAccount("99999")

	// Then: Ensure an error is returned and the result is nil
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "account not found")
	mockRepo.AssertCalled(t, "GetAccount", "99999")
}
