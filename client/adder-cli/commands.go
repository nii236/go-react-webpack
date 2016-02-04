package main

import (
	"fmt"
	"github.com/nii236/go-react-webpack/client"
	"github.com/spf13/cobra"
)

type (
	// AddOperandsCommand is the command line data structure for the add action of operands
	AddOperandsCommand struct {
	}

	// MultiplyOperandsCommand is the command line data structure for the multiply action of operands
	MultiplyOperandsCommand struct {
	}
)

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
