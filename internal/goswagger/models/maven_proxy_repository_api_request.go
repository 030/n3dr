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

// MavenProxyRepositoryAPIRequest maven proxy repository Api request
//
// swagger:model MavenProxyRepositoryApiRequest
type MavenProxyRepositoryAPIRequest struct {

	// cleanup
	Cleanup *CleanupPolicyAttributes `json:"cleanup,omitempty"`

	// http client
	// Required: true
	HTTPClient *HTTPClientAttributesWithPreemptiveAuth `json:"httpClient"`

	// maven
	// Required: true
	Maven *MavenAttributes `json:"maven"`

	// A unique identifier for this repository
	// Example: internal
	// Required: true
	// Pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Name *string `json:"name"`

	// negative cache
	// Required: true
	NegativeCache *NegativeCacheAttributes `json:"negativeCache"`

	// Whether this repository accepts incoming requests
	// Example: true
	// Required: true
	Online *bool `json:"online"`

	// proxy
	// Required: true
	Proxy *ProxyAttributes `json:"proxy"`

	// replication
	Replication *ReplicationAttributes `json:"replication,omitempty"`

	// routing rule
	RoutingRule string `json:"routingRule,omitempty"`

	// storage
	// Required: true
	Storage *StorageAttributes `json:"storage"`
}

// Validate validates this maven proxy repository Api request
func (m *MavenProxyRepositoryAPIRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCleanup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHTTPClient(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaven(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNegativeCache(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOnline(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProxy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReplication(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MavenProxyRepositoryAPIRequest) validateCleanup(formats strfmt.Registry) error {
	if swag.IsZero(m.Cleanup) { // not required
		return nil
	}

	if m.Cleanup != nil {
		if err := m.Cleanup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cleanup")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cleanup")
			}
			return err
		}
	}

	return nil
}

func (m *MavenProxyRepositoryAPIRequest) validateHTTPClient(formats strfmt.Registry) error {

	if err := validate.Required("httpClient", "body", m.HTTPClient); err != nil {
		return err
	}

	if m.HTTPClient != nil {
		if err := m.HTTPClient.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("httpClient")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("httpClient")
			}
			return err
		}
	}

	return nil
}

func (m *MavenProxyRepositoryAPIRequest) validateMaven(formats strfmt.Registry) error {

	if err := validate.Required("maven", "body", m.Maven); err != nil {
		return err
	}

	if m.Maven != nil {
		if err := m.Maven.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("maven")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("maven")
			}
			return err
		}
	}

	return nil
}

func (m *MavenProxyRepositoryAPIRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", *m.Name, `^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$`); err != nil {
		return err
	}

	return nil
}

func (m *MavenProxyRepositoryAPIRequest) validateNegativeCache(formats strfmt.Registry) error {

	if err := validate.Required("negativeCache", "body", m.NegativeCache); err != nil {
		return err
	}

	if m.NegativeCache != nil {
		if err := m.NegativeCache.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("negativeCache")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("negativeCache")
			}
			return err
		}
	}

	return nil
}

func (m *MavenProxyRepositoryAPIRequest) validateOnline(formats strfmt.Registry) error {

	if err := validate.Required("online", "body", m.Online); err != nil {
		return err
	}

	return nil
}

func (m *MavenProxyRepositoryAPIRequest) validateProxy(formats strfmt.Registry) error {

	if err := validate.Required("proxy", "body", m.Proxy); err != nil {
		return err
	}

	if m.Proxy != nil {
		if err := m.Proxy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("proxy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("proxy")
			}
			return err
		}
	}

	return nil
}

func (m *MavenProxyRepositoryAPIRequest) validateReplication(formats strfmt.Registry) error {
	if swag.IsZero(m.Replication) { // not required
		return nil
	}

	if m.Replication != nil {
		if err := m.Replication.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("replication")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("replication")
			}
			return err
		}
	}

	return nil
}

func (m *MavenProxyRepositoryAPIRequest) validateStorage(formats strfmt.Registry) error {

	if err := validate.Required("storage", "body", m.Storage); err != nil {
		return err
	}

	if m.Storage != nil {
		if err := m.Storage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("storage")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this maven proxy repository Api request based on the context it is used
func (m *MavenProxyRepositoryAPIRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCleanup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHTTPClient(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMaven(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNegativeCache(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProxy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReplication(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStorage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MavenProxyRepositoryAPIRequest) contextValidateCleanup(ctx context.Context, formats strfmt.Registry) error {

	if m.Cleanup != nil {
		if err := m.Cleanup.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cleanup")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cleanup")
			}
			return err
		}
	}

	return nil
}

func (m *MavenProxyRepositoryAPIRequest) contextValidateHTTPClient(ctx context.Context, formats strfmt.Registry) error {

	if m.HTTPClient != nil {
		if err := m.HTTPClient.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("httpClient")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("httpClient")
			}
			return err
		}
	}

	return nil
}

func (m *MavenProxyRepositoryAPIRequest) contextValidateMaven(ctx context.Context, formats strfmt.Registry) error {

	if m.Maven != nil {
		if err := m.Maven.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("maven")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("maven")
			}
			return err
		}
	}

	return nil
}

func (m *MavenProxyRepositoryAPIRequest) contextValidateNegativeCache(ctx context.Context, formats strfmt.Registry) error {

	if m.NegativeCache != nil {
		if err := m.NegativeCache.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("negativeCache")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("negativeCache")
			}
			return err
		}
	}

	return nil
}

func (m *MavenProxyRepositoryAPIRequest) contextValidateProxy(ctx context.Context, formats strfmt.Registry) error {

	if m.Proxy != nil {
		if err := m.Proxy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("proxy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("proxy")
			}
			return err
		}
	}

	return nil
}

func (m *MavenProxyRepositoryAPIRequest) contextValidateReplication(ctx context.Context, formats strfmt.Registry) error {

	if m.Replication != nil {
		if err := m.Replication.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("replication")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("replication")
			}
			return err
		}
	}

	return nil
}

func (m *MavenProxyRepositoryAPIRequest) contextValidateStorage(ctx context.Context, formats strfmt.Registry) error {

	if m.Storage != nil {
		if err := m.Storage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("storage")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MavenProxyRepositoryAPIRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MavenProxyRepositoryAPIRequest) UnmarshalBinary(b []byte) error {
	var res MavenProxyRepositoryAPIRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
