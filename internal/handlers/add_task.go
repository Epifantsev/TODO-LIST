package handlers

import (
	"bytes"
	"encoding/json"
	"final_task/internal/config"
	repetitionrule "final_task/internal/repetitionRule"
	"log"
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
		answer.Err = err.Error()
		RespondWithError(w, http.StatusBadRequest, answer)
		return
	}

	if task.Title == "" {
		answer.Err = "empty title field"
		RespondWithError(w, http.StatusBadRequest, answer)
		return
	}

	if task.Date == "" {
		task.Date = time.Now().Format("20060102")
	}

	date, err := time.Parse("20060102", task.Date)
	if err != nil {
		answer.Err = "wrong date format"
		RespondWithError(w, http.StatusBadRequest, answer)
		return
	}

	if date.Before(currentTime) && currentTime.Sub(date) >= time.Hour*24 {
		nextDate, err := repetitionrule.RepetitionRule(currentTime, task.Date, task.Repeat)
		if err != nil {
			if err.Error() == "empty repeat field" {
				task.Date = currentTime.Format("20060102")
			} else {
				answer.Err = err.Error()
				RespondWithError(w, http.StatusBadRequest, answer)
				return
			}
		} else {
			task.Date = nextDate
		}
	}

	lastId, err := h.repo.AddTask(task.Date, task.Title, task.Comment, task.Repeat)
	if err != nil {
		answer.Err = "error adding task to database"
		RespondWithError(w, http.StatusInternalServerError, answer)
		return
	}

	id.Id = lastId

	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(id); err != nil {
		answer.Err = err.Error()
		RespondWithError(w, http.StatusInternalServerError, answer)
		return
	}

	if _, err := w.Write(buf.Bytes()); err != nil {
		log.Printf("Error sending response \"add task\": %v", err)
	}
}
