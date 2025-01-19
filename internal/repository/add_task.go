package repository

import (
	"database/sql"
)

func (r *Repository) AddTask(date, title, comment, repeat string) (int64, error) {
	res, err := r.db.Exec("INSERT INTO scheduler ( date, title, comment, repeat) VALUES ( :date, :title, :comment, :repeat)",
		sql.Named("date", date),
		sql.Named("title", title),
		sql.Named("comment", comment),
		sql.Named("repeat", repeat))
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
