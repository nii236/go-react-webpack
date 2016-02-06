package main

import (
	"fmt"
	"github.com/nii236/go-react-webpack/client"
	"github.com/spf13/cobra"
)

type (
	// LoginAuthenticationCommand is the command line data structure for the login action of authentication
	LoginAuthenticationCommand struct {
		// Password
		Password string
	}

	// LogoutAuthenticationCommand is the command line data structure for the logout action of authentication
	LogoutAuthenticationCommand struct {
		// Password
		Password string
	}

	// SignupAuthenticationCommand is the command line data structure for the signup action of authentication
	SignupAuthenticationCommand struct {
		// Password
		Password string
	}

	// AddOperandsCommand is the command line data structure for the add action of operands
	AddOperandsCommand struct {
	}

	// MultiplyOperandsCommand is the command line data structure for the multiply action of operands
	MultiplyOperandsCommand struct {
	}
)

// Run makes the HTTP request corresponding to the LoginAuthenticationCommand command.
func (cmd *LoginAuthenticationCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		return fmt.Errorf("missing path argument")
	}
	resp, err := c.LoginAuthentication(path, cmd.Password)
	if err != nil {
		return err
	}
	HandleResponse(c, resp)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *LoginAuthenticationCommand) RegisterFlags(cc *cobra.Command) {
	var tmp6 string
	cc.Flags().StringVar(&cmd.Password, "password", tmp6, "Password")
}

// Run makes the HTTP request corresponding to the LogoutAuthenticationCommand command.
func (cmd *LogoutAuthenticationCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		return fmt.Errorf("missing path argument")
	}
	resp, err := c.LogoutAuthentication(path, cmd.Password)
	if err != nil {
		return err
	}
	HandleResponse(c, resp)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *LogoutAuthenticationCommand) RegisterFlags(cc *cobra.Command) {
	var tmp7 string
	cc.Flags().StringVar(&cmd.Password, "password", tmp7, "Password")
}

// Run makes the HTTP request corresponding to the SignupAuthenticationCommand command.
func (cmd *SignupAuthenticationCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		return fmt.Errorf("missing path argument")
	}
	resp, err := c.SignupAuthentication(path, cmd.Password)
	if err != nil {
		return err
	}
	HandleResponse(c, resp)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *SignupAuthenticationCommand) RegisterFlags(cc *cobra.Command) {
	var tmp8 string
	cc.Flags().StringVar(&cmd.Password, "password", tmp8, "Password")
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
