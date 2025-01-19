package handlers

import (
	"encoding/json"
	"final_task/internal/config"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idString := r.FormValue("id")
	answer := config.Err{}
	task := config.Task{}

	if idString == "" {
		answer.Err = "id is required"
		ResponWithError(w, http.StatusBadRequest, answer)
		return
	}

	id, err := strconv.Atoi(idString)
	if err != nil {
		answer.Err = "conversion error"
		ResponWithError(w, http.StatusBadRequest, answer)
		return
	}

	if err = h.repo.GetTask(id, &task); err != nil {
		answer.Err = err.Error()
		ResponWithError(w, http.StatusBadRequest, answer)
		return
	}

	if err = h.repo.DeleteTask(id); err != nil {
		answer.Err = err.Error()
		ResponWithError(w, http.StatusInternalServerError, answer)
		return
	}

	if err := json.NewEncoder(w).Encode(struct{}{}); err != nil {
		http.Error(w, "serialization error", http.StatusInternalServerError)
		return
	}
}
