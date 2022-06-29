package datastore

import (
	"context"
	"database/sql"
	"log"

	"github.com/samverrall/spacecms-api/database"
	"github.com/samverrall/spacecms-api/datastore/sqlc"
	"github.com/samverrall/spacecms-api/gen/invoice"
)

type DataStore struct {
	db      *sql.DB
	querier sqlc.Querier
}

// https://earthly.dev/blog/golang-sqlite/
func New() *DataStore {
	db, err := database.Get()
	if err != nil {
		log.Fatal("failed to get database, err: ", err)
		return nil
	}

	return &DataStore{
		db:      db,
		querier: sqlc.New(db),
	}
}

type DBInterface interface {
	CreateUser(ctx context.Context, id string, email, password, name string) error
	GetUserByEmail(ctx context.Context, email string) (*invoice.User, error)
}
