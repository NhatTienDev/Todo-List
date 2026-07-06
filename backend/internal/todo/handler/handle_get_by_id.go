package handler

import (
	"net/http"

	"github.com/nhattiendev/todo-list/response"
)

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