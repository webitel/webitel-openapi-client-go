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

// EngineCreateSystemSettingRequest engine create system setting request
//
// swagger:model engineCreateSystemSettingRequest
type EngineCreateSystemSettingRequest struct {

	// name
	Name *EngineSystemSettingName `json:"name,omitempty"`

	// value
	Value interface{} `json:"value,omitempty"`
}

// Validate validates this engine create system setting request
func (m *EngineCreateSystemSettingRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EngineCreateSystemSettingRequest) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if m.Name != nil {
		if err := m.Name.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("name")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("name")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this engine create system setting request based on the context it is used
func (m *EngineCreateSystemSettingRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EngineCreateSystemSettingRequest) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if m.Name != nil {

		if swag.IsZero(m.Name) { // not required
			return nil
		}

		if err := m.Name.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("name")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("name")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EngineCreateSystemSettingRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EngineCreateSystemSettingRequest) UnmarshalBinary(b []byte) error {
	var res EngineCreateSystemSettingRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
