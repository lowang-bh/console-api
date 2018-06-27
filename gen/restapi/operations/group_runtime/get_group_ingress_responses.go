// Code generated by go-swagger; DO NOT EDIT.

package group_runtime

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/laincloud/console-api/gen/models"
)

// GetGroupIngressOKCode is the HTTP code returned for type GetGroupIngressOK
const GetGroupIngressOKCode int = 200

/*GetGroupIngressOK Get the ingress

swagger:response getGroupIngressOK
*/
type GetGroupIngressOK struct {

	/*
	  In: Body
	*/
	Payload *models.Ingress `json:"body,omitempty"`
}

// NewGetGroupIngressOK creates GetGroupIngressOK with default headers values
func NewGetGroupIngressOK() *GetGroupIngressOK {

	return &GetGroupIngressOK{}
}

// WithPayload adds the payload to the get group ingress o k response
func (o *GetGroupIngressOK) WithPayload(payload *models.Ingress) *GetGroupIngressOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get group ingress o k response
func (o *GetGroupIngressOK) SetPayload(payload *models.Ingress) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetGroupIngressOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetGroupIngressDefault Failed

swagger:response getGroupIngressDefault
*/
type GetGroupIngressDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetGroupIngressDefault creates GetGroupIngressDefault with default headers values
func NewGetGroupIngressDefault(code int) *GetGroupIngressDefault {
	if code <= 0 {
		code = 500
	}

	return &GetGroupIngressDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get group ingress default response
func (o *GetGroupIngressDefault) WithStatusCode(code int) *GetGroupIngressDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get group ingress default response
func (o *GetGroupIngressDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get group ingress default response
func (o *GetGroupIngressDefault) WithPayload(payload *models.Error) *GetGroupIngressDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get group ingress default response
func (o *GetGroupIngressDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetGroupIngressDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
