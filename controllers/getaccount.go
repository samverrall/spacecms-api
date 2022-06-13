package controllers

import (
	"context"

	"github.com/samverrall/invoice-app/gen/invoice"
)

func (s *invoicesrvc) GetAccount(ctx context.Context, p *invoice.GetAccountPayload) (*invoice.User, error) {
	return nil, nil
}
