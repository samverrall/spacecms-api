package controllers

import (
	"context"

	"github.com/google/uuid"
	"github.com/samverrall/spacecms-api/apierror"
	"github.com/samverrall/spacecms-api/gen/cms"
	"github.com/sirupsen/logrus"
)

func (s *cmsservice) CreateTemplate(ctx context.Context, p *cms.CreateTemplatePayload) error {
	s.logger.Info("cmsservice.CreateTemplate")

	blockID := uuid.New().String()
	if p.BlockID != nil {
		blockID = *p.BlockID
	}

	err := s.dbi.CreateTemplateWithBlock(ctx, p.ID, p.Name, blockID)
	if err != nil {
		s.logger.WithFields(logrus.Fields{
			"error": err.Error(),
			"id":    p.ID,
		}).Error("Failed to create template")
		return apierror.NewResponseError("servererror", "Failed to create template")
	}

	return nil
}
