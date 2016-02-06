package main

import (
	"github.com/goadesign/goa"
	"github.com/nii236/go-react-webpack/app"
)

// AuthenticationController implements theauthentication resource.
type AuthenticationController struct {
	goa.Controller
}

// NewAuthenticationController creates a authentication controller.
func NewAuthenticationController(service goa.Service) app.AuthenticationController {
	return &AuthenticationController{Controller: service.NewController("authentication")}
}

// Login runs the login action.
func (c *AuthenticationController) Login(ctx *app.LoginAuthenticationContext) error {
	return nil
}

// Logout runs the logout action.
func (c *AuthenticationController) Logout(ctx *app.LogoutAuthenticationContext) error {
	return nil
}

// Signup runs the signup action.
func (c *AuthenticationController) Signup(ctx *app.SignupAuthenticationContext) error {
	return nil
}
