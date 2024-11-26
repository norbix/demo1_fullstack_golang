package configs

import (
	"errors"
	"os"
)

type Config struct {
	BaseURL string
	APIKey  string
}

func LoadConfig() (*Config, error) {
	baseURL := os.Getenv("CN_BASE_URL")
	apiKey := os.Getenv("CN_API_KEY")

	if baseURL == "" || apiKey == "" {
		return nil, errors.New("missing required environment variables")
	}

	return &Config{
		BaseURL: baseURL,
		APIKey:  apiKey,
	}, nil
}
