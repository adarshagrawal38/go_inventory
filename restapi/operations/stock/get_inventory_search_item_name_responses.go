// Code generated by go-swagger; DO NOT EDIT.

package stock

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"inventory-management/models"
)

// GetInventorySearchItemNameOKCode is the HTTP code returned for type GetInventorySearchItemNameOK
const GetInventorySearchItemNameOKCode int = 200

/*GetInventorySearchItemNameOK Item fetched

swagger:response getInventorySearchItemNameOK
*/
type GetInventorySearchItemNameOK struct {

	/*
	  In: Body
	*/
	Payload models.Items `json:"body,omitempty"`
}

// NewGetInventorySearchItemNameOK creates GetInventorySearchItemNameOK with default headers values
func NewGetInventorySearchItemNameOK() *GetInventorySearchItemNameOK {

	return &GetInventorySearchItemNameOK{}
}

// WithPayload adds the payload to the get inventory search item name o k response
func (o *GetInventorySearchItemNameOK) WithPayload(payload models.Items) *GetInventorySearchItemNameOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get inventory search item name o k response
func (o *GetInventorySearchItemNameOK) SetPayload(payload models.Items) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetInventorySearchItemNameOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.Items{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
