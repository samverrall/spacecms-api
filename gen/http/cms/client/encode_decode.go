// Code generated by goa v3.7.3, DO NOT EDIT.
//
// cms HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/samverrall/spacecms-api/spacecms-api/design

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"

	cms "github.com/samverrall/spacecms-api/gen/cms"
	goahttp "goa.design/goa/v3/http"
)

// BuildCreatePageRequest instantiates a HTTP request object with method and
// path set to call the "cms" service "CreatePage" endpoint
func (c *Client) BuildCreatePageRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CreatePageCmsPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("cms", "CreatePage", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCreatePageRequest returns an encoder for requests sent to the cms
// CreatePage server.
func EncodeCreatePageRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*cms.CreatePagePayload)
		if !ok {
			return goahttp.ErrInvalidType("cms", "CreatePage", "*cms.CreatePagePayload", v)
		}
		if p.Token != nil {
			head := *p.Token
			req.Header.Set("X-Authorization", head)
		}
		body := NewCreatePageRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("cms", "CreatePage", err)
		}
		return nil
	}
}

// DecodeCreatePageResponse returns a decoder for responses returned by the cms
// CreatePage endpoint. restoreBody controls whether the response body should
// be restored after having been read.
// DecodeCreatePageResponse may return the following errors:
//	- "unauthorized" (type *goa.ServiceError): http.StatusUnauthorized
//	- "servererror" (type *goa.ServiceError): http.StatusInternalServerError
//	- "badrequest" (type *goa.ServiceError): http.StatusBadRequest
//	- "notfound" (type *goa.ServiceError): http.StatusNotFound
//	- error: internal error
func DecodeCreatePageResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
			var (
				body CreatePageResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("cms", "CreatePage", err)
			}
			err = ValidateCreatePageResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("cms", "CreatePage", err)
			}
			res := NewCreatePagePageCreated(&body)
			return res, nil
		case http.StatusUnauthorized:
			var (
				body CreatePageUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("cms", "CreatePage", err)
			}
			err = ValidateCreatePageUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("cms", "CreatePage", err)
			}
			return nil, NewCreatePageUnauthorized(&body)
		case http.StatusInternalServerError:
			var (
				body CreatePageServererrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("cms", "CreatePage", err)
			}
			err = ValidateCreatePageServererrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("cms", "CreatePage", err)
			}
			return nil, NewCreatePageServererror(&body)
		case http.StatusBadRequest:
			var (
				body CreatePageBadrequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("cms", "CreatePage", err)
			}
			err = ValidateCreatePageBadrequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("cms", "CreatePage", err)
			}
			return nil, NewCreatePageBadrequest(&body)
		case http.StatusNotFound:
			var (
				body CreatePageNotfoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("cms", "CreatePage", err)
			}
			err = ValidateCreatePageNotfoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("cms", "CreatePage", err)
			}
			return nil, NewCreatePageNotfound(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("cms", "CreatePage", resp.StatusCode, string(body))
		}
	}
}

// marshalCmsMetaToMetaRequestBody builds a value of type *MetaRequestBody from
// a value of type *cms.Meta.
func marshalCmsMetaToMetaRequestBody(v *cms.Meta) *MetaRequestBody {
	res := &MetaRequestBody{
		Title:       v.Title,
		Description: v.Description,
	}

	return res
}

// marshalMetaRequestBodyToCmsMeta builds a value of type *cms.Meta from a
// value of type *MetaRequestBody.
func marshalMetaRequestBodyToCmsMeta(v *MetaRequestBody) *cms.Meta {
	res := &cms.Meta{
		Title:       v.Title,
		Description: v.Description,
	}

	return res
}

// unmarshalMetaResponseBodyToCmsMeta builds a value of type *cms.Meta from a
// value of type *MetaResponseBody.
func unmarshalMetaResponseBodyToCmsMeta(v *MetaResponseBody) *cms.Meta {
	res := &cms.Meta{
		Title:       v.Title,
		Description: v.Description,
	}

	return res
}
