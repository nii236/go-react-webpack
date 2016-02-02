//************************************************************************//
// adder Swagger Spec
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/nii236/goa-adder
// --design=github.com/nii236/goa-adder/design
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package swagger

import "github.com/goadesign/goa"

// MountController mounts the swagger spec controller under "/swagger.json".
func MountController(service goa.Service) {
	service.ServeFiles("/swagger.json", "swagger/swagger.json")
}
