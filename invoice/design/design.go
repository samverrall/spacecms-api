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

	Error("unauthorized")
	Error("servererror")
	Error("badrequest")
	Error("notfound")

	HTTP(func() {
		Response("unauthorized", StatusUnauthorized)
		Response("servererror", StatusInternalServerError)
		Response("badrequest", StatusBadRequest)
		Response("notfound", StatusNotFound)
	})

	Method("create-account", func() {
		// The payload is the JSON request body with the fields of the type 'user'.
		Payload(user)

		HTTP(func() {
			POST("/create-account")

			// When a account has been created we need to use a http.StatusCreated status code.
			// See https://developer.mozilla.org/en-US/docs/Web/HTTP/Status for status codes.
			Response(StatusCreated)
		})
	})

	Method("get-account", func() {
		Result(user)

		Payload(func() {
			Attribute("userID", String)
		})

		HTTP(func() {
			GET("/create-account/{userID}")
			Params(func() {
				Param("userID")
			})

			// When a account has been created we need to use a http.StatusCreated status code.
			// See https://developer.mozilla.org/en-US/docs/Web/HTTP/Status for status codes.
			Response(StatusOK)
		})
	})

	Files("/openapi.json", "./gen/http/openapi.json")
})
