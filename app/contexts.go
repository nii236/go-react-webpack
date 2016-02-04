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
