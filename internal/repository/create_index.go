package repository

func (r *Repository) IndexDate() error {
	if _, err := r.db.Exec("CREATE INDEX IF NOT EXISTS index_date ON scheduler (date);"); err != nil {
		return err
	}

	return nil
}
