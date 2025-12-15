.PHONY: help build run test lint docker-up docker-down clean tidy fmt vet coverage

# Variables
BINARY_NAME=cadastro-pessoas
MAIN_PATH=./cmd/api
DOCKER_COMPOSE=docker compose

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'

build: ## Build the application binary
	@echo "Building $(BINARY_NAME)..."
	@go build -o bin/$(BINARY_NAME) $(MAIN_PATH)

run: ## Run the application locally
	@echo "Running $(BINARY_NAME)..."
	@go run $(MAIN_PATH)

test: ## Run all tests
	@echo "Running tests..."
	@go test -v ./...

test-coverage: ## Run tests with coverage
	@echo "Running tests with coverage..."
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

test-unit: ## Run only unit tests
	@echo "Running unit tests..."
	@go test -v ./tests/unit/...

test-integration: ## Run only integration tests
	@echo "Running integration tests..."
	@go test -v ./tests/integration/...

lint: ## Run golangci-lint
	@echo "Running linter..."
	@golangci-lint run

fmt: ## Format code with gofmt
	@echo "Formatting code..."
	@gofmt -s -w .

vet: ## Run go vet
	@echo "Running go vet..."
	@go vet ./...

tidy: ## Tidy go modules
	@echo "Tidying modules..."
	@go mod tidy

clean: ## Clean build artifacts
	@echo "Cleaning..."
	@rm -rf bin/
	@rm -f coverage.out coverage.html

docker-up: ## Start Docker services (MongoDB)
	@echo "Starting Docker services..."
	@$(DOCKER_COMPOSE) up -d

docker-down: ## Stop Docker services
	@echo "Stopping Docker services..."
	@$(DOCKER_COMPOSE) down

docker-logs: ## Show Docker logs
	@$(DOCKER_COMPOSE) logs -f

docker-build: ## Build Docker image
	@echo "Building Docker image..."
	@docker build -t $(BINARY_NAME):latest .

docker-run: ## Run application in Docker
	@echo "Running in Docker..."
	@docker run --rm -p 8080:8080 --env-file .env $(BINARY_NAME):latest

dev: ## Run with hot reload (requires air)
	@echo "Starting development server with hot reload..."
	@air

install-tools: ## Install development tools
	@echo "Installing development tools..."
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@go install github.com/air-verse/air@latest

.DEFAULT_GOAL := help
