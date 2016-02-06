//************************************************************************//
// API "adder": Application Contexts
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/nii236/go-react-webpack
// --design=github.com/nii236/go-react-webpack/design
// --pkg=app
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"github.com/goadesign/goa"
	"strconv"
)

// LoginAuthenticationContext provides the authentication login action context.
type LoginAuthenticationContext struct {
	*goa.Context
	Pass     string
	Password *string
	User     string
}

// NewLoginAuthenticationContext parses the incoming request URL and body, performs validations and creates the
// context used by the authentication controller login action.
func NewLoginAuthenticationContext(c *goa.Context) (*LoginAuthenticationContext, error) {
	var err error
	ctx := LoginAuthenticationContext{Context: c}
	rawPass := c.Get("pass")
	if rawPass != "" {
		ctx.Pass = rawPass
	}
	rawPassword := c.Get("password")
	if rawPassword != "" {
		ctx.Password = &rawPassword
	}
	rawUser := c.Get("user")
	if rawUser != "" {
		ctx.User = rawUser
	}
	return &ctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *LoginAuthenticationContext) OK(resp []byte) error {
	ctx.Header().Set("Content-Type", "plain/text")
	return ctx.RespondBytes(200, resp)
}

// LogoutAuthenticationContext provides the authentication logout action context.
type LogoutAuthenticationContext struct {
	*goa.Context
	Pass     string
	Password *string
	User     string
}

// NewLogoutAuthenticationContext parses the incoming request URL and body, performs validations and creates the
// context used by the authentication controller logout action.
func NewLogoutAuthenticationContext(c *goa.Context) (*LogoutAuthenticationContext, error) {
	var err error
	ctx := LogoutAuthenticationContext{Context: c}
	rawPass := c.Get("pass")
	if rawPass != "" {
		ctx.Pass = rawPass
	}
	rawPassword := c.Get("password")
	if rawPassword != "" {
		ctx.Password = &rawPassword
	}
	rawUser := c.Get("user")
	if rawUser != "" {
		ctx.User = rawUser
	}
	return &ctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *LogoutAuthenticationContext) OK(resp []byte) error {
	ctx.Header().Set("Content-Type", "plain/text")
	return ctx.RespondBytes(200, resp)
}

// SignupAuthenticationContext provides the authentication signup action context.
type SignupAuthenticationContext struct {
	*goa.Context
	Pass     string
	Password *string
	User     string
}

// NewSignupAuthenticationContext parses the incoming request URL and body, performs validations and creates the
// context used by the authentication controller signup action.
func NewSignupAuthenticationContext(c *goa.Context) (*SignupAuthenticationContext, error) {
	var err error
	ctx := SignupAuthenticationContext{Context: c}
	rawPass := c.Get("pass")
	if rawPass != "" {
		ctx.Pass = rawPass
	}
	rawPassword := c.Get("password")
	if rawPassword != "" {
		ctx.Password = &rawPassword
	}
	rawUser := c.Get("user")
	if rawUser != "" {
		ctx.User = rawUser
	}
	return &ctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *SignupAuthenticationContext) OK(resp []byte) error {
	ctx.Header().Set("Content-Type", "plain/text")
	return ctx.RespondBytes(200, resp)
}

// AddOperandsContext provides the operands add action context.
type AddOperandsContext struct {
	*goa.Context
	Left  int
	Right int
}

// NewAddOperandsContext parses the incoming request URL and body, performs validations and creates the
// context used by the operands controller add action.
func NewAddOperandsContext(c *goa.Context) (*AddOperandsContext, error) {
	var err error
	ctx := AddOperandsContext{Context: c}
	rawLeft := c.Get("left")
	if rawLeft != "" {
		if left, err2 := strconv.Atoi(rawLeft); err2 == nil {
			ctx.Left = int(left)
		} else {
			err = goa.InvalidParamTypeError("left", rawLeft, "integer", err)
		}
	}
	rawRight := c.Get("right")
	if rawRight != "" {
		if right, err2 := strconv.Atoi(rawRight); err2 == nil {
			ctx.Right = int(right)
		} else {
			err = goa.InvalidParamTypeError("right", rawRight, "integer", err)
		}
	}
	return &ctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *AddOperandsContext) OK(resp []byte) error {
	ctx.Header().Set("Content-Type", "plain/text")
	return ctx.RespondBytes(200, resp)
}

// MultiplyOperandsContext provides the operands multiply action context.
type MultiplyOperandsContext struct {
	*goa.Context
	Left  int
	Right int
}

// NewMultiplyOperandsContext parses the incoming request URL and body, performs validations and creates the
// context used by the operands controller multiply action.
func NewMultiplyOperandsContext(c *goa.Context) (*MultiplyOperandsContext, error) {
	var err error
	ctx := MultiplyOperandsContext{Context: c}
	rawLeft := c.Get("left")
	if rawLeft != "" {
		if left, err2 := strconv.Atoi(rawLeft); err2 == nil {
			ctx.Left = int(left)
		} else {
			err = goa.InvalidParamTypeError("left", rawLeft, "integer", err)
		}
	}
	rawRight := c.Get("right")
	if rawRight != "" {
		if right, err2 := strconv.Atoi(rawRight); err2 == nil {
			ctx.Right = int(right)
		} else {
			err = goa.InvalidParamTypeError("right", rawRight, "integer", err)
		}
	}
	return &ctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *MultiplyOperandsContext) OK(resp []byte) error {
	ctx.Header().Set("Content-Type", "plain/text")
	return ctx.RespondBytes(200, resp)
}
