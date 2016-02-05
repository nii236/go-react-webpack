package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/middleware"
	"github.com/goadesign/middleware/cors"
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
	cspec, err := cors.New(func() {
		cors.Origin("*", func() {
			cors.Resource("/*", func() {
				cors.Headers("Accept", "Content-Type", "Origin", "Authorization")
				cors.Methods("GET", "POST", "PUT", "DELETE", "OPTIONS")
				cors.MaxAge(600)
				cors.Credentials(true)
				cors.Vary("Http-Origin")
			})
		})
	})
	if err != nil {
		panic(err)
	}
	// mount the cors controller
	service.Use(cors.Middleware(cspec))
	cors.MountPreflightController(service, cspec)

	// Mount "operands" controller
	c := NewOperandsController(service)
	app.MountOperandsController(service, c)
	// Mount Swagger spec provider controller
	swagger.MountController(service)
	js.MountController(service)

	// Start service, listen on port 8080
	service.ListenAndServe(":8080")
}
