package database

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func getMigrationsDir() (string, error) {
	migrationsDir := os.Getenv("MIGRATIONS_DIR")
	if strings.TrimSpace(migrationsDir) == "" {
		return "", errors.New("no MIGRATIONS_DIR set")
	}
	return migrationsDir, nil
}

func getDatabaseDir() (string, error) {
	dbDir := os.Getenv("DATABASE_DIR")
	if strings.TrimSpace(dbDir) == "" {
		return "", errors.New("no DATABASE_DIR set")
	}
	return dbDir, nil
}

func Get() (*sql.DB, error) {
	dbDirectory, err := getDatabaseDir()
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("sqlite3", dbDirectory)
	if err != nil {
		return nil, err
	}
	if err := migrate(db); err != nil {
		return nil, err
	}
	return db, nil
}

func getQueryFromFile(filename string) (string, error) {
	migrationsDir, err := getMigrationsDir()
	if err != nil {
		return "", err
	}

	file := filepath.Join(migrationsDir, filename)
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
	migrationDir, err := getMigrationsDir()
	if err != nil {
		return err
	}

	migrationFiles, rErr := os.ReadDir(migrationDir)
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
