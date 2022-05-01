package invoiceapi

import (
	"context"
	"log"

	invoice "github.com/samverrall/invoice-app/gen/invoice"
)

// invoice service example implementation.
// The example methods log the requests and return zero values.
type invoicesrvc struct {
	logger *log.Logger
}

// NewInvoice returns the invoice service implementation.
func NewInvoice(logger *log.Logger) invoice.Service {
	return &invoicesrvc{logger}
}

// CreateAccount implements create-account.
func (s *invoicesrvc) CreateAccount(ctx context.Context, p *invoice.User) (*invoice.User, error) {
	s.logger.Print("invoice.create-account")

	return p, nil
}
