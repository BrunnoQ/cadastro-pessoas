# Quick Start Guide: Person Management REST API

**Feature**: Person Management REST API  
**Branch**: `001-pessoa-crud-api`  
**Date**: 2025-12-14

## Prerequisites

Before you begin, ensure you have the following installed:

- **Go**: Version 1.21 or higher ([download](https://go.dev/dl/))
- **Docker**: Version 20.10 or higher ([download](https://docs.docker.com/get-docker/))
- **Docker Compose**: Version 2.0 or higher (usually included with Docker Desktop)
- **Git**: For cloning the repository
- **Make**: For running build commands (optional but recommended)
- **curl** or **Postman**: For testing API endpoints

Verify installations:

```bash
go version        # Should show 1.21+
docker --version  # Should show 20.10+
docker compose version  # Should show 2.0+
```

## Local Development Setup

### 1. Clone Repository

```bash
git clone https://github.com/BrunnoQ/cadastro-pessoas.git
cd cadastro-pessoas
git checkout 001-pessoa-crud-api
```

### 2. Install Dependencies

```bash
# Download Go modules
go mod download

# Verify dependencies
go mod verify
```

### 3. Start MongoDB

Using Docker Compose (recommended):

```bash
# Start MongoDB in background
docker compose up -d mongodb

# Verify MongoDB is running
docker compose ps
```

MongoDB will be available at `mongodb://localhost:27017`

**Alternative - Manual Docker**:

```bash
docker run -d \
  --name cadastro-pessoas-mongo \
  -p 27017:27017 \
  -e MONGO_INITDB_DATABASE=cadastro_pessoas_local \
  mongo:6.0
```

### 4. Configure Application

The application uses YAML configuration files in the `configs/` directory.

**For Local Development** (configs/config.local.yaml):

```yaml
server:
  port: 8080
  environment: LOCAL
  
database:
  uri: mongodb://localhost:27017
  name: cadastro_pessoas_local
  timeout: 10s
  pool:
    min_size: 5
    max_size: 100
  
logging:
  level: debug
  format: json
  
pagination:
  default_size: 20
  max_size: 100

validation:
  max_name_length: 100
  max_age_years: 150
```

**Set Environment Variable**:

```bash
# Linux/Mac
export APP_ENV=local

# Windows (PowerShell)
$env:APP_ENV="local"

# Windows (CMD)
set APP_ENV=local
```

### 5. Run Application

**Option A: Using Make** (recommended):

```bash
# Run with hot reload
make run

# Or build and run
make build
./bin/cadastro-pessoas
```

**Option B: Direct Go Command**:

```bash
# Run directly
go run ./cmd/api

# Or build first
go build -o bin/cadastro-pessoas ./cmd/api
./bin/cadastro-pessoas
```

**Option C: Using Docker** (full containerization):

```bash
# Build and run everything (app + MongoDB)
docker compose up --build

# Run in background
docker compose up -d --build
```

### 6. Verify Application is Running

```bash
# Health check
curl http://localhost:8080/api/v1/health

# Expected response:
# {
#   "status": "healthy",
#   "version": "1.0.0",
#   "timestamp": "2025-12-14T10:30:00Z",
#   "dependencies": {
#     "database": "connected"
#   }
# }
```

## Quick API Tour

### Create a Person

```bash
curl -X POST http://localhost:8080/api/v1/persons \
  -H "Content-Type: application/json" \
  -d '{
    "name": "João",
    "surname": "Silva",
    "sex": "masculino",
    "birthdate": "1990-01-15",
    "addresses": [
      {
        "street": "Rua das Flores, 123",
        "city": "São Paulo",
        "state": "SP",
        "country": "Brasil"
      }
    ],
    "contacts": [
      {
        "phone": "+5511999999999",
        "email": "joao.silva@example.com"
      }
    ]
  }'

# Response (201 Created):
# {
#   "id": "507f1f77bcf86cd799439011",
#   "name": "João",
#   "surname": "Silva",
#   "sex": "masculino",
#   "birthdate": "1990-01-15",
#   "addresses": [...],
#   "contacts": [...],
#   "created_at": "2025-12-14T10:30:00Z",
#   "updated_at": "2025-12-14T10:30:00Z"
# }
```

**Save the `id` from the response** - you'll need it for the next requests!

### Get Person by ID

```bash
# Replace {id} with the actual ID from create response
curl http://localhost:8080/api/v1/persons/507f1f77bcf86cd799439011

# Response (200 OK):
# {
#   "id": "507f1f77bcf86cd799439011",
#   "name": "João",
#   ...full person data...
# }
```

### List All Persons

```bash
# First page with default size (20)
curl http://localhost:8080/api/v1/persons

# With pagination parameters
curl "http://localhost:8080/api/v1/persons?page=1&pageSize=10"

# Response (200 OK):
# {
#   "data": [
#     {
#       "id": "507f1f77bcf86cd799439011",
#       "name": "João",
#       "surname": "Silva",
#       "sex": "masculino",
#       "birthdate": "1990-01-15",
#       "address_count": 1,
#       "contact_count": 1
#     }
#   ],
#   "pagination": {
#     "page": 1,
#     "page_size": 20,
#     "total_records": 1,
#     "total_pages": 1
#   }
# }
```

### Update Person

```bash
curl -X PUT http://localhost:8080/api/v1/persons/507f1f77bcf86cd799439011 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "João",
    "surname": "Silva Santos",
    "sex": "masculino",
    "birthdate": "1990-01-15",
    "addresses": [
      {
        "street": "Nova Avenida, 456",
        "city": "Rio de Janeiro",
        "state": "RJ",
        "country": "Brasil"
      }
    ],
    "contacts": [
      {
        "phone": "+5521987654321",
        "email": "joao.silva.santos@example.com"
      }
    ]
  }'

# Response (200 OK): Updated person data
```

### Delete Person

```bash
curl -X DELETE http://localhost:8080/api/v1/persons/507f1f77bcf86cd799439011

# Response (200 OK):
# {
#   "message": "Person deleted successfully"
# }
```

## Testing

### Run Unit Tests

```bash
# Run all tests
go test ./...

# Run with verbose output
go test -v ./...

# Run specific package tests
go test ./internal/domain/...
go test ./internal/application/...
go test ./internal/infrastructure/...
```

### Run Tests with Coverage

```bash
# Generate coverage report
go test -coverprofile=coverage.out ./...

# View coverage percentage
go tool cover -func=coverage.out

# View coverage in browser (HTML report)
go tool cover -html=coverage.out
```

### Run Integration Tests

Integration tests use Testcontainers to spin up real MongoDB:

```bash
# Run integration tests (requires Docker)
go test -tags=integration ./tests/integration/...

# Skip integration tests
go test -short ./...
```

### Run Linter

```bash
# Install golangci-lint (if not already installed)
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Run linter
golangci-lint run

# Auto-fix some issues
golangci-lint run --fix
```

## Development Workflow

### 1. Hot Reload with Air

For automatic recompilation on file changes:

```bash
# Install Air
go install github.com/cosmtrek/air@latest

# Run with hot reload
air
```

Configuration in `.air.toml` (auto-generated on first run).

### 2. Code Formatting

```bash
# Format all Go files
gofmt -s -w .

# Format imports
go install golang.org/x/tools/cmd/goimports@latest
goimports -w .
```

### 3. Database Management

**View MongoDB Collections**:

```bash
# Connect to MongoDB container
docker exec -it cadastro-pessoas-mongo mongosh

# Inside mongosh:
use cadastro_pessoas_local
db.persons.find().pretty()
db.persons.countDocuments()
```

**Reset Database**:

```bash
# Drop database and start fresh
docker exec -it cadastro-pessoas-mongo mongosh cadastro_pessoas_local --eval "db.dropDatabase()"
```

### 4. View Logs

**Application Logs**:

```bash
# If running with Docker Compose
docker compose logs -f api

# View MongoDB logs
docker compose logs -f mongodb
```

**Log Format** (JSON structured logging):

```json
{
  "level": "info",
  "ts": "2025-12-14T10:30:00.123Z",
  "msg": "Request completed",
  "method": "POST",
  "path": "/api/v1/persons",
  "status": 201,
  "duration": "45ms"
}
```

## API Documentation

### Interactive Documentation

**Option 1: Swagger UI** (recommended):

```bash
# Install Swagger UI locally
docker run -p 8081:8080 \
  -e SWAGGER_JSON=/api/openapi.yaml \
  -v $(pwd)/specs/001-pessoa-crud-api/contracts:/api \
  swaggerapi/swagger-ui

# Open browser: http://localhost:8081
```

**Option 2: Redoc**:

```bash
docker run -p 8082:80 \
  -e SPEC_URL=/specs/openapi.yaml \
  -v $(pwd)/specs/001-pessoa-crud-api/contracts:/usr/share/nginx/html/specs \
  redocly/redoc

# Open browser: http://localhost:8082
```

### OpenAPI Specification

OpenAPI 3.0 specification available at:

```
specs/001-pessoa-crud-api/contracts/openapi.yaml
```

**Generate Client Libraries**:

```bash
# Install OpenAPI Generator
npm install -g @openapitools/openapi-generator-cli

# Generate Go client
openapi-generator-cli generate \
  -i specs/001-pessoa-crud-api/contracts/openapi.yaml \
  -g go \
  -o clients/go

# Generate TypeScript client
openapi-generator-cli generate \
  -i specs/001-pessoa-crud-api/contracts/openapi.yaml \
  -g typescript-axios \
  -o clients/typescript
```

## Environment-Specific Configuration

### Local Development (LOCAL)

- Config: `configs/config.local.yaml`
- Database: `cadastro_pessoas_local`
- Log Level: `debug`
- Port: `8080`

### Beta Environment (BETA)

- Config: `configs/config.beta.yaml`
- Database: Remote MongoDB (beta cluster)
- Log Level: `info`
- Port: `8080`
- TLS: Enabled

### Production Environment (PROD)

- Config: `configs/config.prod.yaml`
- Database: Remote MongoDB (production cluster)
- Log Level: `warn`
- Port: `8080`
- TLS: Enabled
- Connection Pool: Larger (min: 10, max: 200)

**Switch Environments**:

```bash
# Set environment variable
export APP_ENV=beta  # or prod

# Application automatically loads configs/config.{APP_ENV}.yaml
```

## Makefile Commands

Quick reference for common tasks:

```bash
# Development
make run           # Run application with hot reload
make build         # Build binary to bin/
make clean         # Remove build artifacts

# Testing
make test          # Run all tests
make test-coverage # Run tests with coverage report
make test-integration  # Run integration tests only

# Code Quality
make lint          # Run linter
make fmt           # Format code

# Docker
make docker-build  # Build Docker image
make docker-up     # Start all containers
make docker-down   # Stop all containers
make docker-logs   # View container logs

# Database
make db-up         # Start MongoDB container
make db-down       # Stop MongoDB container
make db-reset      # Drop and recreate database

# Documentation
make docs          # Generate API documentation
make swagger       # Start Swagger UI

# All-in-one
make setup         # Initial setup (deps, db, config)
make dev           # Start development environment (db + app with hot reload)
```

## Troubleshooting

### Port Already in Use

```bash
# Check what's using port 8080
lsof -i :8080  # Mac/Linux
netstat -ano | findstr :8080  # Windows

# Kill the process or change port in config
```

### MongoDB Connection Failed

```bash
# Verify MongoDB is running
docker ps | grep mongo

# Check MongoDB logs
docker logs cadastro-pessoas-mongo

# Restart MongoDB
docker compose restart mongodb
```

### Module Download Issues

```bash
# Clear module cache
go clean -modcache

# Re-download dependencies
go mod download
```

### Tests Failing

```bash
# Run tests with verbose output to see details
go test -v ./...

# Run specific test
go test -v -run TestCreatePerson ./internal/application/...

# Check if Docker is running (for integration tests)
docker info
```

### Application Won't Start

Check:

1. Is MongoDB running? `docker ps`
2. Is port 8080 available? `lsof -i :8080`
3. Is config file present? `ls configs/config.local.yaml`
4. Is APP_ENV set correctly? `echo $APP_ENV`

View application logs for specific error message.

## Next Steps

1. **Read the Documentation**:
   - [Feature Specification](spec.md) - What the API does
   - [Implementation Plan](plan.md) - Technical architecture
   - [Data Model](data-model.md) - Entity structures
   - [API Contracts](contracts/openapi.yaml) - Detailed API spec

2. **Explore the Code**:
   - Domain entities: `internal/domain/entities/`
   - Use cases: `internal/application/usecases/`
   - HTTP handlers: `internal/presentation/http/handlers/`
   - Repository implementation: `internal/infrastructure/persistence/mongodb/`

3. **Run Tests**:
   - Unit tests: `go test ./internal/domain/... ./internal/application/...`
   - Integration tests: `go test ./tests/integration/...`

4. **Experiment with API**:
   - Use Swagger UI for interactive testing
   - Try error cases (invalid data, missing fields)
   - Test pagination with large datasets

5. **Start Development**:
   - Follow TDD: Write tests first, then implementation
   - Follow constitution principles (see `.specify/memory/constitution.md`)
   - Use feature branch workflow
   - Run linter before committing

## Support & Resources

- **Constitution**: `.specify/memory/constitution.md` - Coding principles
- **AGENTS.md**: `AGENTS.md` - Operational guidelines for Go development
- **Gin Framework**: <https://gin-gonic.com/docs/>
- **MongoDB Go Driver**: <https://www.mongodb.com/docs/drivers/go/current/>
- **Go Testing**: <https://go.dev/doc/tutorial/add-a-test>

## Questions?

Create an issue in the repository or contact the development team.
