// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"inventory-management/models"
)

// GetUserIDOKCode is the HTTP code returned for type GetUserIDOK
const GetUserIDOKCode int = 200

/*GetUserIDOK get user details

swagger:response getUserIdOK
*/
type GetUserIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Users `json:"body,omitempty"`
}

// NewGetUserIDOK creates GetUserIDOK with default headers values
func NewGetUserIDOK() *GetUserIDOK {

	return &GetUserIDOK{}
}

// WithPayload adds the payload to the get user Id o k response
func (o *GetUserIDOK) WithPayload(payload *models.Users) *GetUserIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user Id o k response
func (o *GetUserIDOK) SetPayload(payload *models.Users) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
