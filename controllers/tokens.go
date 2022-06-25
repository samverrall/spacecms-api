package controllers

import (
	"context"
	"database/sql"
	"errors"

	"github.com/samverrall/invoice-api-service/apierror"
	"github.com/samverrall/invoice-api-service/gen/invoice"
	"github.com/sirupsen/logrus"
)

func (s *invoicesrvc) AuthoriseLogin(ctx context.Context, p *invoice.AuthoriseLoginPayload) (*invoice.Token, error) {
	s.logger.Info("invoice.AuthoriseLogin")

	user, err := s.dbi.GetUserByEmail(ctx, p.Email)
	switch {
	case err != nil && !errors.Is(err, sql.ErrNoRows):
		s.logger.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to determind if user exists")
		return nil, apierror.NewResponseError("servererror", "Failed to check if user exists")
	case user == nil:
		s.logger.Debug("User already exists")
		return nil, apierror.NewResponseError("badrequest", "No user exists for the supplied email address")
	}

	match, err := s.hasher.Compare(user.Password, p.Password)
	switch {
	case err != nil:
		s.logger.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to compare password hash")
		return nil, apierror.NewResponseError("servererror", "Failed to compare hashes")
	case !match:
		s.logger.Debug("Failed to authorise user")
		return nil, apierror.NewResponseError("badrequest", "Incorrect password supplied")
	}

	token, err := s.tokener.CreateTokenPair(ctx, user.ID)
	if err != nil {
		s.logger.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to create token pair")
		return nil, apierror.NewResponseError("servererror", "Failed to create token pair for user.")
	}

	return token, nil
}