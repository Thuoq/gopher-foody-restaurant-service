# Gopher Foody Restaurant Service

The Restaurant Service manages restaurant listings and their menus for the Gopher Foody ecosystem. It handles restaurant creation, food item management, and inventory tracking.

## 📂 Project Structure

```text
.
├── cmd/server/main.go       # Application entry point
├── internal/
│   ├── application/         # Business logic (Use cases)
│   ├── core/                # Domain entities (Restaurant, Food, Category)
│   ├── infrastructure/      # DB Repositories (GORM/Postgres)
│   └── presentation/        # HTTP Layer
│       └── http/
│           ├── handlers/    # API Handlers
│           │   ├── admin/   # Merchant APIs (with DTOs)
│           │   └── user/    # Customer APIs (with DTOs)
│           └── middleware/  # Gateway Auth integration
├── migrations/              # SQL Up/Down migration files
└── pkg/                     # Common utilities (Response, Logger)
```

## 🚀 Getting Started

### 1. Environment Setup
Create a `.env` file from `.env.example`:
```env
APP_HTTP_PORT=8081
DATABASE_URL=postgres://user:pass@localhost:5432/foody_restaurant_db?sslmode=disable
LOGGER_LEVEL=debug
```

### 2. Database Migrations
Run migrations to set up the database schema:
```bash
make migrate-up
```

### 3. Run the Service
```bash
go mod tidy
go run cmd/server/main.go
```

The service will be available at `http://localhost:8081`.

## 🛠 Features
- **Admin/Merchant API**: Create and manage restaurants.
- **Menu Management**: Add, update, and remove food items with inventory tracking.
- **Identity Awareness**: Integrated with Gopher Gateway to identify owners via `X-User-Id` header.
