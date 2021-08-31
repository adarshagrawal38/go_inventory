// Code generated by go-swagger; DO NOT EDIT.

package stock

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"inventory-management/models"
)

// PutInventoryItemIDOKCode is the HTTP code returned for type PutInventoryItemIDOK
const PutInventoryItemIDOKCode int = 200

/*PutInventoryItemIDOK Successfully updated the record

swagger:response putInventoryItemIdOK
*/
type PutInventoryItemIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Item `json:"body,omitempty"`
}

// NewPutInventoryItemIDOK creates PutInventoryItemIDOK with default headers values
func NewPutInventoryItemIDOK() *PutInventoryItemIDOK {

	return &PutInventoryItemIDOK{}
}

// WithPayload adds the payload to the put inventory item Id o k response
func (o *PutInventoryItemIDOK) WithPayload(payload *models.Item) *PutInventoryItemIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put inventory item Id o k response
func (o *PutInventoryItemIDOK) SetPayload(payload *models.Item) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutInventoryItemIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}