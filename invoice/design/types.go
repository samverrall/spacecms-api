package design

import (
	. "goa.design/goa/v3/dsl"
)

// user defines the fields that can exist on a user.
var user = Type("User", func() {
	Attribute("id", String)
	Attribute("email", String)
	Attribute("name", String)
	Attribute("password", String)

	// These fields have to be on the user or the API will error.
	Required("email", "name", "password")
})
