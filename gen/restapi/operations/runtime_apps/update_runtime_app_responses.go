// Code generated by go-swagger; DO NOT EDIT.

package runtime_apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/laincloud/console-api/gen/models"
)

// UpdateRuntimeAppOKCode is the HTTP code returned for type UpdateRuntimeAppOK
const UpdateRuntimeAppOKCode int = 200

/*UpdateRuntimeAppOK OK

swagger:response updateRuntimeAppOK
*/
type UpdateRuntimeAppOK struct {

	/*
	  In: Body
	*/
	Payload *models.RuntimeApp `json:"body,omitempty"`
}

// NewUpdateRuntimeAppOK creates UpdateRuntimeAppOK with default headers values
func NewUpdateRuntimeAppOK() *UpdateRuntimeAppOK {

	return &UpdateRuntimeAppOK{}
}

// WithPayload adds the payload to the update runtime app o k response
func (o *UpdateRuntimeAppOK) WithPayload(payload *models.RuntimeApp) *UpdateRuntimeAppOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update runtime app o k response
func (o *UpdateRuntimeAppOK) SetPayload(payload *models.RuntimeApp) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateRuntimeAppOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*UpdateRuntimeAppDefault Failed

swagger:response updateRuntimeAppDefault
*/
type UpdateRuntimeAppDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateRuntimeAppDefault creates UpdateRuntimeAppDefault with default headers values
func NewUpdateRuntimeAppDefault(code int) *UpdateRuntimeAppDefault {
	if code <= 0 {
		code = 500
	}

	return &UpdateRuntimeAppDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the update runtime app default response
func (o *UpdateRuntimeAppDefault) WithStatusCode(code int) *UpdateRuntimeAppDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the update runtime app default response
func (o *UpdateRuntimeAppDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the update runtime app default response
func (o *UpdateRuntimeAppDefault) WithPayload(payload *models.Error) *UpdateRuntimeAppDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update runtime app default response
func (o *UpdateRuntimeAppDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateRuntimeAppDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}