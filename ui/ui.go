package ui

import "github.com/goadesign/goa"

// MountController mounts the JavaScript example controller under "/js".
func MountController(service *goa.Service) {
	// Serve static files under js
	service.ServeFiles("/ui/*filepath", "/Users/nii236/Go/src/github.com/nii236/go-react-webpack/ui/dist")
	goa.Info(goa.RootContext, "mount", goa.KV{"ctrl", "ui"}, goa.KV{"action", "ServeFiles"}, goa.KV{"route", "GET /ui/*"})
}
