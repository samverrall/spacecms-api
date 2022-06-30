package controllers

import (
	"context"

	"github.com/golang-jwt/jwt/v4"
)

type authInfo struct {
	user   string
	claims jwt.MapClaims
}

func (auth authInfo) String() string {
	if auth.user != "" {
		return "AuthInfo: Username + Password"
	} else if auth.claims != nil {
		return "AuthInfo: JWT/OAuth"
	} else {
		return "AuthInfo: none"
	}
}

type ctxValue int

const (
	ctxValueClaims ctxValue = iota
)

// contextWithAuthInfo adds the given JWT claims to the context and returns it.
func contextWithAuthInfo(ctx context.Context, auth authInfo) context.Context {
	return context.WithValue(ctx, ctxValueClaims, auth)
}
