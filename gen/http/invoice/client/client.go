// Code generated by goa v3.7.3, DO NOT EDIT.
//
// invoice client HTTP transport
//
// Command:
// $ goa gen github.com/samverrall/invoice-app/invoice/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the invoice service endpoint HTTP clients.
type Client struct {
	// CreateAccount Doer is the HTTP client used to make requests to the
	// create-account endpoint.
	CreateAccountDoer goahttp.Doer

	// GetAccount Doer is the HTTP client used to make requests to the get-account
	// endpoint.
	GetAccountDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the invoice service servers.
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
		GetAccountDoer:      doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// CreateAccount returns an endpoint that makes HTTP requests to the invoice
// service create-account server.
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
			return nil, goahttp.ErrRequestError("invoice", "create-account", err)
		}
		return decodeResponse(resp)
	}
}

// GetAccount returns an endpoint that makes HTTP requests to the invoice
// service get-account server.
func (c *Client) GetAccount() goa.Endpoint {
	var (
		decodeResponse = DecodeGetAccountResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildGetAccountRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.GetAccountDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("invoice", "get-account", err)
		}
		return decodeResponse(resp)
	}
}
