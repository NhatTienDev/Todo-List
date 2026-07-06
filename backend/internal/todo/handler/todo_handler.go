package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/nhattiendev/todo-list/internal/todo/domain"
)

type TodoHandler struct {
	todoSV domain.TodoService
}

func NewTodoHandler(todoSV domain.TodoService) *TodoHandler {
	return &TodoHandler{
		todoSV: todoSV,
	}
}

func parseID(r *http.Request) (int64, error) {
	idStr := chi.URLParam(r, "id")
	return strconv.ParseInt(idStr, 10, 64)
}