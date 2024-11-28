package configs

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// ConfigTestSuite defines the test suite for the Config package
type ConfigTestSuite struct {
	suite.Suite
	originalBaseURL string
	originalAPIKey  string
}

// SetupSuite runs once before the tests are run
func (suite *ConfigTestSuite) SetupSuite() {
	// Store the original values of the environment variables
	suite.originalBaseURL = os.Getenv("BASE_URL")
	suite.originalAPIKey = os.Getenv("API_KEY")
}

// TearDownSuite runs once after all tests in the suite are run
func (suite *ConfigTestSuite) TearDownTest() {
	// Restore the original values of the environment variables
	err := os.Setenv("BASE_URL", suite.originalBaseURL)
	if err != nil {
		suite.Fail("Unable to restore original BASE_URL")
	}

	err = os.Setenv("API_KEY", suite.originalAPIKey)
	if err != nil {
		suite.Fail("Unable to restore original API_KEY")
	}
}

// SetupTest runs before each test in the suite
func (suite *ConfigTestSuite) SetupTest() {
	// Clear the environment variables before each test
	err := os.Unsetenv("BASE_URL")
	if err != nil {
		suite.Fail("Unable to unset BASE_URL")
	}

	err = os.Unsetenv("API_KEY")
	if err != nil {
		suite.Fail("Unable to unset API_KEY")
	}
}

// TestLoadConfig_OK tests the LoadConfig function when the environment variables are properly set
func (suite *ConfigTestSuite) TestLoadConfig_OK() {
	// Given: Environment variables are properly set
	err := os.Setenv("BASE_URL", "http://example.com")
	if err != nil {
		suite.Fail("Unable to set BASE_URL")
	}

	err = os.Setenv("API_KEY", "test-api-key")
	if err != nil {
		suite.Fail("Unable to set API_KEY")
	}

	// When: LoadConfig is called
	config, err := LoadConfig()

	// Then: It should return the expected Config struct with no error
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), config)
	assert.Equal(suite.T(), "http://example.com", config.BaseURL)
	assert.Equal(suite.T(), "test-api-key", config.APIKey)
}

// TestLoadConfig_MissingAPIKey tests the LoadConfig function when the API_KEY environment variable is missing
func (suite *ConfigTestSuite) TestLoadConfig_MissingBaseURL() {
	// Given: BASE_URL is not set
	err := os.Setenv("API_KEY", "test-api-key")
	if err != nil {
		suite.Fail("Unable to set API_KEY")
	}

	// When: LoadConfig is called
	config, err := LoadConfig()

	// Then: It should return an error
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), config)
	assert.Equal(suite.T(), "missing required environment variables", err.Error())
}

// TestLoadConfig_MissingAPIKey tests LoadConfig function when the API_KEY environment variable is missing
func (suite *ConfigTestSuite) TestLoadConfig_MissingAPIKey() {
	// Given: API_KEY is not set
	err := os.Setenv("BASE_URL", "http://example.com")
	if err != nil {
		suite.Fail("Unable to set BASE_URL")
	}

	// When: LoadConfig is called
	config, err := LoadConfig()

	// Then: It should return an error
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), config)
	assert.Equal(suite.T(), "missing required environment variables", err.Error())
}

// TestLoadConfig_NoEnvVariables tests LoadConfig function when no environment variables are set
func (suite *ConfigTestSuite) TestLoadConfig_NoEnvVariables() {
	// Given: No environment variables are set

	// When: LoadConfig is called
	config, err := LoadConfig()

	// Then: It should return an error
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), config)
	assert.Equal(suite.T(), "missing required environment variables", err.Error())
}

// Test Suite Runner
func TestConfigTestSuite(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}
