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

// ContactGroupsUpdateContactGroup2ParamsBody NEW Update of the group link.
//
// swagger:model contactGroupsUpdateContactGroup2ParamsBody
type ContactGroupsUpdateContactGroup2ParamsBody struct {

	// Group of contacts associated.
	Group *WebitelContactsLookup `json:"group,omitempty"`
}

// Validate validates this contact groups update contact group2 params body
func (m *ContactGroupsUpdateContactGroup2ParamsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGroup(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContactGroupsUpdateContactGroup2ParamsBody) validateGroup(formats strfmt.Registry) error {
	if swag.IsZero(m.Group) { // not required
		return nil
	}

	if m.Group != nil {
		if err := m.Group.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("group")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("group")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this contact groups update contact group2 params body based on the context it is used
func (m *ContactGroupsUpdateContactGroup2ParamsBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContactGroupsUpdateContactGroup2ParamsBody) contextValidateGroup(ctx context.Context, formats strfmt.Registry) error {

	if m.Group != nil {

		if swag.IsZero(m.Group) { // not required
			return nil
		}

		if err := m.Group.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("group")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("group")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContactGroupsUpdateContactGroup2ParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContactGroupsUpdateContactGroup2ParamsBody) UnmarshalBinary(b []byte) error {
	var res ContactGroupsUpdateContactGroup2ParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
