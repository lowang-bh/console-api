// Code generated by go-swagger; DO NOT EDIT.

package build_apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetBuildAppsParams creates a new GetBuildAppsParams object
// no default values defined in spec.
func NewGetBuildAppsParams() GetBuildAppsParams {

	return GetBuildAppsParams{}
}

// GetBuildAppsParams contains all the bound params for the get build apps operation
// typically these are obtained from a http.Request
//
// swagger:parameters getBuildApps
type GetBuildAppsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

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
// To ensure default values, the struct must have been initialized with NewGetBuildAppsParams() beforehand.
func (o *GetBuildAppsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rGroupname, rhkGroupname, _ := route.Params.GetOK("groupname")
	if err := o.bindGroupname(rGroupname, rhkGroupname, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetBuildAppsParams) bindGroupname(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *GetBuildAppsParams) validateGroupname(formats strfmt.Registry) error {

	if err := validate.MinLength("groupname", "path", o.Groupname, 1); err != nil {
		return err
	}

	return nil
}
