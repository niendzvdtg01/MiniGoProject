# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

MiniGoProject is a full-stack web application with a Go backend and React frontend. The backend is a Gin-based REST API with user and product management endpoints, while the frontend is a React + Vite application.

## Backend Architecture (Go)

### Directory Structure
- **`Backend/cmd/main.go`** - Entry point that initializes the Gin router and starts the server on port 8085
- **`Backend/api/`** - Router setup and HTTP configuration
- **`Backend/internal/handler/`** - HTTP request handlers (UserHandler, ProductHandler)
- **`Backend/internal/service/`** - Business logic layer (currently minimal; interfaces defined for future implementation)
- **`Backend/internal/repository/`** - Data access layer (database.go has stub implementation)
- **`Backend/internal/model/`** - Domain models (User, Product) with Public() methods for API responses
- **`Backend/pkg/dto/`** - Data Transfer Objects with validation rules (UserRequest, PostProduct, ProductRequest)
- **`Backend/pkg/utils/`** - Shared utilities, especially validation.go which handles custom validation rules

### API Routes
```
POST /api/user/info         - Get user by name (expects UserRequest DTO)
GET  /api/user/:id          - Get user by ID (path parameter)
POST /api/product/all       - Create/list products (expects PostProduct DTO)
POST /api/product/product_info - Get product info (expects ProductRequest DTO)
```

### Key Dependencies
- **Gin v1.12.0** - Web framework
- **go-playground/validator/v10** - Validation library with custom rules
- **go-sql-driver/mysql** - MySQL driver (imported but not yet integrated)

### Custom Validation Rules
The project implements custom validators in `validation.go`:
- `min_int` / `max_int` - Integer range constraints
- `file_extension` - File type validation (jpg, png, mp4)
- Standard rules: min, max, required, gt, uuid, slug, search

### Current State
- **Database**: MySQL configuration exists (.env) but ConnectDB() in repository/database.go is not wired into handlers
- **Middleware**: Empty middlewares directory; no auth/logging middleware implemented
- **Services**: Service layer interfaces exist but are not integrated with handlers
- **Error Handling**: Basic error responses using gin.H; no structured error handling or logging

## Frontend Architecture (React + Vite)

Located in `Frontend/TeamSharingProject/`:
- **Framework**: React 19.2.5 with Vite as build tool
- **Scripts**: `npm run dev` (development), `npm run build` (production), `npm run lint`, `npm run preview`
- **Entry point**: index.html
- **Config**: vite.config.js

## Development Commands

### Backend (Go)

**Run the server**:
```bash
cd Backend && go run ./cmd/main.go
```
Server starts on `http://localhost:8085`

**Build executable**:
```bash
cd Backend && go build -o app ./cmd/main.go
```

**Run tests**:
```bash
cd Backend && go test ./...
```

**Run specific test**:
```bash
cd Backend && go test -run TestName ./path/to/package
```

**Lint/Format**:
```bash
cd Backend && go fmt ./...
```

**Check dependencies**:
```bash
cd Backend && go mod tidy
```

**Debug in VS Code**:
Use the "Launch Program" configuration in `.vscode/launch.json` (already configured)

### Frontend (React)

**Start dev server** (Vite with HMR):
```bash
cd Frontend/TeamSharingProject && npm run dev
```

**Build for production**:
```bash
cd Frontend/TeamSharingProject && npm run build
```

**Lint code**:
```bash
cd Frontend/TeamSharingProject && npm run lint
```

**Preview production build**:
```bash
cd Frontend/TeamSharingProject && npm run preview
```

## Configuration

**Backend (.env)**:
```
DB_HOST = localhost 
DB_PORT = 3310
DB_USER = admin
DB_PASSWORD = 1234
DB_NAME = 
```
Note: Database name is not configured; this needs to be set before the application can use the database.

## Known Limitations

- Database layer is not integrated with handlers
- No middleware for auth, logging, or error handling
- Service layer interfaces defined but not used by handlers
- No structured error responses across API
- Frontend React structure is basic/untested

