// Code generated by go-swagger; DO NOT EDIT.

package group_runtime

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/laincloud/console-api/gen/models"
)

// GetGroupSecretOKCode is the HTTP code returned for type GetGroupSecretOK
const GetGroupSecretOKCode int = 200

/*GetGroupSecretOK Get the secret of the group

swagger:response getGroupSecretOK
*/
type GetGroupSecretOK struct {

	/*
	  In: Body
	*/
	Payload *models.Secret `json:"body,omitempty"`
}

// NewGetGroupSecretOK creates GetGroupSecretOK with default headers values
func NewGetGroupSecretOK() *GetGroupSecretOK {

	return &GetGroupSecretOK{}
}

// WithPayload adds the payload to the get group secret o k response
func (o *GetGroupSecretOK) WithPayload(payload *models.Secret) *GetGroupSecretOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get group secret o k response
func (o *GetGroupSecretOK) SetPayload(payload *models.Secret) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetGroupSecretOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetGroupSecretDefault Failed

swagger:response getGroupSecretDefault
*/
type GetGroupSecretDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetGroupSecretDefault creates GetGroupSecretDefault with default headers values
func NewGetGroupSecretDefault(code int) *GetGroupSecretDefault {
	if code <= 0 {
		code = 500
	}

	return &GetGroupSecretDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get group secret default response
func (o *GetGroupSecretDefault) WithStatusCode(code int) *GetGroupSecretDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get group secret default response
func (o *GetGroupSecretDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get group secret default response
func (o *GetGroupSecretDefault) WithPayload(payload *models.Error) *GetGroupSecretDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get group secret default response
func (o *GetGroupSecretDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetGroupSecretDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
