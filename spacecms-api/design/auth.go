package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("auth", func() {
	Title("Space CMS API Service")
	Description("Auth service for SpaceCMS. An open source content management system.")
	Server("invoice", func() {
		Host("default", func() {
			URI("http://localhost:8000")
		})
	})
})

var _ = Service("auth", func() {
	Description("RESTFUL API Backend for Invoicify. An open source invoicing web app.")

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

	// Auth
	Method("CreateAccount", func() {
		Description("Create an account by email address and password.")

		Payload(user)

		HTTP(func() {
			POST("/user/signup")
			Response(StatusCreated)
		})
	})

	Method("GrantToken", func() {
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
			POST("/user/tokens")
			Response(StatusOK, func() {
				Cookie("token:__Host-token", String)
				CookieMaxAge(90000)
				CookieDomain("SameSite")
				CookiePath("/")
				CookieSecure()
				CookieHTTPOnly()
			})
		})
	})
	// Auth end

	Files("/openapi.json", "./gen/http/openapi.json")
})
