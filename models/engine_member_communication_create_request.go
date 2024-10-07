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

// EngineMemberCommunicationCreateRequest engine member communication create request
//
// swagger:model engineMemberCommunicationCreateRequest
type EngineMemberCommunicationCreateRequest struct {

	// description
	Description string `json:"description,omitempty"`

	// destination
	Destination string `json:"destination,omitempty"`

	// display
	Display string `json:"display,omitempty"`

	// dtmf
	Dtmf string `json:"dtmf,omitempty"`

	// priority
	Priority int32 `json:"priority,omitempty"`

	// resource
	Resource *EngineLookup `json:"resource,omitempty"`

	// stop at
	StopAt string `json:"stop_at,omitempty"`

	// type
	Type *EngineLookup `json:"type,omitempty"`
}

// Validate validates this engine member communication create request
func (m *EngineMemberCommunicationCreateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EngineMemberCommunicationCreateRequest) validateResource(formats strfmt.Registry) error {
	if swag.IsZero(m.Resource) { // not required
		return nil
	}

	if m.Resource != nil {
		if err := m.Resource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resource")
			}
			return err
		}
	}

	return nil
}

func (m *EngineMemberCommunicationCreateRequest) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this engine member communication create request based on the context it is used
func (m *EngineMemberCommunicationCreateRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateResource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EngineMemberCommunicationCreateRequest) contextValidateResource(ctx context.Context, formats strfmt.Registry) error {

	if m.Resource != nil {

		if swag.IsZero(m.Resource) { // not required
			return nil
		}

		if err := m.Resource.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resource")
			}
			return err
		}
	}

	return nil
}

func (m *EngineMemberCommunicationCreateRequest) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {

		if swag.IsZero(m.Type) { // not required
			return nil
		}

		if err := m.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EngineMemberCommunicationCreateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EngineMemberCommunicationCreateRequest) UnmarshalBinary(b []byte) error {
	var res EngineMemberCommunicationCreateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
