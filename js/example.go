//************************************************************************//
// adder JavaScript Client Example
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/nii236/go-react-webpack
// --design=github.com/nii236/go-react-webpack/design
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package js

import "github.com/goadesign/goa"

// MountController mounts the JavaScript example controller under "/js".
func MountController(service *goa.Service) {
	// Serve static files under js
	service.ServeFiles("/js/*filepath", "/Users/nii236/Go/src/github.com/nii236/go-react-webpack/js")
	goa.Info(goa.RootContext, "mount", goa.KV{"ctrl", "JS"}, goa.KV{"action", "ServeFiles"}, goa.KV{"route", "GET /js/*"})
}
