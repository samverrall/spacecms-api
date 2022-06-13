package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("invoice", func() {
	Title("Invoice App API Service")
	Description("RESTFUL API Backend for Invoicify. An open source invoicing web app.")
	Server("invoice", func() {
		Host("localhost", func() {
			URI("http://localhost:8000")
			URI("grpc://localhost:8080")
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
	})

	Method("create-account", func() {
		Payload(user)

		HTTP(func() {
			POST("/create-account")
			Response(StatusCreated)
		})
	})

	Files("/openapi.json", "./gen/http/openapi.json")
})
