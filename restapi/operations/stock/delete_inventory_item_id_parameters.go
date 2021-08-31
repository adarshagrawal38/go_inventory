// Code generated by go-swagger; DO NOT EDIT.

package stock

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewDeleteInventoryItemIDParams creates a new DeleteInventoryItemIDParams object
//
// There are no default values defined in the spec.
func NewDeleteInventoryItemIDParams() DeleteInventoryItemIDParams {

	return DeleteInventoryItemIDParams{}
}

// DeleteInventoryItemIDParams contains all the bound params for the delete inventory item ID operation
// typically these are obtained from a http.Request
//
// swagger:parameters DeleteInventoryItemID
type DeleteInventoryItemIDParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	ItemID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeleteInventoryItemIDParams() beforehand.
func (o *DeleteInventoryItemIDParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rItemID, rhkItemID, _ := route.Params.GetOK("itemId")
	if err := o.bindItemID(rItemID, rhkItemID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindItemID binds and validates parameter ItemID from path.
func (o *DeleteInventoryItemIDParams) bindItemID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("itemId", "path", "int64", raw)
	}
	o.ItemID = value

	return nil
}
