package handlers

import (
	"database/sql"
	"encoding/json"
	"final_task/internal/config"
	"net/http"
)

func (h *Handler) GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var rows *sql.Rows
	var err error
	var tasks config.Tasks
	answer := config.Err{}
	tasks.ListOfTasks = []config.Task{}

	searchTerm := r.URL.Query().Get("search")

	if searchTerm != "" {
		rows, err = h.repo.GetTaskFromSearch(searchTerm)
	} else {
		rows, err = h.repo.GetTasks()
	}

	if err != nil {
		answer.Err = err.Error()
		ResponWithError(w, http.StatusInternalServerError, answer)
		return
	}

	defer rows.Close()

	for rows.Next() {
		task := config.Task{}
		err := rows.Scan(&task.Id, &task.Date, &task.Title, &task.Comment, &task.Repeat)
		if err != nil {
			answer.Err = "error scanning task"
			ResponWithError(w, http.StatusInternalServerError, answer)
			return
		}
		tasks.ListOfTasks = append(tasks.ListOfTasks, task)
	}

	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		http.Error(w, "serialization error", http.StatusInternalServerError)
		return
	}
}
