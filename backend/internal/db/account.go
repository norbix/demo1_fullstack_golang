package db

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/norbix/demo1_fullstack_golang/backend/configs"
	"github.com/norbix/demo1_fullstack_golang/backend/internal/db/dbmodels"
)

// AccountRepo interacts with the downstream immudb Vault API.
type AccountRepo struct {
	config *configs.Config
	client *http.Client
}

// NewAccountRepo initializes the AccountRepo with the given config and HTTP client.
// If no client is provided, it defaults to http.DefaultClient.
func NewAccountRepo(config *configs.Config, client *http.Client) *AccountRepo {
	if client == nil {
		client = http.DefaultClient
	}
	return &AccountRepo{
		config: config,
		client: client,
	}
}

//
// // CreateAccount sends account data to the immudb Vault for storage.
// func (repo *AccountRepo) CreateAccount(account dbmodels.Account) error {
// 	url := fmt.Sprintf("%s/document", repo.config.BaseURL)
// 	payload, err := json.Marshal(account)
// 	if err != nil {
// 		return fmt.Errorf("failed to serialize account data: %w", err)
// 	}
//
// 	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(payload))
// 	if err != nil {
// 		return fmt.Errorf("failed to create request: %w", err)
// 	}
// 	req.Header.Set("Content-Type", "application/json")
// 	req.Header.Set("accept", "application/json")
// 	req.Header.Set("X-API-Key", repo.config.APIKey)
//
// 	resp, err := repo.client.Do(req) // Use the injected HTTP client
// 	if err != nil {
// 		return fmt.Errorf("failed to send request: %w", err)
// 	}
// 	defer resp.Body.Close()
//
// 	if resp.StatusCode != http.StatusOK {
// 		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
// 	}
//
// 	return nil
// }
//
// // GetAccount retrieves account data from the immudb Vault by account number.
// func (repo *AccountRepo) GetAccount(accountNumber string) (*dbmodels.Account, error) {
// 	url := fmt.Sprintf("%s/documents/search", repo.config.BaseURL)
// 	query := map[string]interface{}{
// 		"query": map[string]interface{}{
// 			"account_number": accountNumber,
// 		},
// 	}
// 	payload, err := json.Marshal(query)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to serialize query: %w", err)
// 	}
//
// 	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to create request: %w", err)
// 	}
// 	req.Header.Set("Content-Type", "application/json")
// 	req.Header.Set("accept", "application/json")
// 	req.Header.Set("X-API-Key", repo.config.APIKey)
//
// 	resp, err := repo.client.Do(req) // Use the injected HTTP client
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to send request: %w", err)
// 	}
// 	defer resp.Body.Close()
//
// 	if resp.StatusCode != http.StatusOK {
// 		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
// 	}
//
// 	var response struct {
// 		Documents []dbmodels.Account `json:"documents"`
// 	}
// 	err = json.NewDecoder(resp.Body).Decode(&response)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to decode response: %w", err)
// 	}
//
// 	if len(response.Documents) == 0 {
// 		return nil, errors.New("account not found")
// 	}
//
// 	return &response.Documents[0], nil
// }

// CreateAccount sends account data to the immudb Vault for storage.
func (repo *AccountRepo) CreateAccount(account dbmodels.Account) error {
	url := fmt.Sprintf("%s/document", repo.config.BaseURL)
	payload, err := json.Marshal(account)
	if err != nil {
		return fmt.Errorf("failed to serialize account data: %w", err)
	}

	// Change the request to use POST instead of PUT
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("accept", "application/json")
	req.Header.Set("X-API-Key", repo.config.APIKey)

	resp, err := repo.client.Do(req) // Use the injected HTTP client
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

// GetAccount retrieves account data from the immudb Vault by account number using POST.
func (repo *AccountRepo) GetAccount(accountNumber string) (*dbmodels.Account, error) {
	url := fmt.Sprintf("%s/documents/search", repo.config.BaseURL)
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"account_number": accountNumber,
		},
	}
	payload, err := json.Marshal(query)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize query: %w", err)
	}

	// Use POST for retrieval as required
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("accept", "application/json")
	req.Header.Set("X-API-Key", repo.config.APIKey)

	resp, err := repo.client.Do(req) // Use the injected HTTP client
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var response struct {
		Documents []dbmodels.Account `json:"documents"`
	}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	if len(response.Documents) == 0 {
		return nil, errors.New("account not found")
	}

	return &response.Documents[0], nil
}
