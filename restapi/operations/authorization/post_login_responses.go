// Code generated by go-swagger; DO NOT EDIT.

package authorization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostLoginOKCode is the HTTP code returned for type PostLoginOK
const PostLoginOKCode int = 200

/*PostLoginOK login sucessful

swagger:response postLoginOK
*/
type PostLoginOK struct {
}

// NewPostLoginOK creates PostLoginOK with default headers values
func NewPostLoginOK() *PostLoginOK {

	return &PostLoginOK{}
}

// WriteResponse to the client
func (o *PostLoginOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// PostLoginUnauthorizedCode is the HTTP code returned for type PostLoginUnauthorized
const PostLoginUnauthorizedCode int = 401

/*PostLoginUnauthorized invalid credentials

swagger:response postLoginUnauthorized
*/
type PostLoginUnauthorized struct {
}

// NewPostLoginUnauthorized creates PostLoginUnauthorized with default headers values
func NewPostLoginUnauthorized() *PostLoginUnauthorized {

	return &PostLoginUnauthorized{}
}

// WriteResponse to the client
func (o *PostLoginUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
