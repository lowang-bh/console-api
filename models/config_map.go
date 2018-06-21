// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ConfigMap config map
// swagger:model configMap
type ConfigMap struct {

	// Unix timstamp (Unit: second)
	CreatedAt int64 `json:"createdAt,omitempty"`

	// 创建者邮箱
	// Min Length: 1
	CreatedBy string `json:"createdBy,omitempty"`

	// id
	// Required: true
	// Min Length: 1
	ID *string `json:"id"`

	// 所有者邮箱
	// Min Length: 1
	OwnedBy string `json:"ownedBy,omitempty"`

	// Unix timstamp (Unit: second)
	UpdatedAt int64 `json:"updatedAt,omitempty"`
}

// Validate validates this config map
func (m *ConfigMap) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwnedBy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigMap) validateCreatedBy(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedBy) { // not required
		return nil
	}

	if err := validate.MinLength("createdBy", "body", string(m.CreatedBy), 1); err != nil {
		return err
	}

	return nil
}

func (m *ConfigMap) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.MinLength("id", "body", string(*m.ID), 1); err != nil {
		return err
	}

	return nil
}

func (m *ConfigMap) validateOwnedBy(formats strfmt.Registry) error {

	if swag.IsZero(m.OwnedBy) { // not required
		return nil
	}

	if err := validate.MinLength("ownedBy", "body", string(m.OwnedBy), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigMap) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigMap) UnmarshalBinary(b []byte) error {
	var res ConfigMap
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
