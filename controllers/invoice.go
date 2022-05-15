package controllers

import (
	"log"

	"github.com/samverrall/invoice-app/datastore"
	invoice "github.com/samverrall/invoice-app/gen/invoice"
	goa "goa.design/goa/v3/pkg"
)

// invoice service example implementation.
// The example methods log the requests and return zero values.
type invoicesrvc struct {
	logger *log.Logger
	dbi    datastore.DBInterface
}

// NewInvoice returns the invoice service implementation.
func NewInvoice(logger *log.Logger, dbi *datastore.DataStore) invoice.Service {
	if dbi == nil {
		dbi = datastore.New()
	}

	return &invoicesrvc{
		logger: logger,
		dbi:    dbi,
	}
}

func newResponseError(name, message string) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    name,
		ID:      goa.NewErrorID(),
		Message: message,
	}
}
