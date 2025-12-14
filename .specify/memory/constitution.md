<!--
Sync Impact Report:
- Version: 0.0.0 → 1.0.0
- Ratification Date: 2025-12-14
- Modified Principles: N/A (initial version)
- Added Sections: All core principles established
- Removed Sections: None
- Templates Status: All templates compatible ✅
- Follow-up TODOs: None
-->

# Cadastro Pessoas Constitution

## Core Principles

### I. Clean Code (NON-NEGOTIABLE)

Code MUST prioritize readability and maintainability above all else:

- **Meaningful Names**: Variables, functions, and classes MUST have descriptive, intention-revealing names
- **Single Responsibility**: Every function/method MUST do one thing and do it well
- **Small Functions**: Functions MUST be short (ideally < 20 lines), focused, and operate at a single level of abstraction
- **No Magic Numbers**: Use named constants instead of literal values
- **Comments Only When Necessary**: Code MUST be self-documenting; comments explain "why", not "what"
- **Error Handling**: Use exceptions, not error codes; provide meaningful error messages
- **DRY Principle**: Don't Repeat Yourself - eliminate code duplication through abstraction

**Rationale**: Clean code reduces bugs, accelerates onboarding, and makes maintenance predictable. Technical debt compounds exponentially; preventing it at the source is non-negotiable.

### II. Clean Architecture

System architecture MUST follow dependency inversion and separation of concerns:

- **Layer Independence**: Business logic MUST NOT depend on frameworks, UI, databases, or external services
- **Dependency Rule**: Dependencies MUST point inward - outer layers depend on inner layers, never the reverse
- **Core Layers Structure**:
  - **Entities**: Enterprise business rules (domain models)
  - **Use Cases**: Application-specific business rules (interactors/services)
  - **Interface Adapters**: Controllers, presenters, gateways (adapters between use cases and external world)
  - **Frameworks & Drivers**: External tools, frameworks, databases (outermost layer)
- **Interfaces at Boundaries**: Use interfaces/protocols to define contracts between layers
- **Independent Testability**: Core business logic MUST be testable without frameworks, databases, or external dependencies

**Rationale**: Clean Architecture ensures the system is maintainable, testable, and adaptable to change. Business rules remain stable while technologies evolve.

### III. Code Reusability & DRY

Code reuse MUST be maximized through proper abstraction:

- **Extract Common Logic**: Identify and extract repeated patterns into reusable functions, classes, or modules
- **Utility Functions**: Create utility libraries for cross-cutting concerns (validation, formatting, data transformation)
- **Composition Over Inheritance**: Favor composition and interfaces over deep inheritance hierarchies
- **Generic Solutions**: Design components to be generic and configurable rather than specific and hardcoded
- **Shared Kernel**: Establish a shared core library for domain-agnostic utilities used across features
- **No Copy-Paste**: If code is copied, it MUST be refactored into a shared abstraction

**Rationale**: Reusability reduces maintenance burden, ensures consistency, and accelerates feature development. Changes propagate through one source of truth.

### IV. Design Patterns

Proven design patterns MUST be applied appropriately to solve common problems:

- **Creational Patterns**: Factory, Builder, Singleton (use sparingly) for object creation
- **Structural Patterns**: Adapter, Facade, Decorator, Composite for organizing code structure
- **Behavioral Patterns**: Strategy, Observer, Command, Template Method for defining interactions
- **Pattern Appropriateness**: Patterns MUST solve actual problems, not be applied for their own sake
- **Document Pattern Usage**: When using a pattern, document it in code comments or architecture docs
- **Repository Pattern**: Data access MUST go through repository interfaces
- **Dependency Injection**: Dependencies MUST be injected, not instantiated internally

**Rationale**: Design patterns provide battle-tested solutions to recurring problems, improve communication through shared vocabulary, and make code more maintainable.

### V. Function & Method Cohesion

Functions and methods MUST exhibit high cohesion and low coupling:

- **Do One Thing**: Each function MUST have a single, well-defined purpose
- **Single Level of Abstraction**: All operations within a function MUST be at the same level of abstraction
- **Minimal Parameters**: Functions SHOULD have ≤3 parameters; use objects for complex parameter sets
- **No Side Effects**: Functions MUST NOT have hidden side effects; state changes MUST be explicit
- **Command-Query Separation**: Functions either change state (command) or return data (query), not both
- **Pure Functions Preferred**: Favor pure functions (same input → same output, no side effects) when possible
- **Testable Units**: Every function MUST be independently testable

**Rationale**: Cohesive functions are easier to understand, test, reuse, and maintain. They reduce cognitive load and minimize the risk of bugs.

## Code Quality Standards

### Mandatory Quality Gates

All code MUST pass these gates before merging:

- **Linting**: Zero linting errors using project-configured linter
- **Formatting**: Code MUST be auto-formatted with consistent style
- **Unit Tests**: Minimum 80% code coverage for business logic
- **Integration Tests**: All external integrations MUST have contract tests
- **Code Review**: At least one peer review approval required
- **Documentation**: Public APIs and complex algorithms MUST be documented
- **Performance**: No regression in performance benchmarks

### Code Review Checklist

Reviewers MUST verify:

1. Adherence to all Core Principles (I-V)
2. No violations of Clean Code practices
3. Proper layer separation (Clean Architecture)
4. Appropriate use of design patterns
5. Elimination of code duplication
6. Function cohesion and naming quality
7. Test coverage and quality
8. Documentation completeness

## Development Workflow

### Feature Development Process

1. **Specification**: Start with clear requirements using `/speckit.specify`
2. **Planning**: Define architecture using `/speckit.plan` - verify Clean Architecture principles
3. **Task Breakdown**: Create actionable tasks using `/speckit.tasks`
4. **Implementation**: Follow TDD approach - write tests first, then implement
5. **Refactoring**: Continuously refactor to maintain Clean Code standards
6. **Review**: Peer review against constitution principles
7. **Integration**: Merge only after all quality gates pass

### Refactoring Policy

Continuous refactoring is MANDATORY:

- **Boy Scout Rule**: Leave code cleaner than you found it
- **Scheduled Refactoring**: Allocate 20% of sprint time to technical debt reduction
- **Refactoring Triggers**: Refactor when you see duplication, unclear names, long functions, or violation of principles
- **Safe Refactoring**: Refactor with test coverage; tests MUST pass before and after

## Governance

### Constitution Authority

This constitution supersedes all other development practices and guidelines. All code and architectural decisions MUST be evaluated against these principles.

### Amendment Process

Amendments to this constitution require:

1. **Proposal**: Written proposal with rationale and impact analysis
2. **Review**: Team review and discussion
3. **Approval**: Consensus or majority vote
4. **Documentation**: Update this document with version increment
5. **Migration Plan**: For changes affecting existing code, create migration plan
6. **Communication**: Announce changes to all team members

### Compliance & Enforcement

- All pull requests MUST verify compliance with constitution principles
- Constitution violations MUST be justified and documented in `plan.md` Complexity Tracking section
- Use `.specify/templates/agent-file-template.md` for runtime development guidance
- Periodic audits to ensure ongoing compliance
- New team members MUST review and acknowledge this constitution

### Versioning

This constitution follows semantic versioning:

- **MAJOR**: Backward-incompatible principle changes or removals
- **MINOR**: New principles or material expansions to existing ones
- **PATCH**: Clarifications, wording improvements, non-semantic refinements

**Version**: 1.0.0 | **Ratified**: 2025-12-14 | **Last Amended**: 2025-12-14
