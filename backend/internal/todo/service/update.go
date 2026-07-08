package service

import (
	"context"
	"strings"

	"github.com/nhattiendev/todo-list/internal/todo/domain"
)

func (s *todoService) Update(ctx context.Context, id int64, title *string, description *string, isCompleted *bool) (*domain.Todo, error) {
	existing, err := s.todoRepo.GetTodoByID(ctx, id)
	if err != nil {
		return nil, domain.ErrTodoNotFound
	}

	newTitle := existing.Title
	if title != nil {
		trimmed := strings.TrimSpace(*title)
		if trimmed == "" {
			return nil, domain.ErrTitleRequired
		}
		newTitle = trimmed
	}

	newDescription := existing.Description
	if description != nil {
		newDescription = strings.TrimSpace(*description)
	}

	newIsCompleted := existing.IsCompleted
	if isCompleted != nil {
		newIsCompleted = *isCompleted
	}

	return s.todoRepo.UpdateTodo(ctx, id, newTitle, newDescription, newIsCompleted)
}