// Code generated by go-swagger; DO NOT EDIT.

package runtime_apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/laincloud/console-api/gen/models"
)

// CreateRuntimeAppCreatedCode is the HTTP code returned for type CreateRuntimeAppCreated
const CreateRuntimeAppCreatedCode int = 201

/*CreateRuntimeAppCreated Created

swagger:response createRuntimeAppCreated
*/
type CreateRuntimeAppCreated struct {

	/*
	  In: Body
	*/
	Payload *models.RuntimeApp `json:"body,omitempty"`
}

// NewCreateRuntimeAppCreated creates CreateRuntimeAppCreated with default headers values
func NewCreateRuntimeAppCreated() *CreateRuntimeAppCreated {

	return &CreateRuntimeAppCreated{}
}

// WithPayload adds the payload to the create runtime app created response
func (o *CreateRuntimeAppCreated) WithPayload(payload *models.RuntimeApp) *CreateRuntimeAppCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create runtime app created response
func (o *CreateRuntimeAppCreated) SetPayload(payload *models.RuntimeApp) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateRuntimeAppCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*CreateRuntimeAppDefault Failed

swagger:response createRuntimeAppDefault
*/
type CreateRuntimeAppDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateRuntimeAppDefault creates CreateRuntimeAppDefault with default headers values
func NewCreateRuntimeAppDefault(code int) *CreateRuntimeAppDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateRuntimeAppDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create runtime app default response
func (o *CreateRuntimeAppDefault) WithStatusCode(code int) *CreateRuntimeAppDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create runtime app default response
func (o *CreateRuntimeAppDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the create runtime app default response
func (o *CreateRuntimeAppDefault) WithPayload(payload *models.Error) *CreateRuntimeAppDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create runtime app default response
func (o *CreateRuntimeAppDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateRuntimeAppDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
