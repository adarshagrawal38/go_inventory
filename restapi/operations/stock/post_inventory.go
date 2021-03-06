// Code generated by go-swagger; DO NOT EDIT.

package stock

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostInventoryHandlerFunc turns a function with the right signature into a post inventory handler
type PostInventoryHandlerFunc func(PostInventoryParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostInventoryHandlerFunc) Handle(params PostInventoryParams) middleware.Responder {
	return fn(params)
}

// PostInventoryHandler interface for that can handle valid post inventory params
type PostInventoryHandler interface {
	Handle(PostInventoryParams) middleware.Responder
}

// NewPostInventory creates a new http.Handler for the post inventory operation
func NewPostInventory(ctx *middleware.Context, handler PostInventoryHandler) *PostInventory {
	return &PostInventory{Context: ctx, Handler: handler}
}

/*PostInventory swagger:route POST /inventory Stock postInventory

saves an item to inventory

*/
type PostInventory struct {
	Context *middleware.Context
	Handler PostInventoryHandler
}

func (o *PostInventory) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostInventoryParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
