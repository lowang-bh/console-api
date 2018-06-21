// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ComponentsStatus components status
// swagger:model componentsStatus
type ComponentsStatus struct {

	// ceph
	Ceph interface{} `json:"ceph,omitempty"`

	// etcd
	Etcd interface{} `json:"etcd,omitempty"`

	// harbor
	Harbor interface{} `json:"harbor,omitempty"`

	// mysql
	Mysql interface{} `json:"mysql,omitempty"`

	// redis
	Redis interface{} `json:"redis,omitempty"`
}

// Validate validates this components status
func (m *ComponentsStatus) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ComponentsStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ComponentsStatus) UnmarshalBinary(b []byte) error {
	var res ComponentsStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
