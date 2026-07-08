package service

import (
	"context"
	"strings"

	"github.com/nhattiendev/todo-list/internal/todo/domain"
)

func (s *todoService) GetAll(ctx context.Context, search string, filterStatus string, page int, limit int) (*domain.PaginatedResult, error) {
	search = strings.TrimSpace(search)

	filterStatus = strings.ToLower(strings.TrimSpace(filterStatus))

	switch filterStatus {
	case "true", "completed":
		filterStatus = "completed"
	case "false", "pending":
		filterStatus = "pending"
	case "":
		filterStatus = "all"
	default:
		return nil, domain.ErrInvalidFilterValue
	}

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	return s.todoRepo.GetTodos(ctx, search, filterStatus, page, limit)
}
