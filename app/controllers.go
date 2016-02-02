//************************************************************************//
// API "adder": Application Controllers
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/nii236/goa-adder
// --design=github.com/nii236/goa-adder/design
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import "github.com/goadesign/goa"

// OperandsController is the controller interface for the Operands actions.
type OperandsController interface {
	goa.Controller
	Add(*AddOperandsContext) error
}

// MountOperandsController "mounts" a Operands resource controller on the given service.
func MountOperandsController(service goa.Service, ctrl OperandsController) {
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
}
