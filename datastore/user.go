package datastore

import (
	"context"

	"github.com/samverrall/invoice-app/gen/invoice"
)

func (d *DataStore) CreateUser(ctx context.Context, id string, email, password, name string) error {
	const sqlstr = `INSERT INTO users (id, email, name, pass_hash) VALUES (?,?,?,?)`

	if _, err := d.db.Exec(sqlstr, id, email, name, password); err != nil {
		return err
	}

	return nil
}

func (d *DataStore) GetUserByEmail(ctx context.Context, email string) (*invoice.User, error) {
	const sqlstr = `SELECT id, name, email, is_active FROM users WHERE email = ?`

	var user invoice.User
	err := d.db.QueryRowContext(ctx, sqlstr, email).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
