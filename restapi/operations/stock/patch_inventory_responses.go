// Code generated by go-swagger; DO NOT EDIT.

package stock

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"inventory-management/models"
)

// PatchInventoryOKCode is the HTTP code returned for type PatchInventoryOK
const PatchInventoryOKCode int = 200

/*PatchInventoryOK Successfully updated the record

swagger:response patchInventoryOK
*/
type PatchInventoryOK struct {

	/*
	  In: Body
	*/
	Payload *models.Item `json:"body,omitempty"`
}

// NewPatchInventoryOK creates PatchInventoryOK with default headers values
func NewPatchInventoryOK() *PatchInventoryOK {

	return &PatchInventoryOK{}
}

// WithPayload adds the payload to the patch inventory o k response
func (o *PatchInventoryOK) WithPayload(payload *models.Item) *PatchInventoryOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch inventory o k response
func (o *PatchInventoryOK) SetPayload(payload *models.Item) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchInventoryOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}