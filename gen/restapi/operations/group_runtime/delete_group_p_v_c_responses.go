// Code generated by go-swagger; DO NOT EDIT.

package group_runtime

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/laincloud/console-api/gen/models"
)

// DeleteGroupPVCNoContentCode is the HTTP code returned for type DeleteGroupPVCNoContent
const DeleteGroupPVCNoContentCode int = 204

/*DeleteGroupPVCNoContent Deleted

swagger:response deleteGroupPVCNoContent
*/
type DeleteGroupPVCNoContent struct {
}

// NewDeleteGroupPVCNoContent creates DeleteGroupPVCNoContent with default headers values
func NewDeleteGroupPVCNoContent() *DeleteGroupPVCNoContent {

	return &DeleteGroupPVCNoContent{}
}

// WriteResponse to the client
func (o *DeleteGroupPVCNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

/*DeleteGroupPVCDefault Failed

swagger:response deleteGroupPVCDefault
*/
type DeleteGroupPVCDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteGroupPVCDefault creates DeleteGroupPVCDefault with default headers values
func NewDeleteGroupPVCDefault(code int) *DeleteGroupPVCDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteGroupPVCDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete group p v c default response
func (o *DeleteGroupPVCDefault) WithStatusCode(code int) *DeleteGroupPVCDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete group p v c default response
func (o *DeleteGroupPVCDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete group p v c default response
func (o *DeleteGroupPVCDefault) WithPayload(payload *models.Error) *DeleteGroupPVCDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete group p v c default response
func (o *DeleteGroupPVCDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteGroupPVCDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
