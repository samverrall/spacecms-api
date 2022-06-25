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

var tokenResponse = Type("Token", func() {
	Attribute("token", String) // Setting this value sets a HTTP Cookie

	Attribute("accessToken", String, func() {
		Example("eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzaWQiOiJkYmViMjlhYTMyYjg0MTMxYTA0NjY4MDAyNzAxNWEwZSIsInJvbGUiOlsiQWRtaW5pc3RyYXRvcnMiLCJSZWdpc3RlcmVkIFVzZXJzIiwiU3Vic2NyaWJlcnMiXSwiaXNzIjoidGVzdHNpdGVjZS5sdmgubWUiLCJleHAiOjE0NTA4MzU2ODMsIm5iZiI6MTQ1MDgzMTc4M30.Yf3mmBJ8nV_IozqvvLc8L34dDklU2J7z0uXn3jsICp0")
	})
	Attribute("refreshToken", String)
	Attribute("accessExpiryTime", Int64)
	Attribute("refreshExpiryTime", Int64)
	Attribute("accessExpiryTimeStamp", String)
	Attribute("refreshExpiryTimeStamp", String)
	Required("accessToken", "refreshToken", "accessExpiryTime", "refreshExpiryTime", "accessExpiryTimeStamp", "refreshExpiryTimeStamp")
})
