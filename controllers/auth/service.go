package controllers

import (
	"github.com/samverrall/spacecms-api/datastore"
	"github.com/samverrall/spacecms-api/gen/auth"
	"github.com/samverrall/spacecms-api/hasher"
	"github.com/samverrall/spacecms-api/hasher/argon2id"
	"github.com/samverrall/spacecms-api/tokens"
	"github.com/samverrall/spacecms-api/tokens/jwttoken"
	log "github.com/sirupsen/logrus"
)

const (
	accessTokenExpiryMinutes = 15
)

// auth service example implementation.
// The example methods log the requests and return zero values.
type authservice struct {
	logger  *log.Logger
	dbi     datastore.DBInterface
	hasher  hasher.Hasher
	tokener tokens.Tokener
}

// NewAuthService returns the auth service implementation.
func NewAuthService(logger *log.Logger, dbi *datastore.DataStore) auth.Service {
	if dbi == nil {
		dbi = datastore.New()
	}

	return &authservice{
		logger:  logger,
		dbi:     dbi,
		hasher:  argon2id.New(),
		tokener: jwttoken.New(accessTokenExpiryMinutes),
	}
}
