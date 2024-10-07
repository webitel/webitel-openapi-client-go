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

// APIUpdateRoleResponse api update role response
//
// swagger:model api.UpdateRoleResponse
type APIUpdateRoleResponse struct {

	// updated
	Updated *APIRole `json:"updated,omitempty"`
}

// Validate validates this api update role response
func (m *APIUpdateRoleResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUpdated(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIUpdateRoleResponse) validateUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.Updated) { // not required
		return nil
	}

	if m.Updated != nil {
		if err := m.Updated.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updated")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updated")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this api update role response based on the context it is used
func (m *APIUpdateRoleResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateUpdated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIUpdateRoleResponse) contextValidateUpdated(ctx context.Context, formats strfmt.Registry) error {

	if m.Updated != nil {

		if swag.IsZero(m.Updated) { // not required
			return nil
		}

		if err := m.Updated.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updated")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updated")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *APIUpdateRoleResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIUpdateRoleResponse) UnmarshalBinary(b []byte) error {
	var res APIUpdateRoleResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
