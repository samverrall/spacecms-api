package controllers

import (
	"context"
	"database/sql"
	"errors"

	"github.com/samverrall/spacecms-api/apierror"
	"github.com/samverrall/spacecms-api/datastore"
	"github.com/samverrall/spacecms-api/gen/cms"
	"github.com/sirupsen/logrus"
)

func (s *cmsservice) CreatePage(ctx context.Context, p *cms.CreatePagePayload) error {
	s.logger.Info("cmsservice.CreatePage")

	template, err := s.dbi.GetTemplateByID(ctx, p.TemplateID)
	switch {
	case err != nil && errors.Is(err, sql.ErrNoRows) || template == nil:
		s.logger.WithFields(logrus.Fields{
			"templateID": p.TemplateID,
		}).Debug("Supplied TemplateID does not exist")
		return apierror.NewResponseError("badrequest", "Supplied templateID does not exist")
	case err != nil:
		s.logger.WithFields(logrus.Fields{
			"templateID": p.TemplateID,
			"error":      err.Error(),
		}).Error("Failed to lookup templateID")
		return apierror.NewResponseError("servererror", "Failed to lookup templateID")
	}

	existingPage, err := s.dbi.GetPageByURL(ctx, p.URL)
	switch {
	case existingPage != nil:
		s.logger.WithFields(logrus.Fields{
			"url": p.URL,
		}).Debug("A page already exists with the supplied page URL")
		return apierror.NewResponseError("badrequest", "A page already exists with the supplied page URL")
	case err != nil && !errors.Is(err, sql.ErrNoRows):
		s.logger.WithFields(logrus.Fields{
			"url":   p.URL,
			"error": err.Error(),
		}).Error("Failed to lookup page URL")
		return apierror.NewResponseError("servererror", "Failed to lookup page URL")
	}

	if err := s.dbi.CreatePage(ctx, &datastore.CreatePageArgs{
		ID:           p.ID,
		TemplateID:   p.TemplateID,
		URL:          p.URL,
		CanonicalUrl: p.URL,
		Title:        p.Meta.Title,
		Description:  p.Meta.Description,
	}); err != nil {
		s.logger.WithFields(logrus.Fields{
			"url":   p.URL,
			"error": err.Error(),
		}).Error("Failed to create page")
		return apierror.NewResponseError("servererror", "Failed to create page")
	}

	return nil
}
