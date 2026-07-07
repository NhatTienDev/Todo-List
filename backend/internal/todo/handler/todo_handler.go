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

// Declare API routes for Todo module
func (h *TodoHandler) RegisterTodoRoutes(r chi.Router) {
	r.Route("/api/v1/todos", func(router chi.Router) {
		router.Post("/", h.HandleCreate)
		router.Get("/", h.HandleGetAll)
		router.Get("/{id}", h.HandleGetByID)
		router.Put("/{id}", h.HandleUpdate)
		router.Delete("/{id}", h.HandleDelete)
	})
}

func parseID(r *http.Request) (int64, error) {
	idStr := chi.URLParam(r, "id")
	return strconv.ParseInt(idStr, 10, 64)
}