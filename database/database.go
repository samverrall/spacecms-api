package database

import "database/sql"

const (
	defaultDatabaseDir = "../.data/database.db"
)

func Get() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", defaultDatabaseDir)
	if err != nil {
		return nil, err
	}
	return db, nil
}
