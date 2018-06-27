// Code generated by go-swagger; DO NOT EDIT.

package runtime_app

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/laincloud/console-api/gen/models"
)

// ExecRuntimeAppPodOKCode is the HTTP code returned for type ExecRuntimeAppPodOK
const ExecRuntimeAppPodOKCode int = 200

/*ExecRuntimeAppPodOK Enter the pod of the runtime app

swagger:response execRuntimeAppPodOK
*/
type ExecRuntimeAppPodOK struct {
}

// NewExecRuntimeAppPodOK creates ExecRuntimeAppPodOK with default headers values
func NewExecRuntimeAppPodOK() *ExecRuntimeAppPodOK {

	return &ExecRuntimeAppPodOK{}
}

// WriteResponse to the client
func (o *ExecRuntimeAppPodOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

/*ExecRuntimeAppPodDefault Failed

swagger:response execRuntimeAppPodDefault
*/
type ExecRuntimeAppPodDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewExecRuntimeAppPodDefault creates ExecRuntimeAppPodDefault with default headers values
func NewExecRuntimeAppPodDefault(code int) *ExecRuntimeAppPodDefault {
	if code <= 0 {
		code = 500
	}

	return &ExecRuntimeAppPodDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the exec runtime app pod default response
func (o *ExecRuntimeAppPodDefault) WithStatusCode(code int) *ExecRuntimeAppPodDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the exec runtime app pod default response
func (o *ExecRuntimeAppPodDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the exec runtime app pod default response
func (o *ExecRuntimeAppPodDefault) WithPayload(payload *models.Error) *ExecRuntimeAppPodDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the exec runtime app pod default response
func (o *ExecRuntimeAppPodDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ExecRuntimeAppPodDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
