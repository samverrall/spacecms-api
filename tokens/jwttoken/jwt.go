package jwttoken

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/samverrall/spacecms-api/gen/auth"
)

type JWTToken struct {
	accessTokenExpiryMinutes int
}

func New(accessTokenExpiryMinutes int) *JWTToken {
	return &JWTToken{
		accessTokenExpiryMinutes: accessTokenExpiryMinutes,
	}
}

type Claims struct {
	UserID string `json:"userId"`
	jwt.StandardClaims
}

// https://developer.vonage.com/blog/20/03/13/using-jwt-for-authentication-in-a-golang-application-dr#:~:text=Refresh%20Token%3A%20A%20refresh%20token,hit%20(from%20our%20application).
func (jt *JWTToken) CreateTokenPair(ctx context.Context, userID *string) (*auth.Token, error) {
	ttl := time.Duration(jt.accessTokenExpiryMinutes) * time.Minute
	atExpiryTime := time.Now().Add(ttl)
	rtExpiryTime := time.Now().Add(time.Hour * 24 * 7)

	if userID == nil {
		return nil, errors.New("no user id supplied")
	}

	at, err := jt.NewAccessToken(ctx, &Claims{
		*userID,
		jwt.StandardClaims{
			ExpiresAt: atExpiryTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})
	if err != nil {
		return nil, err
	}

	rt, err := jt.NewRefreshToken(ctx, &Claims{
		*userID,
		jwt.StandardClaims{
			ExpiresAt: rtExpiryTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})
	if err != nil {
		return nil, err
	}

	return &auth.Token{
		AccessToken:            at,
		RefreshToken:           rt,
		AccessExpiryTime:       atExpiryTime.Unix(),
		RefreshExpiryTime:      rtExpiryTime.Unix(),
		AccessExpiryTimeStamp:  atExpiryTime.UTC().String(),
		RefreshExpiryTimeStamp: rtExpiryTime.UTC().String(),
	}, nil
}

func (jt *JWTToken) NewAccessToken(ctx context.Context, claims *Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(os.Getenv("SPACECMS_SECRET")))
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func (jt *JWTToken) NewRefreshToken(ctx context.Context, claims *Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshToken, err := token.SignedString([]byte(os.Getenv("SPACECMS_SECRET_REFRESH")))
	if err != nil {
		return "", err
	}
	return refreshToken, nil
}
