package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("cms", func() {
	Description("CMS service for SpaceCMS. An open source content management system.")

	Error("unauthorized")
	Error("servererror")
	Error("badrequest")
	Error("notfound")

	HTTP(func() {
		Response("unauthorized", StatusUnauthorized)
		Response("servererror", StatusInternalServerError)
		Response("badrequest", StatusBadRequest)
		Response("notfound", StatusNotFound)
		Path("/api/v1")
	})

	Method("CreatePage", func() {
		Description("Create a page")
		Security(JWTAuth)
		Payload(func() {
			Extend(page)
			TokenField(1, "token", String, func() {
				Description("JWT used for authentication")
				Example("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ")
			})
		})

		HTTP(func() {
			Header("token:Authorization") // JWT token passed in "X-Authorization" header
			POST("/pages")
			Response(StatusCreated)
		})
	})

	Method("CreateTemplate", func() {
		Description("Create a template.")
		Security(JWTAuth)
		Payload(func() {
			Extend(template)
			TokenField(1, "token", String, func() {
				Description("JWT used for authentication")
				Example("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ")
			})
		})

		HTTP(func() {
			Header("token:Authorization") // JWT token passed in "X-Authorization" header
			POST("/templates")
			Response(StatusCreated)
		})
	})
})
