package domain

import (
	"context"
	"time"
	"errors"
)

// Domain Entity: Represents the Todo data structure within the application logic
type Todo struct {
	ID          int64
	Title       string
	Description string
	IsCompleted bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

var (
	ErrTitleRequired = errors.New("The task title cannot be left blank")
	ErrTodoNotFound  = errors.New("The task could not be found")
)

type TodoRepository interface {
	CreateTodo(ctx context.Context, title string, description string) (*Todo, error)
	GetTodos(ctx context.Context, search string, filterStatus string) ([]Todo, error)
	GetTodoByID(ctx context.Context, id int64) (*Todo, error)
	UpdateTodo(ctx context.Context, id int64, title string, description string, isCompleted bool) (*Todo, error)
	DeleteTodo(ctx context.Context, id int64) error
}

type TodoService interface {
	GetAll(ctx context.Context, search string, filterStatus string) ([]Todo, error)
	GetByID(ctx context.Context, id int64) (*Todo, error)
	Create(ctx context.Context, title string, description string) (*Todo, error)
	Update(ctx context.Context, id int64, title string, description string, isCompleted bool) (*Todo, error)
	Delete(ctx context.Context, id int64) error
}