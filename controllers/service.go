package controllers

import (
	"github.com/samverrall/invoice-api-service/datastore"
	invoice "github.com/samverrall/invoice-api-service/gen/invoice"
	"github.com/samverrall/invoice-api-service/hasher"
	"github.com/samverrall/invoice-api-service/hasher/argon2id"
	log "github.com/sirupsen/logrus"
)

// invoice service example implementation.
// The example methods log the requests and return zero values.
type invoicesrvc struct {
	logger *log.Logger
	dbi    datastore.DBInterface
	hasher hasher.Hasher
}

// NewInvoice returns the invoice service implementation.
func NewInvoice(logger *log.Logger, dbi *datastore.DataStore) invoice.Service {
	if dbi == nil {
		dbi = datastore.New()
	}

	return &invoicesrvc{
		logger: logger,
		dbi:    dbi,
		hasher: argon2id.New(),
	}
}
