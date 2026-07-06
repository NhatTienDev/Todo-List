package handler

import (
	"net/http"

	"github.com/nhattiendev/todo-list/response"
)

func (h *TodoHandler) HandleGetAll(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")
	filterStatus := r.URL.Query().Get("status")

	todos, err := h.todoSV.GetAll(r.Context(), search, filterStatus)
	if err != nil {
		response.WriteErrorJSON(w, http.StatusInternalServerError, "Error retrieving the task list")
		return
	}

	response.WriteSuccessJSON(w, http.StatusOK, "Successfully retrieved the task list", todos)
}