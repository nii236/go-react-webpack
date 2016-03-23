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
	Produces("application/json")
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
		Response(OK, "plain/text")
	})

	Action("Callback response from Google", func() {
		Routing(GET("GoogleCallback"))
		Description("This will be the response code and random string from Google after oauth2")
		Response(OK, "plain/text")
	})
	Action("Log into Facebook", func() {
		Routing(GET("FacebookLogin"))
		Description("This starts the oauth2 handshake process")
		Response(OK, "plain/text")
	})

	Action("Callback response from Facebook", func() {
		Routing(GET("FacebookCallback"))
		Description("This will be the response code and random string from Google after oauth2")
		Response(OK, "plain/text")
	})
	Action("Log into Github", func() {
		Routing(GET("GithubLogin"))
		Description("This starts the oauth2 handshake process")
		Response(OK, "plain/text")
	})

	Action("Callback response from Github", func() {
		Routing(GET("GithubCallback"))
		Description("This will be the response code and random string from Google after oauth2")
		Response(OK, "plain/text")
	})
})
