package main

import (
	"net/http"

	"github.com/nii236/goa-adder/client"
	"gopkg.in/alecthomas/kingpin.v2"
)

type (
	// AddOperandsCommand is the command line data structure for the add action of operands
	AddOperandsCommand struct {
		// Path is the HTTP request path.
		Path string
	}
)

// Run makes the HTTP request corresponding to the AddOperandsCommand command.
func (cmd *AddOperandsCommand) Run(c *client.Client) (*http.Response, error) {
	return c.AddOperands(cmd.Path)
}

// RegisterFlags registers the command flags with the command line.
func (cmd *AddOperandsCommand) RegisterFlags(cc *kingpin.CmdClause) {
	cc.Arg("path", `Request path, format is /add/:left/:right`).Required().StringVar(&cmd.Path)
}
