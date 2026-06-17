# Go Backend Template

A production-ready Go backend starter with clean architecture, sqlc code generation, and PostgreSQL. Use this as a template to bootstrap new Go API services.

## Features

- **Clean Architecture** — handler → module → storage layers with interface-based dependency injection
- **sqlc** — type-safe SQL with Go code generation from raw queries
- **PostgreSQL** — connection pooling via `pgx/v4`, migrations via `golang-migrate`
- **Gin HTTP framework** — routing, middleware (CORS, rate limiter, request logging)
- **Structured logging** — contextual logging via `zap` with request-id/user-id propagation
- **Config management** — YAML config with environment variable overrides (`APPLICATION_` prefix)
- **Docker Compose** — one-command PostgreSQL setup
- **Hot reload** — `air` support for development
- **Swagger** — API documentation generation
- **Mock generation** — `mockgen` for all layers

## Prerequisites

- Go 1.25+
- [sqlc](https://sqlc.dev/) — `go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest`
- [golang-migrate](https://github.com/golang-migrate/migrate) — `go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest`
- [air](https://github.com/air-verse/air) (optional, for hot reload) — `go install github.com/air-verse/air@latest`
- [mockgen](https://github.com/uber-go/mock) — `go install go.uber.org/mock/mockgen@latest`
- Docker & Docker Compose

## Getting Started

### 1. Create a new repository from this template

On GitHub, click **"Use this template"** to create a new repo, then clone it:

```bash
git clone https://github.com/your-org/your-new-project.git
cd your-new-project
```

### 2. Update the module name

```bash
go mod edit -module github.com/your-org/your-new-project
```

Update all imports in the codebase to match your new module path.

### 3. Start the database

```bash
make up
```

### 4. Run migrations

```bash
make migrate-up
```

### 5. Start the server

```bash
make run
# or with hot reload:
make air
```

The server starts at `http://localhost:8097`.

## Project Structure

```
├── cmd/
│   └── main.go              # Entry point
├── config/
│   ├── config.yml           # App configuration (DB, server)
│   └── sqlc.yml             # sqlc code generation config
├── initiator/               # Composition root (wires all layers)
│   ├── initiator.go         # Orchestrator
│   ├── initiator-config.go  # Viper config loader
│   ├── initiator-db.go      # pgx pool initialization
│   ├── initiator-handler.go
│   ├── initiator-logger.go
│   ├── initiator-module.go
│   ├── initiator-persistance.go
│   └── initiator-routing.go
├── internal/
│   ├── constant/
│   │   ├── db/
│   │   │   ├── queries/     # Raw SQL queries (sqlc input)
│   │   │   ├── schemas/     # Migration files
│   │   │   ├── generated/   # sqlc output (do not edit)
│   │   │   └── db_interface/ # sqlc adapter
│   │   ├── dto/             # Request/response types
│   │   ├── errors/          # Typed domain errors
│   │   └── response/        # API response helpers
│   ├── storage/             # Data access layer (repository)
│   ├── module/              # Business logic layer (service)
│   ├── handler/             # HTTP layer (controllers)
│   │   └── middleware/      # CORS, logger, rate limiter
│   └── glue/                # Route registration
├── platform/
│   └── logger/              # Logger abstraction (zap wrapper)
├── docker-compose.yaml      # PostgreSQL service
└── makefile                 # Common tasks
```

## Architecture

```
HTTP Request
  → Gin Router (glue/)
    → Handler (handler/) — parse request, validate, respond
      → Module (module/) — business logic, error wrapping
        → Storage (storage/) — data access
          → sqlc Queries (constant/db/generated/) — type-safe SQL
            → pgxpool → PostgreSQL
```

Each layer defines its interface in a central file (`storage.go`, `module.go`, `handler.go`) with concrete implementations in sub-packages. This makes testing straightforward — mock any layer and inject it.

## Adding a New Feature

1. **SQL queries & migrations** — add `.sql` files to `internal/constant/db/queries/` and migration files to `internal/constant/db/schemas/`
2. **Generate code** — `make sqlc`
3. **Storage** — add interface method in `internal/storage/storage.go`, implement in a new sub-package
4. **Module** — add interface method in `internal/module/module.go`, implement in a new sub-package
5. **Handler** — add interface method in `internal/handler/handler.go`, implement in a new sub-package
6. **Routes** — register in a new file under `internal/glue/`
7. **Wire it up** — add constructor calls in `initiator/`

## Common Tasks

```bash
make up              # Start PostgreSQL
make down            # Stop PostgreSQL
make run             # Run the server
make air             # Run with hot reload
make sqlc            # Regenerate sqlc code
make migrate-up      # Run all pending migrations
make migrate-down    # Rollback migrations
make migrate-create  # Create a new migration (name=your_name)
make generate-mocks  # Regenerate mocks for testing
make test            # Run tests
```

## Customization Checklist

- [ ] Update `go.mod` module name
- [ ] Update imports across the codebase
- [ ] Change `config/config.yml` DB credentials if needed
- [ ] Update `docker-compose.yaml` DB credentials to match
- [ ] Replace the sample feature (`sample`) with your domain
- [ ] Update `makefile` migration paths if restructuring
- [ ] Remove or update the Swagger annotations
- [ ] Configure your CI/CD pipeline

## License

MIT
