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

// EngineCreateChatPlanRequest engine create chat plan request
//
// swagger:model engineCreateChatPlanRequest
type EngineCreateChatPlanRequest struct {

	// description
	Description string `json:"description,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// schema
	Schema *EngineLookup `json:"schema,omitempty"`
}

// Validate validates this engine create chat plan request
func (m *EngineCreateChatPlanRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSchema(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EngineCreateChatPlanRequest) validateSchema(formats strfmt.Registry) error {
	if swag.IsZero(m.Schema) { // not required
		return nil
	}

	if m.Schema != nil {
		if err := m.Schema.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schema")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("schema")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this engine create chat plan request based on the context it is used
func (m *EngineCreateChatPlanRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSchema(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EngineCreateChatPlanRequest) contextValidateSchema(ctx context.Context, formats strfmt.Registry) error {

	if m.Schema != nil {

		if swag.IsZero(m.Schema) { // not required
			return nil
		}

		if err := m.Schema.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schema")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("schema")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EngineCreateChatPlanRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EngineCreateChatPlanRequest) UnmarshalBinary(b []byte) error {
	var res EngineCreateChatPlanRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
