package handler

import (
	"net/http"
	"encoding/json"
	"errors"

	"github.com/nhattiendev/todo-list/internal/todo/domain"
	"github.com/nhattiendev/todo-list/response"
)

type createTodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (h *TodoHandler) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var req createTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteErrorJSON(w, http.StatusBadRequest, "Invalid input data")
		return
	}

	todo, err := h.todoSV.Create(r.Context(), req.Title, req.Description)
	if err != nil {
		if errors.Is(err, domain.ErrTitleRequired) {
			response.WriteErrorJSON(w, http.StatusBadRequest, err.Error())
			return
		}
		response.WriteErrorJSON(w, http.StatusInternalServerError, "Error occurred while creating task")
		return
	}

	response.WriteSuccessJSON(w, http.StatusCreated, "Task created successfully", todo)
}