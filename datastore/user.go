package datastore

import (
	"context"
	"database/sql"

	"github.com/samverrall/spacecms-api/datastore/sqlc"
	"github.com/samverrall/spacecms-api/gen/invoice"
)

func (d *DataStore) CreateUser(ctx context.Context, id string, email, passwordHash, name string) error {
	if err := d.querier.InsertUser(ctx, sqlc.InsertUserParams{
		ID: id,
		Email: sql.NullString{
			String: email,
			Valid:  true,
		},
		Name:     name,
		PassHash: passwordHash,
	}); err != nil {
		return err
	}
	return nil
}

func (d *DataStore) GetUserByEmail(ctx context.Context, email string) (*invoice.User, error) {
	user, err := d.querier.GetUserByEmail(ctx, sql.NullString{
		String: email,
		Valid:  true,
	})
	if err != nil {
		return nil, err
	}

	return &invoice.User{
		ID:       &user.ID,
		Email:    email,
		Name:     user.Name,
		Password: user.PassHash,
	}, nil
}
