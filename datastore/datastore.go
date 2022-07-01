package datastore

import (
	"context"
	"database/sql"
	"log"

	"github.com/samverrall/spacecms-api/database"
	"github.com/samverrall/spacecms-api/datastore/sqlc"
	"github.com/samverrall/spacecms-api/gen/auth"
)

type DataStore struct {
	db      *sql.DB
	querier *sqlc.Queries
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

type AuthDB interface {
	CreateUser(ctx context.Context, id string, email, password, name string) error
	GetUserByEmail(ctx context.Context, email string) (*auth.User, error)
}

type CmsDB interface {
	CreateTemplate(ctx context.Context, querierTx *sqlc.Queries, id, name, entryBlockID string) error
	CreateBlock(ctx context.Context, querierTx *sqlc.Queries, id, name string) error
	CreateTemplateWithBlock(ctx context.Context, id, name, entryBlockID string) error
	CreatePage(ctx context.Context, in *CreatePageArgs) error
	GetTemplateByID(ctx context.Context, templateID string) (*sqlc.GetTemplateByIDRow, error)
	GetPageByURL(ctx context.Context, url string) (*sqlc.LookupPageURLRow, error)
}
