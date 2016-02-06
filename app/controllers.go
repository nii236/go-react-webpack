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

import "github.com/goadesign/goa"

// AuthenticationController is the controller interface for the Authentication actions.
type AuthenticationController interface {
	goa.Controller
	Login(*LoginAuthenticationContext) error
	Logout(*LogoutAuthenticationContext) error
	Signup(*SignupAuthenticationContext) error
}

// MountAuthenticationController "mounts" a Authentication resource controller on the given service.
func MountAuthenticationController(service goa.Service, ctrl AuthenticationController) {
	// Setup encoders and decoders. This is idempotent and is done by each MountXXX function.
	service.SetEncoder(goa.GobEncoderFactory(), false, "application/gob", "application/x-gob")
	service.SetEncoder(goa.JSONEncoderFactory(), true, "application/json")
	service.SetEncoder(goa.XMLEncoderFactory(), false, "application/xml", "text/xml")
	service.SetDecoder(goa.GobDecoderFactory(), false, "application/gob", "application/x-gob")
	service.SetDecoder(goa.JSONDecoderFactory(), true, "application/json")
	service.SetDecoder(goa.XMLDecoderFactory(), false, "application/xml", "text/xml")

	// Setup endpoint handler
	var h goa.Handler
	mux := service.ServeMux()
	h = func(c *goa.Context) error {
		ctx, err := NewLoginAuthenticationContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Login(ctx)
	}
	mux.Handle("POST", "/login/:user/:pass", ctrl.HandleFunc("Login", h, nil))
	service.Info("mount", "ctrl", "Authentication", "action", "Login", "route", "POST /login/:user/:pass")
	h = func(c *goa.Context) error {
		ctx, err := NewLogoutAuthenticationContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Logout(ctx)
	}
	mux.Handle("POST", "/logout/:user/:pass", ctrl.HandleFunc("Logout", h, nil))
	service.Info("mount", "ctrl", "Authentication", "action", "Logout", "route", "POST /logout/:user/:pass")
	h = func(c *goa.Context) error {
		ctx, err := NewSignupAuthenticationContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Signup(ctx)
	}
	mux.Handle("POST", "/signup/:user/:pass", ctrl.HandleFunc("Signup", h, nil))
	service.Info("mount", "ctrl", "Authentication", "action", "Signup", "route", "POST /signup/:user/:pass")
}

// OperandsController is the controller interface for the Operands actions.
type OperandsController interface {
	goa.Controller
	Add(*AddOperandsContext) error
	Multiply(*MultiplyOperandsContext) error
}

// MountOperandsController "mounts" a Operands resource controller on the given service.
func MountOperandsController(service goa.Service, ctrl OperandsController) {
	// Setup encoders and decoders. This is idempotent and is done by each MountXXX function.
	service.SetEncoder(goa.GobEncoderFactory(), false, "application/gob", "application/x-gob")
	service.SetEncoder(goa.JSONEncoderFactory(), true, "application/json")
	service.SetEncoder(goa.XMLEncoderFactory(), false, "application/xml", "text/xml")
	service.SetDecoder(goa.GobDecoderFactory(), false, "application/gob", "application/x-gob")
	service.SetDecoder(goa.JSONDecoderFactory(), true, "application/json")
	service.SetDecoder(goa.XMLDecoderFactory(), false, "application/xml", "text/xml")

	// Setup endpoint handler
	var h goa.Handler
	mux := service.ServeMux()
	h = func(c *goa.Context) error {
		ctx, err := NewAddOperandsContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Add(ctx)
	}
	mux.Handle("GET", "/add/:left/:right", ctrl.HandleFunc("Add", h, nil))
	service.Info("mount", "ctrl", "Operands", "action", "Add", "route", "GET /add/:left/:right")
	h = func(c *goa.Context) error {
		ctx, err := NewMultiplyOperandsContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Multiply(ctx)
	}
	mux.Handle("GET", "/multiply/:left/:right", ctrl.HandleFunc("Multiply", h, nil))
	service.Info("mount", "ctrl", "Operands", "action", "Multiply", "route", "GET /multiply/:left/:right")
}
