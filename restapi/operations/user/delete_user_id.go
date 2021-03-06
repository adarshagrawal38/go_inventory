// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteUserIDHandlerFunc turns a function with the right signature into a delete user ID handler
type DeleteUserIDHandlerFunc func(DeleteUserIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteUserIDHandlerFunc) Handle(params DeleteUserIDParams) middleware.Responder {
	return fn(params)
}

// DeleteUserIDHandler interface for that can handle valid delete user ID params
type DeleteUserIDHandler interface {
	Handle(DeleteUserIDParams) middleware.Responder
}

// NewDeleteUserID creates a new http.Handler for the delete user ID operation
func NewDeleteUserID(ctx *middleware.Context, handler DeleteUserIDHandler) *DeleteUserID {
	return &DeleteUserID{Context: ctx, Handler: handler}
}

/*DeleteUserID swagger:route DELETE /user/{id} User deleteUserId

DeleteUserID delete user ID API

*/
type DeleteUserID struct {
	Context *middleware.Context
	Handler DeleteUserIDHandler
}

func (o *DeleteUserID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteUserIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
