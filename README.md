# Gopher Foody Restaurant Service

A high-performance, scalable microservice for managing restaurants and food menus, built with Golang following Clean Architecture and Domain-Driven Design (DDD) principles.

## 🚀 Features

- **Clean Architecture**: Strict separation of concerns (Domain, Ports, UseCases, Infrastructure, Presentation).
- **Advanced CRUD**: Complete management for Restaurants and Foods with strict ownership validation.
- **Single Purpose UseCases**: Modular application layer where each operation is isolated for maintainability and testability.
- **Media Management**: AWS S3 integration using **Presigned URLs** for secure, direct-to-client uploads.
- **Pagination & Search**: Standardized pagination utility and case-insensitive search (ILIKE).
- **Security**: Robust authorization checks to ensure only restaurant owners can modify their resources.
- **Dependency Injection**: Compile-time/Runtime DI using `uber-go/dig`.
- **Database Migrations**: Versioned SQL migrations for schema management.

## 🏗 Project Structure

```text
.
├── cmd/
│   └── server/             # Application entry point (main.go)
├── internal/
│   ├── application/
│   │   └── usecases/       # Single-purpose business logic
│   │       ├── food/       # AdminCreate, AdminUpdate, UserListMenu, etc.
│   │       ├── media/      # GetUploadURL
│   │       └── restaurant/ # AdminCreate, AdminListMy, UserList, etc.
│   ├── core/
│   │   ├── domain/         # Domain entities (Restaurant, Food, etc.)
│   │   └── ports/          # Interfaces for repositories and usecases
│   ├── infrastructure/
│   │   ├── database/       # GORM implementation and repositories
│   │   └── storage/        # AWS S3 implementation
│   └── presentation/
│       └── http/           # Gin handlers, DTOs, and routing
├── pkg/
│   ├── logger/             # Zap logger wrapper
│   ├── response/           # Standardized API response utility
│   └── utils/              # Pagination and other helpers
├── migrations/             # SQL migration files (Up/Down)
└── .env.example            # Environment variables template
```

## 🛠 Tech Stack

- **Language**: Go (Golang)
- **Web Framework**: [Gin Gonic](https://github.com/gin-gonic/gin)
- **ORM**: [GORM](https://gorm.io/)
- **Database**: PostgreSQL
- **Storage**: AWS S3
- **Configuration**: Viper
- **Dependency Injection**: Uber-go Dig
- **Logging**: Zap Logger

## 🚦 Getting Started

### 1. Prerequisites
- Go 1.21+
- PostgreSQL
- AWS Account (S3 Bucket)

### 2. Environment Setup
Copy the example environment file and fill in your details:
```bash
cp .env.example .env
```

### 3. Database Migrations
Ensure your database is running and apply migrations:
```bash
# Using golang-migrate
migrate -path migrations -database "postgres://user:pass@localhost:5432/dbname?sslmode=disable" up
```

### 4. Running the Service
```bash
go mod tidy
go run cmd/server/main.go
```

