// Code generated by goa v3.7.3, DO NOT EDIT.
//
// cms HTTP client CLI support package
//
// Command:
// $ goa gen github.com/samverrall/spacecms-api/spacecms-api/design

package client

import (
	"encoding/json"
	"fmt"

	cms "github.com/samverrall/spacecms-api/gen/cms"
	goa "goa.design/goa/v3/pkg"
)

// BuildCreatePagePayload builds the payload for the cms CreatePage endpoint
// from CLI flags.
func BuildCreatePagePayload(cmsCreatePageBody string, cmsCreatePageToken string) (*cms.CreatePagePayload, error) {
	var err error
	var body CreatePageRequestBody
	{
		err = json.Unmarshal([]byte(cmsCreatePageBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"createdAt\": \"2014-06-29T21:24:29Z\",\n      \"id\": \"FF1D889F-0741-6290-783B-66E606310D86\",\n      \"isActive\": false,\n      \"meta\": {\n         \"description\": \"Excepturi non occaecati odit voluptates deleniti.\",\n         \"title\": \"Dolores optio.\"\n      },\n      \"templateId\": \"Nisi odio velit ducimus facilis molestiae.\",\n      \"url\": \"Ratione reprehenderit quae voluptas.\"\n   }'")
		}
		if body.Meta == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("meta", "body"))
		}
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", body.ID, goa.FormatUUID))

		err = goa.MergeErrors(err, goa.ValidateFormat("body.createdAt", body.CreatedAt, goa.FormatDateTime))

		if err != nil {
			return nil, err
		}
	}
	var token *string
	{
		if cmsCreatePageToken != "" {
			token = &cmsCreatePageToken
		}
	}
	v := &cms.CreatePagePayload{
		ID:         body.ID,
		URL:        body.URL,
		TemplateID: body.TemplateID,
		IsActive:   body.IsActive,
		CreatedAt:  body.CreatedAt,
	}
	if body.Meta != nil {
		v.Meta = marshalMetaRequestBodyToCmsMeta(body.Meta)
	}
	v.Token = token

	return v, nil
}

// BuildCreateTemplatePayload builds the payload for the cms CreateTemplate
// endpoint from CLI flags.
func BuildCreateTemplatePayload(cmsCreateTemplateBody string, cmsCreateTemplateToken string) (*cms.CreateTemplatePayload, error) {
	var err error
	var body CreateTemplateRequestBody
	{
		err = json.Unmarshal([]byte(cmsCreateTemplateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"blockId\": \"Animi et iure fugit vitae totam.\",\n      \"createdAt\": \"1975-01-21T23:59:55Z\",\n      \"id\": \"74C53540-E974-ABFF-2565-6BF99F9017B2\",\n      \"name\": \"Atque sint ea laborum.\"\n   }'")
		}
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", body.ID, goa.FormatUUID))

		err = goa.MergeErrors(err, goa.ValidateFormat("body.createdAt", body.CreatedAt, goa.FormatDateTime))

		if err != nil {
			return nil, err
		}
	}
	var token *string
	{
		if cmsCreateTemplateToken != "" {
			token = &cmsCreateTemplateToken
		}
	}
	v := &cms.CreateTemplatePayload{
		ID:        body.ID,
		Name:      body.Name,
		BlockID:   body.BlockID,
		CreatedAt: body.CreatedAt,
	}
	v.Token = token

	return v, nil
}
