package apierror

import (
	goa "goa.design/goa/v3/pkg"
)

func NewResponseError(name, message string) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    name,
		ID:      goa.NewErrorID(),
		Message: message,
	}
}
