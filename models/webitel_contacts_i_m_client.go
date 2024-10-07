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

// WebitelContactsIMClient A contact's [I]nstant[M]essaging client.
//
// swagger:model webitel.contacts.IMClient
type WebitelContactsIMClient struct {

	// App (Text-Gateway) used to connect the IM client.
	// Id will be internal id of gateway.
	// Name will be name of the gateway.
	App *WebitelContactsLookup `json:"app,omitempty"`

	// The user who created this Field.
	CreatedAt string `json:"created_at,omitempty"`

	// Timestamp(milli) of the Field creation.
	CreatedBy *WebitelContactsLookup `json:"created_by,omitempty"`

	// Unique ID of the latest version of the update.
	// This ID changes after any update to the underlying value(s).
	Etag string `json:"etag,omitempty"`

	// The unique ID of the association. Never changes.
	ID string `json:"id,omitempty"`

	// Protocol used to connect the IM client.
	Protocol string `json:"protocol,omitempty"`

	// Timestamp(milli) of the last Field update.
	// Take part in Etag generation.
	UpdatedAt string `json:"updated_at,omitempty"`

	// The user who performed last Update.
	UpdatedBy *WebitelContactsLookup `json:"updated_by,omitempty"`

	// External user which contacted to us.
	// Id will be from external service.
	// Name will be from external service.
	// Example: {"id":"596417343","name":"John Doe","type":"telegram"}
	User *WebitelContactsLookup `json:"user,omitempty"`

	// Version of the latest update. Numeric sequence.
	Ver int32 `json:"ver,omitempty"`

	// [Via] App(-specific) peer(-id) to connect[from] the IM client.
	Via string `json:"via,omitempty"`
}

// Validate validates this webitel contacts i m client
func (m *WebitelContactsIMClient) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedBy(formats); err != nil {
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

func (m *WebitelContactsIMClient) validateApp(formats strfmt.Registry) error {
	if swag.IsZero(m.App) { // not required
		return nil
	}

	if m.App != nil {
		if err := m.App.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("app")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("app")
			}
			return err
		}
	}

	return nil
}

func (m *WebitelContactsIMClient) validateCreatedBy(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedBy) { // not required
		return nil
	}

	if m.CreatedBy != nil {
		if err := m.CreatedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("created_by")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("created_by")
			}
			return err
		}
	}

	return nil
}

func (m *WebitelContactsIMClient) validateUpdatedBy(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedBy) { // not required
		return nil
	}

	if m.UpdatedBy != nil {
		if err := m.UpdatedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updated_by")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updated_by")
			}
			return err
		}
	}

	return nil
}

func (m *WebitelContactsIMClient) validateUser(formats strfmt.Registry) error {
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

// ContextValidate validate this webitel contacts i m client based on the context it is used
func (m *WebitelContactsIMClient) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateApp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreatedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdatedBy(ctx, formats); err != nil {
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

func (m *WebitelContactsIMClient) contextValidateApp(ctx context.Context, formats strfmt.Registry) error {

	if m.App != nil {

		if swag.IsZero(m.App) { // not required
			return nil
		}

		if err := m.App.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("app")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("app")
			}
			return err
		}
	}

	return nil
}

func (m *WebitelContactsIMClient) contextValidateCreatedBy(ctx context.Context, formats strfmt.Registry) error {

	if m.CreatedBy != nil {

		if swag.IsZero(m.CreatedBy) { // not required
			return nil
		}

		if err := m.CreatedBy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("created_by")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("created_by")
			}
			return err
		}
	}

	return nil
}

func (m *WebitelContactsIMClient) contextValidateUpdatedBy(ctx context.Context, formats strfmt.Registry) error {

	if m.UpdatedBy != nil {

		if swag.IsZero(m.UpdatedBy) { // not required
			return nil
		}

		if err := m.UpdatedBy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updated_by")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updated_by")
			}
			return err
		}
	}

	return nil
}

func (m *WebitelContactsIMClient) contextValidateUser(ctx context.Context, formats strfmt.Registry) error {

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
func (m *WebitelContactsIMClient) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebitelContactsIMClient) UnmarshalBinary(b []byte) error {
	var res WebitelContactsIMClient
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
