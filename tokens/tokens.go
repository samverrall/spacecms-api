package tokens

import (
	"context"

	"github.com/golang-jwt/jwt"
	"github.com/samverrall/invoice-api-service/gen/invoice"
)

type Tokener interface {
	CreateTokenPair(ctx context.Context, userID *string) (*invoice.Token, error)
	NewAccessToken(ctx context.Context, claims jwt.MapClaims) (string, error)
	NewRefreshToken(ctx context.Context, claims jwt.MapClaims) (string, error)
}
