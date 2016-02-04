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
