// Code generated by go-swagger; DO NOT EDIT.

package secret

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/laincloud/console-api/models"
)

// CreateOrUpdateSecretOKCode is the HTTP code returned for type CreateOrUpdateSecretOK
const CreateOrUpdateSecretOKCode int = 200

/*CreateOrUpdateSecretOK OK

swagger:response createOrUpdateSecretOK
*/
type CreateOrUpdateSecretOK struct {

	/*
	  In: Body
	*/
	Payload *models.Secret `json:"body,omitempty"`
}

// NewCreateOrUpdateSecretOK creates CreateOrUpdateSecretOK with default headers values
func NewCreateOrUpdateSecretOK() *CreateOrUpdateSecretOK {

	return &CreateOrUpdateSecretOK{}
}

// WithPayload adds the payload to the create or update secret o k response
func (o *CreateOrUpdateSecretOK) WithPayload(payload *models.Secret) *CreateOrUpdateSecretOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create or update secret o k response
func (o *CreateOrUpdateSecretOK) SetPayload(payload *models.Secret) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateOrUpdateSecretOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*CreateOrUpdateSecretDefault Failed

swagger:response createOrUpdateSecretDefault
*/
type CreateOrUpdateSecretDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateOrUpdateSecretDefault creates CreateOrUpdateSecretDefault with default headers values
func NewCreateOrUpdateSecretDefault(code int) *CreateOrUpdateSecretDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateOrUpdateSecretDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create or update secret default response
func (o *CreateOrUpdateSecretDefault) WithStatusCode(code int) *CreateOrUpdateSecretDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create or update secret default response
func (o *CreateOrUpdateSecretDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the create or update secret default response
func (o *CreateOrUpdateSecretDefault) WithPayload(payload *models.Error) *CreateOrUpdateSecretDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create or update secret default response
func (o *CreateOrUpdateSecretDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateOrUpdateSecretDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
