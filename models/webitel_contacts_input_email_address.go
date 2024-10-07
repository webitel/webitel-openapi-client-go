// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WebitelContactsInputEmailAddress Input of the Contact's email address.
// Example: {"email":"user@domain","etag":"1679792219687","primary":true,"type":{"name":"personal"},"verified":false}
//
// swagger:model webitel.contacts.InputEmailAddress
type WebitelContactsInputEmailAddress struct {

	// The email address.
	// Required: true
	Email *string `json:"email"`

	// Unique ID of the latest version of an existing resorce.
	Etag string `json:"etag,omitempty"`

	// Indicates whether this phone number is default within other channels of the same type(phone).
	Primary bool `json:"primary,omitempty"`

	// The type of the email address.
	// Lookup value from CommunicationType dictionary.
	// The type can be custom or one of these predefined values:
	// - home
	// - work
	// - other
	Type *WebitelContactsLookup `json:"type,omitempty"`

	// Indicate whether Contact, as a Person, realy owns this associated phone number.
	// In other words: whether Contact is reachable thru this 'email' communication channel ?
	Verified bool `json:"verified,omitempty"`
}

// Validate validates this webitel contacts input email address
func (m *WebitelContactsInputEmailAddress) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmail(formats); err != nil {
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

func (m *WebitelContactsInputEmailAddress) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	return nil
}

func (m *WebitelContactsInputEmailAddress) validateType(formats strfmt.Registry) error {
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

// ContextValidate validate this webitel contacts input email address based on the context it is used
func (m *WebitelContactsInputEmailAddress) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebitelContactsInputEmailAddress) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

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
func (m *WebitelContactsInputEmailAddress) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebitelContactsInputEmailAddress) UnmarshalBinary(b []byte) error {
	var res WebitelContactsInputEmailAddress
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
