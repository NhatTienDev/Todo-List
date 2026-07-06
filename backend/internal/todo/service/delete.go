package service

import "context"

func (s *todoService) Delete(ctx context.Context, id int64) error {
	return s.todoRepo.DeleteTodo(ctx, id)
}