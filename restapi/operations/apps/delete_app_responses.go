// Code generated by go-swagger; DO NOT EDIT.

package apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/laincloud/console-api/models"
)

// DeleteAppNoContentCode is the HTTP code returned for type DeleteAppNoContent
const DeleteAppNoContentCode int = 204

/*DeleteAppNoContent Deleted

swagger:response deleteAppNoContent
*/
type DeleteAppNoContent struct {
}

// NewDeleteAppNoContent creates DeleteAppNoContent with default headers values
func NewDeleteAppNoContent() *DeleteAppNoContent {

	return &DeleteAppNoContent{}
}

// WriteResponse to the client
func (o *DeleteAppNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

/*DeleteAppDefault Failed

swagger:response deleteAppDefault
*/
type DeleteAppDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteAppDefault creates DeleteAppDefault with default headers values
func NewDeleteAppDefault(code int) *DeleteAppDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteAppDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete app default response
func (o *DeleteAppDefault) WithStatusCode(code int) *DeleteAppDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete app default response
func (o *DeleteAppDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete app default response
func (o *DeleteAppDefault) WithPayload(payload *models.Error) *DeleteAppDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete app default response
func (o *DeleteAppDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteAppDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}