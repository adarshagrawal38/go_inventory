// Code generated by go-swagger; DO NOT EDIT.

package stock

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetInventoryHandlerFunc turns a function with the right signature into a get inventory handler
type GetInventoryHandlerFunc func(GetInventoryParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetInventoryHandlerFunc) Handle(params GetInventoryParams) middleware.Responder {
	return fn(params)
}

// GetInventoryHandler interface for that can handle valid get inventory params
type GetInventoryHandler interface {
	Handle(GetInventoryParams) middleware.Responder
}

// NewGetInventory creates a new http.Handler for the get inventory operation
func NewGetInventory(ctx *middleware.Context, handler GetInventoryHandler) *GetInventory {
	return &GetInventory{Context: ctx, Handler: handler}
}

/*GetInventory swagger:route GET /inventory Stock getInventory

fetches all inventory data

*/
type GetInventory struct {
	Context *middleware.Context
	Handler GetInventoryHandler
}

func (o *GetInventory) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetInventoryParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
