package jwttoken

import (
	"context"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/samverrall/invoice-api-service/gen/invoice"
)

type JWTToken struct {
	accessTokenExpiryMinutes int
}

func New(accessTokenExpiryMinutes int) *JWTToken {
	return &JWTToken{
		accessTokenExpiryMinutes: accessTokenExpiryMinutes,
	}
}

// https://developer.vonage.com/blog/20/03/13/using-jwt-for-authentication-in-a-golang-application-dr#:~:text=Refresh%20Token%3A%20A%20refresh%20token,hit%20(from%20our%20application).
func (jt *JWTToken) CreateTokenPair(ctx context.Context, userID *string) (*invoice.Token, error) {
	atExpiryTime := time.Now().Add(time.Minute * time.Duration(jt.accessTokenExpiryMinutes))
	rtExpiryTime := time.Now().Add(time.Hour * 24 * 7)

	at, err := jt.NewAccessToken(ctx, jwt.MapClaims{
		"expiresAt": atExpiryTime.Unix(),
		"issuedAt":  time.Now().Unix(),
		"userID":    userID,
	})
	if err != nil {
		return nil, err
	}

	rt, err := jt.NewRefreshToken(ctx, jwt.MapClaims{
		"expiresAt": rtExpiryTime.Unix(),
		"issuedAt":  time.Now().Unix(),
		"userID":    userID,
	})
	if err != nil {
		return nil, err
	}

	return &invoice.Token{
		AccessToken:            at,
		RefreshToken:           rt,
		AccessExpiryTime:       atExpiryTime.Unix(),
		RefreshExpiryTime:      rtExpiryTime.Unix(),
		AccessExpiryTimeStamp:  atExpiryTime.UTC().String(),
		RefreshExpiryTimeStamp: rtExpiryTime.UTC().String(),
	}, nil
}

func (jt *JWTToken) NewAccessToken(ctx context.Context, claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(os.Getenv("INVOICE_SECRET")))
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func (jt *JWTToken) NewRefreshToken(ctx context.Context, claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshToken, err := token.SignedString([]byte(os.Getenv("INVOICE_SECRET_REFRESH")))
	if err != nil {
		return "", err
	}
	return refreshToken, nil
}

// func (jt *JWTToken) Verify(token string) error {
// 	jwt.Parse(token, func() {

// 	})
// }
