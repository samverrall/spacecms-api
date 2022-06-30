package tokens

import (
	"context"

	"github.com/golang-jwt/jwt"
	"github.com/samverrall/spacecms-api/gen/auth"
)

type Tokener interface {
	CreateTokenPair(ctx context.Context, userID *string) (*auth.Token, error)
	NewAccessToken(ctx context.Context, claims jwt.MapClaims) (string, error)
	NewRefreshToken(ctx context.Context, claims jwt.MapClaims) (string, error)
}
