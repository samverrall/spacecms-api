// Code generated by goa v3.7.3, DO NOT EDIT.
//
// auth HTTP client types
//
// Command:
// $ goa gen github.com/samverrall/spacecms-api/invoice/design

package client

import (
	auth "github.com/samverrall/spacecms-api/gen/auth"
	goa "goa.design/goa/v3/pkg"
)

// CreateAccountRequestBody is the type of the "auth" service "CreateAccount"
// endpoint HTTP request body.
type CreateAccountRequestBody struct {
	// ID of the user
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Email address of the user
	Email string `form:"email" json:"email" xml:"email"`
	// Full name of the user
	Name string `form:"name" json:"name" xml:"name"`
	// Password of the user. This is never returned to the client.
	Password string `form:"password" json:"password" xml:"password"`
}

// GrantTokenRequestBody is the type of the "auth" service "GrantToken"
// endpoint HTTP request body.
type GrantTokenRequestBody struct {
	// User email
	Email string `form:"email" json:"email" xml:"email"`
	// User password
	Password string `form:"password" json:"password" xml:"password"`
}

// GrantTokenResponseBody is the type of the "auth" service "GrantToken"
// endpoint HTTP response body.
type GrantTokenResponseBody struct {
	AccessToken            *string `form:"accessToken,omitempty" json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	RefreshToken           *string `form:"refreshToken,omitempty" json:"refreshToken,omitempty" xml:"refreshToken,omitempty"`
	AccessExpiryTime       *int64  `form:"accessExpiryTime,omitempty" json:"accessExpiryTime,omitempty" xml:"accessExpiryTime,omitempty"`
	RefreshExpiryTime      *int64  `form:"refreshExpiryTime,omitempty" json:"refreshExpiryTime,omitempty" xml:"refreshExpiryTime,omitempty"`
	AccessExpiryTimeStamp  *string `form:"accessExpiryTimeStamp,omitempty" json:"accessExpiryTimeStamp,omitempty" xml:"accessExpiryTimeStamp,omitempty"`
	RefreshExpiryTimeStamp *string `form:"refreshExpiryTimeStamp,omitempty" json:"refreshExpiryTimeStamp,omitempty" xml:"refreshExpiryTimeStamp,omitempty"`
}

// CreateAccountUnauthorizedResponseBody is the type of the "auth" service
// "CreateAccount" endpoint HTTP response body for the "unauthorized" error.
type CreateAccountUnauthorizedResponseBody struct {
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

// CreateAccountServererrorResponseBody is the type of the "auth" service
// "CreateAccount" endpoint HTTP response body for the "servererror" error.
type CreateAccountServererrorResponseBody struct {
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

// CreateAccountBadrequestResponseBody is the type of the "auth" service
// "CreateAccount" endpoint HTTP response body for the "badrequest" error.
type CreateAccountBadrequestResponseBody struct {
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

// CreateAccountNotfoundResponseBody is the type of the "auth" service
// "CreateAccount" endpoint HTTP response body for the "notfound" error.
type CreateAccountNotfoundResponseBody struct {
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

// GrantTokenUnauthorizedResponseBody is the type of the "auth" service
// "GrantToken" endpoint HTTP response body for the "unauthorized" error.
type GrantTokenUnauthorizedResponseBody struct {
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

// GrantTokenServererrorResponseBody is the type of the "auth" service
// "GrantToken" endpoint HTTP response body for the "servererror" error.
type GrantTokenServererrorResponseBody struct {
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

// GrantTokenBadrequestResponseBody is the type of the "auth" service
// "GrantToken" endpoint HTTP response body for the "badrequest" error.
type GrantTokenBadrequestResponseBody struct {
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

// GrantTokenNotfoundResponseBody is the type of the "auth" service
// "GrantToken" endpoint HTTP response body for the "notfound" error.
type GrantTokenNotfoundResponseBody struct {
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

// NewCreateAccountRequestBody builds the HTTP request body from the payload of
// the "CreateAccount" endpoint of the "auth" service.
func NewCreateAccountRequestBody(p *auth.User) *CreateAccountRequestBody {
	body := &CreateAccountRequestBody{
		ID:       p.ID,
		Email:    p.Email,
		Name:     p.Name,
		Password: p.Password,
	}
	return body
}

// NewGrantTokenRequestBody builds the HTTP request body from the payload of
// the "GrantToken" endpoint of the "auth" service.
func NewGrantTokenRequestBody(p *auth.GrantTokenPayload) *GrantTokenRequestBody {
	body := &GrantTokenRequestBody{
		Email:    p.Email,
		Password: p.Password,
	}
	return body
}

// NewCreateAccountUnauthorized builds a auth service CreateAccount endpoint
// unauthorized error.
func NewCreateAccountUnauthorized(body *CreateAccountUnauthorizedResponseBody) *goa.ServiceError {
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

// NewCreateAccountServererror builds a auth service CreateAccount endpoint
// servererror error.
func NewCreateAccountServererror(body *CreateAccountServererrorResponseBody) *goa.ServiceError {
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

// NewCreateAccountBadrequest builds a auth service CreateAccount endpoint
// badrequest error.
func NewCreateAccountBadrequest(body *CreateAccountBadrequestResponseBody) *goa.ServiceError {
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

// NewCreateAccountNotfound builds a auth service CreateAccount endpoint
// notfound error.
func NewCreateAccountNotfound(body *CreateAccountNotfoundResponseBody) *goa.ServiceError {
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

// NewGrantTokenTokenOK builds a "auth" service "GrantToken" endpoint result
// from a HTTP "OK" response.
func NewGrantTokenTokenOK(body *GrantTokenResponseBody, token *string) *auth.Token {
	v := &auth.Token{
		AccessToken:            *body.AccessToken,
		RefreshToken:           *body.RefreshToken,
		AccessExpiryTime:       *body.AccessExpiryTime,
		RefreshExpiryTime:      *body.RefreshExpiryTime,
		AccessExpiryTimeStamp:  *body.AccessExpiryTimeStamp,
		RefreshExpiryTimeStamp: *body.RefreshExpiryTimeStamp,
	}
	v.Token = token

	return v
}

// NewGrantTokenUnauthorized builds a auth service GrantToken endpoint
// unauthorized error.
func NewGrantTokenUnauthorized(body *GrantTokenUnauthorizedResponseBody) *goa.ServiceError {
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

// NewGrantTokenServererror builds a auth service GrantToken endpoint
// servererror error.
func NewGrantTokenServererror(body *GrantTokenServererrorResponseBody) *goa.ServiceError {
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

// NewGrantTokenBadrequest builds a auth service GrantToken endpoint badrequest
// error.
func NewGrantTokenBadrequest(body *GrantTokenBadrequestResponseBody) *goa.ServiceError {
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

// NewGrantTokenNotfound builds a auth service GrantToken endpoint notfound
// error.
func NewGrantTokenNotfound(body *GrantTokenNotfoundResponseBody) *goa.ServiceError {
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

// ValidateGrantTokenResponseBody runs the validations defined on
// GrantTokenResponseBody
func ValidateGrantTokenResponseBody(body *GrantTokenResponseBody) (err error) {
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

// ValidateCreateAccountUnauthorizedResponseBody runs the validations defined
// on CreateAccount_unauthorized_Response_Body
func ValidateCreateAccountUnauthorizedResponseBody(body *CreateAccountUnauthorizedResponseBody) (err error) {
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

// ValidateCreateAccountServererrorResponseBody runs the validations defined on
// CreateAccount_servererror_Response_Body
func ValidateCreateAccountServererrorResponseBody(body *CreateAccountServererrorResponseBody) (err error) {
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

// ValidateCreateAccountBadrequestResponseBody runs the validations defined on
// CreateAccount_badrequest_Response_Body
func ValidateCreateAccountBadrequestResponseBody(body *CreateAccountBadrequestResponseBody) (err error) {
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

// ValidateCreateAccountNotfoundResponseBody runs the validations defined on
// CreateAccount_notfound_Response_Body
func ValidateCreateAccountNotfoundResponseBody(body *CreateAccountNotfoundResponseBody) (err error) {
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

// ValidateGrantTokenUnauthorizedResponseBody runs the validations defined on
// GrantToken_unauthorized_Response_Body
func ValidateGrantTokenUnauthorizedResponseBody(body *GrantTokenUnauthorizedResponseBody) (err error) {
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

// ValidateGrantTokenServererrorResponseBody runs the validations defined on
// GrantToken_servererror_Response_Body
func ValidateGrantTokenServererrorResponseBody(body *GrantTokenServererrorResponseBody) (err error) {
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

// ValidateGrantTokenBadrequestResponseBody runs the validations defined on
// GrantToken_badrequest_Response_Body
func ValidateGrantTokenBadrequestResponseBody(body *GrantTokenBadrequestResponseBody) (err error) {
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

// ValidateGrantTokenNotfoundResponseBody runs the validations defined on
// GrantToken_notfound_Response_Body
func ValidateGrantTokenNotfoundResponseBody(body *GrantTokenNotfoundResponseBody) (err error) {
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