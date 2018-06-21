// Code generated by go-swagger; DO NOT EDIT.

package builds

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetBuildLogOKCode is the HTTP code returned for type GetBuildLogOK
const GetBuildLogOKCode int = 200

/*GetBuildLogOK Get the build log of the build app

swagger:response getBuildLogOK
*/
type GetBuildLogOK struct {
}

// NewGetBuildLogOK creates GetBuildLogOK with default headers values
func NewGetBuildLogOK() *GetBuildLogOK {

	return &GetBuildLogOK{}
}

// WriteResponse to the client
func (o *GetBuildLogOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}
