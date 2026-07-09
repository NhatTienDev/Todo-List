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
- Create a PostgreSQL database: right click to the Databases in pgAdmin4 and name the db for you.
- Run the script, the script is the the same as the script of **backend/db/migration/schema.sql** to make sure the script is executed successfully without errors.

### 2. Backend
- Create your own .env file, the content you can see in the .env.example
- Install direct dependencies (main libraries):
```bash
go get github.com/go-chi/chi/v5
go get github.com/go-chi/cors
go get github.com/joho/godotenv
go get github.com/lib/pq
go get github.com/swaggo/http-swagger
go get github.com/swaggo/swag
```

- Install dev tools:
```bash
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
go install github.com/swaggo/swag/cmd/swag@latest
go install github.com/air-verse/air@latest # Use for hot reload
```

- Use SQLC:<br>
From the backend directory, run the following command if you have changes at **schema.sql** or **todos.sql**:
```bash
sqlc generate
```
- Build and run the backend to make the connection to PostgreSQL
From the backend directory, run the following command:
```bash
# If you want to manually build and run the backend when you have change in source code:
go build cmd/api/main.go
./main.exe

# If you want to automatically build and run whenever changes are detected:
air init # Run this one time only when you first time use it, after that don't need to run this command again, you will see .air.toml created
air # Run this every time you start to code, it handles for you the build and run process automatically
```

All things in the backend source code is ready, so after getting all libraries and installing tools.<br>
Just run **air** to start backend server.

API server starts at `http://localhost:8080` (or the port you set).<br>
Swagger UI: `http://localhost:8080/swagger/index.html`.

### 3. Frontend
From the frontend directory:
```bash
npm i
npm run dev
```

Dev server starts at `http://localhost:5173`.

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