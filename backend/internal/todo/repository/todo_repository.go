package repository

import (
	"context"
	"database/sql"

	"github.com/nhattiendev/todo-list/internal/todo/domain"
	"github.com/nhattiendev/todo-list/internal/todo/repository/sqlc"
)

type todoRepository struct {
	q *sqlc.Queries
}

func NewTodoRepository(db *sql.DB) domain.TodoRepository {
	return &todoRepository{
		q: sqlc.New(db),
	}
}

func mapToTodoDomain(dbTodo sqlc.Todo) *domain.Todo {
	return &domain.Todo{
		ID:          dbTodo.ID,
		Title:       dbTodo.Title,
		Description: dbTodo.Description,
		IsCompleted: dbTodo.IsCompleted,
		CreatedAt:   dbTodo.CreatedAt,
		UpdatedAt:   dbTodo.UpdatedAt,
	}
}

func (r *todoRepository) CreateTodo(ctx context.Context, title string, description string) (*domain.Todo, error) {
	arg := sqlc.CreateTodoParams{
		Title:       title,
		Description: description,
	}

	dbTodo, err := r.q.CreateTodo(ctx, arg)
	if err != nil {
		return nil, err
	}

	return mapToTodoDomain(dbTodo), nil
}

func (r *todoRepository) GetTodos(ctx context.Context, search string, filterStatus string) ([]domain.Todo, error) {
	arg := sqlc.GetTodosParams{
		Search:       search,
		FilterStatus: filterStatus,
	}

	dbTodo, err := r.q.GetTodos(ctx, arg)
	if err != nil {
		return nil, err
	}

	var todos []domain.Todo
	for _, t := range dbTodo {
		todos = append(todos, *mapToTodoDomain(t))
	}

	return todos, nil
}

func (r *todoRepository) GetTodoByID(ctx context.Context, id int64) (*domain.Todo, error) {
	dbTodo, err := r.q.GetTodoByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return mapToTodoDomain(dbTodo), nil
}

func (r *todoRepository) UpdateTodo(ctx context.Context, id int64, title string, description string, isCompleted bool) (*domain.Todo, error) {
	arg := sqlc.UpdateTodoParams{
		ID:          id,
		Title:       title,
		Description: description,
		IsCompleted: isCompleted,
	}

	dbTodo, err := r.q.UpdateTodo(ctx, arg)
	if err != nil {
		return nil, err
	}

	return mapToTodoDomain(dbTodo), nil
}

func (r *todoRepository) DeleteTodo(ctx context.Context, id int64) error {
	return r.q.DeleteTodo(ctx, id)
}