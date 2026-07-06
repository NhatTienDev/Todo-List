package service

import (
	"context"
	"strings"

	"github.com/nhattiendev/todo-list/internal/todo/domain"
)

func (s *todoService) GetAll(ctx context.Context, search string, filterStatus string) ([]domain.Todo, error) {
	search = strings.TrimSpace(search)

	filterStatus = strings.ToLower(strings.TrimSpace(filterStatus))
	
	if filterStatus != "completed" && filterStatus != "pending" {
		filterStatus = "all"
	}

	return s.todoRepo.GetTodos(ctx, search, filterStatus)
}