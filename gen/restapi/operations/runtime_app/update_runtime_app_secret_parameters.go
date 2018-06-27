// Code generated by go-swagger; DO NOT EDIT.

package runtime_app

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/laincloud/console-api/gen/models"
)

// NewUpdateRuntimeAppSecretParams creates a new UpdateRuntimeAppSecretParams object
// no default values defined in spec.
func NewUpdateRuntimeAppSecretParams() UpdateRuntimeAppSecretParams {

	return UpdateRuntimeAppSecretParams{}
}

// UpdateRuntimeAppSecretParams contains all the bound params for the update runtime app secret operation
// typically these are obtained from a http.Request
//
// swagger:parameters updateRuntimeAppSecret
type UpdateRuntimeAppSecretParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  Min Length: 1
	  In: path
	*/
	Appname string
	/*
	  In: body
	*/
	Body *models.Secret
	/*
	  Required: true
	  Min Length: 1
	  In: path
	*/
	Groupname string
	/*
	  Required: true
	  Min Length: 1
	  In: path
	*/
	Secretname string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUpdateRuntimeAppSecretParams() beforehand.
func (o *UpdateRuntimeAppSecretParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rAppname, rhkAppname, _ := route.Params.GetOK("appname")
	if err := o.bindAppname(rAppname, rhkAppname, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Secret
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("body", "body", "", err))
		} else {

			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = &body
			}
		}
	}
	rGroupname, rhkGroupname, _ := route.Params.GetOK("groupname")
	if err := o.bindGroupname(rGroupname, rhkGroupname, route.Formats); err != nil {
		res = append(res, err)
	}

	rSecretname, rhkSecretname, _ := route.Params.GetOK("secretname")
	if err := o.bindSecretname(rSecretname, rhkSecretname, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateRuntimeAppSecretParams) bindAppname(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *UpdateRuntimeAppSecretParams) validateAppname(formats strfmt.Registry) error {

	if err := validate.MinLength("appname", "path", o.Appname, 1); err != nil {
		return err
	}

	return nil
}

func (o *UpdateRuntimeAppSecretParams) bindGroupname(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *UpdateRuntimeAppSecretParams) validateGroupname(formats strfmt.Registry) error {

	if err := validate.MinLength("groupname", "path", o.Groupname, 1); err != nil {
		return err
	}

	return nil
}

func (o *UpdateRuntimeAppSecretParams) bindSecretname(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Secretname = raw

	if err := o.validateSecretname(formats); err != nil {
		return err
	}

	return nil
}

func (o *UpdateRuntimeAppSecretParams) validateSecretname(formats strfmt.Registry) error {

	if err := validate.MinLength("secretname", "path", o.Secretname, 1); err != nil {
		return err
	}

	return nil
}
