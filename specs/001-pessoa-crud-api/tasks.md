# Tasks: Person Management REST API

**Feature Branch**: `001-pessoa-crud-api`  
**Input**: Design documents from `/specs/001-pessoa-crud-api/`  
**Prerequisites**: plan.md ✅, spec.md ✅, research.md ✅, data-model.md ✅, contracts/ ✅

**Generated**: 2025-12-14  
**Total Tasks**: 123  
**Parallelizable**: 85

## Task Format

Each task follows the format: `- [ ] [TaskID] [P?] [Story?] Description with file path`

- **[P]**: Parallelizable (can run simultaneously with other [P] tasks)
- **[Story]**: User story label (US1, US2, US3, US4, US5)
- Task IDs are sequential (T001, T002, T003...)

## Path Conventions

Based on Clean Architecture single-project structure (Go):

- `cmd/api/` - Application entry point
- `internal/domain/entities/` - Domain entities and value objects
- `internal/domain/repositories/` - Repository interfaces
- `internal/application/usecases/` - Business use cases
- `internal/application/dto/` - Application DTOs
- `internal/infrastructure/persistence/mongodb/` - MongoDB implementation
- `internal/infrastructure/config/` - Configuration management
- `internal/infrastructure/logger/` - Logging setup
- `internal/presentation/http/handlers/` - HTTP request handlers
- `internal/presentation/http/middleware/` - HTTP middleware
- `internal/presentation/dto/` - HTTP request/response DTOs
- `pkg/validator/` - Shared validation utilities
- `pkg/errors/` - Custom error types
- `tests/unit/` - Unit tests
- `tests/integration/` - Integration tests
- `configs/` - YAML configuration files
- `docker/` - Docker and compose files

---

## Phase 1: Setup (Project Initialization)

**Purpose**: Create project structure and initialize development environment

**Estimated Time**: 1-2 hours

- [X] T001 Initialize Go module with `go mod init github.com/BrunnoQ/cadastro-pessoas`
- [X] T002 Create directory structure per plan.md (cmd/, internal/, pkg/, tests/, configs/, docker/)
- [X] T003 [P] Create .gitignore for Go project (bin/, vendor/, *.log, configs/*.yaml except examples)
- [X] T004 [P] Create go.mod with required dependencies (Gin v1.9+, MongoDB driver v1.13+, Viper, Zap, Validator)
- [X] T005 [P] Create Makefile with commands (build, run, test, lint, docker-up, docker-down)
- [X] T006 [P] Create docker-compose.yml with MongoDB 6.0 service configuration
- [X] T007 [P] Create Dockerfile with multi-stage build (builder + scratch base)
- [X] T008 [P] Create README.md with project overview and setup instructions
- [X] T009 [P] Setup golangci-lint configuration (.golangci.yml)
- [X] T010 [P] Create .air.toml for hot reload configuration

**Checkpoint**: ✅ Project structure created, dependencies defined

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: Core infrastructure that MUST be complete before ANY user story implementation

**⚠️ CRITICAL**: No user story work can begin until this phase is complete

**Estimated Time**: 4-6 hours

### Configuration & Environment

- [ ] T011 Create config struct in internal/infrastructure/config/config.go
- [ ] T012 Implement Viper-based YAML config loader in internal/infrastructure/config/config.go
- [ ] T013 [P] Create configs/config.local.yaml with local development settings
- [ ] T014 [P] Create configs/config.beta.yaml with beta environment settings
- [ ] T015 [P] Create configs/config.prod.yaml with production settings
- [ ] T016 [P] Create configs/config.local.yaml.example (template for git)

### Logging

- [ ] T017 [P] Setup Zap logger in internal/infrastructure/logger/logger.go
- [ ] T018 [P] Create logger factory with environment-specific configuration

### Database Connection

- [ ] T019 Create MongoDB connection factory in internal/infrastructure/persistence/mongodb/connection.go
- [ ] T020 Implement connection pooling configuration (min: 5, max: 100)
- [ ] T021 Create database initialization with index setup in internal/infrastructure/persistence/mongodb/indexes.go
- [ ] T022 [P] Create health check function for MongoDB connection

### Error Handling

- [ ] T023 [P] Define custom error types in pkg/errors/errors.go (NotFoundError, ValidationError, InternalError)
- [ ] T024 [P] Define error codes constants in pkg/errors/codes.go
- [ ] T025 [P] Create error response formatter in internal/presentation/http/handlers/error_handler.go

### Validation

- [ ] T026 [P] Create email validator in pkg/validator/email.go
- [ ] T027 [P] Create phone validator (E.164 format) in pkg/validator/phone.go
- [ ] T028 [P] Create date validator in pkg/validator/date.go (past dates, age limits)
- [ ] T029 [P] Create name validator in pkg/validator/name.go (letters, spaces, hyphens)

### HTTP Infrastructure

- [ ] T030 Setup Gin router in internal/presentation/http/router.go
- [ ] T031 [P] Create logging middleware in internal/presentation/http/middleware/logging.go
- [ ] T032 [P] Create recovery middleware in internal/presentation/http/middleware/recovery.go
- [ ] T033 [P] Create CORS middleware in internal/presentation/http/middleware/cors.go
- [ ] T034 [P] Create request validation middleware in internal/presentation/http/middleware/validation.go
- [ ] T035 [P] Create health check handler in internal/presentation/http/handlers/health_handler.go

### Application Entry Point

- [ ] T036 Create main.go in cmd/api/ with dependency injection and server startup
- [ ] T037 Implement graceful shutdown handling in cmd/api/main.go

**Checkpoint**: Foundation complete - all user stories can now proceed in parallel

---

## Phase 3: User Story 1 - Create New Person (Priority: P1) 🎯 MVP

**Goal**: Enable creation of person records with optional addresses and contacts

**Independent Test**: POST /api/v1/persons with valid data returns 201 with person ID

**Estimated Time**: 6-8 hours

### Domain Layer (Entities)

- [ ] T038 [P] [US1] Create Person entity in internal/domain/entities/person.go
- [ ] T039 [P] [US1] Create Address value object in internal/domain/entities/address.go
- [ ] T040 [P] [US1] Create Contact value object in internal/domain/entities/contact.go
- [ ] T041 [P] [US1] Implement Person.NewPerson() constructor with validation
- [ ] T042 [P] [US1] Implement Person.AddAddress() method with validation
- [ ] T043 [P] [US1] Implement Person.AddContact() method with validation
- [ ] T044 [P] [US1] Implement Address.Validate() method
- [ ] T045 [P] [US1] Implement Contact.Validate() method

### Domain Layer (Repository Interface)

- [ ] T046 [US1] Create PersonRepository interface in internal/domain/repositories/person_repository.go with Create method

### Application Layer (Use Case)

- [ ] T047 [US1] Create CreatePersonRequest DTO in internal/application/dto/person_request.go
- [ ] T048 [US1] Create PersonResponse DTO in internal/application/dto/person_response.go
- [ ] T049 [US1] Implement CreatePersonUseCase in internal/application/usecases/create_person.go

### Infrastructure Layer (MongoDB Repository)

- [ ] T050 [US1] Implement MongoPersonRepository.Create() in internal/infrastructure/persistence/mongodb/person_repository.go
- [ ] T051 [US1] Create MongoDB indexes (_id, created_at) in internal/infrastructure/persistence/mongodb/indexes.go

### Presentation Layer (HTTP)

- [ ] T052 [US1] Create HTTP CreatePersonRequest DTO in internal/presentation/dto/request.go
- [ ] T053 [US1] Create HTTP PersonResponse DTO in internal/presentation/dto/response.go
- [ ] T054 [US1] Implement PersonHandler.Create() in internal/presentation/http/handlers/person_handler.go
- [ ] T055 [US1] Register POST /api/v1/persons route in internal/presentation/http/router.go

### Tests

- [ ] T056 [P] [US1] Unit test: Person entity validation in tests/unit/entities/person_test.go
- [ ] T057 [P] [US1] Unit test: Address validation in tests/unit/entities/address_test.go
- [ ] T058 [P] [US1] Unit test: Contact validation in tests/unit/entities/contact_test.go
- [ ] T059 [P] [US1] Unit test: CreatePersonUseCase in tests/unit/usecases/create_person_test.go
- [ ] T060 [P] [US1] Integration test: Create person API endpoint in tests/integration/api/create_person_test.go

**Story 1 Checkpoint**: Can create persons via API, stored in MongoDB, retrievable by ID

---

## Phase 4: User Story 2 - Query Person by Identifier (Priority: P1) 🎯 MVP

**Goal**: Retrieve complete person information by unique identifier

**Independent Test**: GET /api/v1/persons/{id} returns 200 with full person data

**Estimated Time**: 3-4 hours

### Domain Layer (Repository Interface)

- [ ] T061 [US2] Add FindByID method to PersonRepository interface in internal/domain/repositories/person_repository.go

### Application Layer (Use Case)

- [ ] T062 [US2] Implement GetPersonUseCase in internal/application/usecases/get_person.go

### Infrastructure Layer (MongoDB Repository)

- [ ] T063 [US2] Implement MongoPersonRepository.FindByID() in internal/infrastructure/persistence/mongodb/person_repository.go

### Presentation Layer (HTTP)

- [ ] T064 [US2] Implement PersonHandler.GetByID() in internal/presentation/http/handlers/person_handler.go
- [ ] T065 [US2] Register GET /api/v1/persons/:id route in internal/presentation/http/router.go

### Tests

- [ ] T066 [P] [US2] Unit test: GetPersonUseCase in tests/unit/usecases/get_person_test.go
- [ ] T067 [P] [US2] Integration test: Get person API endpoint in tests/integration/api/get_person_test.go
- [ ] T068 [P] [US2] Integration test: Get non-existent person returns 404 in tests/integration/api/get_person_test.go

**Story 2 Checkpoint**: MVP Complete - Can create and retrieve persons (full read-write cycle)

---

## Phase 5: User Story 3 - List All Persons (Priority: P2)

**Goal**: Retrieve paginated list of all persons with summary data

**Independent Test**: GET /api/v1/persons returns 200 with paginated list and metadata

**Estimated Time**: 4-5 hours

### Domain Layer (Repository Interface)

- [ ] T069 [US3] Add FindAll method with pagination to PersonRepository interface in internal/domain/repositories/person_repository.go

### Application Layer (DTOs & Use Case)

- [ ] T070 [US3] Create PersonSummaryResponse DTO in internal/application/dto/person_response.go
- [ ] T071 [US3] Create PaginationResponse DTO in internal/application/dto/pagination.go
- [ ] T072 [US3] Create ListPersonsResponse DTO in internal/application/dto/person_response.go
- [ ] T073 [US3] Implement ListPersonsUseCase in internal/application/usecases/list_persons.go

### Infrastructure Layer (MongoDB Repository)

- [ ] T074 [US3] Implement MongoPersonRepository.FindAll() with pagination in internal/infrastructure/persistence/mongodb/person_repository.go

### Presentation Layer (HTTP)

- [ ] T075 [US3] Create pagination request parser in internal/presentation/dto/pagination.go
- [ ] T076 [US3] Implement PersonHandler.List() in internal/presentation/http/handlers/person_handler.go
- [ ] T077 [US3] Register GET /api/v1/persons route (list) in internal/presentation/http/router.go

### Tests

- [ ] T078 [P] [US3] Unit test: ListPersonsUseCase in tests/unit/usecases/list_persons_test.go
- [ ] T079 [P] [US3] Integration test: List persons API with pagination in tests/integration/api/list_persons_test.go
- [ ] T080 [P] [US3] Integration test: Empty list returns correct structure in tests/integration/api/list_persons_test.go

**Story 3 Checkpoint**: Can browse all persons with pagination

---

## Phase 6: User Story 4 - Update Person Information (Priority: P2)

**Goal**: Modify existing person data including nested addresses and contacts

**Independent Test**: PUT /api/v1/persons/{id} with updated data returns 200 with updated person

**Estimated Time**: 4-5 hours

### Domain Layer (Repository Interface)

- [ ] T081 [US4] Add Update method to PersonRepository interface in internal/domain/repositories/person_repository.go

### Application Layer (DTOs & Use Case)

- [ ] T082 [US4] Create UpdatePersonRequest DTO in internal/application/dto/person_request.go
- [ ] T083 [US4] Implement UpdatePersonUseCase in internal/application/usecases/update_person.go

### Infrastructure Layer (MongoDB Repository)

- [ ] T084 [US4] Implement MongoPersonRepository.Update() with optimistic locking in internal/infrastructure/persistence/mongodb/person_repository.go

### Presentation Layer (HTTP)

- [ ] T085 [US4] Implement PersonHandler.Update() in internal/presentation/http/handlers/person_handler.go
- [ ] T086 [US4] Register PUT /api/v1/persons/:id route in internal/presentation/http/router.go

### Tests

- [ ] T087 [P] [US4] Unit test: UpdatePersonUseCase in tests/unit/usecases/update_person_test.go
- [ ] T088 [P] [US4] Integration test: Update person API endpoint in tests/integration/api/update_person_test.go
- [ ] T089 [P] [US4] Integration test: Update non-existent person returns 404 in tests/integration/api/update_person_test.go
- [ ] T090 [P] [US4] Integration test: Concurrent update detection in tests/integration/api/update_person_test.go

**Story 4 Checkpoint**: Can update person records and nested collections

---

## Phase 7: User Story 5 - Delete Person (Priority: P3)

**Goal**: Remove person records with cascade delete of addresses and contacts

**Independent Test**: DELETE /api/v1/persons/{id} returns 200, subsequent GET returns 404

**Estimated Time**: 3-4 hours

### Domain Layer (Repository Interface)

- [ ] T091 [US5] Add Delete method to PersonRepository interface in internal/domain/repositories/person_repository.go

### Application Layer (Use Case)

- [ ] T092 [US5] Implement DeletePersonUseCase in internal/application/usecases/delete_person.go

### Infrastructure Layer (MongoDB Repository)

- [ ] T093 [US5] Implement MongoPersonRepository.Delete() in internal/infrastructure/persistence/mongodb/person_repository.go

### Presentation Layer (HTTP)

- [ ] T094 [US5] Implement PersonHandler.Delete() in internal/presentation/http/handlers/person_handler.go
- [ ] T095 [US5] Register DELETE /api/v1/persons/:id route in internal/presentation/http/router.go

### Tests

- [ ] T096 [P] [US5] Unit test: DeletePersonUseCase in tests/unit/usecases/delete_person_test.go
- [ ] T097 [P] [US5] Integration test: Delete person API endpoint in tests/integration/api/delete_person_test.go
- [ ] T098 [P] [US5] Integration test: Delete cascades to addresses and contacts in tests/integration/api/delete_person_test.go
- [ ] T099 [P] [US5] Integration test: Delete non-existent person is idempotent in tests/integration/api/delete_person_test.go

**Story 5 Checkpoint**: Complete CRUD lifecycle with cascade delete

---

## Phase 8: Polish & Cross-Cutting Concerns

**Purpose**: Final refinements, documentation, and production readiness

**Estimated Time**: 4-6 hours

### Documentation

- [ ] T100 [P] Generate API documentation from OpenAPI spec (Swagger UI setup)
- [ ] T101 [P] Update README.md with complete setup and usage instructions
- [ ] T102 [P] Create CONTRIBUTING.md with development guidelines
- [ ] T103 [P] Document environment variables in .env.example

### Security Hardening

- [ ] T104 [P] Add rate limiting middleware in internal/presentation/http/middleware/rate_limit.go
- [ ] T105 [P] Add security headers middleware (HSTS, CSP, X-Frame-Options)
- [ ] T106 [P] Implement request ID tracking for audit logs
- [ ] T107 [P] Add input sanitization for log output (prevent log injection)

### Performance Optimization

- [ ] T108 [P] Verify MongoDB indexes are created and used (EXPLAIN queries)
- [ ] T109 [P] Add context timeouts to all database operations (5s writes, 3s reads)
- [ ] T110 [P] Configure connection pool parameters for production load
- [ ] T111 [P] Add response compression middleware (gzip)

### Monitoring & Observability

- [ ] T112 [P] Add request metrics (response times, status codes) to logs
- [ ] T113 [P] Create performance profiling endpoints (pprof)
- [ ] T114 [P] Add structured logging for all error paths
- [ ] T115 [P] Create startup/shutdown log messages with version info

### Testing & Quality

- [ ] T116 Run full test suite and verify 80%+ coverage
- [ ] T117 Run golangci-lint and fix all issues
- [ ] T118 [P] Create integration test suite runner script
- [ ] T119 [P] Add load test scenarios with expected performance benchmarks

### Deployment Preparation

- [ ] T120 [P] Verify Docker image builds correctly and is <50MB
- [ ] T121 [P] Test docker-compose.yml with all services
- [ ] T122 [P] Create CI/CD pipeline configuration (.github/workflows/ci.yml)
- [ ] T123 [P] Create deployment documentation for BETA and PROD environments

**Final Checkpoint**: Production-ready application with complete CRUD API

---

## Dependency Graph (User Story Completion Order)

```
Phase 1 (Setup) → Phase 2 (Foundation)
                         ↓
        ┌────────────────┴────────────────┐
        ↓                                  ↓
Phase 3 (US1: Create) ←→ Phase 4 (US2: Get)  ← MVP Milestone
        ↓                                  ↓
        └────────────────┬────────────────┘
                         ↓
             Phase 5 (US3: List)
                         ↓
             Phase 6 (US4: Update)
                         ↓
             Phase 7 (US5: Delete)
                         ↓
             Phase 8 (Polish)
```

**Critical Path**: Setup → Foundation → US1+US2 (MVP) → US3 → US4 → US5 → Polish

**Parallel Opportunities**:

- US1 and US2 can be developed simultaneously after Foundation
- Within each phase, tasks marked [P] can run in parallel
- US3, US4, US5 are independent and could be developed in parallel after MVP

---

## Implementation Strategy

### MVP First (Recommended)

1. **Week 1**: Phases 1-4 (Setup + Foundation + US1 + US2)
   - Deliverable: Working API that can create and retrieve persons
   - Value: Core read-write cycle functional
   - Team: Can be single developer or pair

2. **Week 2**: Phases 5-7 (US3 + US4 + US5)
   - Deliverable: Complete CRUD operations
   - Value: Full feature set
   - Team: Can parallelize across stories

3. **Week 3**: Phase 8 (Polish)
   - Deliverable: Production-ready application
   - Value: Security, performance, documentation
   - Team: Cross-functional (dev + ops + docs)

### Parallel Execution Example (US1 Implementation)

```
Developer 1: T038-T045 (Domain entities)
Developer 2: T047-T049 (Use case + DTOs)
Developer 3: T052-T054 (HTTP handlers + DTOs)

After all complete:
Developer 1: T050-T051 (MongoDB repository)
Developer 2: T056-T058 (Entity tests)
Developer 3: T059-T060 (Use case + API tests)
```

---

## Task Summary by Phase

| Phase | Tasks | Parallelizable | Estimated Time | Priority |
|-------|-------|----------------|----------------|----------|
| 1: Setup | 10 | 9 | 1-2 hours | High |
| 2: Foundation | 26 | 21 | 4-6 hours | Critical |
| 3: US1 (Create) | 23 | 16 | 6-8 hours | P1 (MVP) |
| 4: US2 (Get) | 8 | 4 | 3-4 hours | P1 (MVP) |
| 5: US3 (List) | 12 | 5 | 4-5 hours | P2 |
| 6: US4 (Update) | 10 | 5 | 4-5 hours | P2 |
| 7: US5 (Delete) | 9 | 5 | 3-4 hours | P3 |
| 8: Polish | 24 | 20 | 4-6 hours | High |
| **Total** | **123** | **85** | **29-40 hours** | - |

**MVP Completion**: Phases 1-4 = ~14-20 hours (2-3 days)  
**Full Feature Set**: Phases 1-7 = ~25-34 hours (4-5 days)  
**Production Ready**: All phases = ~29-40 hours (5-7 days)

---

## Constitution Validation Checkpoints

After completing each phase, verify constitution principles:

- [ ] **Clean Code**: Functions <20 lines, meaningful names, no magic numbers
- [ ] **Clean Architecture**: Proper layer separation maintained
- [ ] **Code Reusability**: No duplication, shared utilities used
- [ ] **Design Patterns**: Repository, Use Case, DTO patterns applied correctly
- [ ] **Function Cohesion**: Each function does one thing
- [ ] **Performance**: Indexes used, pagination implemented, timeouts configured
- [ ] **Security**: Input validated, no secrets in code, safe error messages
- [ ] **Testability**: 80%+ coverage, integration tests passing
- [ ] **Quality Gates**: Linter passing, tests green

---

## References

- [Feature Specification](spec.md) - User stories and requirements
- [Implementation Plan](plan.md) - Technical architecture
- [Data Model](data-model.md) - Entity structures
- [API Contracts](contracts/openapi.yaml) - Endpoint specifications
- [Quick Start](quickstart.md) - Development setup
- [Research](research.md) - Design decisions
- [Constitution](../../.specify/memory/constitution.md) - Coding principles
- [AGENTS.md](../../AGENTS.md) - Go development guidelines
