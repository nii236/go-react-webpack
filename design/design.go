package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("adder", func() {
	Title("The adder API")
	Description("A teaser for goa")
	Host("localhost:8080")
	Scheme("http")
})

var _ = Resource("operands", func() {
	Action("add", func() {
		Routing(GET("add/:left/:right"))
		Description("add returns the sum of the left and right parameters in the response body")
		Params(func() {
			Param("left", Integer, "Left operand")
			Param("right", Integer, "Right operand")
		})
		Response(OK, "plain/text")
	})

	Action("multiply", func() {
		Routing(GET("multiply/:left/:right"))
		Description("multiply returns the multiplication of the left and right parameters in the response body")
		Params(func() {
			Param("left", Integer, "Left operand")
			Param("right", Integer, "Right operand")
		})
		Response(OK, "plain/text")
	})
})

var _ = Resource("authentication", func() {
	Action("Log into Google", func() {
		Routing(GET("GoogleLogin"))
		Description("This starts the oauth2 handshake process")
		Params(func() {
			Param("user", String, "Username")
			Param("password", String, "Password")
		})
		Response(OK, "plain/text")
	})

	Action("Callback response from Google", func() {
		Routing(POST("GoogleCallback"))
		Description("This will be the response code and random string from Google after oauth2")
		Params(func() {
			Param("user", String, "Left operand")
			Param("password", String, "Password")
		})
		Response(OK, "plain/text")
	})
})
