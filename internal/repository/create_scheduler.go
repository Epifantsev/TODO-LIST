package repository

import "log"

func (r *Repository) CreateScheduler() error {
	query := `CREATE TABLE IF NOT EXISTS scheduler(
			  id INTEGER PRIMARY KEY AUTOINCREMENT,
			  date CHAR(8) NOT NULL DEFAULT "",
			  title VARCHAR(256) NOT NULL DEFAULT "",
			  comment TEXT NOT NULL DEFAULT "",
			  repeat VARCHAR(128) NOT NULL DEFAULT "");`

	if _, err := r.db.Exec(query); err != nil {
		log.Printf("Error creating scheduler table: %v", err)
		return err
	}

	return nil
}
