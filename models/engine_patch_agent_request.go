// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EnginePatchAgentRequest engine patch agent request
//
// swagger:model enginePatchAgentRequest
type EnginePatchAgentRequest struct {

	// allow channels
	AllowChannels []string `json:"allow_channels"`

	// auditor
	Auditor []*EngineLookup `json:"auditor"`

	// chat count
	ChatCount int64 `json:"chat_count,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// fields
	Fields []string `json:"fields"`

	// greeting media
	GreetingMedia *EngineLookup `json:"greeting_media,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// is supervisor
	IsSupervisor bool `json:"is_supervisor,omitempty"`

	// progressive count
	ProgressiveCount int32 `json:"progressive_count,omitempty"`

	// region
	Region *EngineLookup `json:"region,omitempty"`

	// supervisor
	Supervisor []*EngineLookup `json:"supervisor"`

	// team
	Team *EngineLookup `json:"team,omitempty"`

	// user
	User *EngineLookup `json:"user,omitempty"`
}

// Validate validates this engine patch agent request
func (m *EnginePatchAgentRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuditor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGreetingMedia(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSupervisor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTeam(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EnginePatchAgentRequest) validateAuditor(formats strfmt.Registry) error {
	if swag.IsZero(m.Auditor) { // not required
		return nil
	}

	for i := 0; i < len(m.Auditor); i++ {
		if swag.IsZero(m.Auditor[i]) { // not required
			continue
		}

		if m.Auditor[i] != nil {
			if err := m.Auditor[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("auditor" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("auditor" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EnginePatchAgentRequest) validateGreetingMedia(formats strfmt.Registry) error {
	if swag.IsZero(m.GreetingMedia) { // not required
		return nil
	}

	if m.GreetingMedia != nil {
		if err := m.GreetingMedia.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("greeting_media")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("greeting_media")
			}
			return err
		}
	}

	return nil
}

func (m *EnginePatchAgentRequest) validateRegion(formats strfmt.Registry) error {
	if swag.IsZero(m.Region) { // not required
		return nil
	}

	if m.Region != nil {
		if err := m.Region.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("region")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("region")
			}
			return err
		}
	}

	return nil
}

func (m *EnginePatchAgentRequest) validateSupervisor(formats strfmt.Registry) error {
	if swag.IsZero(m.Supervisor) { // not required
		return nil
	}

	for i := 0; i < len(m.Supervisor); i++ {
		if swag.IsZero(m.Supervisor[i]) { // not required
			continue
		}

		if m.Supervisor[i] != nil {
			if err := m.Supervisor[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("supervisor" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("supervisor" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EnginePatchAgentRequest) validateTeam(formats strfmt.Registry) error {
	if swag.IsZero(m.Team) { // not required
		return nil
	}

	if m.Team != nil {
		if err := m.Team.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("team")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("team")
			}
			return err
		}
	}

	return nil
}

func (m *EnginePatchAgentRequest) validateUser(formats strfmt.Registry) error {
	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this engine patch agent request based on the context it is used
func (m *EnginePatchAgentRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAuditor(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGreetingMedia(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRegion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSupervisor(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTeam(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EnginePatchAgentRequest) contextValidateAuditor(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Auditor); i++ {

		if m.Auditor[i] != nil {

			if swag.IsZero(m.Auditor[i]) { // not required
				return nil
			}

			if err := m.Auditor[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("auditor" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("auditor" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EnginePatchAgentRequest) contextValidateGreetingMedia(ctx context.Context, formats strfmt.Registry) error {

	if m.GreetingMedia != nil {

		if swag.IsZero(m.GreetingMedia) { // not required
			return nil
		}

		if err := m.GreetingMedia.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("greeting_media")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("greeting_media")
			}
			return err
		}
	}

	return nil
}

func (m *EnginePatchAgentRequest) contextValidateRegion(ctx context.Context, formats strfmt.Registry) error {

	if m.Region != nil {

		if swag.IsZero(m.Region) { // not required
			return nil
		}

		if err := m.Region.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("region")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("region")
			}
			return err
		}
	}

	return nil
}

func (m *EnginePatchAgentRequest) contextValidateSupervisor(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Supervisor); i++ {

		if m.Supervisor[i] != nil {

			if swag.IsZero(m.Supervisor[i]) { // not required
				return nil
			}

			if err := m.Supervisor[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("supervisor" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("supervisor" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EnginePatchAgentRequest) contextValidateTeam(ctx context.Context, formats strfmt.Registry) error {

	if m.Team != nil {

		if swag.IsZero(m.Team) { // not required
			return nil
		}

		if err := m.Team.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("team")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("team")
			}
			return err
		}
	}

	return nil
}

func (m *EnginePatchAgentRequest) contextValidateUser(ctx context.Context, formats strfmt.Registry) error {

	if m.User != nil {

		if swag.IsZero(m.User) { // not required
			return nil
		}

		if err := m.User.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EnginePatchAgentRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EnginePatchAgentRequest) UnmarshalBinary(b []byte) error {
	var res EnginePatchAgentRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
