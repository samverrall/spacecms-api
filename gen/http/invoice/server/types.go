// Code generated by goa v3.7.3, DO NOT EDIT.
//
// invoice HTTP server types
//
// Command:
// $ goa gen github.com/samverrall/invoice-api-service/invoice/design

package server

import (
	invoice "github.com/samverrall/invoice-api-service/gen/invoice"
	goa "goa.design/goa/v3/pkg"
)

// CreateAccountRequestBody is the type of the "invoice" service
// "create-account" endpoint HTTP request body.
type CreateAccountRequestBody struct {
	// ID of the user
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Email address of the user
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// Full name of the user
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Password of the user. This is never returned to the client.
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
}

// CreateAccountUnauthorizedResponseBody is the type of the "invoice" service
// "create-account" endpoint HTTP response body for the "unauthorized" error.
type CreateAccountUnauthorizedResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// CreateAccountServererrorResponseBody is the type of the "invoice" service
// "create-account" endpoint HTTP response body for the "servererror" error.
type CreateAccountServererrorResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// CreateAccountBadrequestResponseBody is the type of the "invoice" service
// "create-account" endpoint HTTP response body for the "badrequest" error.
type CreateAccountBadrequestResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// CreateAccountNotfoundResponseBody is the type of the "invoice" service
// "create-account" endpoint HTTP response body for the "notfound" error.
type CreateAccountNotfoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// NewCreateAccountUnauthorizedResponseBody builds the HTTP response body from
// the result of the "create-account" endpoint of the "invoice" service.
func NewCreateAccountUnauthorizedResponseBody(res *goa.ServiceError) *CreateAccountUnauthorizedResponseBody {
	body := &CreateAccountUnauthorizedResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewCreateAccountServererrorResponseBody builds the HTTP response body from
// the result of the "create-account" endpoint of the "invoice" service.
func NewCreateAccountServererrorResponseBody(res *goa.ServiceError) *CreateAccountServererrorResponseBody {
	body := &CreateAccountServererrorResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewCreateAccountBadrequestResponseBody builds the HTTP response body from
// the result of the "create-account" endpoint of the "invoice" service.
func NewCreateAccountBadrequestResponseBody(res *goa.ServiceError) *CreateAccountBadrequestResponseBody {
	body := &CreateAccountBadrequestResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewCreateAccountNotfoundResponseBody builds the HTTP response body from the
// result of the "create-account" endpoint of the "invoice" service.
func NewCreateAccountNotfoundResponseBody(res *goa.ServiceError) *CreateAccountNotfoundResponseBody {
	body := &CreateAccountNotfoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewCreateAccountUser builds a invoice service create-account endpoint
// payload.
func NewCreateAccountUser(body *CreateAccountRequestBody) *invoice.User {
	v := &invoice.User{
		ID:       body.ID,
		Email:    *body.Email,
		Name:     *body.Name,
		Password: *body.Password,
	}

	return v
}

// ValidateCreateAccountRequestBody runs the validations defined on
// Create-AccountRequestBody
func ValidateCreateAccountRequestBody(body *CreateAccountRequestBody) (err error) {
	if body.Email == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("email", "body"))
	}
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Password == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("password", "body"))
	}
	return
}
