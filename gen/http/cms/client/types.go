// Code generated by goa v3.7.3, DO NOT EDIT.
//
// cms HTTP client types
//
// Command:
// $ goa gen github.com/samverrall/spacecms-api/spacecms-api/design

package client

import (
	cms "github.com/samverrall/spacecms-api/gen/cms"
	goa "goa.design/goa/v3/pkg"
)

// CreatePageRequestBody is the type of the "cms" service "CreatePage" endpoint
// HTTP request body.
type CreatePageRequestBody struct {
	// User email
	Email string `form:"email" json:"email" xml:"email"`
	// User password
	Password string `form:"password" json:"password" xml:"password"`
}

// CreatePageResponseBody is the type of the "cms" service "CreatePage"
// endpoint HTTP response body.
type CreatePageResponseBody struct {
	Token                  *string `form:"token,omitempty" json:"token,omitempty" xml:"token,omitempty"`
	AccessToken            *string `form:"accessToken,omitempty" json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	RefreshToken           *string `form:"refreshToken,omitempty" json:"refreshToken,omitempty" xml:"refreshToken,omitempty"`
	AccessExpiryTime       *int64  `form:"accessExpiryTime,omitempty" json:"accessExpiryTime,omitempty" xml:"accessExpiryTime,omitempty"`
	RefreshExpiryTime      *int64  `form:"refreshExpiryTime,omitempty" json:"refreshExpiryTime,omitempty" xml:"refreshExpiryTime,omitempty"`
	AccessExpiryTimeStamp  *string `form:"accessExpiryTimeStamp,omitempty" json:"accessExpiryTimeStamp,omitempty" xml:"accessExpiryTimeStamp,omitempty"`
	RefreshExpiryTimeStamp *string `form:"refreshExpiryTimeStamp,omitempty" json:"refreshExpiryTimeStamp,omitempty" xml:"refreshExpiryTimeStamp,omitempty"`
}

// CreatePageUnauthorizedResponseBody is the type of the "cms" service
// "CreatePage" endpoint HTTP response body for the "unauthorized" error.
type CreatePageUnauthorizedResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// CreatePageServererrorResponseBody is the type of the "cms" service
// "CreatePage" endpoint HTTP response body for the "servererror" error.
type CreatePageServererrorResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// CreatePageBadrequestResponseBody is the type of the "cms" service
// "CreatePage" endpoint HTTP response body for the "badrequest" error.
type CreatePageBadrequestResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// CreatePageNotfoundResponseBody is the type of the "cms" service "CreatePage"
// endpoint HTTP response body for the "notfound" error.
type CreatePageNotfoundResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// NewCreatePageRequestBody builds the HTTP request body from the payload of
// the "CreatePage" endpoint of the "cms" service.
func NewCreatePageRequestBody(p *cms.CreatePagePayload) *CreatePageRequestBody {
	body := &CreatePageRequestBody{
		Email:    p.Email,
		Password: p.Password,
	}
	return body
}

// NewCreatePageTokenCreated builds a "cms" service "CreatePage" endpoint
// result from a HTTP "Created" response.
func NewCreatePageTokenCreated(body *CreatePageResponseBody) *cms.Token {
	v := &cms.Token{
		Token:                  body.Token,
		AccessToken:            *body.AccessToken,
		RefreshToken:           *body.RefreshToken,
		AccessExpiryTime:       *body.AccessExpiryTime,
		RefreshExpiryTime:      *body.RefreshExpiryTime,
		AccessExpiryTimeStamp:  *body.AccessExpiryTimeStamp,
		RefreshExpiryTimeStamp: *body.RefreshExpiryTimeStamp,
	}

	return v
}

// NewCreatePageUnauthorized builds a cms service CreatePage endpoint
// unauthorized error.
func NewCreatePageUnauthorized(body *CreatePageUnauthorizedResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewCreatePageServererror builds a cms service CreatePage endpoint
// servererror error.
func NewCreatePageServererror(body *CreatePageServererrorResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewCreatePageBadrequest builds a cms service CreatePage endpoint badrequest
// error.
func NewCreatePageBadrequest(body *CreatePageBadrequestResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewCreatePageNotfound builds a cms service CreatePage endpoint notfound
// error.
func NewCreatePageNotfound(body *CreatePageNotfoundResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// ValidateCreatePageResponseBody runs the validations defined on
// CreatePageResponseBody
func ValidateCreatePageResponseBody(body *CreatePageResponseBody) (err error) {
	if body.AccessToken == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("accessToken", "body"))
	}
	if body.RefreshToken == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("refreshToken", "body"))
	}
	if body.AccessExpiryTime == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("accessExpiryTime", "body"))
	}
	if body.RefreshExpiryTime == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("refreshExpiryTime", "body"))
	}
	if body.AccessExpiryTimeStamp == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("accessExpiryTimeStamp", "body"))
	}
	if body.RefreshExpiryTimeStamp == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("refreshExpiryTimeStamp", "body"))
	}
	return
}

// ValidateCreatePageUnauthorizedResponseBody runs the validations defined on
// CreatePage_unauthorized_Response_Body
func ValidateCreatePageUnauthorizedResponseBody(body *CreatePageUnauthorizedResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateCreatePageServererrorResponseBody runs the validations defined on
// CreatePage_servererror_Response_Body
func ValidateCreatePageServererrorResponseBody(body *CreatePageServererrorResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateCreatePageBadrequestResponseBody runs the validations defined on
// CreatePage_badrequest_Response_Body
func ValidateCreatePageBadrequestResponseBody(body *CreatePageBadrequestResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateCreatePageNotfoundResponseBody runs the validations defined on
// CreatePage_notfound_Response_Body
func ValidateCreatePageNotfoundResponseBody(body *CreatePageNotfoundResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}
