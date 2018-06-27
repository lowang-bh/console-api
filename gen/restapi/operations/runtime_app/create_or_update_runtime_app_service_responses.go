// Code generated by go-swagger; DO NOT EDIT.

package runtime_app

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/laincloud/console-api/gen/models"
)

// CreateOrUpdateRuntimeAppServiceOKCode is the HTTP code returned for type CreateOrUpdateRuntimeAppServiceOK
const CreateOrUpdateRuntimeAppServiceOKCode int = 200

/*CreateOrUpdateRuntimeAppServiceOK OK

swagger:response createOrUpdateRuntimeAppServiceOK
*/
type CreateOrUpdateRuntimeAppServiceOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Service `json:"body,omitempty"`
}

// NewCreateOrUpdateRuntimeAppServiceOK creates CreateOrUpdateRuntimeAppServiceOK with default headers values
func NewCreateOrUpdateRuntimeAppServiceOK() *CreateOrUpdateRuntimeAppServiceOK {

	return &CreateOrUpdateRuntimeAppServiceOK{}
}

// WithPayload adds the payload to the create or update runtime app service o k response
func (o *CreateOrUpdateRuntimeAppServiceOK) WithPayload(payload []*models.Service) *CreateOrUpdateRuntimeAppServiceOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create or update runtime app service o k response
func (o *CreateOrUpdateRuntimeAppServiceOK) SetPayload(payload []*models.Service) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateOrUpdateRuntimeAppServiceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*models.Service, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*CreateOrUpdateRuntimeAppServiceDefault Failed

swagger:response createOrUpdateRuntimeAppServiceDefault
*/
type CreateOrUpdateRuntimeAppServiceDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateOrUpdateRuntimeAppServiceDefault creates CreateOrUpdateRuntimeAppServiceDefault with default headers values
func NewCreateOrUpdateRuntimeAppServiceDefault(code int) *CreateOrUpdateRuntimeAppServiceDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateOrUpdateRuntimeAppServiceDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create or update runtime app service default response
func (o *CreateOrUpdateRuntimeAppServiceDefault) WithStatusCode(code int) *CreateOrUpdateRuntimeAppServiceDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create or update runtime app service default response
func (o *CreateOrUpdateRuntimeAppServiceDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the create or update runtime app service default response
func (o *CreateOrUpdateRuntimeAppServiceDefault) WithPayload(payload *models.Error) *CreateOrUpdateRuntimeAppServiceDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create or update runtime app service default response
func (o *CreateOrUpdateRuntimeAppServiceDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateOrUpdateRuntimeAppServiceDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
