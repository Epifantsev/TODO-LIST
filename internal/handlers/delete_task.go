package handlers

import (
	"bytes"
	"encoding/json"
	"final_task/internal/config"
	"log"
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
		RespondWithError(w, http.StatusBadRequest, answer)
		return
	}

	id, err := strconv.Atoi(idString)
	if err != nil {
		answer.Err = "conversion error"
		RespondWithError(w, http.StatusBadRequest, answer)
		return
	}

	if err = h.repo.GetTask(id, &task); err != nil {
		answer.Err = err.Error()
		RespondWithError(w, http.StatusInternalServerError, answer)
		return
	}

	if err = h.repo.DeleteTask(id); err != nil {
		answer.Err = err.Error()
		RespondWithError(w, http.StatusInternalServerError, answer)
		return
	}

	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(struct{}{}); err != nil {
		answer.Err = err.Error()
		RespondWithError(w, http.StatusInternalServerError, answer)
		return
	}

	if _, err := w.Write(buf.Bytes()); err != nil {
		log.Printf("Error sending response \"delete task\": %v", err)
	}
}
