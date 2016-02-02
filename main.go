package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/middleware"
	"github.com/nii236/go-react-webpack/app"
	"github.com/nii236/go-react-webpack/js"
	"github.com/nii236/go-react-webpack/swagger"
)

func main() {
	// Create service
	service := goa.New("API")

	// Setup middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest())
	service.Use(middleware.Recover())

	// Mount "operands" controller
	c := NewOperandsController(service)
	app.MountOperandsController(service, c)
	// Mount Swagger spec provider controller
	swagger.MountController(service)
	js.MountController(service)

	// Start service, listen on port 8080
	service.ListenAndServe(":8080")
}
