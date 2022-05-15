package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const (
	defaultDatabaseDir = "../.data/database.db"
	migrationsDir      = "./migration"
)

func Get() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", defaultDatabaseDir)
	if err != nil {
		return nil, err
	}
	if err := migrate(db); err != nil {
		return nil, err
	}
	return db, nil
}

func getQueryFromFile(filename string) (string, error) {
	bytes, err := os.ReadFile(fmt.Sprintf("%s/%s", migrationsDir, filename))
	if err != nil {
		return "", err
	}
	if len(bytes) == 0 {
		return "", fmt.Errorf("got empty migration file, file: %s", filename)
	}
	return string(bytes), nil
}

// migrate will read the migrations directory, and attempt to exec any .SQL files.
func migrate(db *sql.DB) error {
	migrationFiles, rErr := os.ReadDir(migrationsDir)
	if rErr != nil {
		return rErr
	}

	for _, file := range migrationFiles {
		if file.IsDir() {
			continue
		}

		query, qErr := getQueryFromFile(file.Name())
		if qErr != nil {
			return qErr
		}

		if _, err := db.Exec(query); err != nil {
			return err
		}
	}

	return nil
}
