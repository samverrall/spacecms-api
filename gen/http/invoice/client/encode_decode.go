// Code generated by goa v3.7.3, DO NOT EDIT.
//
// invoice HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/samverrall/invoice-app/invoice/design

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"

	invoice "github.com/samverrall/invoice-app/gen/invoice"
	goahttp "goa.design/goa/v3/http"
)

// BuildCreateAccountRequest instantiates a HTTP request object with method and
// path set to call the "invoice" service "create-account" endpoint
func (c *Client) BuildCreateAccountRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CreateAccountInvoicePath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("invoice", "create-account", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCreateAccountRequest returns an encoder for requests sent to the
// invoice create-account server.
func EncodeCreateAccountRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*invoice.User)
		if !ok {
			return goahttp.ErrInvalidType("invoice", "create-account", "*invoice.User", v)
		}
		body := NewCreateAccountRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("invoice", "create-account", err)
		}
		return nil
	}
}

// DecodeCreateAccountResponse returns a decoder for responses returned by the
// invoice create-account endpoint. restoreBody controls whether the response
// body should be restored after having been read.
// DecodeCreateAccountResponse may return the following errors:
//	- "unauthorized" (type *goa.ServiceError): http.StatusUnauthorized
//	- "servererror" (type *goa.ServiceError): http.StatusInternalServerError
//	- "badrequest" (type *goa.ServiceError): http.StatusBadRequest
//	- "notfound" (type *goa.ServiceError): http.StatusNotFound
//	- error: internal error
func DecodeCreateAccountResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusCreated:
			return nil, nil
		case http.StatusUnauthorized:
			var (
				body CreateAccountUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("invoice", "create-account", err)
			}
			err = ValidateCreateAccountUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("invoice", "create-account", err)
			}
			return nil, NewCreateAccountUnauthorized(&body)
		case http.StatusInternalServerError:
			var (
				body CreateAccountServererrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("invoice", "create-account", err)
			}
			err = ValidateCreateAccountServererrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("invoice", "create-account", err)
			}
			return nil, NewCreateAccountServererror(&body)
		case http.StatusBadRequest:
			var (
				body CreateAccountBadrequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("invoice", "create-account", err)
			}
			err = ValidateCreateAccountBadrequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("invoice", "create-account", err)
			}
			return nil, NewCreateAccountBadrequest(&body)
		case http.StatusNotFound:
			var (
				body CreateAccountNotfoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("invoice", "create-account", err)
			}
			err = ValidateCreateAccountNotfoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("invoice", "create-account", err)
			}
			return nil, NewCreateAccountNotfound(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("invoice", "create-account", resp.StatusCode, string(body))
		}
	}
}

// BuildGetAccountRequest instantiates a HTTP request object with method and
// path set to call the "invoice" service "get-account" endpoint
func (c *Client) BuildGetAccountRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		userID string
	)
	{
		p, ok := v.(*invoice.GetAccountPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("invoice", "get-account", "*invoice.GetAccountPayload", v)
		}
		if p.UserID != nil {
			userID = *p.UserID
		}
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetAccountInvoicePath(userID)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("invoice", "get-account", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeGetAccountResponse returns a decoder for responses returned by the
// invoice get-account endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeGetAccountResponse may return the following errors:
//	- "unauthorized" (type *goa.ServiceError): http.StatusUnauthorized
//	- "servererror" (type *goa.ServiceError): http.StatusInternalServerError
//	- "badrequest" (type *goa.ServiceError): http.StatusBadRequest
//	- "notfound" (type *goa.ServiceError): http.StatusNotFound
//	- error: internal error
func DecodeGetAccountResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body GetAccountResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("invoice", "get-account", err)
			}
			err = ValidateGetAccountResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("invoice", "get-account", err)
			}
			res := NewGetAccountUserOK(&body)
			return res, nil
		case http.StatusUnauthorized:
			var (
				body GetAccountUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("invoice", "get-account", err)
			}
			err = ValidateGetAccountUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("invoice", "get-account", err)
			}
			return nil, NewGetAccountUnauthorized(&body)
		case http.StatusInternalServerError:
			var (
				body GetAccountServererrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("invoice", "get-account", err)
			}
			err = ValidateGetAccountServererrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("invoice", "get-account", err)
			}
			return nil, NewGetAccountServererror(&body)
		case http.StatusBadRequest:
			var (
				body GetAccountBadrequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("invoice", "get-account", err)
			}
			err = ValidateGetAccountBadrequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("invoice", "get-account", err)
			}
			return nil, NewGetAccountBadrequest(&body)
		case http.StatusNotFound:
			var (
				body GetAccountNotfoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("invoice", "get-account", err)
			}
			err = ValidateGetAccountNotfoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("invoice", "get-account", err)
			}
			return nil, NewGetAccountNotfound(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("invoice", "get-account", resp.StatusCode, string(body))
		}
	}
}
