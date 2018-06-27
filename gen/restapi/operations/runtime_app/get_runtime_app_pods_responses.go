// Code generated by go-swagger; DO NOT EDIT.

package runtime_app

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/laincloud/console-api/gen/models"
)

// GetRuntimeAppPodsOKCode is the HTTP code returned for type GetRuntimeAppPodsOK
const GetRuntimeAppPodsOKCode int = 200

/*GetRuntimeAppPodsOK Get the pods of the runtime app

swagger:response getRuntimeAppPodsOK
*/
type GetRuntimeAppPodsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Pod `json:"body,omitempty"`
}

// NewGetRuntimeAppPodsOK creates GetRuntimeAppPodsOK with default headers values
func NewGetRuntimeAppPodsOK() *GetRuntimeAppPodsOK {

	return &GetRuntimeAppPodsOK{}
}

// WithPayload adds the payload to the get runtime app pods o k response
func (o *GetRuntimeAppPodsOK) WithPayload(payload []*models.Pod) *GetRuntimeAppPodsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get runtime app pods o k response
func (o *GetRuntimeAppPodsOK) SetPayload(payload []*models.Pod) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRuntimeAppPodsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*models.Pod, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*GetRuntimeAppPodsDefault Failed

swagger:response getRuntimeAppPodsDefault
*/
type GetRuntimeAppPodsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetRuntimeAppPodsDefault creates GetRuntimeAppPodsDefault with default headers values
func NewGetRuntimeAppPodsDefault(code int) *GetRuntimeAppPodsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetRuntimeAppPodsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get runtime app pods default response
func (o *GetRuntimeAppPodsDefault) WithStatusCode(code int) *GetRuntimeAppPodsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get runtime app pods default response
func (o *GetRuntimeAppPodsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get runtime app pods default response
func (o *GetRuntimeAppPodsDefault) WithPayload(payload *models.Error) *GetRuntimeAppPodsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get runtime app pods default response
func (o *GetRuntimeAppPodsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRuntimeAppPodsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}