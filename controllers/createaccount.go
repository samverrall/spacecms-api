package controllers

import (
	"context"

	invoice "github.com/samverrall/invoice-app/gen/invoice"
)

// CreateAccount implements create-account.
func (s *invoicesrvc) CreateAccount(ctx context.Context, p *invoice.User) (*invoice.User, error) {
	s.logger.Print("invoice.create-account")

	// TODO: Lookup user by email address, and determind whether they exist in the systems.

	// Create the user as it does not exist.
	createdUser, err := s.dbi.CreateUser(ctx, p.Email, p.Password, p.Name)
	switch {
	case err != nil:
		s.logger.Println("failed to create user, error: ", err)
		return nil, newResponseError("servererror", "Failed to create user")
	case createdUser == nil:
		s.logger.Println("got nil created user")
		return nil, newResponseError("servererror", "Failed to determind if user was created.")
	}

	return createdUser, nil
}
