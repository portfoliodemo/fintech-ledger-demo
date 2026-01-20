# Fintech Ledger Demo (Go)

This project is a small backend-focused demo that models a simplified financial ledger.
It is designed to demonstrate backend fundamentals relevant to production fintech systems,
including clear separation of concerns, testable business logic, and deterministic state handling.

The scope is intentionally small, but the patterns reflect real-world backend design.

---

## Architecture Overview

The application is structured into four main layers:

- cmd/server
- internal/models
- internal/storage
- internal/service

### Models
The `models` package defines core domain entities such as `Transaction`.
Balances are not stored directly; instead, they are derived from immutable transaction records.
This mirrors real-world ledger systems and supports auditability and correctness.

### Storage
The `storage` package defines a repository interface along with an in-memory implementation.
This abstraction decouples business logic from persistence concerns and enables fast,
deterministic unit testing without external dependencies.

### Service
The `service` package contains the core business logic and enforces domain invariants, such as:
- Credits must be positive
- Debits cannot exceed the available balance

This layer is intentionally unaware of HTTP, JSON, or transport concerns.

### Transport (HTTP)
The HTTP server acts as a thin transport layer responsible only for:
- Request parsing
- Response encoding
- Routing

All business rules are delegated to the service layer.

---

## API Endpoints

- `POST /credit`
- `POST /debit`
- `GET /balance?user_id=1`

---

## Testing Strategy

Business logic is tested using table-driven unit tests against an in-memory repository.
This approach validates both successful and failure scenarios while keeping tests fast,
isolated, and deterministic.

---

## Tooling & Workflow

- Go modules for dependency management
- Table-driven unit testing
- Git for version control
- Linux-based development environment

---

## Notes

AI tooling was used to accelerate initial scaffolding and iteration.
All generated code was reviewed, refactored, and tested manually to ensure a clear
understanding of design decisions and tradeoffs.
