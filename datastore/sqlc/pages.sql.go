// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: pages.sql

package sqlc

import (
	"context"
	"database/sql"
	"time"
)

const insertPage = `-- name: InsertPage :exec
INSERT INTO pages
(id, template_id, url, canonical_url, title, description, thumbnail, is_active,
is_in_sitemap, no_follow, no_archive, created_at, updated_at, deleted_at)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
`

type InsertPageParams struct {
	ID           string
	TemplateID   string
	Url          string
	CanonicalUrl string
	Title        string
	Description  string
	Thumbnail    string
	IsActive     int32
	IsInSitemap  int32
	NoFollow     int32
	NoArchive    int32
	CreatedAt    time.Time
	UpdatedAt    sql.NullTime
	DeletedAt    sql.NullTime
}

func (q *Queries) InsertPage(ctx context.Context, arg InsertPageParams) error {
	_, err := q.db.ExecContext(ctx, insertPage,
		arg.ID,
		arg.TemplateID,
		arg.Url,
		arg.CanonicalUrl,
		arg.Title,
		arg.Description,
		arg.Thumbnail,
		arg.IsActive,
		arg.IsInSitemap,
		arg.NoFollow,
		arg.NoArchive,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.DeletedAt,
	)
	return err
}

const lookupPageURL = `-- name: LookupPageURL :one
SELECT id, url FROM pages WHERE url = ?
`

type LookupPageURLRow struct {
	ID  string
	Url string
}

func (q *Queries) LookupPageURL(ctx context.Context, url string) (LookupPageURLRow, error) {
	row := q.db.QueryRowContext(ctx, lookupPageURL, url)
	var i LookupPageURLRow
	err := row.Scan(&i.ID, &i.Url)
	return i, err
}
