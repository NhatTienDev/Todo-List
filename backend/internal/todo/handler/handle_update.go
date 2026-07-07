package handler

import (
	"net/http"
	"encoding/json"
	"errors"

	"github.com/nhattiendev/todo-list/internal/todo/domain"
	"github.com/nhattiendev/todo-list/response"
)

type updateTodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	IsCompleted bool   `json:"is_completed"`
}

// @Summary      Update an existing task
// @Tags         Todos
// @Accept       json
// @Produce      json
// @Param        id path int true "Task ID"
// @Param        request body updateTodoRequest true "Task update payload"
// @Router       /api/v1/todos/{id} [put]
func (h *TodoHandler) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		response.WriteErrorJSON(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	var req updateTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteErrorJSON(w, http.StatusBadRequest, "Invalid input data")
		return
	}

	todo, err := h.todoSV.Update(r.Context(), id, req.Title, req.Description, req.IsCompleted)
	if err != nil {
		if errors.Is(err, domain.ErrTitleRequired) {
			response.WriteErrorJSON(w, http.StatusBadRequest, err.Error())
			return
		}
		response.WriteErrorJSON(w, http.StatusInternalServerError, "Error occurred while updating task")
		return
	}

	response.WriteSuccessJSON(w, http.StatusOK, "Task updated successfully", todo)
}