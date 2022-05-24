package controllers

import (
	"context"
	"database/sql"
	"errors"

	invoice "github.com/samverrall/invoice-app/gen/invoice"
)

// CreateAccount implements create-account.
func (s *invoicesrvc) CreateAccount(ctx context.Context, p *invoice.User) error {
	s.logger.Print("invoice.create-account")

	// Lookup the email to make sure it does not already exist in the system.
	// If it does, return an error.
	user, err := s.dbi.GetUserByEmail(ctx, p.Email)
	switch {
	case err != nil && !errors.Is(err, sql.ErrNoRows):
		s.logger.Println("Failed to determind if user exists")
		return newResponseError("servererror", "Failed to check if user exists")
	case user != nil:
		s.logger.Println("User already exists")
		return newResponseError("badrequest", "This email address has already registered for an account.")
	}

	// Create the user in the database.
	cErr := s.dbi.CreateUser(ctx, p.Email, p.Password, p.Name)
	switch {
	case cErr != nil:
		s.logger.Println("failed to create user, error: ", err)
		return newResponseError("servererror", "Failed to create user")
	}

	s.logger.Println("Account successfully created")

	return nil
}
