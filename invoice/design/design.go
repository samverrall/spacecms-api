package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("invoice", func() {
	Title("Invoice App API Service")
	Description("REST API for the invoice app.")
	Server("invoice", func() {
		Host("localhost", func() {
			URI("http://localhost:8000")
			URI("grpc://localhost:8080")
		})
	})
})

var _ = Service("invoice", func() {
	Description("REST API for the invoice app.")

	HTTP(func() {
		Response("unauthorized", StatusUnauthorized)
		Response("servererror", StatusInternalServerError)
		Response("badrequest", StatusBadRequest)
		Response("notfound", StatusNotFound)
	})

	Method("create-account", func() {
		// The payload is the JSON request body with the fields of the type 'user'.
		Payload(user)

		Result(user)

		HTTP(func() {
			POST("/create-account")

			// When a account has been created we need to use a http.StatusCreated status code.
			// See https://developer.mozilla.org/en-US/docs/Web/HTTP/Status for status codes.
			Response(StatusCreated)
		})
	})

	Method("get-account", func() {
		Result(user)

		HTTP(func() {
			GET("/create-account/{userID}")

			// When a account has been created we need to use a http.StatusCreated status code.
			// See https://developer.mozilla.org/en-US/docs/Web/HTTP/Status for status codes.
			Response(StatusOK)
		})
	})

	Files("/openapi.json", "./gen/http/openapi.json")
})

// user defines the fields that can exist on a user.
var user = Type("User", func() {
	Attribute("email", String)
	Attribute("name", String)
	Attribute("password", String)

	// These fields have to be on the user or the API will error.
	Required("email", "name", "password")
})
