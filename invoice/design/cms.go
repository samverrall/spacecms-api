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

	Files("/openapi.json", "./gen/http/openapi.json")
})
