package database

func (db *Database) createTables() error {
	if _, err := db.database.Exec(CREATE_BOT_TABLE, nil); err != nil {
		return err
	}

	if _, err := db.database.Exec(CREATE_BOT_INDEX, nil); err != nil {
		return err
	}

	if _, err := db.database.Exec(CREATE_BOT_ALLOWED_DOMAINS_TABLE, nil); err != nil {
		return err
	}

	if _, err := db.database.Exec(CREATE_BOT_LOG_TABLE, nil); err != nil {
		return err
	}

	if _, err := db.database.Exec(CREATE_HTTP_LOG_TABLE, nil); err != nil {
		return err
	}

	return nil
}
