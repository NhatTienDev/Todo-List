package domain

import (
	"context"
	"time"
	"errors"
)

// Domain Entity: Represents the Todo data structure within the application logic
type Todo struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	IsCompleted bool      `json:"is_completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type PaginatedResult struct {
	Data  []Todo `json:"data"`
	Total int64  `json:"total"`
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
}

var (
	ErrTitleRequired      = errors.New("The task title cannot be left blank")
	ErrTodoNotFound       = errors.New("The task could not be found")
	ErrInvalidFilterValue = errors.New("Invalid filter status value")
)

type TodoRepository interface {
	CreateTodo(ctx context.Context, title string, description string) (*Todo, error)
	GetTodos(ctx context.Context, search string, filterStatus string, page int, limit int) (*PaginatedResult, error)
	GetTodoByID(ctx context.Context, id int64) (*Todo, error)
	UpdateTodo(ctx context.Context, id int64, title string, description string, isCompleted bool) (*Todo, error)
	DeleteTodo(ctx context.Context, id int64) error
}

type TodoService interface {
	GetAll(ctx context.Context, search string, filterStatus string, page int, limit int) (*PaginatedResult, error)
	GetByID(ctx context.Context, id int64) (*Todo, error)
	Create(ctx context.Context, title string, description string) (*Todo, error)
	Update(ctx context.Context, id int64, title *string, description *string, isCompleted *bool) (*Todo, error)
	Delete(ctx context.Context, id int64) error
}
