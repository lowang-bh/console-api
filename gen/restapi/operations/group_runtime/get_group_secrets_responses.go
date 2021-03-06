// Code generated by go-swagger; DO NOT EDIT.

package group_runtime

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/laincloud/console-api/gen/models"
)

// GetGroupSecretsOKCode is the HTTP code returned for type GetGroupSecretsOK
const GetGroupSecretsOKCode int = 200

/*GetGroupSecretsOK List the secrets

swagger:response getGroupSecretsOK
*/
type GetGroupSecretsOK struct {

	/*
	  In: Body
	*/
	Payload []models.Pvc `json:"body,omitempty"`
}

// NewGetGroupSecretsOK creates GetGroupSecretsOK with default headers values
func NewGetGroupSecretsOK() *GetGroupSecretsOK {

	return &GetGroupSecretsOK{}
}

// WithPayload adds the payload to the get group secrets o k response
func (o *GetGroupSecretsOK) WithPayload(payload []models.Pvc) *GetGroupSecretsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get group secrets o k response
func (o *GetGroupSecretsOK) SetPayload(payload []models.Pvc) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetGroupSecretsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]models.Pvc, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*GetGroupSecretsDefault Failed

swagger:response getGroupSecretsDefault
*/
type GetGroupSecretsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetGroupSecretsDefault creates GetGroupSecretsDefault with default headers values
func NewGetGroupSecretsDefault(code int) *GetGroupSecretsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetGroupSecretsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get group secrets default response
func (o *GetGroupSecretsDefault) WithStatusCode(code int) *GetGroupSecretsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get group secrets default response
func (o *GetGroupSecretsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get group secrets default response
func (o *GetGroupSecretsDefault) WithPayload(payload *models.Error) *GetGroupSecretsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get group secrets default response
func (o *GetGroupSecretsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetGroupSecretsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
