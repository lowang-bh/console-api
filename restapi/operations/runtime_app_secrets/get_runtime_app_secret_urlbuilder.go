// Code generated by go-swagger; DO NOT EDIT.

package runtime_app_secrets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// GetRuntimeAppSecretURL generates an URL for the get runtime app secret operation
type GetRuntimeAppSecretURL struct {
	Appname    string
	Groupname  string
	Secretname string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetRuntimeAppSecretURL) WithBasePath(bp string) *GetRuntimeAppSecretURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetRuntimeAppSecretURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetRuntimeAppSecretURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/api/v1/groups/{groupname}/runtime/apps/{appname}/secrets/{secretname}"

	appname := o.Appname
	if appname != "" {
		_path = strings.Replace(_path, "{appname}", appname, -1)
	} else {
		return nil, errors.New("Appname is required on GetRuntimeAppSecretURL")
	}

	groupname := o.Groupname
	if groupname != "" {
		_path = strings.Replace(_path, "{groupname}", groupname, -1)
	} else {
		return nil, errors.New("Groupname is required on GetRuntimeAppSecretURL")
	}

	secretname := o.Secretname
	if secretname != "" {
		_path = strings.Replace(_path, "{secretname}", secretname, -1)
	} else {
		return nil, errors.New("Secretname is required on GetRuntimeAppSecretURL")
	}

	_basePath := o._basePath
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetRuntimeAppSecretURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetRuntimeAppSecretURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetRuntimeAppSecretURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetRuntimeAppSecretURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetRuntimeAppSecretURL")
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
func (o *GetRuntimeAppSecretURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
