//************************************************************************//
// API "adder": Application Controllers
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
	"net/http"

	"github.com/goadesign/goa"
	"golang.org/x/net/context"
)

// inited is true if initService has been called
var inited = false

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	if inited {
		return
	}
	inited = true

	// Setup encoders and decoders
	service.Encoder(goa.NewJSONEncoder, "application/json")
	service.Decoder(goa.NewJSONDecoder, "application/json")
	service.Decoder(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder(goa.NewJSONEncoder, "*/*")
	service.Decoder(goa.NewJSONDecoder, "*/*")
}

// AuthenticationController is the controller interface for the Authentication actions.
type AuthenticationController interface {
	goa.Muxer
	CallbackResponseFromFacebook(*CallbackResponseFromFacebookAuthenticationContext) error
	CallbackResponseFromGithub(*CallbackResponseFromGithubAuthenticationContext) error
	CallbackResponseFromGoogle(*CallbackResponseFromGoogleAuthenticationContext) error
	LogIntoFacebook(*LogIntoFacebookAuthenticationContext) error
	LogIntoGithub(*LogIntoGithubAuthenticationContext) error
	LogIntoGoogle(*LogIntoGoogleAuthenticationContext) error
}

// MountAuthenticationController "mounts" a Authentication resource controller on the given service.
func MountAuthenticationController(service *goa.Service, ctrl AuthenticationController) {
	initService(service)
	var h goa.Handler
	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewCallbackResponseFromFacebookAuthenticationContext(ctx)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.CallbackResponseFromFacebook(rctx)
	}
	service.Mux.Handle("GET", "/FacebookCallback", ctrl.MuxHandler("CallbackResponseFromFacebook", h, nil))
	goa.Info(goa.RootContext, "mount", goa.KV{"ctrl", "Authentication"}, goa.KV{"action", "CallbackResponseFromFacebook"}, goa.KV{"route", "GET /FacebookCallback"})
	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewCallbackResponseFromGithubAuthenticationContext(ctx)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.CallbackResponseFromGithub(rctx)
	}
	service.Mux.Handle("GET", "/GithubCallback", ctrl.MuxHandler("CallbackResponseFromGithub", h, nil))
	goa.Info(goa.RootContext, "mount", goa.KV{"ctrl", "Authentication"}, goa.KV{"action", "CallbackResponseFromGithub"}, goa.KV{"route", "GET /GithubCallback"})
	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewCallbackResponseFromGoogleAuthenticationContext(ctx)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.CallbackResponseFromGoogle(rctx)
	}
	service.Mux.Handle("GET", "/GoogleCallback", ctrl.MuxHandler("CallbackResponseFromGoogle", h, nil))
	goa.Info(goa.RootContext, "mount", goa.KV{"ctrl", "Authentication"}, goa.KV{"action", "CallbackResponseFromGoogle"}, goa.KV{"route", "GET /GoogleCallback"})
	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewLogIntoFacebookAuthenticationContext(ctx)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.LogIntoFacebook(rctx)
	}
	service.Mux.Handle("GET", "/FacebookLogin", ctrl.MuxHandler("LogIntoFacebook", h, nil))
	goa.Info(goa.RootContext, "mount", goa.KV{"ctrl", "Authentication"}, goa.KV{"action", "LogIntoFacebook"}, goa.KV{"route", "GET /FacebookLogin"})
	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewLogIntoGithubAuthenticationContext(ctx)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.LogIntoGithub(rctx)
	}
	service.Mux.Handle("GET", "/GithubLogin", ctrl.MuxHandler("LogIntoGithub", h, nil))
	goa.Info(goa.RootContext, "mount", goa.KV{"ctrl", "Authentication"}, goa.KV{"action", "LogIntoGithub"}, goa.KV{"route", "GET /GithubLogin"})
	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewLogIntoGoogleAuthenticationContext(ctx)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.LogIntoGoogle(rctx)
	}
	service.Mux.Handle("GET", "/GoogleLogin", ctrl.MuxHandler("LogIntoGoogle", h, nil))
	goa.Info(goa.RootContext, "mount", goa.KV{"ctrl", "Authentication"}, goa.KV{"action", "LogIntoGoogle"}, goa.KV{"route", "GET /GoogleLogin"})
}

// OperandsController is the controller interface for the Operands actions.
type OperandsController interface {
	goa.Muxer
	Add(*AddOperandsContext) error
	Multiply(*MultiplyOperandsContext) error
}

// MountOperandsController "mounts" a Operands resource controller on the given service.
func MountOperandsController(service *goa.Service, ctrl OperandsController) {
	initService(service)
	var h goa.Handler
	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewAddOperandsContext(ctx)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Add(rctx)
	}
	service.Mux.Handle("GET", "/add/:left/:right", ctrl.MuxHandler("Add", h, nil))
	goa.Info(goa.RootContext, "mount", goa.KV{"ctrl", "Operands"}, goa.KV{"action", "Add"}, goa.KV{"route", "GET /add/:left/:right"})
	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewMultiplyOperandsContext(ctx)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Multiply(rctx)
	}
	service.Mux.Handle("GET", "/multiply/:left/:right", ctrl.MuxHandler("Multiply", h, nil))
	goa.Info(goa.RootContext, "mount", goa.KV{"ctrl", "Operands"}, goa.KV{"action", "Multiply"}, goa.KV{"route", "GET /multiply/:left/:right"})
}
