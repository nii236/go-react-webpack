package main

import (
	"fmt"

	"github.com/goadesign/goa"
	"github.com/nii236/go-react-webpack/app"
)

// AuthenticationController implements the authentication resource.
type AuthenticationController struct {
	*goa.Controller
}

// NewAuthenticationController creates a authentication controller.
func NewAuthenticationController(service *goa.Service) *AuthenticationController {
	return &AuthenticationController{Controller: service.NewController("authentication")}
}

// CallbackResponseFromGoogle runs the Callback response from Google action.
func (c *AuthenticationController) CallbackResponseFromGoogle(ctx *app.CallbackResponseFromGoogleAuthenticationContext) error {
	return nil
}

// LogIntoGoogle runs the Log into Google action.
func (c *AuthenticationController) LogIntoGoogle(ctx *app.LogIntoGoogleAuthenticationContext) error {
	goa.Info(ctx, "HELLO!", goa.KV{"Key and a: ", "Value"})
	r := goa.Response(ctx)
	fmt.Fprintf(r, "TESTIN")
	return nil
}
