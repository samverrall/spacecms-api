package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("invoice", func() {
	Title("Invoice App API Service")
	Description("RESTFUL API Backend for Invoicify. An open source invoicing web app.")
	Server("invoice", func() {
		Host("default", func() {
			URI("http://localhost:8000")
		})
	})
})

var _ = Service("invoice", func() {
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

	Method("CreateAccount", func() {
		Description("Create an account by email address and password.")

		Payload(user)

		HTTP(func() {
			POST("/create-account")
			Response(StatusCreated)
		})
	})

	Method("AuthoriseLogin", func() {
		Description("Create an account by email address and password.")

		Payload(func() {
			Attribute("grant_type", String, func() {
				Default("access_token")
				Enum("access_token", "refresh_token")
			})
			Attribute("email", String, "User email")
			Attribute("password", String, "User password")
			Required("email", "password")
		})

		Result(tokenResponse)

		HTTP(func() {
			POST("/tokens")
			Params(func() {
				Param("grant_type")
			})
			Response(StatusOK)
		})
	})

	Files("/openapi.json", "./gen/http/openapi.json")
})
