package repository

import (
	"database/sql"
	"fmt"
)

func (r *Repository) DeleteTask(id int) error {
	result, err := r.db.Exec("DELETE FROM scheduler WHERE id = :id",
		sql.Named("id", id))
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows deleted for id: %v", id)
	}

	return nil
}
