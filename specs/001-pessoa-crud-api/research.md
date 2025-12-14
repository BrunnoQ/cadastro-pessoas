# Research & Design Decisions: Person Management REST API

**Feature**: Person Management REST API  
**Branch**: `001-pessoa-crud-api`  
**Date**: 2025-12-14

## Purpose

This document captures all technical research, design decisions, and rationale for the Person Management REST API implementation. It serves as the authoritative reference for understanding why specific technologies, patterns, and approaches were selected.

## Technology Stack Decisions

### 1. Programming Language: Go 1.21+

**Decision**: Use Go (Golang) version 1.21 or higher as the primary programming language.

**Rationale**:
- **Performance**: Go provides excellent performance with compiled binaries, efficient memory management, and built-in concurrency via goroutines - critical for meeting <500ms P95 latency targets
- **Simplicity**: Go's simple syntax and standard library align with Clean Code principles (readability, maintainability)
- **Strong Typing**: Static typing catches errors at compile time, reducing runtime bugs
- **Standard Tooling**: Built-in testing framework, formatting (gofmt), and package management (go mod)
- **Container-Friendly**: Produces small, statically-linked binaries ideal for Docker containers
- **REST API Ecosystem**: Mature HTTP frameworks and libraries available

**Alternatives Considered**:
- **Python**: Rejected due to slower runtime performance (GIL limitations), would struggle with 100 concurrent request requirement
- **Node.js**: Rejected due to JavaScript's dynamic typing increasing bug risk, less robust for CPU-intensive operations
- **Java**: Rejected due to larger memory footprint, slower startup times, more complex dependency management

### 2. HTTP Framework: Gin-Gonic v1.9+

**Decision**: Use Gin-Gonic framework for HTTP routing and request handling.

**Rationale**:
- **Performance**: Gin is one of the fastest Go HTTP frameworks (~40x faster than standard net/http with martini-like routing)
- **Middleware Support**: Built-in middleware chain for logging, recovery, CORS - aligns with security requirements
- **JSON Handling**: Excellent JSON binding and validation support - critical for REST API
- **Context Management**: Gin's context handling integrates well with Go's context package for timeouts
- **Active Community**: Well-maintained, extensive documentation, large ecosystem
- **Clean Architecture Compatible**: Doesn't force architectural decisions, can be isolated in presentation layer

**Alternatives Considered**:
- **Echo**: Similar performance, but Gin has larger community and more middleware options
- **Chi**: More minimalist, but less out-of-box features (would require more custom middleware)
- **Standard net/http**: Too low-level, would require significant boilerplate for routing, middleware

### 3. Database: MongoDB 6.0+

**Decision**: Use MongoDB as the primary data store.

**Rationale**:
- **Schema Flexibility**: Document model naturally represents person with nested addresses/contacts arrays - no complex JOIN operations needed
- **Performance**: Excellent read performance with proper indexing, meets <200ms P95 read latency target
- **Scalability**: Horizontal scaling via sharding supports future growth beyond initial 100 concurrent users
- **Nested Documents**: Native support for embedded arrays (addresses, contacts) eliminates N+1 query problems
- **Go Driver**: Official mongo-go-driver is mature, well-documented, supports all MongoDB features
- **Development Speed**: Schema-less design allows rapid iteration during development

**Alternatives Considered**:
- **PostgreSQL**: Rejected because relational model requires 3 tables (persons, addresses, contacts) with JOINs, complicating queries and potentially impacting performance. JSONB columns would work but lose relational benefits.
- **MySQL**: Similar issues to PostgreSQL, slightly less mature JSON support
- **DynamoDB**: Rejected due to complexity of nested querying, AWS vendor lock-in, cost unpredictability

**Schema Design Decision**:
```javascript
// Embedded document approach (chosen)
{
  _id: ObjectId("..."),
  name: "João",
  surname: "Silva",
  sex: "masculino",
  birthdate: ISODate("1990-01-15"),
  addresses: [
    {street: "Rua A", city: "São Paulo", state: "SP", country: "Brasil"}
  ],
  contacts: [
    {phone: "+5511999999999", email: "joao@example.com"}
  ],
  createdAt: ISODate("..."),
  updatedAt: ISODate("...")
}
```

**Rationale for Embedded**: 
- Addresses and contacts have no meaning without a person (dependent lifecycle)
- Typical queries always fetch person with full addresses/contacts (no need for separate queries)
- Update patterns involve modifying entire person document (atomic updates)
- Bounded collections: reasonable maximum of ~10 addresses, ~10 contacts per person

### 4. Configuration Management: Viper

**Decision**: Use Viper library for loading YAML configuration files with environment-specific overrides.

**Rationale**:
- **Multi-Format Support**: Reads YAML, JSON, TOML, environment variables - flexible for different deployment scenarios
- **Environment-Specific Config**: Easy pattern for config.local.yaml, config.beta.yaml, config.prod.yaml
- **Type-Safe Access**: Strongly-typed config value retrieval with defaults
- **Hot Reload**: Supports config file watching for runtime updates (useful for feature flags, rate limits)
- **Standard in Go**: De facto standard for configuration in Go applications

**Configuration Structure**:
```yaml
# config.local.yaml
server:
  port: 8080
  environment: "LOCAL"
  
database:
  uri: "mongodb://localhost:27017"
  name: "cadastro_pessoas_local"
  timeout: 10s
  
logging:
  level: "debug"
  format: "json"
  
pagination:
  defaultSize: 20
  maxSize: 100
```

**Alternatives Considered**:
- **Environment Variables Only**: Rejected because complex nested config is harder to manage, no type safety
- **JSON Config**: Rejected because YAML is more human-readable, supports comments

### 5. Logging: Uber Zap

**Decision**: Use Uber's Zap library for structured logging.

**Rationale**:
- **Performance**: Zero-allocation logging in production, critical for high-throughput scenarios
- **Structured Logging**: JSON-formatted logs with fields - easily parseable by log aggregators (ELK, Splunk)
- **Log Levels**: Debug/Info/Warn/Error levels support environment-specific verbosity
- **Context Integration**: Works well with Go context for request tracing
- **Constitution Compliance**: Meets security requirement for detailed audit logging

**Alternatives Considered**:
- **Logrus**: Rejected due to slower performance (allocations), deprecated by maintainers
- **Standard log**: Too basic, no structured logging, no log levels

### 6. Validation: Go Playground Validator v10

**Decision**: Use go-playground/validator for struct tag-based validation.

**Rationale**:
- **Declarative**: Validation rules defined in struct tags - clear, self-documenting
- **Comprehensive**: Built-in rules for email, required, min/max length, custom validators
- **Integration**: Works seamlessly with Gin binding
- **Performance**: Fast validation execution, minimal overhead
- **Constitution Compliance**: Satisfies security requirement for input validation

**Example**:
```go
type CreatePersonRequest struct {
    Name      string    `json:"name" validate:"required,min=2,max=100"`
    Surname   string    `json:"surname" validate:"required,min=2,max=100"`
    Sex       string    `json:"sex" validate:"required,oneof=masculino feminino"`
    Birthdate time.Time `json:"birthdate" validate:"required,past,maxage=150"`
}
```

**Alternatives Considered**:
- **Manual Validation**: Rejected due to code duplication, error-prone, violates DRY
- **Ozzo-validation**: Rejected because struct tags are more idiomatic in Go

### 7. Testing: Go Testing + Testify + Testcontainers

**Decision**: Use Go's standard testing package with testify assertions and testcontainers-go for integration tests.

**Rationale**:
- **Standard Testing**: Go's built-in testing is sufficient, no need for additional framework overhead
- **Testify Assertions**: Cleaner assertion syntax (`assert.Equal(t, expected, actual)`) improves test readability
- **Testcontainers**: Spins up real MongoDB in Docker for integration tests - eliminates mocking complexity, tests real behavior
- **Coverage**: Built-in coverage reporting (`go test -cover`) tracks 80%+ target
- **TDD-Friendly**: Fast test execution supports test-first development workflow

**Test Strategy**:
- **Unit Tests**: Use cases, entities, validators (with mocked repositories)
- **Integration Tests**: Repository implementations (with Testcontainers MongoDB)
- **API Tests**: HTTP handlers (with in-memory MongoDB or test database)

**Alternatives Considered**:
- **Ginkgo/Gomega**: Rejected due to additional learning curve, BDD style not required
- **Mocking MongoDB**: Rejected because testcontainers provides real database behavior

### 8. Containerization: Docker Multi-Stage Build

**Decision**: Use Docker multi-stage builds for creating production images.

**Rationale**:
- **Small Image Size**: Multi-stage build produces <20MB final image (Go binary + scratch base)
- **Security**: Minimal attack surface with distroless or scratch base image
- **Fast Startup**: Compiled binary starts in milliseconds
- **Reproducible Builds**: Same Dockerfile produces identical images across environments
- **Constitution Compliance**: Meets security requirement for container deployment

**Dockerfile Strategy**:
```dockerfile
# Stage 1: Build
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o api ./cmd/api

# Stage 2: Production
FROM scratch
COPY --from=builder /app/api /api
COPY --from=builder /app/configs /configs
ENTRYPOINT ["/api"]
```

**Alternatives Considered**:
- **Single-Stage Build**: Rejected due to large image size (~800MB with build tools)
- **Alpine Final Image**: Rejected because scratch is smaller and more secure

## Design Pattern Decisions

### 9. Repository Pattern

**Decision**: Implement Repository pattern with interface in domain layer, MongoDB implementation in infrastructure layer.

**Rationale**:
- **Clean Architecture**: Decouples business logic from database implementation
- **Testability**: Use cases can be tested with mock repositories
- **Flexibility**: Can swap MongoDB for another database without changing use cases
- **Single Responsibility**: Repository handles only data access concerns

**Interface Design**:
```go
type PersonRepository interface {
    Create(ctx context.Context, person *entities.Person) error
    FindByID(ctx context.Context, id string) (*entities.Person, error)
    FindAll(ctx context.Context, page, pageSize int) ([]*entities.Person, int, error)
    Update(ctx context.Context, person *entities.Person) error
    Delete(ctx context.Context, id string) error
}
```

### 10. Use Case Pattern

**Decision**: Implement each business operation as a separate use case struct with Execute method.

**Rationale**:
- **Single Responsibility**: Each use case does one thing (Create, Get, List, Update, Delete)
- **Testability**: Use cases are independently testable
- **Clarity**: Clear mapping between user stories and use case files
- **Reusability**: Use cases can be called from HTTP handlers, CLI, gRPC, etc.

**Example Structure**:
```go
type CreatePersonUseCase struct {
    repo      repositories.PersonRepository
    validator *validator.Validate
    logger    *zap.Logger
}

func (uc *CreatePersonUseCase) Execute(ctx context.Context, req dto.CreatePersonRequest) (*entities.Person, error) {
    // 1. Validate input
    // 2. Create entity
    // 3. Save via repository
    // 4. Return result
}
```

### 11. DTO Pattern

**Decision**: Use separate DTOs for HTTP requests/responses, distinct from domain entities.

**Rationale**:
- **API Stability**: Can change internal entities without breaking API contracts
- **Validation**: DTOs contain validation tags for API layer
- **Security**: Prevents exposing internal entity details (e.g., database IDs, timestamps)
- **Transformation**: Clean separation between API format and domain format

### 12. Dependency Injection

**Decision**: Use constructor injection for all dependencies.

**Rationale**:
- **Testability**: Dependencies can be mocked/stubbed in tests
- **Clarity**: All dependencies explicit in constructor signature
- **Constitution Compliance**: Satisfies dependency inversion principle
- **Go Idiom**: Standard practice in Go (no DI framework needed)

**Example**:
```go
func NewPersonHandler(
    createUC *usecases.CreatePersonUseCase,
    getUC *usecases.GetPersonUseCase,
    logger *zap.Logger,
) *PersonHandler {
    return &PersonHandler{
        createUC: createUC,
        getUC:    getUC,
        logger:   logger,
    }
}
```

## API Design Decisions

### 13. REST Endpoint Design

**Decision**: Follow RESTful conventions with resource-oriented URLs.

**Endpoints**:
- `POST /api/v1/persons` - Create person
- `GET /api/v1/persons/:id` - Get person by ID
- `GET /api/v1/persons` - List persons (with ?page=1&pageSize=20)
- `PUT /api/v1/persons/:id` - Update person (full update)
- `DELETE /api/v1/persons/:id` - Delete person

**Rationale**:
- **Industry Standard**: REST conventions well-understood by API consumers
- **Versioning**: /v1/ prefix allows future API versions without breaking changes
- **HTTP Semantics**: Proper verb usage (POST=create, GET=read, PUT=update, DELETE=delete)
- **Plural Resources**: /persons (plural) is REST convention for collections

**Alternatives Considered**:
- **GraphQL**: Rejected as overkill for simple CRUD operations, adds complexity
- **gRPC**: Rejected as requirement specifies REST API, not RPC

### 14. HTTP Status Codes

**Decision**: Follow RFC 7231 HTTP status code best practices.

**Status Code Mapping**:
- `201 Created` - Successful person creation (with Location header)
- `200 OK` - Successful read, update, delete, list
- `400 Bad Request` - Validation errors, malformed input
- `404 Not Found` - Person ID doesn't exist
- `500 Internal Server Error` - Unexpected server errors (don't expose details)

**Rationale**:
- **Specification Requirement**: FR-019 explicitly requires proper HTTP status codes
- **Client Experience**: Clear status codes allow proper error handling
- **RESTful Practice**: Semantic use of HTTP status codes

### 15. Error Response Format

**Decision**: Standardize JSON error responses.

**Format**:
```json
{
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "Invalid input data",
    "details": [
      {"field": "email", "message": "must be valid email format"},
      {"field": "birthdate", "message": "must be in the past"}
    ]
  }
}
```

**Rationale**:
- **Constitution Compliance**: Meets security requirement for safe error messages (no internal details)
- **Specification Requirement**: FR-020 requires clear, actionable error messages
- **Consistency**: Same structure for all error responses
- **Debuggability**: Structured details allow specific field-level errors

### 16. Pagination Response Format

**Decision**: Include pagination metadata in list responses.

**Format**:
```json
{
  "data": [...],
  "pagination": {
    "page": 1,
    "pageSize": 20,
    "totalRecords": 145,
    "totalPages": 8
  }
}
```

**Rationale**:
- **Specification Requirement**: FR-013 requires pagination support
- **User Experience**: Clients need total count for UI rendering (e.g., page numbers)
- **Performance**: Prevents loading all records into memory

## Performance Optimization Decisions

### 17. MongoDB Indexing Strategy

**Decision**: Create indexes on frequently queried fields.

**Indexes**:
- `_id`: Unique index (automatic, primary key)
- `createdAt`: Index for sorting by creation date in list queries
- Compound index optional based on query patterns discovered in testing

**Rationale**:
- **Specification Requirement**: SC-002 requires <200ms P95 read latency
- **Query Performance**: Index on _id ensures O(1) lookups by ID
- **List Performance**: Index on createdAt enables efficient sorting

### 18. Connection Pooling

**Decision**: Configure MongoDB connection pool with reasonable limits.

**Configuration**:
- Min Pool Size: 5 connections
- Max Pool Size: 100 connections
- Connection Timeout: 10 seconds

**Rationale**:
- **Concurrency**: Supports 100 concurrent users requirement
- **Resource Management**: Pool prevents connection exhaustion
- **Constitution Compliance**: Meets performance standards for resource management

### 19. Context Timeouts

**Decision**: Use context.WithTimeout for all database operations.

**Timeout Values**:
- Write operations (create, update, delete): 5 seconds
- Read operations (get, list): 3 seconds

**Rationale**:
- **Performance Guarantee**: Ensures operations don't hang indefinitely
- **Resource Protection**: Prevents slow queries from exhausting connections
- **Specification Compliance**: Helps meet P95 latency targets

## Security Decisions

### 20. Input Validation Strategy

**Decision**: Validate at multiple layers (presentation, application).

**Layers**:
- **Presentation Layer**: Gin binding + validator tags (basic structure validation)
- **Application Layer**: Use case validates business rules (age limits, date ranges)
- **Domain Layer**: Entity constructor validates invariants

**Rationale**:
- **Defense in Depth**: Multiple validation layers catch different error types
- **Constitution Compliance**: Meets security requirement for comprehensive input validation
- **Specification Requirement**: FR-004 requires age validation, FR-007/008 require format validation

### 21. NoSQL Injection Prevention

**Decision**: Use MongoDB driver's BSON encoding, never string concatenation for queries.

**Safe Pattern**:
```go
// ✅ Safe - uses BSON encoding
filter := bson.M{"_id": objectID}
collection.FindOne(ctx, filter)

// ❌ Unsafe - NEVER do this
query := fmt.Sprintf(`{"_id": "%s"}`, userInput)
```

**Rationale**:
- **Constitution Compliance**: Meets security requirement for injection prevention
- **MongoDB Best Practice**: Official driver prevents injection by design

### 22. Secret Management

**Decision**: Store secrets in environment-specific YAML files, load via Viper, never commit to git.

**Pattern**:
- Config files in `configs/` directory
- Git ignore: `configs/*.yaml` (commit only `configs/*.yaml.example`)
- Deployment: Inject actual configs via ConfigMaps (Kubernetes) or secrets management

**Rationale**:
- **Constitution Compliance**: Meets security requirement "NEVER commit secrets"
- **Environment Isolation**: Separate credentials for LOCAL/BETA/PROD
- **Specification Requirement**: Environment-specific configuration per user requirement

## Development Workflow Decisions

### 23. TDD Approach

**Decision**: Follow Test-Driven Development - write tests before implementation.

**Workflow**:
1. Write failing test for new feature
2. Implement minimal code to pass test
3. Refactor while keeping tests green

**Rationale**:
- **AGENTS.md Requirement**: TDD explicitly required in operational guidelines
- **Quality**: TDD ensures testable code by design
- **Constitution Compliance**: Supports 80%+ coverage requirement

### 24. Git Workflow

**Decision**: Feature branch workflow with pull requests.

**Pattern**:
- Feature branches: `###-feature-name` (e.g., `001-pessoa-crud-api`)
- Commits follow conventional commits (feat:, fix:, refactor:, test:, docs:)
- PR requires passing CI (lint, tests, build)

**Rationale**:
- **AGENTS.md Standard**: Documented in operational guidelines
- **Quality Gates**: CI enforces constitution principles before merge
- **Traceability**: Feature branches map to specifications

## Open Questions & Future Research

None at this time. All technical decisions have been made with clear rationale. If new unknowns emerge during implementation, they will be documented here with research tasks.

## References

- [Feature Specification](spec.md)
- [Implementation Plan](plan.md)
- [Constitution](../../.specify/memory/constitution.md)
- [AGENTS.md](../../AGENTS.md)
- [Gin Documentation](https://gin-gonic.com/docs/)
- [MongoDB Go Driver](https://www.mongodb.com/docs/drivers/go/current/)
- [Go Clean Architecture Examples](https://github.com/bxcodec/go-clean-arch)
