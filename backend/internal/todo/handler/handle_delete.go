package handler

import (
	"net/http"
	"errors"

	"github.com/nhattiendev/todo-list/internal/todo/domain"
	"github.com/nhattiendev/todo-list/response"
)

// @Summary      Delete a task
// @Tags         Todos
// @Accept       json
// @Produce      json
// @Param        id path int true "Task ID"
// @Router       /api/v1/todos/{id} [delete]
func (h *TodoHandler) HandleDelete(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		response.WriteErrorJSON(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	if err := h.todoSV.Delete(r.Context(), id); err != nil {
		if errors.Is(err, domain.ErrTodoNotFound) {
			response.WriteErrorJSON(w, http.StatusNotFound, err.Error())
			return
		}
		response.WriteErrorJSON(w, http.StatusInternalServerError, "Error occurred while deleting task")
		return
	}

	response.WriteSuccessJSON(w, http.StatusOK, "Task deleted successfully", nil)
}