// Code generated by go-swagger; DO NOT EDIT.

package runtime_app

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/laincloud/console-api/gen/models"
)

// UpdateRuntimeAppSecretOKCode is the HTTP code returned for type UpdateRuntimeAppSecretOK
const UpdateRuntimeAppSecretOKCode int = 200

/*UpdateRuntimeAppSecretOK OK

swagger:response updateRuntimeAppSecretOK
*/
type UpdateRuntimeAppSecretOK struct {

	/*
	  In: Body
	*/
	Payload *models.Secret `json:"body,omitempty"`
}

// NewUpdateRuntimeAppSecretOK creates UpdateRuntimeAppSecretOK with default headers values
func NewUpdateRuntimeAppSecretOK() *UpdateRuntimeAppSecretOK {

	return &UpdateRuntimeAppSecretOK{}
}

// WithPayload adds the payload to the update runtime app secret o k response
func (o *UpdateRuntimeAppSecretOK) WithPayload(payload *models.Secret) *UpdateRuntimeAppSecretOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update runtime app secret o k response
func (o *UpdateRuntimeAppSecretOK) SetPayload(payload *models.Secret) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateRuntimeAppSecretOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*UpdateRuntimeAppSecretDefault Failed

swagger:response updateRuntimeAppSecretDefault
*/
type UpdateRuntimeAppSecretDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateRuntimeAppSecretDefault creates UpdateRuntimeAppSecretDefault with default headers values
func NewUpdateRuntimeAppSecretDefault(code int) *UpdateRuntimeAppSecretDefault {
	if code <= 0 {
		code = 500
	}

	return &UpdateRuntimeAppSecretDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the update runtime app secret default response
func (o *UpdateRuntimeAppSecretDefault) WithStatusCode(code int) *UpdateRuntimeAppSecretDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the update runtime app secret default response
func (o *UpdateRuntimeAppSecretDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the update runtime app secret default response
func (o *UpdateRuntimeAppSecretDefault) WithPayload(payload *models.Error) *UpdateRuntimeAppSecretDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update runtime app secret default response
func (o *UpdateRuntimeAppSecretDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateRuntimeAppSecretDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
