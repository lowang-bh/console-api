// Code generated by go-swagger; DO NOT EDIT.

package runtime_app

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/laincloud/console-api/gen/models"
)

// GetRuntimeAppServiceOKCode is the HTTP code returned for type GetRuntimeAppServiceOK
const GetRuntimeAppServiceOKCode int = 200

/*GetRuntimeAppServiceOK Get the service of the runtime app

swagger:response getRuntimeAppServiceOK
*/
type GetRuntimeAppServiceOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Service `json:"body,omitempty"`
}

// NewGetRuntimeAppServiceOK creates GetRuntimeAppServiceOK with default headers values
func NewGetRuntimeAppServiceOK() *GetRuntimeAppServiceOK {

	return &GetRuntimeAppServiceOK{}
}

// WithPayload adds the payload to the get runtime app service o k response
func (o *GetRuntimeAppServiceOK) WithPayload(payload []*models.Service) *GetRuntimeAppServiceOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get runtime app service o k response
func (o *GetRuntimeAppServiceOK) SetPayload(payload []*models.Service) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRuntimeAppServiceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*models.Service, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*GetRuntimeAppServiceDefault Failed

swagger:response getRuntimeAppServiceDefault
*/
type GetRuntimeAppServiceDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetRuntimeAppServiceDefault creates GetRuntimeAppServiceDefault with default headers values
func NewGetRuntimeAppServiceDefault(code int) *GetRuntimeAppServiceDefault {
	if code <= 0 {
		code = 500
	}

	return &GetRuntimeAppServiceDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get runtime app service default response
func (o *GetRuntimeAppServiceDefault) WithStatusCode(code int) *GetRuntimeAppServiceDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get runtime app service default response
func (o *GetRuntimeAppServiceDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get runtime app service default response
func (o *GetRuntimeAppServiceDefault) WithPayload(payload *models.Error) *GetRuntimeAppServiceDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get runtime app service default response
func (o *GetRuntimeAppServiceDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRuntimeAppServiceDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}