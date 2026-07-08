# Simpe Todo List App
A full-stack Todo List application built with **Go** (backend) and **React** (frontend).

## Features
- Create, read, update and delete tasks
- Mark tasks as completed / pending
- Search by title
- Filter by status (completed / pending / all)
- Pagination (10 tasks per page)
- Responsive design (desktop, tablet, mobile)

## Tech Stack
| Layer    | Technology                     |
| -------- | ------------------------------ |
| Backend  | Go 1.26.1, Chi router          |
| Database | PostgreSQL 18.3, sqlc          |
| Frontend | React 19, Vite                 |

## Prerequisites
- Go 1.22+ ([Download Go here](https://go.dev/doc/install))
- PostgreSQL 14+ ([Download PostgreSQL (Core engine + pgAdmin4) here](https://www.enterprisedb.com/downloads/postgres-postgresql-downloads)), choose the OS and PostgreSQL version that requires for the project
- Node.js 18+ ([Download Node.js here](https://nodejs.org/en/download))

## Setup
### 1. Database
Create a PostgreSQL database and run the migration:

```bash
psql -U your_user -d your_db -f backend/db/migration/schema.sql
```

### 2. Backend

```bash
cd backend
cp .env.example .env      # then edit DB_URL and BE_PORT
go run ./cmd/api
```

The API server starts at `http://localhost:8081` (or the port you set).
Swagger UI: `http://localhost:8081/swagger/index.html`

### 3. Frontend

```bash
cd frontend
npm install
npm run dev
```

The dev server starts at `http://localhost:5173` and proxies `/api` requests to the backend.

### API Endpoints

| Method | Endpoint               | Description              |
| ------ | ---------------------- | ------------------------ |
| GET    | `/api/v1/todos`        | List tasks (paginated)   |
| GET    | `/api/v1/todos/{id}`   | Get task by ID           |
| POST   | `/api/v1/todos`        | Create a task            |
| PUT    | `/api/v1/todos/{id}`   | Update a task            |
| DELETE | `/api/v1/todos/{id}`   | Delete a task            |

**Query parameters for GET `/api/v1/todos`:**

| Param    | Type   | Default | Description                              |
| -------- | ------ | ------- | ---------------------------------------- |
| search   | string | ""      | Search by title (ILIKE)                  |
| status   | string | ""      | Filter: `completed`, `pending`, `true`, `false` |
| page     | int    | 1       | Page number                              |
| limit    | int    | 10      | Items per page                           |

### Project Structure

```
backend/
├── cmd/api/           # Entry point
├── db/
│   ├── migration/     # SQL schema
│   └── query/         # SQL queries (sqlc)
├── internal/
│   └── todo/
│       ├── domain/    # Entities & interfaces
│       ├── handler/   # HTTP handlers
│       ├── repository/ # Data access (sqlc generated)
│       └── service/   # Business logic
└── response/          # JSON response helpers

frontend/
├── src/
│   ├── api.js         # API client
│   ├── components/
│   │   ├── TodoForm.jsx
│   │   ├── TodoItem.jsx
│   │   ├── TodoList.jsx
│   │   ├── SearchBar.jsx
│   │   ├── Pagination.jsx
│   │   ├── ConfirmModal.jsx
│   │   └── Toast.jsx
│   ├── App.jsx
│   └── App.css
└── vite.config.js
```
