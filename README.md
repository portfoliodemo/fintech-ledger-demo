# Fintech Ledger Demo (Go)

This project is a small backend-focused demo built to model a simplified financial ledger.
It was designed to demonstrate backend fundamentals relevant to production fintech systems:
clear separation of concerns, testable business logic, and deterministic state handling.

The scope is intentionally small but representative of real-world patterns.

---

## Architecture Overview

The application is structured into four main layers:

cmd/server
internal/models
internal/storage
internal/service

### Models
The `models` package defines core domain entities such as `Transaction`.
Balances are not stored directly; instead, they are derived from immutable transactions.
This mirrors real-world ledger and auditability practices.

### Storage
The `storage` package defines a repository interface and an in-memory implementation.
The repository abstraction allows the business logic to remain independent of persistence
details and enables fast, deterministic testing.

### Service
The `service` package contains the core business logic.
It enforces invariants such as:
- Credits must be positive
- Debits cannot exceed available balance

This layer is intentionally unaware of HTTP or JSON.

### Transport (HTTP)
The HTTP server is a thin layer responsible only for:
- Request parsing
- Response encoding
- Routing

All business rules live in the service layer.

---

## API Endpoints

POST /credit  
POST /debit  
GET  /balance?user_id=1

---

## Testing Strategy

Business logic is tested using table-driven tests against an in-memory repository.
This allows validation of both successful and failure scenarios without external dependencies.

---

## Tooling & Workflow

- Go modules for dependency management
- Table-driven unit tests
- Git for version control
- Linux development environment

---

## Notes

AI tooling was used to accelerate initial scaffolding.
All code was reviewed, refactored, and tested manually to ensure full understanding
of design decisions and tradeoffs.
