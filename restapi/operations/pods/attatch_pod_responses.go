// Code generated by go-swagger; DO NOT EDIT.

package pods

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/laincloud/console-api/models"
)

// AttatchPodOKCode is the HTTP code returned for type AttatchPodOK
const AttatchPodOKCode int = 200

/*AttatchPodOK Attatch the pod of the runtime app

swagger:response attatchPodOK
*/
type AttatchPodOK struct {
}

// NewAttatchPodOK creates AttatchPodOK with default headers values
func NewAttatchPodOK() *AttatchPodOK {

	return &AttatchPodOK{}
}

// WriteResponse to the client
func (o *AttatchPodOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

/*AttatchPodDefault Failed

swagger:response attatchPodDefault
*/
type AttatchPodDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAttatchPodDefault creates AttatchPodDefault with default headers values
func NewAttatchPodDefault(code int) *AttatchPodDefault {
	if code <= 0 {
		code = 500
	}

	return &AttatchPodDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the attatch pod default response
func (o *AttatchPodDefault) WithStatusCode(code int) *AttatchPodDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the attatch pod default response
func (o *AttatchPodDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the attatch pod default response
func (o *AttatchPodDefault) WithPayload(payload *models.Error) *AttatchPodDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the attatch pod default response
func (o *AttatchPodDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AttatchPodDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
