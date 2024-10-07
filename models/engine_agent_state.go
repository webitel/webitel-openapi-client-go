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

// EngineAgentState engine agent state
//
// swagger:model engineAgentState
type EngineAgentState struct {

	// agent
	Agent *EngineLookup `json:"agent,omitempty"`

	// channel
	Channel string `json:"channel,omitempty"`

	// duration
	Duration string `json:"duration,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// joined at
	JoinedAt string `json:"joined_at,omitempty"`

	// pause cause
	PauseCause string `json:"pause_cause,omitempty"`

	// payload
	Payload string `json:"payload,omitempty"`

	// queue
	Queue *EngineLookup `json:"queue,omitempty"`

	// state
	State string `json:"state,omitempty"`
}

// Validate validates this engine agent state
func (m *EngineAgentState) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAgent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EngineAgentState) validateAgent(formats strfmt.Registry) error {
	if swag.IsZero(m.Agent) { // not required
		return nil
	}

	if m.Agent != nil {
		if err := m.Agent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("agent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("agent")
			}
			return err
		}
	}

	return nil
}

func (m *EngineAgentState) validateQueue(formats strfmt.Registry) error {
	if swag.IsZero(m.Queue) { // not required
		return nil
	}

	if m.Queue != nil {
		if err := m.Queue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("queue")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("queue")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this engine agent state based on the context it is used
func (m *EngineAgentState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAgent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateQueue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EngineAgentState) contextValidateAgent(ctx context.Context, formats strfmt.Registry) error {

	if m.Agent != nil {

		if swag.IsZero(m.Agent) { // not required
			return nil
		}

		if err := m.Agent.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("agent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("agent")
			}
			return err
		}
	}

	return nil
}

func (m *EngineAgentState) contextValidateQueue(ctx context.Context, formats strfmt.Registry) error {

	if m.Queue != nil {

		if swag.IsZero(m.Queue) { // not required
			return nil
		}

		if err := m.Queue.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("queue")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("queue")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EngineAgentState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EngineAgentState) UnmarshalBinary(b []byte) error {
	var res EngineAgentState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
