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
		Description("Create an account by email address and password.")
		Payload(func() {
			Attribute("token", String)
			Attribute("email", String, "User email")
			Attribute("password", String, "User password")
			Required("email", "password")
		})

		Result(tokenResponse)

		HTTP(func() {
			Cookie("token:__Host-token", String)
			POST("/pages")
			Response(StatusCreated)
		})
	})

	Files("/openapi.json", "./gen/http/openapi.json")
})
