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
├── /tests                     # End-to-end or integration tests
└── README.md
```

## Setup

### Prerequisites

- Go SDK 1.23 or higher from [https://golang.org/dl/](https://golang.org/dl/)
- Docker from [https://docs.docker.com/get-docker/](https://docs.docker.com/get-docker/)
- Taskfile from [https://taskfile.dev](https://taskfile.dev)
- docker-compose from [https://docs.docker.com/compose/install/](https://docs.docker.com/compose/install/)

### Installation

1. Clone the repository:

```bash
git clone
```

2. Install dependencies:

```bash
go mod download
```

3. Start the backend service:

    ```text
    go run ./backend/cmd/api
    ```
    
    The backend service will start on `http://localhost:8080`.
    
    ### Environment variables
    
    | Variable  | Description                     | Default Value                                                                |
    |-----------|---------------------------------|------------------------------------------------------------------------------|
    | BASE_URL  | Base URL for downstream services | `https://vault.immudb.io/ics/api/v1/ledger/default/collection/default`      |
    | API_KEY   | API key for authentication       | `your-api-key`                                                              |
    | SKIP_TLS  | Skip TLS verification (true/false) | `false`                                                                     |
    
    ```text
    export BASE_URL="https://vault.immudb.io/ics/api/v1/ledger/default/collection/default"
    export API_KEY="<replace>"
    export SKIP_TLS="false"
    ```

    ### Endpoints
    
    | **Endpoint**            | **Method** | **Description**                                | **Request Body**                                                                                               | **Response**                          |
    |--------------------------|------------|------------------------------------------------|---------------------------------------------------------------------------------------------------------------|---------------------------------------|
    | `/healthz`              | GET        | Health check to verify the service is running | None                                                                                                          | `200 OK`: `"Backend is healthy!"`    |
    | `/swagger/`             | GET        | Access Swagger documentation                  | None                                                                                                          | Swagger UI                            |
    | `/accounts`             | PUT        | Create a new account                          | ```json { "id": "string", "name": "string", "email": "string" } ```                                           | `200 OK`: Account created            |
    | `/accounts/retrieve`    | POST       | Retrieve accounts with pagination             | ```json { "page": 1, "perPage": 100 } ```                                                                     | `200 OK`: List of accounts            |

## Manual E2E Testing

Please use `IntelliJ IDEA` or any other REST client to test the API endpoints manually.

Code is available in the [http-client](backend/docs/http-client) directory.
    
    
## Development Workflow

### Backend

1. Use the following developer workflows in `Taskfile`

    ```text
    task
    ```
   
    Example output:
    
    ```text
    task: [default] task --list
    task: Available tasks for this project:
    * clean:        Clean up unused Docker resources
    * clean-all:    Delete all Docker objects including images, containers, volumes, and networks
    * default:      Show available Docker tasks
    * down:         Stop all services with Docker Compose
    * logs:         View logs from Docker Compose services
    * rebuild:      Rebuild and restart services with Docker Compose
    * up:           Start all services with Docker Compose
      ```   

### Frontend

    WIP
