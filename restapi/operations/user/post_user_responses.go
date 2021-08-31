// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostUserCreatedCode is the HTTP code returned for type PostUserCreated
const PostUserCreatedCode int = 201

/*PostUserCreated Item saved successfully

swagger:response postUserCreated
*/
type PostUserCreated struct {
}

// NewPostUserCreated creates PostUserCreated with default headers values
func NewPostUserCreated() *PostUserCreated {

	return &PostUserCreated{}
}

// WriteResponse to the client
func (o *PostUserCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}
