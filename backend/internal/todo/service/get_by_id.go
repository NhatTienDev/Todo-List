package service

import (
	"context"

	"github.com/nhattiendev/todo-list/internal/todo/domain"
)

func (s *todoService) GetByID(ctx context.Context, id int64) (*domain.Todo, error) {
	return s.todoRepo.GetTodoByID(ctx, id)
}