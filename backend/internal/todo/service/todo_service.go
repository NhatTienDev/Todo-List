package service

import "github.com/nhattiendev/todo-list/internal/todo/domain"

type todoService struct {
	todoRepo domain.TodoRepository
}

func NewTodoService(todoRepo domain.TodoRepository) domain.TodoService {
	return &todoService{
		todoRepo: todoRepo,
	}
}