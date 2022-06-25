// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package sqlc

import (
	"context"
	"database/sql"
)

type Querier interface {
	GetUserByEmail(ctx context.Context, email sql.NullString) (GetUserByEmailRow, error)
	InsertUser(ctx context.Context, arg InsertUserParams) error
}

var _ Querier = (*Queries)(nil)
