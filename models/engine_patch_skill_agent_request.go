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

// EnginePatchSkillAgentRequest engine patch skill agent request
//
// swagger:model enginePatchSkillAgentRequest
type EnginePatchSkillAgentRequest struct {

	// agent id
	AgentID []string `json:"agent_id"`

	// capacity
	Capacity int32 `json:"capacity,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// fields
	Fields []string `json:"fields"`

	// id
	ID []string `json:"id"`

	// skill
	Skill *EngineLookup `json:"skill,omitempty"`

	// skill id
	SkillID string `json:"skill_id,omitempty"`
}

// Validate validates this engine patch skill agent request
func (m *EnginePatchSkillAgentRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSkill(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EnginePatchSkillAgentRequest) validateSkill(formats strfmt.Registry) error {
	if swag.IsZero(m.Skill) { // not required
		return nil
	}

	if m.Skill != nil {
		if err := m.Skill.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("skill")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("skill")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this engine patch skill agent request based on the context it is used
func (m *EnginePatchSkillAgentRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSkill(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EnginePatchSkillAgentRequest) contextValidateSkill(ctx context.Context, formats strfmt.Registry) error {

	if m.Skill != nil {

		if swag.IsZero(m.Skill) { // not required
			return nil
		}

		if err := m.Skill.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("skill")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("skill")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EnginePatchSkillAgentRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EnginePatchSkillAgentRequest) UnmarshalBinary(b []byte) error {
	var res EnginePatchSkillAgentRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
