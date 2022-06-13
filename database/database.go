package database

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

const (
	defaultDatabaseDir = "/Users/samverrall/projects/invoice-app/.data/database.db"
	migrationsDir      = "/Users/samverrall/projects/invoice-app/database/migration"
)

func getMigrationsDir() string {
	migrationsDir := os.Getenv("MIGRATIONS_DIR")
	if strings.TrimSpace(migrationsDir) == "" {
		return defaultDatabaseDir
	}
	return migrationsDir
}

func getDatabaseDir() string {
	dbDir := os.Getenv("DATABASE_DIR")
	if strings.TrimSpace(dbDir) == "" {
		return defaultDatabaseDir
	}
	return dbDir
}

func Get() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", getDatabaseDir())
	if err != nil {
		return nil, err
	}
	if err := migrate(db); err != nil {
		return nil, err
	}
	return db, nil
}

func getQueryFromFile(filename string) (string, error) {
	file := filepath.Join(getMigrationsDir(), filename)
	bytes, err := os.ReadFile(file)
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
	migrationFiles, rErr := os.ReadDir(getMigrationsDir())
	if rErr != nil {
		return rErr
	}

	for _, file := range migrationFiles {
		fileName := file.Name()
		if file.IsDir() {
			continue
		}

		query, qErr := getQueryFromFile(fileName)
		if qErr != nil {
			return qErr
		}

		if _, err := db.Exec(query); err != nil {
			fmt.Println("failed to run migration in file:", fileName)
			return err
		}
	}

	return nil
}
