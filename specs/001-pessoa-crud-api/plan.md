# Implementation Plan: Person Management REST API

**Branch**: `001-pessoa-crud-api` | **Date**: 2025-12-14 | **Spec**: [spec.md](spec.md)
**Input**: Feature specification from `/specs/001-pessoa-crud-api/spec.md`

**Note**: This template is filled in by the `/speckit.plan` command. See `.specify/templates/commands/plan.md` for the execution workflow.

## Summary

Build a REST API for managing person records (CRUD operations) with nested addresses and contacts. The API will support creating, retrieving, listing (with pagination), updating, and deleting person records. Each person contains required fields (name, surname, sex, birthdate) and optional collections of addresses (street, city, state, country) and contacts (phone, email). The system must validate all inputs, ensure data integrity, and meet performance targets (<500ms create, <200ms read at P95). Implementation follows Clean Architecture principles using Go with Gin framework, MongoDB for persistence, containerized deployment via Docker, and environment-specific YAML configuration.

## Technical Context

**Language/Version**: Go 1.21+  
**Primary Dependencies**: 
- gin-gonic/gin v1.9+ (HTTP framework)
- mongodb/mongo-go-driver v1.13+ (MongoDB client)
- go-playground/validator/v10 (input validation)
- spf13/viper (YAML configuration management)
- uber-go/zap (structured logging)

**Storage**: MongoDB 6.0+ (document database for flexible person/address/contact schema)  
**Testing**: Go standard testing package + testify/assert for assertions, testcontainers-go for integration tests  
**Target Platform**: Linux containers (Docker), deployable to LOCAL/BETA/PROD environments  
**Project Type**: Single backend service (REST API server)  
**Performance Goals**: 
- P95 latency: <500ms for write operations (create, update, delete)
- P95 latency: <200ms for read operations (get by ID, list)
- Throughput: 100+ concurrent requests without degradation

**Constraints**: 
- Response time: <500ms P95 for creates, <200ms P95 for reads
- Pagination: Max 100 records per page, default 20
- Validation: RFC 5322 email, international phone formats
- Age limit: Max 150 years from birthdate

**Scale/Scope**: 
- Initial MVP: P1 user stories (Create + Query by ID)
- Phase 2: P2 user stories (List with pagination + Update)
- Phase 3: P3 user story (Delete with cascade)
- Expected load: 100 concurrent users, thousands of person records

## Constitution Check

*GATE: Must pass before Phase 0 research. Re-check after Phase 1 design.*

Verify compliance with constitution principles:

- [x] **Clean Code**: Code structure follows naming conventions, single responsibility, DRY
  - Go idioms enforced (gofmt, goimports, golangci-lint)
  - Functions <20 lines, meaningful names (e.g., `CreatePerson`, `ValidateEmail`)
  - No magic numbers (constants for page sizes, age limits, status codes)
  
- [x] **Clean Architecture**: Proper layer separation (Entities → Use Cases → Interface Adapters → Frameworks)
  - Domain layer: Person, Address, Contact entities + repository interfaces
  - Application layer: CreatePersonUseCase, GetPersonUseCase, etc.
  - Infrastructure layer: MongoDB repository implementations, Gin router
  - Presentation layer: HTTP handlers, DTOs, middleware
  
- [x] **Code Reusability**: Common logic extracted, no duplication, composition over inheritance
  - Shared validation package (email, phone, date validators)
  - Common error handling utilities
  - Reusable pagination logic
  - Composition via embedded structs (Go idiom)
  
- [x] **Design Patterns**: Appropriate patterns applied (Repository, Factory, Strategy, etc.)
  - Repository pattern: PersonRepository interface in domain, MongoPersonRepository in infrastructure
  - Factory pattern: Database connection factory, validator factory
  - Strategy pattern: Validation strategies for different field types
  - Dependency Injection: Constructor injection for all dependencies
  
- [x] **Function Cohesion**: Functions are small, focused, single-purpose with minimal parameters
  - Each handler handles one HTTP operation
  - Use cases perform single business operation
  - Parameter objects for complex inputs (CreatePersonRequest struct)
  
- [x] **Performance Standards**: Performance requirements defined, algorithms efficient, caching strategy, resource management
  - MongoDB indexes on person ID for O(1) lookups
  - Pagination prevents loading all records
  - Context with timeout for all database operations
  - Connection pooling configured
  - Defer for resource cleanup (database connections, file handles)
  
- [x] **Security Requirements**: Input validation, authentication/authorization, data encryption, injection prevention, no secrets in code
  - Input validation via go-playground/validator tags
  - MongoDB parameterized queries prevent NoSQL injection
  - YAML config for secrets (database credentials, API keys)
  - Environment-specific configs (LOCAL/BETA/PROD) loaded via Viper
  - TLS for MongoDB connections in PROD
  - No secrets in code or version control
  - Error messages don't expose internal details
  
- [x] **Testability**: Business logic independent of frameworks and external dependencies
  - Domain entities have no external dependencies
  - Use cases depend on interfaces, not concrete implementations
  - Mock repository implementations for unit tests
  - Testcontainers for integration tests with real MongoDB
  
- [x] **Quality Gates**: Linting, formatting, test coverage (80%+), documentation
  - golangci-lint enforced in CI
  - gofmt/goimports for formatting
  - Target 80%+ test coverage
  - GoDoc comments on exported types and functions

*Note: No violations - all constitution principles satisfied by design*

## Project Structure

### Documentation (this feature)

```text
specs/[###-feature]/
├── plan.md              # This file (/speckit.plan command output)
├── research.md          # Phase 0 output (/speckit.plan command)
├── data-model.md        # Phase 1 output (/speckit.plan command)
├── quickstart.md        # Phase 1 output (/speckit.plan command)
├── contracts/           # Phase 1 output (/speckit.plan command)
└── tasks.md             # Phase 2 output (/speckit.tasks command - NOT created by /speckit.plan)
```

### Source Code (repository root)

```text
cadastro-pessoas/
├── cmd/
│   └── api/
│       └── main.go              # Application entry point, wire up dependencies
├── internal/
│   ├── domain/                  # Entities layer (business rules)
│   │   ├── entities/
│   │   │   ├── person.go       # Person entity with validation rules
│   │   │   ├── address.go      # Address value object
│   │   │   └── contact.go      # Contact value object
│   │   └── repositories/
│   │       └── person_repository.go  # Repository interface
│   ├── application/             # Use Cases layer (application logic)
│   │   ├── usecases/
│   │   │   ├── create_person.go      # Create person use case
│   │   │   ├── get_person.go         # Get person by ID use case
│   │   │   ├── list_persons.go       # List persons with pagination
│   │   │   ├── update_person.go      # Update person use case
│   │   │   └── delete_person.go      # Delete person use case
│   │   └── dto/
│   │       ├── person_request.go     # Request DTOs
│   │       └── person_response.go    # Response DTOs
│   ├── infrastructure/          # Frameworks & Drivers layer
│   │   ├── persistence/
│   │   │   ├── mongodb/
│   │   │   │   ├── connection.go     # MongoDB connection factory
│   │   │   │   ├── person_repository.go  # MongoDB implementation of PersonRepository
│   │   │   │   └── indexes.go        # Index definitions and setup
│   │   │   └── migrations/
│   │   │       └── init_indexes.go   # Initial index setup
│   │   ├── config/
│   │   │   ├── config.go            # Viper-based YAML config loader
│   │   │   └── environments.go      # Environment-specific settings
│   │   └── logger/
│   │       └── logger.go            # Zap logger setup
│   └── presentation/            # Interface Adapters layer
│       ├── http/
│       │   ├── handlers/
│       │   │   ├── person_handler.go     # HTTP handlers for person endpoints
│       │   │   ├── health_handler.go     # Health check endpoint
│       │   │   └── error_handler.go      # Error response formatting
│       │   ├── middleware/
│       │   │   ├── logging.go           # Request logging middleware
│       │   │   ├── recovery.go          # Panic recovery middleware
│       │   │   └── validation.go        # Input validation middleware
│       │   └── router.go                # Gin router setup
│       └── dto/
│           ├── request.go              # HTTP request DTOs
│           ├── response.go             # HTTP response DTOs
│           └── pagination.go           # Pagination response structure
├── pkg/                         # Public reusable packages
│   ├── validator/
│   │   ├── email.go            # Email validation utilities
│   │   ├── phone.go            # Phone validation utilities
│   │   └── date.go             # Date validation utilities
│   └── errors/
│       ├── errors.go           # Custom error types
│       └── codes.go            # Error codes
├── tests/
│   ├── unit/                   # Unit tests (use cases, entities)
│   │   ├── usecases/
│   │   └── entities/
│   ├── integration/            # Integration tests (MongoDB, API)
│   │   ├── repository/
│   │   └── api/
│   └── mocks/                  # Mock implementations
│       └── person_repository_mock.go
├── configs/                    # Configuration files
│   ├── config.local.yaml       # Local development config
│   ├── config.beta.yaml        # Beta environment config
│   └── config.prod.yaml        # Production environment config
├── docker/
│   ├── Dockerfile              # Multi-stage build for production
│   └── docker-compose.yml      # Local development with MongoDB
├── scripts/
│   ├── setup-local.sh          # Local environment setup
│   └── run-tests.sh            # Test runner script
├── docs/
│   └── api/                    # API documentation (generated from contracts/)
├── .github/
│   └── workflows/
│       ├── ci.yml              # CI pipeline (lint, test, build)
│       └── cd.yml              # CD pipeline (deploy to environments)
├── go.mod
├── go.sum
├── Makefile                    # Build, test, run commands
└── README.md                   # Project documentation
```

**Structure Decision**: Selected Clean Architecture single-project structure as this is a backend REST API service. The structure follows Go conventions with `cmd/` for entry points, `internal/` for application code (not importable by external packages), and `pkg/` for reusable utilities. Clear layer separation ensures business logic (domain) is independent of HTTP framework (Gin) and database (MongoDB).

## Complexity Tracking

> **Fill ONLY if Constitution Check has violations that must be justified**

*No violations - all constitution principles are satisfied by the chosen architecture and technology stack. The Clean Architecture approach with Go, Gin, MongoDB, and Docker aligns with all requirements without requiring exceptions.*
