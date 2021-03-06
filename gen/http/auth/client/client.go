// Code generated by goa v3.7.3, DO NOT EDIT.
//
// auth client HTTP transport
//
// Command:
// $ goa gen github.com/samverrall/spacecms-api/spacecms-api/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the auth service endpoint HTTP clients.
type Client struct {
	// CreateAccount Doer is the HTTP client used to make requests to the
	// CreateAccount endpoint.
	CreateAccountDoer goahttp.Doer

	// GrantToken Doer is the HTTP client used to make requests to the GrantToken
	// endpoint.
	GrantTokenDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the auth service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		CreateAccountDoer:   doer,
		GrantTokenDoer:      doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// CreateAccount returns an endpoint that makes HTTP requests to the auth
// service CreateAccount server.
func (c *Client) CreateAccount() goa.Endpoint {
	var (
		encodeRequest  = EncodeCreateAccountRequest(c.encoder)
		decodeResponse = DecodeCreateAccountResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildCreateAccountRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.CreateAccountDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("auth", "CreateAccount", err)
		}
		return decodeResponse(resp)
	}
}

// GrantToken returns an endpoint that makes HTTP requests to the auth service
// GrantToken server.
func (c *Client) GrantToken() goa.Endpoint {
	var (
		encodeRequest  = EncodeGrantTokenRequest(c.encoder)
		decodeResponse = DecodeGrantTokenResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildGrantTokenRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.GrantTokenDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("auth", "GrantToken", err)
		}
		return decodeResponse(resp)
	}
}
