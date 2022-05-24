package datastore

import (
	"context"
	"database/sql"
	"log"

	"github.com/samverrall/invoice-app/database"
	"github.com/samverrall/invoice-app/gen/invoice"
)

type DataStore struct {
	db *sql.DB
}

// https://earthly.dev/blog/golang-sqlite/
func New() *DataStore {
	db, err := database.Get()
	if err != nil {
		log.Fatal("failed to get database, err: ", err)
		return nil
	}

	return &DataStore{
		db: db,
	}
}

type DBInterface interface {
	CreateUser(ctx context.Context, email, password, name string) error
	GetUserByEmail(ctx context.Context, email string) (*invoice.User, error)
}
