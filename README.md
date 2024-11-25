# demo1_fullstack_golang

Demo application containing fullstack solution in pure Golang.

## Codebase structure

```text
/demo1_fullstack_golang
├── /backend
│   ├── /cmd
│   │   ├── /api
│   │   │   └── main.go        # Main entry point for the backend service
│   ├── /internal              # Internal packages (business logic, database, etc.)
│   │   ├── /auth
│   │   ├── /db
│   │   └── /services
│   ├── /pkg                   # Shared reusable libraries
│   ├── /configs               # Configuration files
│   ├── /docs                  # API documentation
│   └── go.mod
│
├── /frontend
│   ├── /assets                # Static assets (CSS, images, etc.)
│   ├── /cmd
│   │   └── /ui
│   │       └── main.go        # Entry point for WebAssembly-based frontend
│   ├── /internal              # Frontend-specific logic
│   │   └── /components        # UI components
│   ├── /pkg                   # Shared reusable frontend libraries
│   ├── /templates             # HTML templates if using SSR
│   └── go.mod
│
├── /scripts                   # Automation scripts (e.g., build, deploy)
│   ├── build.sh
├── /docker                    # Docker configuration for backend and frontend
│   ├── /backend
│   │   └── Dockerfile
│   ├── /frontend
│   │   └── Dockerfile
│   └── docker-compose.yml
├── /configs                   # Shared configuration files
├── /tests                     # End-to-end or integration tests
└── README.md
```
