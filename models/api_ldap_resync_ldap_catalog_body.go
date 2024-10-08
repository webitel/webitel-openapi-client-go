// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APILDAPResyncLDAPCatalogBody api LDAP resync LDAP catalog body
//
// swagger:model api.LDAP.ResyncLDAPCatalogBody
type APILDAPResyncLDAPCatalogBody struct {

	// Options
	Options *APILDAProcessOptions `json:"options,omitempty"`
}

// Validate validates this api LDAP resync LDAP catalog body
func (m *APILDAPResyncLDAPCatalogBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOptions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APILDAPResyncLDAPCatalogBody) validateOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.Options) { // not required
		return nil
	}

	if m.Options != nil {
		if err := m.Options.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("options")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("options")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this api LDAP resync LDAP catalog body based on the context it is used
func (m *APILDAPResyncLDAPCatalogBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APILDAPResyncLDAPCatalogBody) contextValidateOptions(ctx context.Context, formats strfmt.Registry) error {

	if m.Options != nil {

		if swag.IsZero(m.Options) { // not required
			return nil
		}

		if err := m.Options.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("options")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("options")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *APILDAPResyncLDAPCatalogBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APILDAPResyncLDAPCatalogBody) UnmarshalBinary(b []byte) error {
	var res APILDAPResyncLDAPCatalogBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
