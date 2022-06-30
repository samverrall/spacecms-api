package controllers

import (
	"context"
	"os"

	"github.com/golang-jwt/jwt/v4"
	log "github.com/sirupsen/logrus"
	"goa.design/goa/v3/security"

	"github.com/samverrall/spacecms-api/datastore"
	"github.com/samverrall/spacecms-api/gen/cms"
)

// auth service example implementation.
// The example methods log the requests and return zero values.
type cmsservice struct {
	logger *log.Logger
	dbi    datastore.DBInterface
}

// NewCMSService returns the auth service implementation.
func NewCMSService(logger *log.Logger, dbi *datastore.DataStore) cms.Service {
	if dbi == nil {
		dbi = datastore.New()
	}

	return &cmsservice{
		logger: logger,
		dbi:    dbi,
	}
}

// JWTAuth implements the authorization logic for service "secured_service" for
// the "jwt" security scheme.
func (s *cmsservice) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
	claims := make(jwt.MapClaims)

	_, err := jwt.ParseWithClaims(token, claims, func(_ *jwt.Token) (interface{}, error) { return os.Getenv("INVOICE_SECRET"), nil })
	if err != nil {
		return ctx, err
	}

	if claims["scopes"] == nil {
		return ctx, err
	}
	scopes, ok := claims["scopes"].([]interface{})
	if !ok {
		return ctx, err
	}
	scopesInToken := make([]string, len(scopes))
	for _, scp := range scopes {
		scopesInToken = append(scopesInToken, scp.(string))
	}
	if err := scheme.Validate(scopesInToken); err != nil {
		return ctx, err
	}

	ctx = contextWithAuthInfo(ctx, authInfo{
		claims: claims,
	})
	return ctx, nil
}
