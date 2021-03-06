// Code generated by goa v3.7.3, DO NOT EDIT.
//
// auth HTTP client CLI support package
//
// Command:
// $ goa gen github.com/samverrall/spacecms-api/spacecms-api/design

package client

import (
	"encoding/json"
	"fmt"

	auth "github.com/samverrall/spacecms-api/gen/auth"
)

// BuildCreateAccountPayload builds the payload for the auth CreateAccount
// endpoint from CLI flags.
func BuildCreateAccountPayload(authCreateAccountBody string) (*auth.User, error) {
	var err error
	var body CreateAccountRequestBody
	{
		err = json.Unmarshal([]byte(authCreateAccountBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"email\": \"Consequatur aut eos est dolor.\",\n      \"id\": \"Ea alias minus possimus aut.\",\n      \"name\": \"Ea minus cupiditate.\",\n      \"password\": \"Omnis ex.\"\n   }'")
		}
	}
	v := &auth.User{
		ID:       body.ID,
		Email:    body.Email,
		Name:     body.Name,
		Password: body.Password,
	}

	return v, nil
}

// BuildGrantTokenPayload builds the payload for the auth GrantToken endpoint
// from CLI flags.
func BuildGrantTokenPayload(authGrantTokenBody string, authGrantTokenToken string) (*auth.GrantTokenPayload, error) {
	var err error
	var body GrantTokenRequestBody
	{
		err = json.Unmarshal([]byte(authGrantTokenBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"email\": \"Quia illo voluptatem.\",\n      \"password\": \"Quas quas autem.\"\n   }'")
		}
	}
	var token *string
	{
		if authGrantTokenToken != "" {
			token = &authGrantTokenToken
		}
	}
	v := &auth.GrantTokenPayload{
		Email:    body.Email,
		Password: body.Password,
	}
	v.Token = token

	return v, nil
}
