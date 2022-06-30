package design

import (
	. "goa.design/goa/v3/dsl"
)

var page = Type("Page", func() {
	Attribute("id", String, "Page UUID", func() {
		Format(FormatUUID)
	})
	Attribute("url", String, "Page URL")
	Attribute("templateId", String, "Page template UUID")
	Attribute("isActive", Boolean, "Page active")
	Attribute("meta", pageMeta)
	Attribute("createdAt", String, func() {
		Format(FormatDateTime)
	})

	Required("id", "url", "templateId", "isActive", "meta", "createdAt")
})

var pageMeta = Type("Meta", func() {
	Attribute("title", String, "Page meta title")
	Attribute("description", String, "Page meta description")
})
