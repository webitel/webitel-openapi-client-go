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

// APILDAPUpdateLDAPCatalogBody api LDAP update LDAP catalog body
//
// swagger:model api.LDAP.UpdateLDAPCatalogBody
type APILDAPUpdateLDAPCatalogBody struct {

	// catalog
	Catalog *APILDAPUpdateLDAPCatalogBodyCatalog `json:"catalog,omitempty"`

	// Fields for partial update. PATCH
	Fields []string `json:"fields"`
}

// Validate validates this api LDAP update LDAP catalog body
func (m *APILDAPUpdateLDAPCatalogBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCatalog(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APILDAPUpdateLDAPCatalogBody) validateCatalog(formats strfmt.Registry) error {
	if swag.IsZero(m.Catalog) { // not required
		return nil
	}

	if m.Catalog != nil {
		if err := m.Catalog.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("catalog")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("catalog")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this api LDAP update LDAP catalog body based on the context it is used
func (m *APILDAPUpdateLDAPCatalogBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCatalog(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APILDAPUpdateLDAPCatalogBody) contextValidateCatalog(ctx context.Context, formats strfmt.Registry) error {

	if m.Catalog != nil {

		if swag.IsZero(m.Catalog) { // not required
			return nil
		}

		if err := m.Catalog.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("catalog")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("catalog")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *APILDAPUpdateLDAPCatalogBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APILDAPUpdateLDAPCatalogBody) UnmarshalBinary(b []byte) error {
	var res APILDAPUpdateLDAPCatalogBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
