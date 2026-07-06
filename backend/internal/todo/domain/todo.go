package domain

import (
	"context"
	"time"
)

// Domain Entity: Represents the Todo data structure within the application logic
type Todo struct {
	ID          int32     
	Title       string    
	Description string    
	IsCompleted bool      
	CreatedAt   time.Time 
	UpdatedAt   time.Time 
}

type TodoRepository interface {
	GetTodos(ctx context.Context, search string, filterStatus string) ([]Todo, error)
	GetTodoByID(ctx context.Context, id int32) (*Todo, error)
	CreateTodo(ctx context.Context, title string, description string) (*Todo, error)
	UpdateTodo(ctx context.Context, id int32, title string, description string, isCompleted bool) (*Todo, error)
	DeleteTodo(ctx context.Context, id int32) error
}

type TodoService interface {
	GetAll(ctx context.Context, search string, filterStatus string) ([]Todo, error)
	GetByID(ctx context.Context, id int32) (*Todo, error)
	Create(ctx context.Context, title string, description string) (*Todo, error)
	Update(ctx context.Context, id int32, title string, description string, isCompleted bool) (*Todo, error)
	Delete(ctx context.Context, id int32) error
}