package handler

import (
	"net/http"

	"github.com/nhattiendev/todo-list/response"
)

// @Summary      Get task by ID
// @Tags         Todos
// @Accept       json
// @Produce      json
// @Param        id path int true "Task ID"
// @Router       /api/v1/todos/{id} [get]
func (h *TodoHandler) HandleGetByID(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		response.WriteErrorJSON(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	todo, err := h.todoSV.GetByID(r.Context(), id)
	if err != nil {
		response.WriteErrorJSON(w, http.StatusNotFound, "Task not found")
		return
	}

	response.WriteSuccessJSON(w, http.StatusOK, "Task retrieved successfully", todo)
}