package repository

import (
	"database/sql"
	"final_task/internal/config"
)

func (r *Repository) GetTask(id int, task *config.Task) error {
	row := r.db.QueryRow("SELECT id, date, title, comment, repeat FROM scheduler WHERE id = :id",
		sql.Named("id", id))

	if err := row.Scan(&task.Id, &task.Date, &task.Title, &task.Comment, &task.Repeat); err != nil {
		return err
	}

	return nil
}
