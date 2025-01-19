package repository

import (
	"database/sql"
	"time"
)

func (r *Repository) GetTaskFromSearch(search string) (*sql.Rows, error) {
	var result, query string

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

	return rows, nil
}
