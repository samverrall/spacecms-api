package datastore

import (
	"context"

	"github.com/samverrall/invoice-app/gen/invoice"
)

func (d *DataStore) CreateUser(ctx context.Context, email, password, name string) (*invoice.User, error) {
	return nil, nil
}
