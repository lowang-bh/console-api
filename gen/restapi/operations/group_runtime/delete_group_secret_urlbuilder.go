// Code generated by go-swagger; DO NOT EDIT.

package group_runtime

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// DeleteGroupSecretURL generates an URL for the delete group secret operation
type DeleteGroupSecretURL struct {
	Groupname  string
	Secretname string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *DeleteGroupSecretURL) WithBasePath(bp string) *DeleteGroupSecretURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *DeleteGroupSecretURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *DeleteGroupSecretURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/api/v1/groups/{groupname}/runtime/secrets/{secretname}"

	groupname := o.Groupname
	if groupname != "" {
		_path = strings.Replace(_path, "{groupname}", groupname, -1)
	} else {
		return nil, errors.New("Groupname is required on DeleteGroupSecretURL")
	}

	secretname := o.Secretname
	if secretname != "" {
		_path = strings.Replace(_path, "{secretname}", secretname, -1)
	} else {
		return nil, errors.New("Secretname is required on DeleteGroupSecretURL")
	}

	_basePath := o._basePath
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *DeleteGroupSecretURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *DeleteGroupSecretURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *DeleteGroupSecretURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on DeleteGroupSecretURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on DeleteGroupSecretURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *DeleteGroupSecretURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
