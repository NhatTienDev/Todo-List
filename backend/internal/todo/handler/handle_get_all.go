package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/nhattiendev/todo-list/internal/todo/domain"
	"github.com/nhattiendev/todo-list/response"
)

// @Summary      Get all tasks
// @Tags         Todos
// @Accept       json
// @Produce      json
// @Param        search query string false "Search by title"
// @Param        status query string false "Filter by status"
// @Param        page query int false "Page number"
// @Param        limit query int false "Items per page"
// @Router       /api/v1/todos [get]
func (h *TodoHandler) HandleGetAll(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")
	filterStatus := r.URL.Query().Get("status")

	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page < 1 {
		page = 1
	}

	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if limit < 1 {
		limit = 10
	}

	result, err := h.todoSV.GetAll(r.Context(), search, filterStatus, page, limit)
	if err != nil {
		if errors.Is(err, domain.ErrInvalidFilterValue) {
			response.WriteErrorJSON(w, http.StatusBadRequest, err.Error())
			return
		}
		response.WriteErrorJSON(w, http.StatusInternalServerError, "Error retrieving the task list")
		return
	}

	response.WriteSuccessJSON(w, http.StatusOK, "Successfully retrieved the task list", result)
}
