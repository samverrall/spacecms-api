// Code generated by goa v3.7.3, DO NOT EDIT.
//
// invoice service
//
// Command:
// $ goa gen github.com/samverrall/invoice-api-service/invoice/design

package invoice

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// RESTFUL API Backend for Invoicify. An open source invoicing web app.
type Service interface {
	// Create an account by email address and password.
	CreateAccount(context.Context, *User) (err error)
	// Create an account by email address and password.
	GrantToken(context.Context, *GrantTokenPayload) (res *Token, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "invoice"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [2]string{"CreateAccount", "GrantToken"}

// GrantTokenPayload is the payload type of the invoice service GrantToken
// method.
type GrantTokenPayload struct {
	Token     *string
	GrantType string
	// User email
	Email string
	// User password
	Password string
}

// Token is the result type of the invoice service GrantToken method.
type Token struct {
	Token                  *string
	AccessToken            string
	RefreshToken           string
	AccessExpiryTime       int64
	RefreshExpiryTime      int64
	AccessExpiryTimeStamp  string
	RefreshExpiryTimeStamp string
}

// User is the payload type of the invoice service CreateAccount method.
type User struct {
	// ID of the user
	ID *string
	// Email address of the user
	Email string
	// Full name of the user
	Name string
	// Password of the user. This is never returned to the client.
	Password string
}

// MakeUnauthorized builds a goa.ServiceError from an error.
func MakeUnauthorized(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "unauthorized",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeServererror builds a goa.ServiceError from an error.
func MakeServererror(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "servererror",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeBadrequest builds a goa.ServiceError from an error.
func MakeBadrequest(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "badrequest",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeNotfound builds a goa.ServiceError from an error.
func MakeNotfound(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "notfound",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}
