// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RoleXORequest role x o request
//
// swagger:model RoleXORequest
type RoleXORequest struct {

	// The description of this role.
	Description string `json:"description,omitempty"`

	// The id of the role.
	ID string `json:"id,omitempty"`

	// The name of the role.
	Name string `json:"name,omitempty"`

	// The list of privileges assigned to this role.
	// Unique: true
	Privileges []string `json:"privileges"`

	// The list of roles assigned to this role.
	// Unique: true
	Roles []string `json:"roles"`
}

// Validate validates this role x o request
func (m *RoleXORequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePrivileges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoles(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RoleXORequest) validatePrivileges(formats strfmt.Registry) error {
	if swag.IsZero(m.Privileges) { // not required
		return nil
	}

	if err := validate.UniqueItems("privileges", "body", m.Privileges); err != nil {
		return err
	}

	return nil
}

func (m *RoleXORequest) validateRoles(formats strfmt.Registry) error {
	if swag.IsZero(m.Roles) { // not required
		return nil
	}

	if err := validate.UniqueItems("roles", "body", m.Roles); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this role x o request based on context it is used
func (m *RoleXORequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RoleXORequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RoleXORequest) UnmarshalBinary(b []byte) error {
	var res RoleXORequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
