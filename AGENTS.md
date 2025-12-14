# AGENTS.md

## Project Overview

Sistema de cadastro de pessoas desenvolvido em Go seguindo princípios de Clean Code e Clean Architecture.

**Core Principles**: Este projeto segue rigorosamente a constitution definida em `.specify/memory/constitution.md` (v1.1.0), que estabelece 7 princípios fundamentais não-negociáveis.

## Setup Commands

```bash
# Install Go dependencies
go mod download
go mod tidy

# Build the application
go build -o bin/cadastro-pessoas ./cmd/api

# Run the application
go run ./cmd/api

# Run with hot reload (using air)
air

# Run tests
go test ./...

# Run tests with coverage
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# Run tests verbosely
go test -v ./...

# Run specific package tests
go test ./internal/domain/...
go test ./internal/application/...

# Linting
golangci-lint run

# Format code
gofmt -s -w .
# or using goimports
goimports -w .

# Generate mocks (if using mockgen)
go generate ./...

# Check for vulnerabilities
go list -json -m all | nancy sleuth
```

## Project Structure (Clean Architecture)

```
cadastro-pessoas/
├── cmd/
│   └── api/              # Application entry points
│       └── main.go
├── internal/
│   ├── domain/           # Entities layer (business rules)
│   │   ├── entities/     # Core business objects (structs)
│   │   └── repositories/ # Repository interfaces
│   ├── application/      # Use Cases layer (application logic)
│   │   ├── usecases/     # Business workflows
│   │   └── services/     # Application services
│   ├── infrastructure/   # Frameworks & Drivers layer
│   │   ├── persistence/  # Database implementations
│   │   ├── config/       # Configuration management
│   │   └── external/     # External service clients
│   └── presentation/     # Interface Adapters layer
│       ├── http/         # HTTP handlers/controllers
│       ├── dto/          # Data transfer objects
│       └── middleware/   # HTTP middleware
├── pkg/                  # Public reusable packages
│   ├── logger/          # Logging utilities
│   ├── validator/       # Input validation
│   └── errors/          # Custom error types
├── tests/
│   ├── unit/            # Unit tests
│   ├── integration/     # Integration tests
│   └── contract/        # API contract tests
├── scripts/             # Build and deployment scripts
├── docs/                # Documentation
├── .specify/            # Spec Kit templates and memory
└── go.mod
```

**Navigation Tips**:
- Use `cd internal/domain` for business logic
- Use `cd internal/application` for use cases
- Use `cd internal/infrastructure` for external dependencies
- Use `cd internal/presentation` for HTTP handlers

## Code Style Guidelines (Go-Specific)

### I. Clean Code Principles

**Naming Conventions**:
```go
// ✅ Good - descriptive, intention-revealing
type UserRepository interface {
    FindByEmail(ctx context.Context, email string) (*User, error)
}

// ❌ Bad - unclear, abbreviated
type UsrRepo interface {
    FndByEml(ctx context.Context, e string) (*User, error)
}

// Constants - use named constants, not magic numbers
const (
    MaxLoginAttempts     = 3
    SessionTimeoutMinutes = 30
    DefaultPageSize      = 20
)

// ❌ Avoid magic numbers
if attempts > 3 { // What is 3?
```

**Function Size & Responsibility**:
```go
// ✅ Good - single responsibility, < 20 lines
func (s *UserService) CreateUser(ctx context.Context, email, password string) (*User, error) {
    if err := s.validator.ValidateEmail(email); err != nil {
        return nil, err
    }
    
    hashedPassword, err := s.hasher.Hash(password)
    if err != nil {
        return nil, err
    }
    
    user := &User{Email: email, Password: hashedPassword}
    return s.repo.Create(ctx, user)
}

// ❌ Bad - doing too many things
func (s *UserService) CreateUserAndSendEmailAndLogAndNotify(...) // Too much!
```

**Error Handling**:
```go
// ✅ Good - use errors.Wrap for context
if err := db.Save(user); err != nil {
    return fmt.Errorf("failed to save user %s: %w", user.Email, err)
}

// ✅ Good - custom error types for domain errors
var ErrUserNotFound = errors.New("user not found")
var ErrInvalidCredentials = errors.New("invalid credentials")

// ❌ Bad - returning error codes or using panic
if err != nil {
    return -1, nil // Don't return error codes
}
panic("something went wrong") // Never panic in business logic
```

**DRY Principle**:
```go
// ✅ Good - extracted common validation
func validateEmailFormat(email string) error {
    // Single source of truth for email validation
}

// ❌ Bad - duplicated validation logic across multiple places
```

### II. Clean Architecture in Go

**Dependency Rule** (dependencies point inward):
```go
// ✅ Good - domain doesn't import anything
package entities

type User struct {
    ID       string
    Email    string
    Password string
}

// ✅ Good - repository interface in domain
package repositories

type UserRepository interface {
    Create(ctx context.Context, user *entities.User) error
    FindByID(ctx context.Context, id string) (*entities.User, error)
}

// ✅ Good - infrastructure implements domain interface
package persistence

import "internal/domain/repositories"

type PostgresUserRepository struct {
    db *sql.DB
}

// Implements repositories.UserRepository
func (r *PostgresUserRepository) Create(ctx context.Context, user *entities.User) error {
    // Implementation using concrete database
}

// ❌ Bad - domain importing infrastructure
package entities

import "github.com/lib/pq" // NEVER import concrete implementations in domain!
```

**Use Cases (Application Layer)**:
```go
// ✅ Good - use case with injected dependencies
type CreateUserUseCase struct {
    userRepo repositories.UserRepository
    hasher   crypto.PasswordHasher
    logger   logger.Logger
}

func NewCreateUserUseCase(
    userRepo repositories.UserRepository,
    hasher crypto.PasswordHasher,
    logger logger.Logger,
) *CreateUserUseCase {
    return &CreateUserUseCase{
        userRepo: userRepo,
        hasher:   hasher,
        logger:   logger,
    }
}

func (uc *CreateUserUseCase) Execute(ctx context.Context, email, password string) (*entities.User, error) {
    // Business logic here
}
```

### III. Code Reusability & DRY

**Utility Packages**:
```go
// Create reusable packages in pkg/
// pkg/validator/email.go
package validator

import "regexp"

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func IsValidEmail(email string) bool {
    return emailRegex.MatchString(email)
}

// pkg/crypto/hasher.go
package crypto

type PasswordHasher interface {
    Hash(password string) (string, error)
    Compare(hashed, password string) error
}
```

**Composition Over Inheritance** (Go uses composition naturally):
```go
// ✅ Good - composition with embedded structs
type BaseRepository struct {
    db *sql.DB
}

type UserRepository struct {
    BaseRepository // Compose, don't inherit
    cache Cache
}

// ✅ Good - interfaces for flexibility
type Logger interface {
    Info(msg string, fields ...Field)
    Error(msg string, err error, fields ...Field)
}
```

### IV. Design Patterns in Go

**Repository Pattern**:
```go
// Domain layer - interface
type UserRepository interface {
    Create(ctx context.Context, user *User) error
    Update(ctx context.Context, user *User) error
    Delete(ctx context.Context, id string) error
    FindByID(ctx context.Context, id string) (*User, error)
    FindByEmail(ctx context.Context, email string) (*User, error)
}

// Infrastructure layer - implementation
type postgresUserRepository struct {
    db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) UserRepository {
    return &postgresUserRepository{db: db}
}
```

**Factory Pattern**:
```go
type Database interface {
    Connect() error
    Close() error
}

func NewDatabase(config *Config) (Database, error) {
    switch config.Type {
    case "postgres":
        return NewPostgresDB(config)
    case "mysql":
        return NewMySQLDB(config)
    default:
        return nil, fmt.Errorf("unsupported database type: %s", config.Type)
    }
}
```

**Dependency Injection** (constructor injection):
```go
// ✅ Good - dependencies injected via constructor
type UserHandler struct {
    createUserUseCase *CreateUserUseCase
    logger            Logger
}

func NewUserHandler(createUserUseCase *CreateUserUseCase, logger Logger) *UserHandler {
    return &UserHandler{
        createUserUseCase: createUserUseCase,
        logger:            logger,
    }
}

// ❌ Bad - creating dependencies inside
func NewUserHandler() *UserHandler {
    repo := NewUserRepository() // Don't do this!
    return &UserHandler{repo: repo}
}
```

**Strategy Pattern**:
```go
type AuthStrategy interface {
    Authenticate(ctx context.Context, credentials Credentials) (*User, error)
}

type EmailPasswordAuth struct{}
type OAuth2Auth struct{}
type JWTAuth struct{}

// Each implements AuthStrategy
```

### V. Function & Method Cohesion

**Single Responsibility**:
```go
// ✅ Good - each function does one thing
func (s *UserService) ValidateUser(user *User) error {
    return s.validator.Validate(user)
}

func (s *UserService) HashPassword(password string) (string, error) {
    return s.hasher.Hash(password)
}

func (s *UserService) SaveUser(ctx context.Context, user *User) error {
    return s.repo.Create(ctx, user)
}

// ❌ Bad - doing multiple things
func (s *UserService) ValidateHashAndSave(ctx context.Context, user *User) error {
    // Doing too much!
}
```

**Minimal Parameters** (≤3 parameters, use structs for more):
```go
// ✅ Good - using struct for multiple parameters
type CreateUserRequest struct {
    Email     string
    Password  string
    FirstName string
    LastName  string
    BirthDate time.Time
}

func (s *UserService) CreateUser(ctx context.Context, req CreateUserRequest) (*User, error) {
    // Implementation
}

// ❌ Bad - too many parameters
func CreateUser(ctx context.Context, email, password, firstName, lastName string, birthDate time.Time) (*User, error) {
    // Hard to read and maintain
}
```

**Pure Functions Preferred**:
```go
// ✅ Good - pure function (same input → same output, no side effects)
func CalculateAge(birthDate time.Time) int {
    return int(time.Since(birthDate).Hours() / 24 / 365)
}

// ✅ Good - side effects are explicit (context, error)
func (r *UserRepository) Save(ctx context.Context, user *User) error {
    // Side effect is clear from signature
}
```

### VI. Performance Standards

**Context Usage**:
```go
// ✅ Always pass context as first parameter
func (s *Service) FetchData(ctx context.Context, id string) (*Data, error) {
    // Can be cancelled, has timeout, carries values
}

// Set timeouts
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()
```

**Database Optimization**:
```go
// ✅ Good - use prepared statements
stmt, err := db.PrepareContext(ctx, "SELECT * FROM users WHERE email = $1")
defer stmt.Close()

// ✅ Good - batch operations
tx, err := db.BeginTx(ctx, nil)
for _, user := range users {
    tx.ExecContext(ctx, "INSERT INTO users...", user.Email)
}
tx.Commit()

// ❌ Bad - N+1 query problem
for _, userID := range userIDs {
    user, _ := repo.FindByID(ctx, userID) // Multiple queries!
}

// ✅ Good - single query with IN clause
users, _ := repo.FindByIDs(ctx, userIDs)
```

**Concurrency** (use goroutines wisely):
```go
// ✅ Good - parallel processing with sync.WaitGroup
var wg sync.WaitGroup
results := make(chan Result, len(items))

for _, item := range items {
    wg.Add(1)
    go func(item Item) {
        defer wg.Done()
        result := process(item)
        results <- result
    }(item)
}

go func() {
    wg.Wait()
    close(results)
}()

// ✅ Good - use worker pool for controlled concurrency
semaphore := make(chan struct{}, maxWorkers)
```

**Memory Management**:
```go
// ✅ Good - close resources
defer file.Close()
defer db.Close()
defer resp.Body.Close()

// ✅ Good - use sync.Pool for frequently allocated objects
var bufferPool = sync.Pool{
    New: func() interface{} {
        return new(bytes.Buffer)
    },
}

// ✅ Good - avoid memory leaks
ctx, cancel := context.WithTimeout(context.Background(), time.Second)
defer cancel() // Always defer cancel!
```

**Caching**:
```go
// ✅ Use caching for expensive operations
type CachedUserRepository struct {
    repo  UserRepository
    cache Cache
}

func (r *CachedUserRepository) FindByID(ctx context.Context, id string) (*User, error) {
    if cached, ok := r.cache.Get(id); ok {
        return cached.(*User), nil
    }
    
    user, err := r.repo.FindByID(ctx, id)
    if err == nil {
        r.cache.Set(id, user, 5*time.Minute)
    }
    return user, err
}
```

### VII. Security Requirements

**Input Validation**:
```go
// ✅ Good - validate all inputs
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
    var req CreateUserRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }
    
    // Validate
    if err := validator.Validate(req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    
    // Sanitize
    req.Email = strings.TrimSpace(strings.ToLower(req.Email))
    
    // Process...
}

// ✅ Use allowlists for validation
var allowedRoles = map[string]bool{
    "admin": true,
    "user":  true,
    "guest": true,
}

func isValidRole(role string) bool {
    return allowedRoles[role]
}
```

**SQL Injection Prevention**:
```go
// ✅ Good - parameterized queries
query := "SELECT * FROM users WHERE email = $1 AND active = $2"
row := db.QueryRowContext(ctx, query, email, true)

// ❌ NEVER do string concatenation for SQL
query := fmt.Sprintf("SELECT * FROM users WHERE email = '%s'", email) // VULNERABLE!
```

**Password Handling**:
```go
// ✅ Good - use bcrypt or argon2
import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(bytes), err
}

func CheckPassword(hashed, password string) error {
    return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
}

// ❌ NEVER store plain text passwords
user.Password = password // NEVER!
```

**Secret Management**:
```go
// ✅ Good - use environment variables
import "os"

dbPassword := os.Getenv("DB_PASSWORD")
apiKey := os.Getenv("API_KEY")

// ✅ Good - use a config struct
type Config struct {
    DBPassword string `env:"DB_PASSWORD,required"`
    APIKey     string `env:"API_KEY,required"`
}

// ❌ NEVER hardcode secrets
const apiKey = "sk_live_123456789" // NEVER!
```

**Secure Error Handling**:
```go
// ✅ Good - generic error to user, detailed log internally
func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
    user, err := h.service.GetUser(ctx, id)
    if err != nil {
        h.logger.Error("Failed to get user", err, Fields{"user_id": id})
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return // Don't expose stack trace!
    }
    json.NewEncoder(w).Encode(user)
}

// ❌ Bad - exposing internal details
http.Error(w, err.Error(), 500) // Might expose database structure!
```

**HTTPS & TLS**:
```go
// ✅ Good - enforce HTTPS
func main() {
    server := &http.Server{
        Addr:         ":443",
        Handler:      handler,
        TLSConfig:    &tls.Config{MinVersion: tls.VersionTLS12},
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 10 * time.Second,
        IdleTimeout:  120 * time.Second,
    }
    log.Fatal(server.ListenAndServeTLS("cert.pem", "key.pem"))
}
```

**Rate Limiting**:
```go
// ✅ Implement rate limiting middleware
import "golang.org/x/time/rate"

func RateLimitMiddleware(limiter *rate.Limiter) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            if !limiter.Allow() {
                http.Error(w, "Too many requests", http.StatusTooManyRequests)
                return
            }
            next.ServeHTTP(w, r)
        })
    }
}
```

## Testing Instructions

### Test Structure
```bash
# Run all tests
go test ./...

# Run tests with race detection
go test -race ./...

# Run tests with coverage (minimum 80% required)
go test -coverprofile=coverage.out ./...
go tool cover -func=coverage.out

# View coverage in browser
go tool cover -html=coverage.out
```

### Test Naming Convention
```go
// Format: Test<FunctionName>_<Scenario>_<ExpectedResult>
func TestCreateUser_ValidInput_ReturnsUser(t *testing.T) {}
func TestCreateUser_InvalidEmail_ReturnsError(t *testing.T) {}
func TestCreateUser_DuplicateEmail_ReturnsError(t *testing.T) {}
```

### TDD Workflow (MANDATORY)
1. **Write test first** - test should FAIL initially
2. **Run test** - verify it fails for the right reason
3. **Write minimal code** to make test pass
4. **Refactor** while keeping tests green
5. **Repeat**

### Test Examples
```go
// Unit test example
func TestCreateUser_ValidInput_ReturnsUser(t *testing.T) {
    // Arrange
    mockRepo := &MockUserRepository{}
    service := NewUserService(mockRepo)
    req := CreateUserRequest{Email: "test@example.com"}
    
    // Act
    user, err := service.CreateUser(context.Background(), req)
    
    // Assert
    assert.NoError(t, err)
    assert.NotNil(t, user)
    assert.Equal(t, req.Email, user.Email)
}

// Table-driven tests
func TestValidateEmail(t *testing.T) {
    tests := []struct {
        name    string
        email   string
        wantErr bool
    }{
        {"valid email", "user@example.com", false},
        {"invalid email", "invalid", true},
        {"empty email", "", true},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := ValidateEmail(tt.email)
            if (err != nil) != tt.wantErr {
                t.Errorf("ValidateEmail() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}
```

### Integration Tests
```go
// Use testcontainers for real database
func TestUserRepository_Integration(t *testing.T) {
    if testing.Short() {
        t.Skip("skipping integration test")
    }
    
    // Setup test database
    container := setupPostgresContainer(t)
    defer container.Terminate(context.Background())
    
    // Run tests against real database
}
```

## Build & Deployment

```bash
# Build for production
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/app ./cmd/api

# Build with version info
go build -ldflags "-X main.Version=$(git describe --tags)" -o bin/app ./cmd/api

# Cross-compile
GOOS=darwin GOARCH=amd64 go build -o bin/app-darwin
GOOS=windows GOARCH=amd64 go build -o bin/app-windows.exe

# Docker build
docker build -t cadastro-pessoas:latest .

# Run in Docker
docker run -p 8080:8080 cadastro-pessoas:latest
```

## Dependencies

### Required Go Version
- Go 1.21+

### Core Dependencies
```go
// Framework & HTTP
github.com/gorilla/mux          // HTTP router
github.com/go-chi/chi           // Alternative router

// Database
github.com/lib/pq               // PostgreSQL driver
github.com/jmoiron/sqlx         // SQL extensions

// Configuration
github.com/spf13/viper          // Configuration management
github.com/joho/godotenv        // .env file support

// Validation
github.com/go-playground/validator/v10  // Struct validation

// Logging
go.uber.org/zap                 // Structured logging
github.com/sirupsen/logrus      // Alternative logging

// Testing
github.com/stretchr/testify     // Test assertions
github.com/golang/mock          // Mocking framework

// Security
golang.org/x/crypto/bcrypt      // Password hashing
github.com/golang-jwt/jwt       // JWT tokens
```

## Git Workflow

### Commit Messages
Follow conventional commits:
```
feat: add user creation endpoint
fix: resolve nil pointer in user service
refactor: extract validation to separate package
test: add integration tests for user repository
docs: update AGENTS.md with security guidelines
perf: optimize database queries with batch operations
security: add rate limiting middleware
```

### Pre-commit Checklist
- [ ] Code formatted with `gofmt` or `goimports`
- [ ] All tests pass (`go test ./...`)
- [ ] Test coverage ≥ 80%
- [ ] Linter passes (`golangci-lint run`)
- [ ] No security vulnerabilities (`go list -json -m all | nancy sleuth`)
- [ ] Constitution principles followed (see `.specify/memory/constitution.md`)
- [ ] No hardcoded secrets
- [ ] Proper error handling

### Pull Request Template
```markdown
## Description
[Describe what this PR does]

## Type of Change
- [ ] Bug fix
- [ ] New feature
- [ ] Breaking change
- [ ] Documentation update

## Constitution Compliance
- [ ] Clean Code: Functions < 20 lines, meaningful names
- [ ] Clean Architecture: Proper layer separation
- [ ] Code Reusability: No duplication
- [ ] Design Patterns: Appropriate patterns used
- [ ] Function Cohesion: Single responsibility
- [ ] Performance: No regressions, optimizations justified
- [ ] Security: Input validated, no secrets, proper error handling

## Testing
- [ ] Unit tests added/updated
- [ ] Integration tests added/updated
- [ ] All tests passing
- [ ] Coverage ≥ 80%

## Checklist
- [ ] Code reviewed by at least one peer
- [ ] Documentation updated
- [ ] CHANGELOG.md updated (if applicable)
```

## Common Commands Reference

```bash
# Development
go run ./cmd/api                              # Run application
go build -o bin/app ./cmd/api                # Build binary
air                                          # Hot reload during development

# Testing
go test ./...                                # Run all tests
go test -v ./internal/domain/...            # Run domain tests
go test -cover ./...                        # Run with coverage
go test -bench=.                            # Run benchmarks
go test -race ./...                         # Run with race detector

# Code Quality
gofmt -s -w .                               # Format code
goimports -w .                              # Format imports
golangci-lint run                           # Comprehensive linting
go vet ./...                                # Go vet checks
staticcheck ./...                           # Static analysis

# Dependencies
go mod download                             # Download dependencies
go mod tidy                                 # Clean up dependencies
go mod vendor                               # Vendor dependencies
go get -u ./...                             # Update dependencies

# Documentation
godoc -http=:6060                           # Start godoc server
go doc <package>                            # View package docs

# Profiling
go test -cpuprofile=cpu.prof -bench=.      # CPU profiling
go test -memprofile=mem.prof -bench=.      # Memory profiling
go tool pprof cpu.prof                     # Analyze profile

# Security
go list -json -m all | nancy sleuth        # Check vulnerabilities
gosec ./...                                 # Security scanning
```

## Environment Variables

```bash
# Required
export DB_HOST=localhost
export DB_PORT=5432
export DB_NAME=cadastro_pessoas
export DB_USER=postgres
export DB_PASSWORD=<secret>
export DB_SSLMODE=disable

export JWT_SECRET=<secret>
export API_PORT=8080

# Optional
export LOG_LEVEL=info
export CACHE_ENABLED=true
export CACHE_TTL=300
export RATE_LIMIT_REQUESTS=100
export RATE_LIMIT_WINDOW=60
```

## Troubleshooting

### Common Issues

**Import cycle**:
```bash
# Check for circular dependencies
go mod graph | grep <package>
```

**Race conditions**:
```bash
# Always run with race detector during development
go test -race ./...
```

**Memory leaks**:
```bash
# Profile memory usage
go test -memprofile=mem.prof -bench=.
go tool pprof mem.prof
```

## Additional Resources

- Constitution: `.specify/memory/constitution.md`
- API Documentation: `docs/api.md`
- Architecture Decisions: `docs/adr/`
- Spec Kit Documentation: https://github.com/github/spec-kit

## Questions or Issues?

1. Check `.specify/memory/constitution.md` for architectural guidance
2. Review existing code in similar packages
3. Ask in team chat or create a GitHub issue
4. Follow the Spec-Driven Development workflow using `/speckit.*` commands

---

**Remember**: This project follows Clean Code, Clean Architecture, and strict security/performance standards. Always refer to `.specify/memory/constitution.md` for the authoritative principles and governance rules.
