# RABC - GO

A robust Role-Based Access Control (RBAC) system implemented in Go using Echo framework, GORM ORM, Echo-JWT middleware, Go-Playground validator, and PostgreSQL database. Designed with best practices and tailored for SaaS applications.

---

## Features

- User signup and signin with JWT authentication and secure cookie management
- Role management with flexible comma-separated roles stored in Postgres
- Role-based authorization middleware using Echo-JWT
- User profile management with relational model (User <-> Profile)
- Pagination, filtering, and search for user listing APIs
- Secure password hashing and update password functionality
- Input validation using `go-playground/validator`
- Modular and clean code structure following Go best practices
- PostgreSQL as the persistent store with GORM for ORM
- Middleware support for CORS, logging, error handling, and recovery
- Easily extensible for complex RBAC rules and SaaS multi-tenant models

---

## Tech Stack

- **Language:** Go (Golang)
- **Web framework:** Echo
- **ORM:** GORM
- **Authentication:** JWT with Echo-JWT middleware
- **Validation:** go-playground/validator
- **Database:** PostgreSQL
- **Dependency Management:** Go modules

---

## Configuration

This project uses a TOML configuration file for local development credentials.

Create a `dev.toml` file in the root directory with the following example content:

```toml
[database]
db_host = "localhost"
db_port = "5432"
db_user = "postgres"
db_password = "sahil"
db_name = "RBAC-GO"

[jwt]
cookie_name = "rbac-cookie"
secret = "K3#v@9$1!pZ^mL2&uQ7*rF4)gT8_W+oB"

origins = ["http://localhost:3000", "https://example.com"]
```

## Prerequisites

- Go 1.20+
- PostgreSQL 13+
- Git

## Installation

1. Run the command to copy the project and download all the packages

   ```bash
   git clone https://github.com/Sahil2k07/RBAC-GO.git
   cd rbac-go
   go mod download
   ```

2. Migrate the Models

   ```bash
   go run ./cmd/migrate/main.go
   ```

3. Start the server

   ```bash
   go run ./cmd/server/main.go
   ```

## Folder Structrue

```graphql
.
├── cmd/            # Entry points (server, migration)
├── internal/
│   ├── config/     # Configuration structs & helpers
│   ├── database/   # DB connection and migration files
│   ├── enum/       # Role enums and helpers
│   ├── handler/    # HTTP handlers (Echo controllers)
│   ├── interface/  # Interfaces for services and repositories
│   ├── model/      # GORM models (User, Profile, etc.)
│   ├── repository/ # DB repositories implementing interfaces
│   ├── service/    # Business logic and service layer
│   ├── util/       # Utilities (validation, pagination, error handling)
│   └── view/       # Request and response structs (DTOs)
├── vendor/         # Vendored dependencies (if any)
├── go.mod
└── go.sum
```
