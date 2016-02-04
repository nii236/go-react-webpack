package main

import (
	"strconv"

	"github.com/goadesign/goa"
	"github.com/nii236/go-react-webpack/app"
)

// OperandsController implements theoperands resource.
type OperandsController struct {
	goa.Controller
}

// NewOperandsController creates a operands controller.
func NewOperandsController(service goa.Service) app.OperandsController {
	return &OperandsController{Controller: service.NewController("operands")}
}

// Add runs the add action.
func (c *OperandsController) Add(ctx *app.AddOperandsContext) error {
	sum := ctx.Left + ctx.Right
	return ctx.OK([]byte(strconv.Itoa(sum)))
}

func (c *OperandsController) Multiply(ctx *app.MultiplyOperandsContext) error {
	multiply := ctx.Left * ctx.Right
	return ctx.OK([]byte(strconv.Itoa(multiply)))
}
