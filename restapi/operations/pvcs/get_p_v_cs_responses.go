// Code generated by go-swagger; DO NOT EDIT.

package pvcs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/laincloud/console-api/models"
)

// GetPVCsOKCode is the HTTP code returned for type GetPVCsOK
const GetPVCsOKCode int = 200

/*GetPVCsOK List the PVCs

swagger:response getPVCsOK
*/
type GetPVCsOK struct {

	/*
	  In: Body
	*/
	Payload []models.Pvc `json:"body,omitempty"`
}

// NewGetPVCsOK creates GetPVCsOK with default headers values
func NewGetPVCsOK() *GetPVCsOK {

	return &GetPVCsOK{}
}

// WithPayload adds the payload to the get p v cs o k response
func (o *GetPVCsOK) WithPayload(payload []models.Pvc) *GetPVCsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get p v cs o k response
func (o *GetPVCsOK) SetPayload(payload []models.Pvc) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPVCsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]models.Pvc, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*GetPVCsDefault Failed

swagger:response getPVCsDefault
*/
type GetPVCsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetPVCsDefault creates GetPVCsDefault with default headers values
func NewGetPVCsDefault(code int) *GetPVCsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetPVCsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get p v cs default response
func (o *GetPVCsDefault) WithStatusCode(code int) *GetPVCsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get p v cs default response
func (o *GetPVCsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get p v cs default response
func (o *GetPVCsDefault) WithPayload(payload *models.Error) *GetPVCsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get p v cs default response
func (o *GetPVCsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPVCsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
