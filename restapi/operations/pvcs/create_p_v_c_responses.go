// Code generated by go-swagger; DO NOT EDIT.

package pvcs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/laincloud/console-api/models"
)

// CreatePVCCreatedCode is the HTTP code returned for type CreatePVCCreated
const CreatePVCCreatedCode int = 201

/*CreatePVCCreated Created

swagger:response createPVCCreated
*/
type CreatePVCCreated struct {

	/*
	  In: Body
	*/
	Payload models.Pvc `json:"body,omitempty"`
}

// NewCreatePVCCreated creates CreatePVCCreated with default headers values
func NewCreatePVCCreated() *CreatePVCCreated {

	return &CreatePVCCreated{}
}

// WithPayload adds the payload to the create p v c created response
func (o *CreatePVCCreated) WithPayload(payload models.Pvc) *CreatePVCCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create p v c created response
func (o *CreatePVCCreated) SetPayload(payload models.Pvc) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreatePVCCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*CreatePVCDefault Failed

swagger:response createPVCDefault
*/
type CreatePVCDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreatePVCDefault creates CreatePVCDefault with default headers values
func NewCreatePVCDefault(code int) *CreatePVCDefault {
	if code <= 0 {
		code = 500
	}

	return &CreatePVCDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create p v c default response
func (o *CreatePVCDefault) WithStatusCode(code int) *CreatePVCDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create p v c default response
func (o *CreatePVCDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the create p v c default response
func (o *CreatePVCDefault) WithPayload(payload *models.Error) *CreatePVCDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create p v c default response
func (o *CreatePVCDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreatePVCDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
