package handlers

import (
	"encoding/json"
	"final_task/internal/config"
	repetitionrule "final_task/internal/repetitionRule"
	"net/http"
	"time"
)

func (h *Handler) AddTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	currentTime := time.Now()

	task := config.Task{}
	answer := config.Err{}
	id := config.Id{}

	dec := json.NewDecoder(r.Body)

	err := dec.Decode(&task)
	if err != nil {
		answer.Err = "deserialization error"
		ResponWithError(w, http.StatusBadRequest, answer)
		return
	}

	if task.Title == "" {
		answer.Err = "empty title field"
		ResponWithError(w, http.StatusBadRequest, answer)
		return
	}

	if task.Date == "" {
		task.Date = time.Now().Format("20060102")
	}

	date, err := time.Parse("20060102", task.Date)
	if err != nil {
		answer.Err = "wrong date format"
		ResponWithError(w, http.StatusBadRequest, answer)
		return
	}

	if date.Before(currentTime) && currentTime.Sub(date) >= time.Hour*24 {
		nextDate, err := repetitionrule.RepetitionRule(currentTime, task.Date, task.Repeat)
		if err != nil {
			if err.Error() == "empty repeat field" {
				task.Date = currentTime.Format("20060102")
			} else {
				answer.Err = err.Error()
				ResponWithError(w, http.StatusBadRequest, answer)
				return
			}
		} else {
			task.Date = nextDate
		}
	}

	lastId, err := h.repo.AddTask(task.Date, task.Title, task.Comment, task.Repeat)
	if err != nil {
		answer.Err = "error adding task to database"
		ResponWithError(w, http.StatusInternalServerError, answer)
		return
	}

	id.Id = lastId
	if err := json.NewEncoder(w).Encode(id); err != nil {
		http.Error(w, "serialization error", http.StatusInternalServerError)
		return
	}
}
