package handlers

import (
	"bytes"
	"encoding/json"
	"final_task/internal/config"
	"log"
	"net/http"
)

func (h *Handler) GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var err error
	var tasks config.Tasks
	answer := config.Err{}
	tasks.ListOfTasks = []config.Task{}

	searchTerm := r.URL.Query().Get("search")

	if searchTerm != "" {
		tasks.ListOfTasks, err = h.repo.GetTaskFromSearch(searchTerm)
	} else {
		tasks.ListOfTasks, err = h.repo.GetTasks()
	}

	if err != nil {
		answer.Err = err.Error()
		RespondWithError(w, http.StatusInternalServerError, answer)
		return
	}

	/*defer rows.Close()

	for rows.Next() {
		task := config.Task{}
		err := rows.Scan(&task.Id, &task.Date, &task.Title, &task.Comment, &task.Repeat)
		if err != nil {
			answer.Err = "error scanning task"
			RespondWithError(w, http.StatusInternalServerError, answer)
			return
		}
		tasks.ListOfTasks = append(tasks.ListOfTasks, task)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}*/

	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(tasks); err != nil {
		answer.Err = err.Error()
		RespondWithError(w, http.StatusInternalServerError, answer)
		return
	}

	if _, err := w.Write(buf.Bytes()); err != nil {
		log.Printf("Error sending response \"get tasks\": %v", err)
	}
}
