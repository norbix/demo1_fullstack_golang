package configs

import (
	"errors"
	"os"
	"strconv"
)

type Config struct {
	BaseURL string
	APIKey  string
	SkipTLS bool
}

func LoadConfig() (*Config, error) {
	baseURL := os.Getenv("BASE_URL")
	apiKey := os.Getenv("API_KEY")
	skipTLS := os.Getenv("SKIP_TLS")

	if baseURL == "" || apiKey == "" {
		return nil, errors.New("missing required environment variables")
	}

	// Convert skipTLS to a boolean
	skipTLSBool, err := strconv.ParseBool(skipTLS)
	if err != nil && skipTLS != "" {
		return nil, errors.New("invalid value for SKIP_TLS, must be true or false")
	}

	return &Config{
		BaseURL: baseURL,
		APIKey:  apiKey,
		SkipTLS: skipTLSBool,
	}, nil
}
