# Nebula Bridge

Nebula Bridge is a modular Web3 infrastructure toolkit implemented in Go using the Gin web framework and GORM ORM. It provides a lightweight HTTP API for managing nodes and ingesting chain events (telemetry/relays), designed for easy extension to additional Web3 services.

Key features
- Clean layered architecture: HTTP handlers, services, repositories, models
- GORM-backed storage with Postgres (recommended) and SQLite fallback for local dev/tests
- Simple migration and CLI entry points
- Unit tests for services and handlers

Architecture

- cmd/
  - server: application entrypoint that wires config, DB, router and handlers
  - migrate: small helper to run auto-migrations (optional)
- internal/
  - config: environment and config loader
  - db: DB initialization (GORM)
  - models: application data models (Node, Event)
  - repository: repository interfaces and GORM implementations
  - service: business logic and validations
  - handlers: HTTP handlers using Gin
- migrations/: SQL migration files (optional)

Request flow

Client -> Gin HTTP handlers -> Service layer -> Repository (GORM) -> Database

Installation

Prerequisites
- Go 1.20+
- Docker & Docker Compose (optional for running Postgres)

Clone repository

  git clone https://github.com/yourorg/nebula-bridge.git
  cd nebula-bridge

Build

  make build

Run (development, SQLite fallback)

  cp .env.example .env
  # Edit .env to configure DATABASE_URL if you want Postgres
  make run

Run with Docker Compose (Postgres)

  docker-compose up --build
  # App will be available on PORT configured in .env (default 8080)

Environment variables

- PORT: HTTP server port (default: 8080)
- DATABASE_URL: Postgres DSN (if empty, uses SQLite file: nebula.db)
- LOG_LEVEL: log verbosity (info|debug)

Usage (HTTP endpoints)

- GET /health - simple health check
- POST /nodes - create a node
  - payload: {"name":"node-1","address":"0x123..."}
- GET /nodes - list nodes
- POST /events - record an event
  - payload: {"node_id":1,"type":"relay","payload":"{...}"}
- GET /events - list events (supports ?node_id=)

Testing

  make test

Docker

- Dockerfile provided for building a container image
- docker-compose.yml sets up Postgres and the app for local integration

Development notes

- Follow the layered architecture: add features in service layer and persist in repository implementations
- Use in-memory SQLite for unit tests and small CI runs

Contributing

Contributions welcome. Please open issues for feature requests or bug reports. For larger changes, open a pull request with tests and an explanation of your design choices.
