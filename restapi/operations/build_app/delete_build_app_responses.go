// Code generated by go-swagger; DO NOT EDIT.

package build_app

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/laincloud/console-api/models"
)

// DeleteBuildAppNoContentCode is the HTTP code returned for type DeleteBuildAppNoContent
const DeleteBuildAppNoContentCode int = 204

/*DeleteBuildAppNoContent Deleted

swagger:response deleteBuildAppNoContent
*/
type DeleteBuildAppNoContent struct {
}

// NewDeleteBuildAppNoContent creates DeleteBuildAppNoContent with default headers values
func NewDeleteBuildAppNoContent() *DeleteBuildAppNoContent {

	return &DeleteBuildAppNoContent{}
}

// WriteResponse to the client
func (o *DeleteBuildAppNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

/*DeleteBuildAppDefault Failed

swagger:response deleteBuildAppDefault
*/
type DeleteBuildAppDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteBuildAppDefault creates DeleteBuildAppDefault with default headers values
func NewDeleteBuildAppDefault(code int) *DeleteBuildAppDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteBuildAppDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete build app default response
func (o *DeleteBuildAppDefault) WithStatusCode(code int) *DeleteBuildAppDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete build app default response
func (o *DeleteBuildAppDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete build app default response
func (o *DeleteBuildAppDefault) WithPayload(payload *models.Error) *DeleteBuildAppDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete build app default response
func (o *DeleteBuildAppDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteBuildAppDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
