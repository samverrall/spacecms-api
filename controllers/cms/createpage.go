package controllers

import (
	"context"

	"github.com/samverrall/spacecms-api/gen/cms"
)

func (s *cmsservice) CreatePage(ctx context.Context, p *cms.CreatePagePayload) (*cms.Page, error) {
	s.logger.Info("cmsservice.CreatePage")

	return &cms.Page{
		ID:         "",
		URL:        "",
		TemplateID: "",
		IsActive:   false,
		Meta:       &cms.Meta{},
		CreatedAt:  "",
	}, nil
}
