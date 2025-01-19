package repository

import (
	"database/sql"
)

func (r *Repository) GetTasks() (*sql.Rows, error) {
	query := `SELECT * FROM scheduler ORDER BY date LIMIT :limit`

	rows, err := r.db.Query(query, sql.Named("limit", 10))
	if err != nil {
		return nil, err
	}

	return rows, nil
}
