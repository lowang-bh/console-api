// Code generated by go-swagger; DO NOT EDIT.

package runtime_apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteRuntimeAppParams creates a new DeleteRuntimeAppParams object
// no default values defined in spec.
func NewDeleteRuntimeAppParams() DeleteRuntimeAppParams {

	return DeleteRuntimeAppParams{}
}

// DeleteRuntimeAppParams contains all the bound params for the delete runtime app operation
// typically these are obtained from a http.Request
//
// swagger:parameters deleteRuntimeApp
type DeleteRuntimeAppParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  Min Length: 1
	  In: path
	*/
	Appname string
	/*
	  Required: true
	  Min Length: 1
	  In: path
	*/
	Groupname string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeleteRuntimeAppParams() beforehand.
func (o *DeleteRuntimeAppParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rAppname, rhkAppname, _ := route.Params.GetOK("appname")
	if err := o.bindAppname(rAppname, rhkAppname, route.Formats); err != nil {
		res = append(res, err)
	}

	rGroupname, rhkGroupname, _ := route.Params.GetOK("groupname")
	if err := o.bindGroupname(rGroupname, rhkGroupname, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteRuntimeAppParams) bindAppname(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Appname = raw

	if err := o.validateAppname(formats); err != nil {
		return err
	}

	return nil
}

func (o *DeleteRuntimeAppParams) validateAppname(formats strfmt.Registry) error {

	if err := validate.MinLength("appname", "path", o.Appname, 1); err != nil {
		return err
	}

	return nil
}

func (o *DeleteRuntimeAppParams) bindGroupname(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Groupname = raw

	if err := o.validateGroupname(formats); err != nil {
		return err
	}

	return nil
}

func (o *DeleteRuntimeAppParams) validateGroupname(formats strfmt.Registry) error {

	if err := validate.MinLength("groupname", "path", o.Groupname, 1); err != nil {
		return err
	}

	return nil
}
