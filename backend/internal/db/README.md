# Info

Package directory structure:

```text
internal/db/
├── account.go          # Implements the AccountService interface
├── account_test.go     # Unit tests for account.go
├── dbmodels/           # Holds shared data models
│   └── account.go      # Defines the Account struct
├── service.go          # Defines the interface and a struct for database services
├── service_test.go     # Unit tests for service.go
└── suite_test.go       # Test suite setup
```
