package repository

import (
	"database/sql"
	"final_task/internal/config"
	"log"
	"time"
)

func (r *Repository) GetTaskFromSearch(search string) ([]config.Task, error) {
	var result, query string
	var tasks config.Tasks
	tasks.ListOfTasks = []config.Task{}

	date, err := time.Parse("02.01.2006", search)
	if err == nil {
		query = "SELECT * FROM scheduler WHERE date = :search LIMIT :limit"
		result = date.Format("20060102")
	} else {
		query = "SELECT * FROM scheduler WHERE title LIKE  :search OR comment LIKE :search ORDER BY date LIMIT :limit"
		result = "%" + search + "%"
	}

	rows, err := r.db.Query(query,
		sql.Named("search", result),
		sql.Named("limit", 10))

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		task := config.Task{}
		err := rows.Scan(&task.Id, &task.Date, &task.Title, &task.Comment, &task.Repeat)
		if err != nil {
			return nil, err
		}
		tasks.ListOfTasks = append(tasks.ListOfTasks, task)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return tasks.ListOfTasks, nil
}
