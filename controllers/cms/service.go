package controllers

import (
	"context"
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v4"
	log "github.com/sirupsen/logrus"
	"goa.design/goa/v3/security"

	"github.com/samverrall/spacecms-api/datastore"
	"github.com/samverrall/spacecms-api/gen/cms"
	"github.com/samverrall/spacecms-api/tokens/jwttoken"
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

// JWTAuth implements the authorization logic for the cms service for
// the "jwt" security scheme.
func (s *cmsservice) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
	s.logger.Info("JWTAuth")

	t, err := jwt.ParseWithClaims(token, &jwttoken.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SPACECMS_SECRET")), nil
	})
	if err != nil {
		return ctx, err
	}

	claims, ok := t.Claims.(*jwttoken.Claims)
	if !ok || !t.Valid {
		return ctx, errors.New("got invalid token")
	}

	s.logger.Info("UserID from token: ", claims.UserID)

	ctx = contextWithAuthInfo(ctx, authInfo{
		user: claims.UserID,
	})

	return ctx, nil
}
