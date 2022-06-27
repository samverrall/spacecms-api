package controllers

import (
	"context"
	"database/sql"
	"errors"

	"github.com/samverrall/invoice-api-service/apierror"
	"github.com/samverrall/invoice-api-service/gen/invoice"
	"github.com/sirupsen/logrus"
)

const (
	GrantTypeRefresh = "refresh_token"
	GrantTypeAccess  = "access_token"
)

func (s *invoicesrvc) GrantToken(ctx context.Context, p *invoice.GrantTokenPayload) (*invoice.Token, error) {
	s.logger.Info("invoice.GrantToken")

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

	token, tokenErr := s.tokener.CreateTokenPair(ctx, user.ID)
	if tokenErr != nil {
		s.logger.WithFields(logrus.Fields{
			"error": tokenErr.Error(),
		}).Error("Failed to create token pair")
		return nil, apierror.NewResponseError("servererror", "Failed to create token pair for user.")
	}

	// Set __Host-token HTTP Cookie
	token.Token = &token.AccessToken

	return token, nil
}
