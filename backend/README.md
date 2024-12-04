## Development Workflow

### Backend

1. Use the following developer workflows in `Taskfile`

    ```text
    task
    ```

   Example output:

    ```text
   ==
   Tasks available 4 this infra KUBE.
   
   task: Available tasks for this project:
   * build:compile:                Compile the Go binary
   * build:default:                Show available tasks
   * build:mockery:clean:          Delete all generated mocks in the 'internal/mocks/' directory
   * build:mockery:generate:       Generate mocks for all interfaces, placing them in 'internal/mocks' directory
   * build:mockery:install:        Install mockery tool for generating mocks
   * build:run:                    Run the Go server
   * build:swagger:clean:          Delete all generated contracts in docs directory.
   * build:swagger:init:           Generate Swagger documentation
   * build:swagger:install:        Install Swagger CLI tool
   * build:test:                   Run tests for the project
   * default:                      List all commands defined.
    ```   
   