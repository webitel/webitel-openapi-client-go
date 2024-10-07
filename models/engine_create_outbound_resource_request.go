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

// EngineCreateOutboundResourceRequest engine create outbound resource request
//
// swagger:model engineCreateOutboundResourceRequest
type EngineCreateOutboundResourceRequest struct {

	// description
	Description string `json:"description,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// error ids
	ErrorIds []string `json:"error_ids"`

	// failure dial delay
	FailureDialDelay int64 `json:"failure_dial_delay,omitempty"`

	// gateway
	Gateway *EngineLookup `json:"gateway,omitempty"`

	// limit
	Limit int32 `json:"limit,omitempty"`

	// max successively errors
	MaxSuccessivelyErrors int32 `json:"max_successively_errors,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// number
	Number string `json:"number,omitempty"`

	// parameters
	Parameters *EngineOutboundResourceParameters `json:"parameters,omitempty"`

	// patterns
	Patterns []string `json:"patterns"`

	// reserve
	Reserve bool `json:"reserve,omitempty"`

	// rps
	Rps int32 `json:"rps,omitempty"`

	// variables
	Variables map[string]string `json:"variables,omitempty"`
}

// Validate validates this engine create outbound resource request
func (m *EngineCreateOutboundResourceRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGateway(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParameters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EngineCreateOutboundResourceRequest) validateGateway(formats strfmt.Registry) error {
	if swag.IsZero(m.Gateway) { // not required
		return nil
	}

	if m.Gateway != nil {
		if err := m.Gateway.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gateway")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gateway")
			}
			return err
		}
	}

	return nil
}

func (m *EngineCreateOutboundResourceRequest) validateParameters(formats strfmt.Registry) error {
	if swag.IsZero(m.Parameters) { // not required
		return nil
	}

	if m.Parameters != nil {
		if err := m.Parameters.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parameters")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("parameters")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this engine create outbound resource request based on the context it is used
func (m *EngineCreateOutboundResourceRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGateway(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateParameters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EngineCreateOutboundResourceRequest) contextValidateGateway(ctx context.Context, formats strfmt.Registry) error {

	if m.Gateway != nil {

		if swag.IsZero(m.Gateway) { // not required
			return nil
		}

		if err := m.Gateway.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gateway")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gateway")
			}
			return err
		}
	}

	return nil
}

func (m *EngineCreateOutboundResourceRequest) contextValidateParameters(ctx context.Context, formats strfmt.Registry) error {

	if m.Parameters != nil {

		if swag.IsZero(m.Parameters) { // not required
			return nil
		}

		if err := m.Parameters.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parameters")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("parameters")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EngineCreateOutboundResourceRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EngineCreateOutboundResourceRequest) UnmarshalBinary(b []byte) error {
	var res EngineCreateOutboundResourceRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
