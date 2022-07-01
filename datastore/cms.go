package datastore

import (
	"context"
	"database/sql"
	"time"

	"github.com/samverrall/spacecms-api/datastore/sqlc"
)

// CreateTemplateWithBlock creates a block and template using the block ID as the
// template entry block ID.
func (d *DataStore) CreateTemplateWithBlock(ctx context.Context, id, name, entryBlockID string) error {
	// TODO: Add logging

	tx, err := d.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	querier := d.querier.WithTx(tx)

	if err := d.CreateBlock(ctx, querier, entryBlockID, name); err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}

		return err
	}

	if err := d.CreateTemplate(ctx, querier, id, name, entryBlockID); err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}

		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (d *DataStore) CreateTemplate(ctx context.Context, querierTx *sqlc.Queries, id, name, entryBlockID string) error {
	querier := d.querier
	if querierTx != nil {
		querier = querierTx
	}

	now := time.Now().UTC()
	if err := querier.InsertTemplate(ctx, sqlc.InsertTemplateParams{
		ID:           id,
		EntryBlockID: entryBlockID,
		Name:         name,
		CreatedAt:    now,
	}); err != nil {
		return err
	}

	return nil
}

func (d *DataStore) CreateBlock(ctx context.Context, querierTx *sqlc.Queries, id, name string) error {
	querier := d.querier
	if querierTx != nil {
		querier = querierTx
	}

	now := time.Now().UTC()
	if err := querier.InsertBlock(ctx, sqlc.InsertBlockParams{
		ID:        id,
		Name:      name,
		CreatedAt: now,
	}); err != nil {
		return err
	}

	return nil
}

type CreatePageArgs struct {
	ID           string
	TemplateID   string
	URL          string
	CanonicalUrl string
	Title        *string
	Description  *string
}

func (d *DataStore) CreatePage(ctx context.Context, in *CreatePageArgs) error {
	now := time.Now().UTC()

	title := ""
	if in.Title != nil {
		title = *in.Title
	}

	description := ""
	if in.Description != nil {
		title = *in.Description
	}

	if err := d.querier.InsertPage(ctx, sqlc.InsertPageParams{
		ID:           in.ID,
		TemplateID:   in.TemplateID,
		Url:          in.URL,
		CanonicalUrl: in.CanonicalUrl,
		Title:        title,
		Description:  description,
		Thumbnail:    "",
		IsActive:     1,
		IsInSitemap:  1,
		CreatedAt:    now,
	}); err != nil {
		return err
	}

	return nil
}

func (d *DataStore) GetTemplateByID(ctx context.Context, templateID string) (*sqlc.GetTemplateByIDRow, error) {
	row, err := d.querier.GetTemplateByID(ctx, templateID)
	if err != nil {
		return nil, err
	}
	return &row, nil
}

func (d *DataStore) GetPageByURL(ctx context.Context, url string) (*sqlc.LookupPageURLRow, error) {
	row, err := d.querier.LookupPageURL(ctx, url)
	if err != nil {
		return nil, err
	}
	return &row, nil
}
