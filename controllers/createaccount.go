package controllers

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/samverrall/invoice-api-service/apierror"
	invoice "github.com/samverrall/invoice-api-service/gen/invoice"
	"github.com/sirupsen/logrus"
)

// CreateAccount implements create-account.
func (s *invoicesrvc) CreateAccount(ctx context.Context, p *invoice.User) error {
	s.logger.Info("invoice.create-account")

	// Lookup the email to make sure it does not already exist in the system.
	// If it does, return an error.
	user, err := s.dbi.GetUserByEmail(ctx, p.Email)
	switch {
	case err != nil && !errors.Is(err, sql.ErrNoRows):
		s.logger.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to determind if user exists")
		return apierror.NewResponseError("servererror", "Failed to check if user exists")
	case user != nil:
		s.logger.Error("User already exists")
		return apierror.NewResponseError("badrequest", "This email address has already registered for an account.")
	}

	s.logger.Info("User does not exist, ready to create account.")

	newUserID := uuid.NewString()

	hashedPassword, err := s.hasher.Hash(p.Password)
	if err != nil {
		s.logger.WithFields(logrus.Fields{
			"email": p.Email,
		}).Error("Failed to hash password")
		return apierror.NewResponseError("servererror", "Failed to create account")
	}

	// Create the user in the database.
	cErr := s.dbi.CreateUser(ctx, newUserID, p.Email, hashedPassword, p.Name)
	if cErr != nil {
		s.logger.Error("failed to create user, error: ", cErr)
		return apierror.NewResponseError("servererror", "Failed to create user")
	}

	s.logger.WithFields(logrus.Fields{
		"email": p.Email,
	}).Info("Account successfully created")

	return nil
}
