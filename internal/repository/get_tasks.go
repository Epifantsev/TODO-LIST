package repository

import (
	"database/sql"
	"final_task/internal/config"
	"log"
)

func (r *Repository) GetTasks() ([]config.Task, error) {
	var tasks config.Tasks
	tasks.ListOfTasks = []config.Task{}

	query := `SELECT * FROM scheduler ORDER BY date LIMIT :limit`

	rows, err := r.db.Query(query, sql.Named("limit", 10))
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
