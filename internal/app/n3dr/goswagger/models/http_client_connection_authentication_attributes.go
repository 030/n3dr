// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HTTPClientConnectionAuthenticationAttributes Http client connection authentication attributes
//
// swagger:model HttpClientConnectionAuthenticationAttributes
type HTTPClientConnectionAuthenticationAttributes struct {

	// ntlm domain
	NtlmDomain string `json:"ntlmDomain,omitempty"`

	// ntlm host
	NtlmHost string `json:"ntlmHost,omitempty"`

	// password
	Password string `json:"password,omitempty"`

	// Authentication type
	// Enum: [username ntlm]
	Type string `json:"type,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this Http client connection authentication attributes
func (m *HTTPClientConnectionAuthenticationAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var httpClientConnectionAuthenticationAttributesTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["username","ntlm"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		httpClientConnectionAuthenticationAttributesTypeTypePropEnum = append(httpClientConnectionAuthenticationAttributesTypeTypePropEnum, v)
	}
}

const (

	// HTTPClientConnectionAuthenticationAttributesTypeUsername captures enum value "username"
	HTTPClientConnectionAuthenticationAttributesTypeUsername string = "username"

	// HTTPClientConnectionAuthenticationAttributesTypeNtlm captures enum value "ntlm"
	HTTPClientConnectionAuthenticationAttributesTypeNtlm string = "ntlm"
)

// prop value enum
func (m *HTTPClientConnectionAuthenticationAttributes) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, httpClientConnectionAuthenticationAttributesTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *HTTPClientConnectionAuthenticationAttributes) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this Http client connection authentication attributes based on context it is used
func (m *HTTPClientConnectionAuthenticationAttributes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HTTPClientConnectionAuthenticationAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HTTPClientConnectionAuthenticationAttributes) UnmarshalBinary(b []byte) error {
	var res HTTPClientConnectionAuthenticationAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}