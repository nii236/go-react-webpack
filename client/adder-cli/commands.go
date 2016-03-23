package main

import (
	"fmt"
	"github.com/nii236/go-react-webpack/client"
	"github.com/spf13/cobra"
)

type (
	// CallbackResponseFromFacebookAuthenticationCommand is the command line data structure for the Callback response from Facebook action of authentication
	CallbackResponseFromFacebookAuthenticationCommand struct {
	}

	// CallbackResponseFromGithubAuthenticationCommand is the command line data structure for the Callback response from Github action of authentication
	CallbackResponseFromGithubAuthenticationCommand struct {
	}

	// CallbackResponseFromGoogleAuthenticationCommand is the command line data structure for the Callback response from Google action of authentication
	CallbackResponseFromGoogleAuthenticationCommand struct {
	}

	// LogIntoFacebookAuthenticationCommand is the command line data structure for the Log into Facebook action of authentication
	LogIntoFacebookAuthenticationCommand struct {
	}

	// LogIntoGithubAuthenticationCommand is the command line data structure for the Log into Github action of authentication
	LogIntoGithubAuthenticationCommand struct {
	}

	// LogIntoGoogleAuthenticationCommand is the command line data structure for the Log into Google action of authentication
	LogIntoGoogleAuthenticationCommand struct {
	}

	// AddOperandsCommand is the command line data structure for the add action of operands
	AddOperandsCommand struct {
	}

	// MultiplyOperandsCommand is the command line data structure for the multiply action of operands
	MultiplyOperandsCommand struct {
	}
)

// Run makes the HTTP request corresponding to the CallbackResponseFromFacebookAuthenticationCommand command.
func (cmd *CallbackResponseFromFacebookAuthenticationCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/FacebookCallback"
	}
	resp, err := c.CallbackResponseFromFacebookAuthentication(path)
	if err != nil {
		return err
	}
	HandleResponse(c, resp)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CallbackResponseFromFacebookAuthenticationCommand) RegisterFlags(cc *cobra.Command) {
}

// Run makes the HTTP request corresponding to the CallbackResponseFromGithubAuthenticationCommand command.
func (cmd *CallbackResponseFromGithubAuthenticationCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/GithubCallback"
	}
	resp, err := c.CallbackResponseFromGithubAuthentication(path)
	if err != nil {
		return err
	}
	HandleResponse(c, resp)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CallbackResponseFromGithubAuthenticationCommand) RegisterFlags(cc *cobra.Command) {
}

// Run makes the HTTP request corresponding to the CallbackResponseFromGoogleAuthenticationCommand command.
func (cmd *CallbackResponseFromGoogleAuthenticationCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/GoogleCallback"
	}
	resp, err := c.CallbackResponseFromGoogleAuthentication(path)
	if err != nil {
		return err
	}
	HandleResponse(c, resp)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CallbackResponseFromGoogleAuthenticationCommand) RegisterFlags(cc *cobra.Command) {
}

// Run makes the HTTP request corresponding to the LogIntoFacebookAuthenticationCommand command.
func (cmd *LogIntoFacebookAuthenticationCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/FacebookLogin"
	}
	resp, err := c.LogIntoFacebookAuthentication(path)
	if err != nil {
		return err
	}
	HandleResponse(c, resp)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *LogIntoFacebookAuthenticationCommand) RegisterFlags(cc *cobra.Command) {
}

// Run makes the HTTP request corresponding to the LogIntoGithubAuthenticationCommand command.
func (cmd *LogIntoGithubAuthenticationCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/GithubLogin"
	}
	resp, err := c.LogIntoGithubAuthentication(path)
	if err != nil {
		return err
	}
	HandleResponse(c, resp)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *LogIntoGithubAuthenticationCommand) RegisterFlags(cc *cobra.Command) {
}

// Run makes the HTTP request corresponding to the LogIntoGoogleAuthenticationCommand command.
func (cmd *LogIntoGoogleAuthenticationCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/GoogleLogin"
	}
	resp, err := c.LogIntoGoogleAuthentication(path)
	if err != nil {
		return err
	}
	HandleResponse(c, resp)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *LogIntoGoogleAuthenticationCommand) RegisterFlags(cc *cobra.Command) {
}

// Run makes the HTTP request corresponding to the AddOperandsCommand command.
func (cmd *AddOperandsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		return fmt.Errorf("missing path argument")
	}
	resp, err := c.AddOperands(path)
	if err != nil {
		return err
	}
	HandleResponse(c, resp)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *AddOperandsCommand) RegisterFlags(cc *cobra.Command) {
}

// Run makes the HTTP request corresponding to the MultiplyOperandsCommand command.
func (cmd *MultiplyOperandsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		return fmt.Errorf("missing path argument")
	}
	resp, err := c.MultiplyOperands(path)
	if err != nil {
		return err
	}
	HandleResponse(c, resp)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *MultiplyOperandsCommand) RegisterFlags(cc *cobra.Command) {
}
