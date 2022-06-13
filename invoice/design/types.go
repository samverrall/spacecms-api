package design

import (
	. "goa.design/goa/v3/dsl"
)

// user defines the fields that can exist on a user.
var user = Type("User", func() {
	Attribute("id", String, "ID of the user")
	Attribute("email", String, "Email address of the user")
	Attribute("name", String, "Full name of the user")
	Attribute("password", String, "Password of the user. This is never returned to the client.")

	// These fields have to be on the user or the API will error.
	Required("email", "name", "password")
})
