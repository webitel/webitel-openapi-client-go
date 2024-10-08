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

// LoggerLog logger log
//
// swagger:model loggerLog
type LoggerLog struct {

	// action
	Action string `json:"action,omitempty"`

	// config id
	ConfigID int32 `json:"config_id,omitempty"`

	// date
	Date string `json:"date,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// new state
	NewState string `json:"new_state,omitempty"`

	// object
	Object *LoggerLookup `json:"object,omitempty"`

	// record
	Record *LoggerLookup `json:"record,omitempty"`

	// user
	User *LoggerLookup `json:"user,omitempty"`

	// user ip
	UserIP string `json:"user_ip,omitempty"`
}

// Validate validates this logger log
func (m *LoggerLog) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateObject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecord(formats); err != nil {
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

func (m *LoggerLog) validateObject(formats strfmt.Registry) error {
	if swag.IsZero(m.Object) { // not required
		return nil
	}

	if m.Object != nil {
		if err := m.Object.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("object")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("object")
			}
			return err
		}
	}

	return nil
}

func (m *LoggerLog) validateRecord(formats strfmt.Registry) error {
	if swag.IsZero(m.Record) { // not required
		return nil
	}

	if m.Record != nil {
		if err := m.Record.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("record")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("record")
			}
			return err
		}
	}

	return nil
}

func (m *LoggerLog) validateUser(formats strfmt.Registry) error {
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

// ContextValidate validate this logger log based on the context it is used
func (m *LoggerLog) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateObject(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRecord(ctx, formats); err != nil {
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

func (m *LoggerLog) contextValidateObject(ctx context.Context, formats strfmt.Registry) error {

	if m.Object != nil {

		if swag.IsZero(m.Object) { // not required
			return nil
		}

		if err := m.Object.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("object")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("object")
			}
			return err
		}
	}

	return nil
}

func (m *LoggerLog) contextValidateRecord(ctx context.Context, formats strfmt.Registry) error {

	if m.Record != nil {

		if swag.IsZero(m.Record) { // not required
			return nil
		}

		if err := m.Record.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("record")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("record")
			}
			return err
		}
	}

	return nil
}

func (m *LoggerLog) contextValidateUser(ctx context.Context, formats strfmt.Registry) error {

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
func (m *LoggerLog) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LoggerLog) UnmarshalBinary(b []byte) error {
	var res LoggerLog
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
