// Code generated by go-swagger; DO NOT EDIT.

package runtime_app_secrets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/laincloud/console-api/models"
)

// GetRuntimeAppSecretOKCode is the HTTP code returned for type GetRuntimeAppSecretOK
const GetRuntimeAppSecretOKCode int = 200

/*GetRuntimeAppSecretOK Get the secret of the runtime app

swagger:response getRuntimeAppSecretOK
*/
type GetRuntimeAppSecretOK struct {

	/*
	  In: Body
	*/
	Payload *models.Secret `json:"body,omitempty"`
}

// NewGetRuntimeAppSecretOK creates GetRuntimeAppSecretOK with default headers values
func NewGetRuntimeAppSecretOK() *GetRuntimeAppSecretOK {

	return &GetRuntimeAppSecretOK{}
}

// WithPayload adds the payload to the get runtime app secret o k response
func (o *GetRuntimeAppSecretOK) WithPayload(payload *models.Secret) *GetRuntimeAppSecretOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get runtime app secret o k response
func (o *GetRuntimeAppSecretOK) SetPayload(payload *models.Secret) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRuntimeAppSecretOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetRuntimeAppSecretDefault Failed

swagger:response getRuntimeAppSecretDefault
*/
type GetRuntimeAppSecretDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetRuntimeAppSecretDefault creates GetRuntimeAppSecretDefault with default headers values
func NewGetRuntimeAppSecretDefault(code int) *GetRuntimeAppSecretDefault {
	if code <= 0 {
		code = 500
	}

	return &GetRuntimeAppSecretDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get runtime app secret default response
func (o *GetRuntimeAppSecretDefault) WithStatusCode(code int) *GetRuntimeAppSecretDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get runtime app secret default response
func (o *GetRuntimeAppSecretDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get runtime app secret default response
func (o *GetRuntimeAppSecretDefault) WithPayload(payload *models.Error) *GetRuntimeAppSecretDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get runtime app secret default response
func (o *GetRuntimeAppSecretDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRuntimeAppSecretDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
