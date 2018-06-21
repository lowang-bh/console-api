// Code generated by go-swagger; DO NOT EDIT.

package runtime_app_secrets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/laincloud/console-api/models"
)

// CreateRuntimeAppSecretOKCode is the HTTP code returned for type CreateRuntimeAppSecretOK
const CreateRuntimeAppSecretOKCode int = 200

/*CreateRuntimeAppSecretOK OK

swagger:response createRuntimeAppSecretOK
*/
type CreateRuntimeAppSecretOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Secret `json:"body,omitempty"`
}

// NewCreateRuntimeAppSecretOK creates CreateRuntimeAppSecretOK with default headers values
func NewCreateRuntimeAppSecretOK() *CreateRuntimeAppSecretOK {

	return &CreateRuntimeAppSecretOK{}
}

// WithPayload adds the payload to the create runtime app secret o k response
func (o *CreateRuntimeAppSecretOK) WithPayload(payload []*models.Secret) *CreateRuntimeAppSecretOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create runtime app secret o k response
func (o *CreateRuntimeAppSecretOK) SetPayload(payload []*models.Secret) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateRuntimeAppSecretOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*models.Secret, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*CreateRuntimeAppSecretDefault Failed

swagger:response createRuntimeAppSecretDefault
*/
type CreateRuntimeAppSecretDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateRuntimeAppSecretDefault creates CreateRuntimeAppSecretDefault with default headers values
func NewCreateRuntimeAppSecretDefault(code int) *CreateRuntimeAppSecretDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateRuntimeAppSecretDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create runtime app secret default response
func (o *CreateRuntimeAppSecretDefault) WithStatusCode(code int) *CreateRuntimeAppSecretDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create runtime app secret default response
func (o *CreateRuntimeAppSecretDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the create runtime app secret default response
func (o *CreateRuntimeAppSecretDefault) WithPayload(payload *models.Error) *CreateRuntimeAppSecretDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create runtime app secret default response
func (o *CreateRuntimeAppSecretDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateRuntimeAppSecretDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
