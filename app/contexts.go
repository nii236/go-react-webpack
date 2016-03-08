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
	"golang.org/x/net/context"
	"strconv"
)

// CallbackResponseFromGoogleAuthenticationContext provides the authentication Callback response from Google action context.
type CallbackResponseFromGoogleAuthenticationContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Password *string
	User     *string
}

// NewCallbackResponseFromGoogleAuthenticationContext parses the incoming request URL and body, performs validations and creates the
// context used by the authentication controller Callback response from Google action.
func NewCallbackResponseFromGoogleAuthenticationContext(ctx context.Context) (*CallbackResponseFromGoogleAuthenticationContext, error) {
	var err error
	req := goa.Request(ctx)
	rctx := CallbackResponseFromGoogleAuthenticationContext{Context: ctx, ResponseData: goa.Response(ctx), RequestData: req}
	rawPassword := req.Params.Get("password")
	if rawPassword != "" {
		rctx.Password = &rawPassword
	}
	rawUser := req.Params.Get("user")
	if rawUser != "" {
		rctx.User = &rawUser
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *CallbackResponseFromGoogleAuthenticationContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "plain/text")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// LogIntoGoogleAuthenticationContext provides the authentication Log into Google action context.
type LogIntoGoogleAuthenticationContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Password *string
	User     *string
}

// NewLogIntoGoogleAuthenticationContext parses the incoming request URL and body, performs validations and creates the
// context used by the authentication controller Log into Google action.
func NewLogIntoGoogleAuthenticationContext(ctx context.Context) (*LogIntoGoogleAuthenticationContext, error) {
	var err error
	req := goa.Request(ctx)
	rctx := LogIntoGoogleAuthenticationContext{Context: ctx, ResponseData: goa.Response(ctx), RequestData: req}
	rawPassword := req.Params.Get("password")
	if rawPassword != "" {
		rctx.Password = &rawPassword
	}
	rawUser := req.Params.Get("user")
	if rawUser != "" {
		rctx.User = &rawUser
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *LogIntoGoogleAuthenticationContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "plain/text")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// AddOperandsContext provides the operands add action context.
type AddOperandsContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Left  int
	Right int
}

// NewAddOperandsContext parses the incoming request URL and body, performs validations and creates the
// context used by the operands controller add action.
func NewAddOperandsContext(ctx context.Context) (*AddOperandsContext, error) {
	var err error
	req := goa.Request(ctx)
	rctx := AddOperandsContext{Context: ctx, ResponseData: goa.Response(ctx), RequestData: req}
	rawLeft := req.Params.Get("left")
	if rawLeft != "" {
		if left, err2 := strconv.Atoi(rawLeft); err2 == nil {
			rctx.Left = left
		} else {
			err = goa.InvalidParamTypeError("left", rawLeft, "integer", err)
		}
	}
	rawRight := req.Params.Get("right")
	if rawRight != "" {
		if right, err2 := strconv.Atoi(rawRight); err2 == nil {
			rctx.Right = right
		} else {
			err = goa.InvalidParamTypeError("right", rawRight, "integer", err)
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *AddOperandsContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "plain/text")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// MultiplyOperandsContext provides the operands multiply action context.
type MultiplyOperandsContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Left  int
	Right int
}

// NewMultiplyOperandsContext parses the incoming request URL and body, performs validations and creates the
// context used by the operands controller multiply action.
func NewMultiplyOperandsContext(ctx context.Context) (*MultiplyOperandsContext, error) {
	var err error
	req := goa.Request(ctx)
	rctx := MultiplyOperandsContext{Context: ctx, ResponseData: goa.Response(ctx), RequestData: req}
	rawLeft := req.Params.Get("left")
	if rawLeft != "" {
		if left, err2 := strconv.Atoi(rawLeft); err2 == nil {
			rctx.Left = left
		} else {
			err = goa.InvalidParamTypeError("left", rawLeft, "integer", err)
		}
	}
	rawRight := req.Params.Get("right")
	if rawRight != "" {
		if right, err2 := strconv.Atoi(rawRight); err2 == nil {
			rctx.Right = right
		} else {
			err = goa.InvalidParamTypeError("right", rawRight, "integer", err)
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *MultiplyOperandsContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "plain/text")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}
