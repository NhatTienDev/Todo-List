package service

import (
	"context"
	"strings"

	"github.com/nhattiendev/todo-list/internal/todo/domain"
)

func (s *todoService) Update(ctx context.Context, id int64, title string, description string, isCompleted bool) (*domain.Todo, error) {
	title = strings.TrimSpace(title)
	
	if title == "" {
		return nil, domain.ErrTitleRequired
	}

	description = strings.TrimSpace(description)

	return s.todoRepo.UpdateTodo(ctx, id, title, description, isCompleted)
}