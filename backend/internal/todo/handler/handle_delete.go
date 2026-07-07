package handler

import (
	"net/http"

	"github.com/nhattiendev/todo-list/response"
)

func (h *TodoHandler) HandleDelete(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		response.WriteErrorJSON(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	if err := h.todoSV.Delete(r.Context(), id); err != nil {
		response.WriteErrorJSON(w, http.StatusInternalServerError, "Error occurred while deleting task")
		return
	}

	response.WriteSuccessJSON(w, http.StatusOK, "Task deleted successfully", nil)
}