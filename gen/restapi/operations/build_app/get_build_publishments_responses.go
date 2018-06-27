// Code generated by go-swagger; DO NOT EDIT.

package build_app

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/laincloud/console-api/gen/models"
)

// GetBuildPublishmentsOKCode is the HTTP code returned for type GetBuildPublishmentsOK
const GetBuildPublishmentsOKCode int = 200

/*GetBuildPublishmentsOK OK

swagger:response getBuildPublishmentsOK
*/
type GetBuildPublishmentsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Publishment `json:"body,omitempty"`
}

// NewGetBuildPublishmentsOK creates GetBuildPublishmentsOK with default headers values
func NewGetBuildPublishmentsOK() *GetBuildPublishmentsOK {

	return &GetBuildPublishmentsOK{}
}

// WithPayload adds the payload to the get build publishments o k response
func (o *GetBuildPublishmentsOK) WithPayload(payload []*models.Publishment) *GetBuildPublishmentsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get build publishments o k response
func (o *GetBuildPublishmentsOK) SetPayload(payload []*models.Publishment) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBuildPublishmentsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*models.Publishment, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*GetBuildPublishmentsDefault Failed

swagger:response getBuildPublishmentsDefault
*/
type GetBuildPublishmentsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetBuildPublishmentsDefault creates GetBuildPublishmentsDefault with default headers values
func NewGetBuildPublishmentsDefault(code int) *GetBuildPublishmentsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetBuildPublishmentsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get build publishments default response
func (o *GetBuildPublishmentsDefault) WithStatusCode(code int) *GetBuildPublishmentsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get build publishments default response
func (o *GetBuildPublishmentsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get build publishments default response
func (o *GetBuildPublishmentsDefault) WithPayload(payload *models.Error) *GetBuildPublishmentsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get build publishments default response
func (o *GetBuildPublishmentsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBuildPublishmentsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}