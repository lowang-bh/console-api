// Code generated by go-swagger; DO NOT EDIT.

package config_map

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/laincloud/console-api/models"
)

// GetConfigMapOKCode is the HTTP code returned for type GetConfigMapOK
const GetConfigMapOKCode int = 200

/*GetConfigMapOK Get the config map of the runtime app

swagger:response getConfigMapOK
*/
type GetConfigMapOK struct {

	/*
	  In: Body
	*/
	Payload *models.ConfigMap `json:"body,omitempty"`
}

// NewGetConfigMapOK creates GetConfigMapOK with default headers values
func NewGetConfigMapOK() *GetConfigMapOK {

	return &GetConfigMapOK{}
}

// WithPayload adds the payload to the get config map o k response
func (o *GetConfigMapOK) WithPayload(payload *models.ConfigMap) *GetConfigMapOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config map o k response
func (o *GetConfigMapOK) SetPayload(payload *models.ConfigMap) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigMapOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetConfigMapDefault Failed

swagger:response getConfigMapDefault
*/
type GetConfigMapDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetConfigMapDefault creates GetConfigMapDefault with default headers values
func NewGetConfigMapDefault(code int) *GetConfigMapDefault {
	if code <= 0 {
		code = 500
	}

	return &GetConfigMapDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get config map default response
func (o *GetConfigMapDefault) WithStatusCode(code int) *GetConfigMapDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get config map default response
func (o *GetConfigMapDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get config map default response
func (o *GetConfigMapDefault) WithPayload(payload *models.Error) *GetConfigMapDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config map default response
func (o *GetConfigMapDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigMapDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
