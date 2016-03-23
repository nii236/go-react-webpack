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

// CallbackResponseFromFacebookAuthenticationContext provides the authentication Callback response from Facebook action context.
type CallbackResponseFromFacebookAuthenticationContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewCallbackResponseFromFacebookAuthenticationContext parses the incoming request URL and body, performs validations and creates the
// context used by the authentication controller Callback response from Facebook action.
func NewCallbackResponseFromFacebookAuthenticationContext(ctx context.Context) (*CallbackResponseFromFacebookAuthenticationContext, error) {
	var err error
	req := goa.Request(ctx)
	rctx := CallbackResponseFromFacebookAuthenticationContext{Context: ctx, ResponseData: goa.Response(ctx), RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *CallbackResponseFromFacebookAuthenticationContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "plain/text")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// CallbackResponseFromGithubAuthenticationContext provides the authentication Callback response from Github action context.
type CallbackResponseFromGithubAuthenticationContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewCallbackResponseFromGithubAuthenticationContext parses the incoming request URL and body, performs validations and creates the
// context used by the authentication controller Callback response from Github action.
func NewCallbackResponseFromGithubAuthenticationContext(ctx context.Context) (*CallbackResponseFromGithubAuthenticationContext, error) {
	var err error
	req := goa.Request(ctx)
	rctx := CallbackResponseFromGithubAuthenticationContext{Context: ctx, ResponseData: goa.Response(ctx), RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *CallbackResponseFromGithubAuthenticationContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "plain/text")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// CallbackResponseFromGoogleAuthenticationContext provides the authentication Callback response from Google action context.
type CallbackResponseFromGoogleAuthenticationContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewCallbackResponseFromGoogleAuthenticationContext parses the incoming request URL and body, performs validations and creates the
// context used by the authentication controller Callback response from Google action.
func NewCallbackResponseFromGoogleAuthenticationContext(ctx context.Context) (*CallbackResponseFromGoogleAuthenticationContext, error) {
	var err error
	req := goa.Request(ctx)
	rctx := CallbackResponseFromGoogleAuthenticationContext{Context: ctx, ResponseData: goa.Response(ctx), RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *CallbackResponseFromGoogleAuthenticationContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "plain/text")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// LogIntoFacebookAuthenticationContext provides the authentication Log into Facebook action context.
type LogIntoFacebookAuthenticationContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewLogIntoFacebookAuthenticationContext parses the incoming request URL and body, performs validations and creates the
// context used by the authentication controller Log into Facebook action.
func NewLogIntoFacebookAuthenticationContext(ctx context.Context) (*LogIntoFacebookAuthenticationContext, error) {
	var err error
	req := goa.Request(ctx)
	rctx := LogIntoFacebookAuthenticationContext{Context: ctx, ResponseData: goa.Response(ctx), RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *LogIntoFacebookAuthenticationContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "plain/text")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// LogIntoGithubAuthenticationContext provides the authentication Log into Github action context.
type LogIntoGithubAuthenticationContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewLogIntoGithubAuthenticationContext parses the incoming request URL and body, performs validations and creates the
// context used by the authentication controller Log into Github action.
func NewLogIntoGithubAuthenticationContext(ctx context.Context) (*LogIntoGithubAuthenticationContext, error) {
	var err error
	req := goa.Request(ctx)
	rctx := LogIntoGithubAuthenticationContext{Context: ctx, ResponseData: goa.Response(ctx), RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *LogIntoGithubAuthenticationContext) OK(resp []byte) error {
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
}

// NewLogIntoGoogleAuthenticationContext parses the incoming request URL and body, performs validations and creates the
// context used by the authentication controller Log into Google action.
func NewLogIntoGoogleAuthenticationContext(ctx context.Context) (*LogIntoGoogleAuthenticationContext, error) {
	var err error
	req := goa.Request(ctx)
	rctx := LogIntoGoogleAuthenticationContext{Context: ctx, ResponseData: goa.Response(ctx), RequestData: req}
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
