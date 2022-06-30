package tokens

import (
	"context"

	"github.com/samverrall/spacecms-api/gen/auth"
	"github.com/samverrall/spacecms-api/tokens/jwttoken"
)

type Tokener interface {
	CreateTokenPair(ctx context.Context, userID *string) (*auth.Token, error)
	NewAccessToken(ctx context.Context, claims *jwttoken.Claims) (string, error)
	NewRefreshToken(ctx context.Context, claims *jwttoken.Claims) (string, error)
}
