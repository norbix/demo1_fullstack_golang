package db

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/norbix/demo1_fullstack_golang/backend/internal/db/dbmodels"
)

// CreateAccount sends account data to the immudb Vault for storage.
func (repo accountRepositoryImpl) CreateAccount(account dbmodels.Account) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/document", repo.config.BaseURL)

	// Serialize account data
	payload, err := json.Marshal(account)
	if err != nil {
		log.Printf("[ERROR] Failed to serialize account data: %v", err)
		return nil, fmt.Errorf("failed to serialize account data: %w", err)
	}

	log.Printf("[INFO] Sending request to immudb Vault. URL: %s, Payload: %s", url, string(payload))

	// Create the HTTP request
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(payload))
	if err != nil {
		log.Printf("[ERROR] Failed to create HTTP request: %v", err)
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("accept", "application/json")
	req.Header.Set("X-API-Key", repo.config.APIKey)

	// Send the request using the injected HTTP client
	resp, err := repo.client.Do(req)
	if err != nil {
		log.Printf("[ERROR] Failed to send request to immudb Vault: %v", err)
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	log.Printf("[INFO] Response Status Code: %d", resp.StatusCode)

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("[ERROR] Failed to read response body: %v", err)
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	log.Printf("[INFO] Raw response from immudb Vault: %s", string(body))

	// Check for non-OK status code
	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		log.Printf("[ERROR] Unexpected status code: %d, Response: %s", resp.StatusCode, string(body))
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Parse the response body into a map
	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		log.Printf("[ERROR] Failed to decode response body: %v", err)
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	log.Printf("[INFO] Account successfully created in immudb Vault. Response: %v", response)

	return response, nil
}

// GetAccounts retrieves a list of accounts from the immudb Vault.
func (repo accountRepositoryImpl) GetAccounts(page, perPage int) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/documents/search", repo.config.BaseURL)
	query := map[string]interface{}{
		"page":    page,
		"perPage": perPage,
	}

	payload, err := json.Marshal(query)
	if err != nil {
		log.Printf("[ERROR] Failed to serialize query: %v", err)
		return nil, fmt.Errorf("failed to serialize query: %w", err)
	}

	log.Printf("[INFO] Sending request to immudb Vault. URL: %s, Payload: %s", url, string(payload))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		log.Printf("[ERROR] Failed to create request: %v", err)
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("accept", "application/json")
	req.Header.Set("X-API-Key", repo.config.APIKey)

	resp, err := repo.client.Do(req)
	if err != nil {
		log.Printf("[ERROR] Failed to send request to immudb Vault: %v", err)
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	log.Printf("[INFO] Response Status Code: %d", resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("[ERROR] Failed to read response body: %v", err)
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	log.Printf("[INFO] Raw response from immudb Vault: %s", string(body))

	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		log.Printf("[ERROR] Failed to decode response: %v", err)
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return response, nil
}
