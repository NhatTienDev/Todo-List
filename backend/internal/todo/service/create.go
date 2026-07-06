package service

import (
	"context"
	"strings"

	"github.com/nhattiendev/todo-list/internal/todo/domain"
)

func (s *todoService) Create(ctx context.Context, title string, description string) (*domain.Todo, error) {
	title = strings.TrimSpace(title)
	
	if title == "" {
		return nil, domain.ErrTitleRequired
	}

	description = strings.TrimSpace(description)

	return s.todoRepo.CreateTodo(ctx, title, description)
}