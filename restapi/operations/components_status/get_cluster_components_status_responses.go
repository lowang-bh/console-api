// Code generated by go-swagger; DO NOT EDIT.

package components_status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/laincloud/console-api/models"
)

// GetClusterComponentsStatusOKCode is the HTTP code returned for type GetClusterComponentsStatusOK
const GetClusterComponentsStatusOKCode int = 200

/*GetClusterComponentsStatusOK Get the cluster components status

swagger:response getClusterComponentsStatusOK
*/
type GetClusterComponentsStatusOK struct {

	/*
	  In: Body
	*/
	Payload *models.ComponentsStatus `json:"body,omitempty"`
}

// NewGetClusterComponentsStatusOK creates GetClusterComponentsStatusOK with default headers values
func NewGetClusterComponentsStatusOK() *GetClusterComponentsStatusOK {

	return &GetClusterComponentsStatusOK{}
}

// WithPayload adds the payload to the get cluster components status o k response
func (o *GetClusterComponentsStatusOK) WithPayload(payload *models.ComponentsStatus) *GetClusterComponentsStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get cluster components status o k response
func (o *GetClusterComponentsStatusOK) SetPayload(payload *models.ComponentsStatus) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetClusterComponentsStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetClusterComponentsStatusDefault Failed

swagger:response getClusterComponentsStatusDefault
*/
type GetClusterComponentsStatusDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetClusterComponentsStatusDefault creates GetClusterComponentsStatusDefault with default headers values
func NewGetClusterComponentsStatusDefault(code int) *GetClusterComponentsStatusDefault {
	if code <= 0 {
		code = 500
	}

	return &GetClusterComponentsStatusDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get cluster components status default response
func (o *GetClusterComponentsStatusDefault) WithStatusCode(code int) *GetClusterComponentsStatusDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get cluster components status default response
func (o *GetClusterComponentsStatusDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get cluster components status default response
func (o *GetClusterComponentsStatusDefault) WithPayload(payload *models.Error) *GetClusterComponentsStatusDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get cluster components status default response
func (o *GetClusterComponentsStatusDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetClusterComponentsStatusDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
